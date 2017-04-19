# Medialize

[![Build Status](http://github.dronehippie.de/api/badges/webhippie/medialize/status.svg)](http://github.dronehippie.de/webhippie/medialize)
[![Go Doc](https://godoc.org/github.com/webhippie/medialize?status.svg)](http://godoc.org/github.com/webhippie/medialize)
[![Go Report](http://goreportcard.com/badge/github.com/webhippie/medialize)](http://goreportcard.com/report/github.com/webhippie/medialize)
[![](https://images.microbadger.com/badges/image/tboerger/medialize.svg)](http://microbadger.com/images/tboerger/medialize "Get your own image badge on microbadger.com")
[![Join the chat at https://gitter.im/webhippie/general](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/webhippie/general)
[![Stories in Ready](https://badge.waffle.io/webhippie/medialize.svg?label=ready&title=Ready)](http://waffle.io/webhippie/medialize)

Medialize is used to sort images and videos by their meta information into a specific folder structure separated by creation year and month. Beside that it's planned to integrate a functionality to find and remove duplicated photos or videos.


## Install

You can download prebuilt binaries from the GitHub releases or from our [download site](http://dl.webhippie.de/misc/medialize). You are a Mac user? Just take a look at our [homebrew formula](https://github.com/webhippie/homebrew-webhippie). If you are missing an architecture just write us on our nice [Gitter](https://gitter.im/webhippie/general) chat. Take a look at the help output, per default we enabled an auto updater for the binary to avoid bugs related to old versions. If you find a security issue please contact thomas@webhippie.de first.


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). As this project relies on vendoring of the dependencies and we are not exporting `GO15VENDOREXPERIMENT=1` within our makefile you have to use a Go version `>= 1.6`. It is also possible to just simply execute the `go get github.com/webhippie/medialize` command, but we prefer to use our `Makefile`:

```bash
go get -d github.com/webhippie/medialize
cd $GOPATH/src/github.com/webhippie/medialize
make clean build

./medialize -h
```


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2015 Thomas Boerger <http://www.webhippie.de>
```
