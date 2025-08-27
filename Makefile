TARGET_PACKAGES := " \
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
ROOT_PASSWORD := ""
APPENDIX_DIRECTORY := "examples/hello"
BUILD_FILE := "hello.efi"

.PHONY: all
all: $(BUILD_FILE)

.PHONY: builder
builder:
	docker buildx create \
		--name shoeshiner \
		--node main

$(BUILD_FILE): builder
	docker buildx build \
		--builder shoeshiner \
		--progress plain \
		--file build/package/Dockerfile \
		--target shoeshiner \
		--output "type=local,dest=." \
		--build-arg TARGET_PACKAGES=$(TARGET_PACKAGES) \
		--build-arg ROOT_PASSWORD=$(ROOT_PASSWORD) \
		--build-arg APPENDIX_DIRECTORY=$(APPENDIX_DIRECTORY) \
		--build-arg BUILD_FILE=$(BUILD_FILE) \
		.

.PHONY: clean
clean:
	-docker buildx rm \
		--force \
		--builder shoeshiner
