# syntax=docker/dockerfile:1

ARG BOOT_DIRECTORY="boots/hello"
ARG BUILD_FILE="hello.efi"
ARG MIRRORLIST_FILE="hack/etc/pacman.d/mirrorlist"

FROM archlinux:base AS archlinux
  ARG MIRRORLIST_FILE
  COPY $MIRRORLIST_FILE /etc/pacman.d/mirrorlist.candidate
  RUN \
    if \
      grep \
        --quiet \
        --invert-match \
        --extended-regexp '^[[:space:]]*([#]|$)' \
        /etc/pacman.d/mirrorlist.candidate; \
    then \
      mv /etc/pacman.d/mirrorlist.candidate /etc/pacman.d/mirrorlist; \
    else \
      rm --force /etc/pacman.d/mirrorlist.candidate; \
    fi

  RUN sed --in-place 's/^#DisableSandbox/DisableSandbox/' /etc/pacman.conf

  RUN \
    rm \
      --recursive \
      --force \
      /etc/pacman.d/gnupg && \
    pacman-key --init && \
    pacman-key --populate archlinux && \
    pacman \
      --sync \
      --refresh \
      --noconfirm \
      archlinux-keyring

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
