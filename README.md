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

* Copy the example of the appendix directory to your others.
```
cp -ar examples/hello ~/my_test1
cp -ar examples/hello ~/my_test2
...
```
* Make any changes to your directories.
* Build your applications.
```
make APPENDIX_DIRECTORY="~/my_test1" BUILD_FILE="~/my_test1.efi"
make APPENDIX_DIRECTORY="~/my_test2" BUILD_FILE="~/my_test2.efi"
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
* TARGET_PACKAGES - please see `build/package/Dockerfile` for understanding.
* ROOT_PASSWORD - empty password means disabled.
* APPENDIX_DIRECTORY - see the building section of this document.
* BUILD_FILE - see the building section of this document.

## Authors

* [Danil Kichai](https://github.com/DanilKichai)

## Version History

* v0.1.0
    * Initial release

## License

This project is licensed under the Boost Software License 1.0 - see the LICENSE file for details

## Acknowledgments

Inspiration, code snippets, etc.
* [DomPizzie/README-Template.md](https://gist.github.com/DomPizzie/7a5ff55ffa9081f2de27c315f5018afc)
* [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
