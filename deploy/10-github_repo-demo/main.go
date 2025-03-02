package main

import (
	"github.com/kemadev/iac-components/pkg/github/repo"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		err := repo.Wrapper(ctx, repo.WrapperArgs{
			Repository: repo.RepositoryArgs{
				Name:        "repo-templating-demo",
				Description: "Yikes! This is a template repository",
				Visibility:  "public",
				Topics: []string{
					"repository-template",
					"pulumi",
					"iac",
					"go",
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
