// File created by repo-as-code, however you can still modify it as you like!
package main

import (
	"github.com/kemadev/iac-components/pkg/github/repo"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		err := repo.Wrapper(ctx, repo.WrapperArgs{
			Repository: repo.RepositoryArgs{
				Name: "repo-as-code-demo",
				// Set description
				Description: "repo-as-code-demo",
				// Set visibility
				Visibility: "public",
				// Optionally set other properties such as teams, direct members, etc.
			},
			Codeowners: repo.CodeownersArgs{
				Codeowners: []repo.CodeownerParam{
					// Set codeowners
					{
						Path:   "*",
						Entity: "@kemadev/maintainers",
					},
				},
			},
			// Optionally set other properties
		})
		if err != nil {
			return err
		}
		return nil
	})
}
