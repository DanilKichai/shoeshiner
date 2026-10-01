BUILDX_BUILDER := "default"
BOOT_DIRECTORY := "boots/hello"
BOOT_BUILD_APPEND :=
BUILD_FILE := "hello.efi"
MIRRORLIST_FILE := "hack/etc/pacman.d/mirrorlist"

.PHONY: all
all: "$(BUILD_FILE)"

.PHONY: archlinux
archlinux:
	docker buildx build \
		--builder "$(BUILDX_BUILDER)" \
		--progress plain \
		--file build/package/Dockerfile \
		--target archlinux \
		--output "type=docker" \
		--tag shoeshiner:archlinux \
		--build-arg MIRRORLIST_FILE="$(MIRRORLIST_FILE)" \
		.

.PHONY: bootstrap
bootstrap: archlinux
	docker buildx build \
		--builder "$(BUILDX_BUILDER)" \
		--progress plain \
		--file build/package/Dockerfile \
		--target bootstrap \
		--output "type=docker" \
		--tag shoeshiner:bootstrap \
		.

.PHONY: boot
boot: archlinux bootstrap
	docker buildx build \
		--builder "$(BUILDX_BUILDER)" \
		--progress plain \
		--file "$(BOOT_DIRECTORY)/Dockerfile" \
		--output "type=docker" \
		--tag shoeshiner:boot \
		$(BOOT_BUILD_APPEND) \
		"$(BOOT_DIRECTORY)"

"$(BUILD_FILE)": archlinux boot
	docker buildx build \
		--builder "$(BUILDX_BUILDER)" \
		--progress plain \
		--file build/package/Dockerfile \
		--target shoeshiner \
		--output "type=local,dest=." \
		--build-arg BOOT_DIRECTORY="$(BOOT_DIRECTORY)" \
		--build-arg BUILD_FILE="$(BUILD_FILE)" \
		.

.PHONY: clean
clean:
	docker rmi \
		shoeshiner:archlinux \
		shoeshiner:bootstrap \
		shoeshiner:boot \
		--force
