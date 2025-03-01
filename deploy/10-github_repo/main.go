package main

import (
	"github.com/kemadev/iac-components/pkg/github/repo"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		err := repo.Wrapper(ctx, repo.WrapperArgs{
			Repository: repo.RepositoryArgs{
				Name:        "repository-template",
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
		})
		if err != nil {
			return err
		}
		return nil
	})
}
