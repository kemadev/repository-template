package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// pulumi import github:index/repository:Repository repo repository-template --provider 'urn=urn:pulumi:dev::github-com-kemadev-repository-template-tmp-iac-github-repo-bootstrap::pulumi:providers:github::github'

var (
	args = WrapperArgs{
		provider: ProviderArgs{
			owner: "kemadev",
		},
		actions: ActionsArgs{
			Actions: []string{
				// Internal workflows and actions
				"kemadev/workflows-and-actions/.github/workflows/*",
				"kemadev/workflows-and-actions/.github/actions/*",
				// Actions from reusable workflows
				"anchore/sbom-action@*",
				"anchore/scan-action@*",
				"aws-actions/configure-aws-credentials@*",
				"DavidAnson/markdownlint-cli2-action@*",
				// "docker://rhysd/actionlint@*", // TODO check if Docker actions are supported
				"golangci/golangci-lint-action@*",
				"googleapis/release-please-action@*",
				"goreleaser/goreleaser-action@*",
				"hadolint/hadolint-action@*",
				"ibiqlik/action-yamllint@*",
				"pulumi/actions@*",
				// "semgrep/semgrep@*", // TODO Check if container workflows are supported
				"trufflesecurity/trufflehog@*",
			},
		},
		branches: BranchesArgs{
			Dev:     "dev",
			Next:    "next",
			Prod:    "main",
			Default: "main",
		},
		envs: EnvsArgs{
			Dev:  "dev",
			Next: "next",
			Prod: "prod",
		},
		rulesets: RulesetsArgs{
			requiredReviewersNext: 1,
			requiredReviewersProd: 1,
		},
		repository: RepositoryArgs{
			Description: "GitHub on steroids!",
			Topics: []string{
				"repository-template",
				"pulumi",
				"iac",
				"go",
			},
		},
	}
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		err := Wrapper(ctx, args)
		if err != nil {
			return err
		}
		return nil
	})

}
