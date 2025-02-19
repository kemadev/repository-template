# `deploy`

> [!NOTE]
> Just use `task` to manage components initialization, it handles everything for you!

## Directories in this directory

- Should be named with a prefix that defines the order, such as `10-`, `20-`, ...
- Are deployed sequentially in alphabetical order, so the prefix is important
- Should be named after the component they deploy (e.g., `deploy/XX-network`, `deploy/XX-database`, `deploy/XX-app1`, `deploy/XX-app2`, ...)
- For functions, following more or less the same structure as the `cmd` directory
- For infrastructure, be named in a meaningful way, such as `deploy/XX-network`, `deploy/XX-database`, `deploy/XX-app`, ...

## Files in this directory

- Are placed in subdirectories, see above
- Are related to deployments
- Should manage application deployment resources for applications, including different environments (e.g., `dev`, `next`, ...)
- Should name their projects according to the URL of their directory, replacing non-alphanumeric characters with `-` (e.g., `github-com-username-repo-deploy-XX-app1` for `deploy/XX-app1`).
- Should name their stacks according to the environment they deploy to (e.g., `next`, `prod`, ...)
- Should implement GitOps best practices
- Should be moved to a separate repository if it can be reused across multiple projects
- Should be as simple as possible, making it possible to use stack references to share resources across projects
- Should manage deployment resources for applications
