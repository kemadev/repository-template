# Intermediate image providing go dependencies
FROM --platform=${BUILDPLATFORM} public.ecr.aws/amazonlinux/amazonlinux:2023 AS go-deps

SHELL ["/bin/bash", "-uo", "pipefail", "-c"]

# Docker buildx ARGs
ARG TARGETOS
ARG TARGETARCH

WORKDIR /deps

# Install packages required for go installation
# hadolint ignore=DL3041
RUN dnf install --assumeyes tar gzip zip && dnf clean all

# Set build environment
ENV GOOS=${TARGETOS} GOARCH=${TARGETARCH}

# Install go version specified in go.mod
COPY go.mod /tmp/go.mod
RUN curl -sSL "https://go.dev/dl/go$(grep -oP 'go\s+\K([0-9.]+)' < /tmp/go.mod).${GOOS}-${GOARCH}.tar.gz" | tar -C /usr/local -xz

# Ensure GOPATH exists, even if no Go command has been run
RUN mkdir -p /root/go



# Intermediate image providing global dependencies, cleaned up
FROM --platform=${BUILDPLATFORM} public.ecr.aws/amazonlinux/amazonlinux:2023 AS clean-deps

# Docker buildx ARGs
ARG TARGETOS
ARG TARGETARCH
# Disable C compatability, pass build arg ENABLE_CGO=1 to enable
ARG ENABLE_CGO=0

WORKDIR /deps

# Copy go
COPY --link --from=go-deps /usr/local/go /usr/local/go
COPY --link --from=go-deps /root/go /root/go

# Set build environment
ENV CGO_ENABLED=${ENABLE_CGO} GOOS=${TARGETOS} GOARCH=${TARGETARCH}
# Add go to PATH (not using ${HOME} as it is not working in ENV directive)
ENV PATH="/root/go/bin:/usr/local/go/bin:${PATH}"


# Intermediate image providing dependencies for development
FROM --platform=${BUILDPLATFORM} clean-deps AS dev-deps

# Install air for live reloading
RUN go install github.com/air-verse/air@latest

# Set environment variables defined in docker compose definition, default to nonexistent to avoid accidental use
ENV GO_MAIN_FILE_DIR="/nonexistent"



# Intermediate image providing dependencies for cloud development
FROM --platform=${BUILDPLATFORM} dev-deps AS dev-cloud-base

# Use volume mounted in docker-compose for live reloading
WORKDIR /app

# Install awscli (can't use a aws-deps image and COPY as awscli uses system libraries like libpython, ...)
# Convert target arch to awscli arch
RUN case ${TARGETARCH} in \
		amd64) echo x86_64 > /tmp/aws_cli_arch ;; \
		arm64) echo arm64 > /tmp/aws_cli_arch ;; \
		*) echo "Unsupported architecture: ${TARGETARCH}"; exit 1 ;; \
	esac

# Install packages required for cloud deployment (provided awscli is obsolete, see https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html#getting-started-install-instructions)
# hadolint ignore=DL3041
RUN dnf remove --assumeyes awscli && dnf install --assumeyes zip && dnf clean all
RUN curl -sSL "https://awscli.amazonaws.com/awscli-exe-${TARGETOS}-$(cat /tmp/aws_cli_arch).zip" -o /tmp/awscliv2.zip && unzip /tmp/awscliv2.zip -d /tmp && /tmp/aws/install

# Set environment variables defined in docker compose definition, default to nonexistent to avoid accidental use
ENV LAMBDA_FUNCTION_NAME="nonexistent"
# Set build environment for cloud deployment, lambda runs on linux
ENV GOOS=linux



# Dev image for cloud development targetting arm64
FROM --platform=${BUILDPLATFORM} dev-cloud-base AS dev-cloud-arm64

ENV GOARCH=arm64

# Start with live reloading (live deployment)
CMD ["air", "-c", "config/air/.air-cloud.toml"]



# Dev image for cloud development targetting amd64
FROM --platform=${BUILDPLATFORM} dev-cloud-base AS dev-cloud-amd64

ENV GOARCH=amd64

# Start with live reloading (live deployment)
CMD ["air", "-c", "config/air/.air-cloud.toml"]



# Dev image for local development
FROM --platform=${BUILDPLATFORM} public.ecr.aws/lambda/provided:al2023 AS dev-local

# Docker buildx ARGs matching local development environment
ARG TARGETOS
ARG TARGETARCH
# Disable C compatibility (ENABLE_CGO=1 to enable)
ARG ENABLE_CGO=0

# Set build environment for local tools
ENV CGO_ENABLED=${ENABLE_CGO} GOOS=${TARGETOS} GOARCH=${TARGETARCH}

# Use volume mounted in docker-compose for live reloading
WORKDIR /app

# Copy go
COPY --link --from=dev-deps /usr/local/go /usr/local/go
COPY --link --from=dev-deps /root/go /root/go

# Add go to PATH (not using ${HOME} as it is not working in ENV directive)
ENV PATH="/root/go/bin:/usr/local/go/bin:${PATH}"

# Set environment variables defined in docker compose definition
ENV GO_MAIN_FILE_DIR="/nonexistent"

# Start with live reloading
ENTRYPOINT ["air", "-c", "config/air/.air-local.toml"]



# Intermediate image providing build dependencies
FROM --platform=${BUILDPLATFORM} clean-deps AS base

# Add cache layer for dependencies, mount from host to avoid COPY layer, verbose output for cache behavior logging
RUN --mount=type=cache,target=/root/go/pkg \
	--mount=type=bind,source=go.mod,target=go.mod \
	--mount=type=bind,source=go.sum,target=go.sum \
	go mod download -x



# Intermediate image providing build artifacts and using build stage cache
FROM base AS build

# Which main file (application) to build, default to nonexistent to avoid accidental use
ARG GO_MAIN_FILE_DIR="/nonexistent"

# Default directory to use
WORKDIR /app

# Add cache layer for dependencies, mount from host to avoid COPY layer
RUN --mount=type=cache,target=/root/go/pkg \
	--mount=type=cache,target=/root/.cache \
	--mount=type=bind,target=. \
	go build -ldflags='-s -w' -tags lambda.norpc -o /usr/bin/app ./cmd/${GO_MAIN_FILE_DIR}



# Zip archive
FROM build AS archive

WORKDIR /app

# Install zip
# hadolint ignore=DL3041
RUN dnf install --assumeyes zip && dnf clean all && zip -j /usr/bin/app.zip /usr/bin/app



# Local zip archive locally using `--output <local path>` flag
FROM scratch AS local-archive

COPY --from=archive /usr/bin/app.zip /
