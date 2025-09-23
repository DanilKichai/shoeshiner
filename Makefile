BUILDX_BUILDER := "default"
BOOT_DIRECTORY := "boots/hello"
BUILD_FILE := "hello.efi"

.PHONY: all
all: $(BUILD_FILE)

.PHONY: archlinux
archlinux:
	docker buildx build \
		--builder $(BUILDX_BUILDER) \
		--progress plain \
		--file build/package/Dockerfile \
		--target archlinux \
		--output "type=docker" \
		--tag shoeshiner:archlinux \
		.

.PHONY: bootstrap
bootstrap: archlinux
	docker buildx build \
		--builder $(BUILDX_BUILDER) \
		--progress plain \
		--file build/package/Dockerfile \
		--target bootstrap \
		--output "type=docker" \
		--tag shoeshiner:bootstrap \
		.

.PHONY: boot
boot: archlinux bootstrap
	docker buildx build \
		--builder default \
		--progress plain \
		--file $(BOOT_DIRECTORY)/Dockerfile \
		--output "type=docker" \
		--tag shoeshiner:boot \
		$(BOOT_DIRECTORY)

$(BUILD_FILE): archlinux boot
	docker buildx build \
		--builder default \
		--progress plain \
		--file build/package/Dockerfile \
		--target shoeshiner \
		--output "type=local,dest=." \
		--build-arg BOOT_DIRECTORY=$(BOOT_DIRECTORY) \
		--build-arg BUILD_FILE=$(BUILD_FILE) \
		.

.PHONY: clean
clean:
	docker rmi \
		shoeshiner:archlinux \
		shoeshiner:bootstrap \
		shoeshiner:boot \
		--force
