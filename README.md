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
cp -ar boots/hello ~/my_test2
...
```
* Make any changes to your boots.
* Build your applications.
```
make BOOT_DIRECTORY="boots/my_test1" BUILD_FILE="my_test1.efi"
make BOOT_DIRECTORY="~/my_test2" BUILD_FILE="~/my_test2.efi"
...
```
* Clean build cache.
```
make clean
```

### Executing program

* Use the EFI file for its intended purpose.

## Help

You can also use build options.
* BUILDX_BUILDER - buildx plugin builder name (only docker driver is supported).
* BOOT_DIRECTORY - see the building section of this document.
* BUILD_FILE - see the building section of this document.

## Authors

* [Danil Kichai](https://github.com/DanilKichai)

## Version History
* v0.2.0
    * Exposed rootfs build to separated Dockerfile (owned by boot directory now)
    * Added rendering of `/shoeshiner/run/*.env` at runtime for the hello boot
* v0.1.0
    * Initial release

## License

This project is licensed under the Boost Software License 1.0 - see the LICENSE file for details

## Acknowledgments

Inspiration, code snippets, etc.
* [DomPizzie/README-Template.md](https://gist.github.com/DomPizzie/7a5ff55ffa9081f2de27c315f5018afc)
* [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
