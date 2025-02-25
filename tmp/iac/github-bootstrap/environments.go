package main

import (
	"github.com/kema-dev/infra-definition/pkg/bootstrap/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnvsArgs struct {
	Dev     string
	Next    string
	Prod    string
	Default string
}

type TEnvironmentsCreated struct {
	dev  *github.RepositoryEnvironment
	next *github.RepositoryEnvironment
	prod *github.RepositoryEnvironment
}

func createEnvironments(ctx *pulumi.Context, provider *github.Provider, repo *github.Repository, argsEnvs EnvsArgs, argsBranches BranchesArgs) (TEnvironmentsCreated, error) {
	deploymentEnvironmentDevName := util.FormatResourceName("Deployment environment dev")
	deploymentEnvironmentDev, err := github.NewRepositoryEnvironment(ctx, deploymentEnvironmentDevName, &github.RepositoryEnvironmentArgs{
		Repository:        repo.Name,
		Environment:       pulumi.String(argsEnvs.Dev),
		CanAdminsBypass:   pulumi.Bool(true),
		PreventSelfReview: pulumi.Bool(false),
		WaitTimer:         pulumi.Int(0),
		// Trust PR reviews
		Reviewers: github.RepositoryEnvironmentReviewerArray{},
		DeploymentBranchPolicy: github.RepositoryEnvironmentDeploymentBranchPolicyArgs{
			CustomBranchPolicies: pulumi.Bool(true),
			ProtectedBranches:    pulumi.Bool(false),
		},
	}, pulumi.Provider(provider))
	if err != nil {
		return TEnvironmentsCreated{}, err
	}

	deploymentEnvironmentNextName := util.FormatResourceName("Deployment environment next")
	deploymentEnvironmentNext, err := github.NewRepositoryEnvironment(ctx, deploymentEnvironmentNextName, &github.RepositoryEnvironmentArgs{
		Repository:        repo.Name,
		Environment:       pulumi.String(argsEnvs.Next),
		CanAdminsBypass:   pulumi.Bool(true),
		PreventSelfReview: pulumi.Bool(false),
		WaitTimer:         pulumi.Int(0),
		// Trust PR reviews
		Reviewers: github.RepositoryEnvironmentReviewerArray{},
		DeploymentBranchPolicy: github.RepositoryEnvironmentDeploymentBranchPolicyArgs{
			CustomBranchPolicies: pulumi.Bool(true),
			ProtectedBranches:    pulumi.Bool(false),
		},
	}, pulumi.Provider(provider))
	if err != nil {
		return TEnvironmentsCreated{}, err
	}

	deploymentEnvironmentProdName := util.FormatResourceName("Deployment environment prod")
	deploymentEnvironmentProd, err := github.NewRepositoryEnvironment(ctx, deploymentEnvironmentProdName, &github.RepositoryEnvironmentArgs{
		Repository:        repo.Name,
		Environment:       pulumi.String(argsEnvs.Prod),
		CanAdminsBypass:   pulumi.Bool(true),
		PreventSelfReview: pulumi.Bool(false),
		WaitTimer:         pulumi.Int(0),
		// Trust PR reviews
		Reviewers: github.RepositoryEnvironmentReviewerArray{},
		DeploymentBranchPolicy: github.RepositoryEnvironmentDeploymentBranchPolicyArgs{
			CustomBranchPolicies: pulumi.Bool(true),
			ProtectedBranches:    pulumi.Bool(false),
		},
	}, pulumi.Provider(provider))
	if err != nil {
		return TEnvironmentsCreated{}, err
	}

	repositoryEnvironmentDeploymentPolicyDevName := util.FormatResourceName("Repository environment deployment policy dev")
	_, err = github.NewRepositoryEnvironmentDeploymentPolicy(ctx, repositoryEnvironmentDeploymentPolicyDevName, &github.RepositoryEnvironmentDeploymentPolicyArgs{
		Repository:    repo.Name,
		Environment:   deploymentEnvironmentDev.Environment,
		BranchPattern: pulumi.String(argsBranches.Dev),
	}, pulumi.Provider(provider))
	if err != nil {
		return TEnvironmentsCreated{}, err
	}

	repositoryEnvironmentDeploymentPolicyNextName := util.FormatResourceName("Repository environment deployment policy next")
	_, err = github.NewRepositoryEnvironmentDeploymentPolicy(ctx, repositoryEnvironmentDeploymentPolicyNextName, &github.RepositoryEnvironmentDeploymentPolicyArgs{
		Repository:    repo.Name,
		Environment:   deploymentEnvironmentNext.Environment,
		BranchPattern: pulumi.String(argsBranches.Next),
	}, pulumi.Provider(provider))
	if err != nil {
		return TEnvironmentsCreated{}, err
	}

	repositoryEnvironmentDeploymentPolicyProdName := util.FormatResourceName("Repository environment deployment policy prod")
	_, err = github.NewRepositoryEnvironmentDeploymentPolicy(ctx, repositoryEnvironmentDeploymentPolicyProdName, &github.RepositoryEnvironmentDeploymentPolicyArgs{
		Repository:    repo.Name,
		Environment:   deploymentEnvironmentProd.Environment,
		BranchPattern: pulumi.String(argsBranches.Prod),
	}, pulumi.Provider(provider))
	if err != nil {
		return TEnvironmentsCreated{}, err
	}

	return TEnvironmentsCreated{
		dev:  deploymentEnvironmentDev,
		next: deploymentEnvironmentNext,
		prod: deploymentEnvironmentProd,
	}, nil
}
