# Medialize

[![Current Tag](https://img.shields.io/github/v/tag/webhippie/medialize?sort=semver)](https://github.com/webhippie/medialize) [![Build Status](https://github.com/webhippie/medialize/actions/workflows/general.yml/badge.svg)](https://github.com/webhippie/medialize/actions) [![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org) [![Docker Size](https://img.shields.io/docker/image-size/webhippie/medialize/latest)](https://hub.docker.com/r/webhippie/medialize) [![Docker Pulls](https://img.shields.io/docker/pulls/webhippie/medialize)](https://hub.docker.com/r/webhippie/medialize) [![Go Reference](https://pkg.go.dev/badge/github.com/webhippie/medialize.svg)](https://pkg.go.dev/github.com/webhippie/medialize) [![Go Report Card](https://goreportcard.com/badge/github.com/webhippie/medialize)](https://goreportcard.com/report/github.com/webhippie/medialize) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/667130ec21cf4c3eb45a7d798fe98322)](https://www.codacy.com/gh/webhippie/medialize/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=webhippie/medialize&amp;utm_campaign=Badge_Grade)

Medialize is used to sort images and videos by their meta information into a
specific folder structure separated by creation year and month. Beside that it's
planned to integrate a functionality to find and remove duplicated photos or
videos.

## Install

You can download prebuilt binaries from our [GitHub releases][releases], or you
can use our Docker images published on [Docker Hub][dockerhub] or [Quay][quay].
If you need further guidance how to install this take a look at our
[documentation][docs].

## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions][golang]. This project requires
Go >= v1.17, at least that's the version we are using.

```console
git clone https://github.com/webhippie/medialize.git
cd medialize

make generate build

./bin/medialize -h
```

## Security

If you find a security issue please contact
[thomas@webhippie.de](mailto:thomas@webhippie.de) first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```

[releases]: https://github.com/webhippie/medialize/releases
[dockerhub]: https://hub.docker.com/r/webhippie/medialize/tags/
[quay]: https://quay.io/repository/webhippie/medialize?tab=tags
[docs]: https://webhippie.github.io/medialize/#getting-started
[golang]: http://golang.org/doc/install.html
