# syntax=docker/dockerfile:1
# check=skip=SecretsUsedInArgOrEnv

ARG TARGET_PACKAGES=" \
	core/bash \
	core/curl \
	core/coreutils \
	core/cryptsetup \
	core/dosfstools \
	core/e2fsprogs \
	core/efibootmgr \
	core/gawk \
	core/grep \
	core/gzip \
	core/iproute2 \
	core/iputils \
	core/less \
  core/libpcap \
	core/lvm2 \
	core/nano \
	core/procps-ng \
	core/sed \
	core/systemd \
	core/udev \
	core/util-linux \
	extra/kexec-tools \
	extra/polkit \
	extra/tcpdump \
	extra/traceroute \
"
ARG ROOT_PASSWORD=""
ARG APPENDIX_DIRECTORY="boots/hello"
ARG BUILD_FILE="hello.efi"

FROM archlinux:base AS archlinux-fresh
  #COPY hack/etc/pacman.d/mirrorlist /etc/pacman.d/
  RUN \
    pacman \
      --sync \
      --sysupgrade \
      --refresh \
      --noconfirm

FROM archlinux-fresh AS go-builder
  RUN \
    pacman \
      --sync \
      --noconfirm \
      --needed \
      core/base-devel \
      core/libpcap \
      extra/go

  WORKDIR /usr/src/shoeshiner
  COPY . .

  ENV CGO_ENABLED=1
  RUN \
    go build \
      -o /usr/local/bin/bootstrap \
      ./cmd/bootstrap
  RUN \
    go build \
      -o /usr/local/bin/inspector \
      ./cmd/inspector

FROM archlinux-fresh AS target-base
  ARG TARGET_PACKAGES
  RUN \
    pacman \
      --sync \
      --noconfirm \
      --needed \
      $TARGET_PACKAGES

  WORKDIR /target
  RUN \
    tar \
      --create \
      --file - \
      --one-file-system \
      --exclude=target \
      --exclude=var/cache \
      --exclude=run/* \
      / | tar --extract --file -

FROM archlinux-fresh AS toolchain
  RUN \
    pacman \
      --sync \
      --noconfirm \
      --needed \
      core/linux

  COPY build/package/toolchain/ /

FROM toolchain AS wrapper
  COPY --from=target-base /target/ /target/
  COPY --from=go-builder /usr/local/bin/bootstrap /target/usr/local/bin/
  COPY --from=go-builder /usr/local/bin/inspector /target/usr/local/bin/

  ARG APPENDIX_DIRECTORY
  COPY $APPENDIX_DIRECTORY/ /target/

  ARG ROOT_PASSWORD
  RUN \
    if [ -n "$ROOT_PASSWORD" ]; then \
      ( \
        echo "$ROOT_PASSWORD"; \
        echo "$ROOT_PASSWORD"; \
      ) | passwd --root /target; \
    fi

  RUN mkinitcpio --preset wrapper

FROM scratch AS shoeshiner
  ARG BUILD_FILE
  COPY --from=wrapper /boot/wrapper.efi $BUILD_FILE
