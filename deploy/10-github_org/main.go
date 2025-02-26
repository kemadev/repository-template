package main

import (
	"github.com/kemadev/iac-components/pkg/github/org"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		err := org.Wrapper(ctx, org.WrapperArgs{
			Settings: org.SettingsArgs{
				Company:      "kemadev",
				Description:  "Making cloud infrastructure a breeze!",
				Email:        "contact@kema.dev",
				BillingEmail: "billing@kema.dev",
				Blog:         "https://www.kema.dev",
				Location:     "France",
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
