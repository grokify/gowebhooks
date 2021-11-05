# Webhook Directory

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

This repo contains a set information on services with webhooks. The primary use of this repo is to provide a set of example webhook event payloads.

Examples are assembed from a variety of sources including service documentation andfrom actual live events.

Some clean up is implemented. All examples are valid HTTP bodies which may require editing. For example, some docuemntation examples have comments which provides nice documentation but results in non-validating JSON.

This project was extracted from the [Chathooks](https://github.com/grokify/chathooks) formatting webhook proxy service.

 [build-status-svg]: https://github.com/grokify/gowebhooks/workflows/go%20build/badge.svg
 [build-status-url]: https://github.com/grokify/gowebhooks/actions
 [coverage-status-svg]: https://coveralls.io/repos/grokify/gowebhooks/badge.svg?branch=master
 [coverage-status-url]: https://coveralls.io/r/grokify/gowebhooks?branch=master
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/gowebhooks
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/gowebhooks
 [codeclimate-status-svg]: https://codeclimate.com/github/grokify/gowebhooks/badges/gpa.svg
 [codeclimate-status-url]: https://codeclimate.com/github/grokify/gowebhooks
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/gowebhooks
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/gowebhooks
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/gowebhooks/blob/master/LICENSE
