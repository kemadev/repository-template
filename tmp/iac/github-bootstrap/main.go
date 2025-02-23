package main

import (
	"fmt"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// pulumi import github:index/repository:Repository repo repository-template --provider 'urn=urn:pulumi:dev::github-com-kemadev-repository-template-tmp-iac-github-repo-bootstrap::pulumi:providers:github::github'

type RepositoryArgs struct {
	// Repository description
	// - required
	Description string
	// Repository homepage URL
	HomepageUrl string
	// Repository topics
	// - required
	Topics []string
	// Repository visibility
	// - required
	Visibility string
	// Repository is a template
	// - false by default
	IsTemplate bool
}

func checkBootstrapGithubRepositoryArgs(args RepositoryArgs) error {
	if args.Description == "" {
		return fmt.Errorf("Description is required")
	}
	if len(args.Topics) == 0 {
		return fmt.Errorf("Topics is required")
	}
	if args.Visibility != "public" && args.Visibility != "private" {
		return fmt.Errorf("Visibility must be 'public' or 'private'")
	}
	return nil
}

func bootstrapGithubRepository(ctx *pulumi.Context, provider *github.Provider, args RepositoryArgs) (*github.Repository, error) {
	err := checkBootstrapGithubRepositoryArgs(args)
	if err != nil {
		return nil, err
	}
	repo, err := github.NewRepository(ctx, "repo", &github.RepositoryArgs{
		// Keep name from import
		// Name:        pulumi.String("repository-template"),

		// Prevent accidental deletion
		ArchiveOnDestroy: pulumi.Bool(true),
		// Allow non-admins read access from pulumi
		IgnoreVulnerabilityAlertsDuringRead: pulumi.Bool(true),

		Description: pulumi.String(args.Description),
		HomepageUrl: pulumi.String(args.HomepageUrl),
		Topics: func() pulumi.StringArray {
			var topics pulumi.StringArray
			for _, topic := range args.Topics {
				topics = append(topics, pulumi.String(topic))
			}
			return topics
		}(),
		Visibility: pulumi.String(args.Visibility),
		IsTemplate: pulumi.Bool(args.IsTemplate == true),

		AllowAutoMerge:           pulumi.Bool(true),
		AllowMergeCommit:         pulumi.Bool(false),
		AllowRebaseMerge:         pulumi.Bool(false),
		AllowSquashMerge:         pulumi.Bool(true),
		SquashMergeCommitTitle:   pulumi.String("PR_TITLE"),
		SquashMergeCommitMessage: pulumi.String("COMMIT_MESSAGES"),
		AllowUpdateBranch:        pulumi.Bool(true),
		DefaultBranch:            pulumi.String("main"),
		DeleteBranchOnMerge:      pulumi.Bool(true),
		HasDiscussions:           pulumi.Bool(true),
		HasIssues:                pulumi.Bool(true),
		HasProjects:              pulumi.Bool(true),
		HasWiki:                  pulumi.Bool(false),
		WebCommitSignoffRequired: pulumi.Bool(false),

		VulnerabilityAlerts: func() pulumi.BoolPtrInput {
			if args.Visibility == "public" {
				return pulumi.BoolPtr(true)
			}
			// Advanced Security is required for private repositories
			return pulumi.BoolPtr(false)
		}(),
		SecurityAndAnalysis: github.RepositorySecurityAndAnalysisArgs{
			AdvancedSecurity: func() github.RepositorySecurityAndAnalysisAdvancedSecurityPtrInput {
				if args.Visibility == "public" {
					return nil
				}
				// Advanced Security is required for private repositories
				return nil
			}(),
			SecretScanning: func() github.RepositorySecurityAndAnalysisSecretScanningPtrInput {
				if args.Visibility == "public" {
					return github.RepositorySecurityAndAnalysisSecretScanningArgs{
						Status: pulumi.String("enabled"),
					}
				}
				// Advanced Security is required for private repositories
				return nil
			}(),
			SecretScanningPushProtection: func() github.RepositorySecurityAndAnalysisSecretScanningPushProtectionPtrInput {
				if args.Visibility == "public" {
					return github.RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs{
						Status: pulumi.String("enabled"),
					}
				}
				// Advanced Security is required for private repositories
				return nil
			}(),
		},
	}, pulumi.Provider(provider))
	if err != nil {
		return nil, err
	}
	_, err = github.NewBranchDefault(ctx, "branch", &github.BranchDefaultArgs{
		Repository: repo.Name,
		Branch:     pulumi.String("main"),
	}, pulumi.Provider(provider))
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		githubProvider, err := github.NewProvider(ctx, "github", &github.ProviderArgs{
			Owner: pulumi.String("kemadev"),
		})
		if err != nil {
			return err
		}
		repo, err := bootstrapGithubRepository(ctx, githubProvider, RepositoryArgs{
			Description: "GitHub on steroids!",
			Topics: []string{
				"repository-template",
				"pulumi",
				"iac",
				"go",
			},
			Visibility: "public",
		})
		if err != nil {
			return err
		}

		_ = repo
		return nil
	})
}
