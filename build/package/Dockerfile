# syntax=docker/dockerfile:1

ARG BOOT_DIRECTORY="boots/hello"
ARG BUILD_FILE="hello.efi"

FROM archlinux:base AS archlinux
  #COPY hack/etc/pacman.d/mirrorlist /etc/pacman.d/
  RUN \
    pacman \
      --sync \
      --sysupgrade \
      --refresh \
      --noconfirm

FROM archlinux AS bootstrap
  RUN \
    pacman \
      --sync \
      --noconfirm \
      --needed \
      core/base-devel \
      extra/go

  WORKDIR /usr/src/shoeshiner
  COPY . .

  ENV CGO_ENABLED=0
  RUN \
    go build \
      -o /bootstrap \
      ./cmd/bootstrap

FROM archlinux AS toolchain
  RUN \
    pacman \
      --sync \
      --noconfirm \
      --needed \
      core/linux

  COPY build/package/toolchain/ /

FROM shoeshiner:boot AS boot

FROM toolchain AS wrapper
  COPY --from=boot / /target/
  RUN mkinitcpio --preset wrapper

FROM scratch AS shoeshiner
  ARG BUILD_FILE
  COPY --from=wrapper /boot/wrapper.efi $BUILD_FILE
