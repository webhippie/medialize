# Medialize

[![Build Status](http://github.dronehippie.de/api/badges/webhippie/medialize/status.svg)](http://github.dronehippie.de/webhippie/medialize)
[![Coverage Status](http://coverage.dronehippie.de/badges/webhippie/medialize/coverage.svg)](http://coverage.dronehippie.de/webhippie/medialize)
[![Go Doc](https://godoc.org/github.com/webhippie/medialize?status.svg)](http://godoc.org/github.com/webhippie/medialize)
[![Go Report](http://goreportcard.com/badge/webhippie/medialize)](http://goreportcard.com/report/webhippie/medialize)

Medialize is used to sort images and videos by their meta information into a
specific folder structure separated by creation year and month. Beside that it's
planned to integrate a functionality to find and remove duplicated photos or
videos.


## Install

You can download prebuilt binaries from the GitHub releases or from our
[download site](http://dl.webhippie.de/medialize). You are a Mac user? Just take
a look at our [homebrew formula](https://github.com/webhippie/homebrew-webhippie).
Take a look at the help output, you can enable auto updates to the binary to
avoid bugs related to old versions. If you find a security issue please contact
thomas@webhippie.de first.


## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions](http://golang.org/doc/install.html).
As this project relies on vendoring of the dependencies and we are not
exporting `GO15VENDOREXPERIMENT=1` within our makefile you have to use a Go
version `>= 1.6`

```bash
go get -d github.com/webhippie/medialize
cd $GOPATH/src/github.com/webhippie/medialize
make deps build

bin/medialize -h
```


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2015-2016 Thomas Boerger <http://www.webhippie.de>
```
