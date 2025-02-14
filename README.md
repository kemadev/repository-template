# Project Name

<!-- Brief description of the project -->

## Welcome

<!-- Project presentation, motivation, and main features -->

## Demo

<!-- If applicable, project demo (video, screenshots, asciicinema, ...) -->

## Getting Started

### Installation

<!-- Prerequisites and installation instructions -->

### Usage

<!-- Basic usage and main commands -->

## Contributing

### Global Guidelines

- You can find contributing guidelines in [CONTRIBUTING.md](/../../../../kemadev/.github/blob/main/.github/CONTRIBUTING.md).
- Feeling like something could be improved? Let's do it! From code to documentation, services to use, or linter rules, everything is discussable and improvable. Let's move forward together!

### Concepts

- A GitHub repository such as this one is representing a project
- A project is basically an application
- An application is composed of a set of microservices that works together to achieve the project's goal
- Microservices are small, loosely coupled, and independently deployable and scalable
- Each microservice should be agnostic of it downstreams. However, it should expose a clear and well-defined API to its downstreams for them to consume (that is, the microservice itself uses its upstreams' API)

### Development Guidelines and Conventions

- All major directories contain a [README.md](README.md) file with a brief description of directory's content
- Code sharing is encouraged, such code should be placed in [pkg](pkg) and [internal/pkg](internal/pkg) directories
- Importing other applications libraries and packages is encouraged, following code sharing encouragement
- First class code documentation (following [Go doc comment guidelines](https://go.dev/doc/comment)) as well as project documentation is encouraged
- Following [Effective Go](https://go.dev/doc/effective_go) standards is encouraged
- Variables, functions, methods, ... should be named in a short and descriptive way, avoiding acronyms usage as much as possible

### Project development

#### Prerequisites

- [Docker](https://github.com/docker/cli) and [Docker Compose](https://github.com/docker/compose) to run applications in containers. You should configure your credentials store and credential helpers for Docker to work with your container registry
- [Go](https://github.com/golang/go) to install applications dependencies as needed
- [Task](https://github.com/go-task/task) to run common commands such as running, testing, linting, building, ...
- [AWS CLI](https://github.com/aws/aws-cli) configured using the company's config files available as a release asset in the company's repository
- Very few other CLI tools such as [curl](https://github.com/curl/curl), [git](https://github.com/git/git), ... that are most likely already installed on your system

#### Running the project

- Common tasks such as running, formatting, linting, testing, ... are done via [Task](https://github.com/go-task/task). [Taskfile.yaml](Taskfile.yaml) is the root Taskfile, and include other taskfiles located in sub-directories
- You should run `task ci` before pushing your changes to ensure your changes are expected to pass the CI pipelines. However, please note that these tasks are meant to be run in a development environment, and might not be on-par with CI/CD pipelines
- You can create your own (non git-tracked) taskfiles in [tool/task/custom](tool/task/custom) directory to extend the project's tasks for your personal needs. Tasks defined in such custom taskfiles will be available in the root Taskfile thanks to the [`includes` directive](tool/task/Taskfile.yaml#L10)
- For specific application setup such as debugging and full-fledged development mode, please refer to `tool/docker-compose.yaml` and its profiles with associated comments

#### CI / CD

- CI / CD pipelines are fully automated and managed by GitHub Actions. You can find the workflows in [.github/workflows](.github/workflows) directory.
- However, they can sometime report false positives. Here is what you can do to remediate (be as specific as possible on silences to avoid shadowing real issues):
  - `golangci-lint`: Add a `nolint:<linter>[,<linter>]` comment. See [this doc](https://golangci-lint.run/usage/false-positives/)
  - `semgrep`: Add a `nosemgrep: <rule-id>` comment. See [this doc](https://semgrep.dev/docs/ignoring-files-folders-code)
  - `trufflehog`: Add a `trufflehog:ignore` comment. See [this doc](https://github.com/trufflesecurity/trufflehog/blob/main/README.md#question-faq). Please note that **any leaked secret should be revoked and replaced as soon as possible**
  - `yamllint`: Add a `yamllint disable-line rule:<rule>` comment. See [this doc](https://yamllint.readthedocs.io/en/stable/disable_with_comments.html)
  - `markdownlint`: Add a `markdownlint-disable <rule>` comment. See [this doc](https://github.com/DavidAnson/markdownlint/blob/main/README.md#configuration)
  - `shellcheck`: Add a `shellcheck disable=<rule>` comment. See [this doc](https://github.com/koalaman/shellcheck/wiki/Ignore)
  - `hadolint`: Add a `hadolint ignore=<rule>` comment. See [this doc](https://github.com/hadolint/hadolint/blob/master/README.md#ignoring-rules)
  - `actionlint`: In case of a `shellcheck` error, refer to the `shellcheck` section. Otherwise, you can pass arguments to the linting action to ignore specific rules. See [this doc](https://github.com/rhysd/actionlint/blob/main/docs/usage.md#ignore-some-errors)

## Acknowledgments

- **[Every contributor](/../../graphs/contributors)**, as well as the open-source community, for making this project possible!
- **[golang-standards/project-layout](https://github.com/golang-standards/project-layout)** broadly inspired the project structure and conventions
- **[Pulumi](https://github.com/pulumi)** inspired the full usage of GitHub's features and the need for a complete GitOps approach with full automation
- **[oh-my-posh](https://github.com/jandedobbeleer/oh-my-posh)** inspired some conventions and the sharing way of thinking, and provides a wonderful shell prompt!
- **[super-linter](https://github.com/super-linter/super-linter)** inspired some of the CI/CD pipeline conventions and the need for a full-fledged linting solution
