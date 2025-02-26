package main

import (
	"github.com/kemadev/iac-components/pkg/github/repo"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		err := repo.Wrapper(ctx, repo.WrapperArgs{
			Repository: repo.RepositoryArgs{
				Description: "GitHub on steroids!",
				Visibility:  "public",
				IsTemplate:  true,
				Topics: []string{
					"repository-template",
					"pulumi",
					"iac",
					"go",
				},
			},
			Actions: repo.ActionsArgs{
				Actions: []string{
					"anchore/sbom-action@*",
					"anchore/scan-action@*",
					"aws-actions/configure-aws-credentials@*",
					"DavidAnson/markdownlint-cli2-action@*",
					"docker://rhysd/actionlint@*",
					"golangci/golangci-lint-action@*",
					"googleapis/release-please-action@*",
					"goreleaser/goreleaser-action@*",
					"hadolint/hadolint-action@*",
					"ibiqlik/action-yamllint@*",
					"pulumi/actions@*",
					"semgrep/semgrep@*",
					"trufflesecurity/trufflehog@*",
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
