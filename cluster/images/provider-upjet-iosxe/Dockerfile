FROM alpine:3.23.4
RUN apk --no-cache add ca-certificates bash

ARG TARGETOS
ARG TARGETARCH

ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider

ENV USER_ID=65532

# This is a no-fork provider: the Cisco IOS-XE Terraform provider is linked into
# the controller binary and invoked in-process via the Terraform Plugin
# Framework. Neither the Terraform CLI nor the provider plugin binary is
# shipped in this image. The variables below are only recorded in the Terraform
# setup used by Upjet (provider source/version metadata).
ARG TERRAFORM_VERSION
ARG TERRAFORM_PROVIDER_SOURCE
ARG TERRAFORM_PROVIDER_VERSION

ENV TERRAFORM_VERSION=${TERRAFORM_VERSION}
ENV TERRAFORM_PROVIDER_SOURCE=${TERRAFORM_PROVIDER_SOURCE}
ENV TERRAFORM_PROVIDER_VERSION=${TERRAFORM_PROVIDER_VERSION}

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
