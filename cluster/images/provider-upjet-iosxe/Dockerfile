FROM gcr.io/distroless/static-debian13:nonroot@sha256:1c2c046bc09ed40fad370b599a0b1ae7987f55b01e247cf27a7c27cd97e5bbc7

ARG TARGETOS
ARG TARGETARCH

# This is a no-fork provider: the Cisco IOS-XE Terraform provider is linked
# into the controller binary and called in-process over the Terraform Plugin
# Framework. Neither the Terraform CLI nor the provider plugin binary ships
# here, so the image only needs the controller and a CA bundle.
ENV USER_ID=65532

USER ${USER_ID}

ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider

EXPOSE 8080

ENTRYPOINT ["provider"]
