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

FROM shoeshiner:bootstrap AS bootstrap

FROM shoeshiner:archlinux
  ARG TARGET_PACKAGES
  RUN \
    pacman \
      --sync \
      --noconfirm \
      --needed \
      $TARGET_PACKAGES
  RUN \
    rm \
      --recursive \
      --force \
      /var/cache

  ARG ROOT_PASSWORD
  RUN \
    if [ -n "$ROOT_PASSWORD" ]; then \
      ( \
        echo "$ROOT_PASSWORD"; \
        echo "$ROOT_PASSWORD"; \
      ) | passwd; \
    fi

  RUN mkdir --parents /shoeshiner/efivars
  COPY --from=bootstrap /bootstrap /shoeshiner/bin/

  COPY . /
