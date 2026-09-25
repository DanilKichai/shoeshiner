# Shoeshiner

Simple Linux wrapper for any purpose.

## Description

Shoeshiner or boot polisher is an occupation in which a person cleans and buffs shoes and then applies a waxy paste to give a shiny appearance and a protective coating. With this tool you can "polish" your own set of scripts and applications to get your own UEFI application.

## Getting Started

### Dependencies

* GNU make utility.
* Docker with buildx plugin.
* Internet connection.

### Building

* Copy the example boot directory to your others.
```
cp -ar boots/hello boots/my_test1
```
* Make any changes to your boot configuration.
* Build your applications.
```
make \
    BOOT_DIRECTORY="boots/my_test1" \
    BUILD_FILE="test/esp/EFI/BOOT/BOOTX64.EFI"
```
* Clean build cache.
```
make clean
```

### Executing program

* Use the EFI file for its intended purpose.
```
qemu-system-x86_64 \
    -enable-kvm \
    -m 2G \
    -bios /usr/share/edk2/x64/OVMF.4m.fd \
    -drive format=raw,file=fat:rw:test/esp \
    -display gtk,gl=on
```

## Help

You can also use build options.
* BUILDX_BUILDER - The buildx plugin builder name (only docker driver is supported).
* BOOT_DIRECTORY - See the building section of this document.
* BOOT_BUILD_APPEND - Build append for boot (see the Makefile).
* BUILD_FILE - See the building section of this document.
* MIRRORLIST_FILE - The pacman mirrorlist (a.k.a. `hack/etc/pacman.d/mirrorlist`).

## Authors

* [Danil Kichai](https://github.com/DanilKichai)

## Version History
* v0.2.2
    * Made `MIRRORLIST_FILE` configurable via Makefile arguments.
    * Improved the `hello` example boot configuration.
    * Improved `README.md`.
* v0.2.1
    * Implemented general improvements and minor bug fixes.
    * Made `BOOT_BUILD_APPEND` configurable via Makefile arguments.
* v0.2.0
    * Moved the rootfs build to a separate Dockerfile (now owned by the `boot` directory).
    * Added rendering of `/shoeshiner/run/*.env` at runtime for the `hello` boot.
* v0.1.0
    * Initial release

## License

This project is licensed under the Boost Software License 1.0 - see the LICENSE file for details

## Acknowledgments

Inspiration, code snippets, etc.
* [DomPizzie/README-Template.md](https://gist.github.com/DomPizzie/7a5ff55ffa9081f2de27c315f5018afc)
* [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
