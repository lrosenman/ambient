# ambient
[![License](https://img.shields.io/badge/License-BSD%202--Clause-orange.svg)](https://opensource.org/licenses/BSD-2-Clause)
[![GoDoc](https://godoc.org/github.com/lrosenman/ambient?status.svg)](https://godoc.org/github.com/lrosenman/ambient)
[![continuous-integration](https://github.com/lrosenman/ambient/actions/workflows/ci-build.yml/badge.svg)](https://github.com/lrosenman/ambient/actions/workflows/ci-build.yml)

ambient is a Go client library for accessing the [Ambient Weather API](https://ambientweather.docs.apiary.io/).

Currently, ambient requires Go version 1.19 or greater.  We do our best not to break older versions of Go if we don't have to, but due to tooling constraints, we don't always test older versions.

## Installation

ambient is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/lrosenman/ambient
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/lrosenman/ambient"
```

and run `go get` without parameters.

## Usage

### List Devices
Lists all devices (weather stations) that are registered for a given api key
```go
key := ambient.NewKey("... your application key ...", "... you api key ...")
devices, err := ambient.Device(key)
```

### Query Device Data
Queries a specific device for its last 10 observations
```go
key := ambient.NewKey("... your application key ...", "... you api key ...")
queryResults, err := ambient.DeviceMac(key, "... device mac address ...", time.Now().UTC(), 10)
```

More examples of how to use this library can be found in the [examples](/examples) directory

| Name                                                     | Purpose                                                                                                                   |
|----------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------|
| [list-devices](/examples/list-devices/main.go)           | Lists all devices (weather stations) that are registered for the account the application and api keys are associated with |
| [query-device](/examples/query-device/main.go)           | Queries a specific device for its observations                                                                            |
| [query-all-devices](/examples/query-all-devices/main.go) | Queries all registered devices for an account for their observations                                                      |
| [print-api](/examples/print-api/main.go)                 | Shows all API calls and the responses to them                                                                             |

## Authentication
The Ambient Weather API uses an application key that identifies a specific application and an api key that grants access to a specific user's devices.  See [Ambient API Authentication documentation](https://ambientweather.docs.apiary.io/#introduction/authentication) for more details on these values and how to generate / manage.

This is represented in the ambient library by the ```Key``` struct and is used for any calls to the API:

```go
key := ambient.NewKey("... your application key ...", "... you api key ...")
devices, err := ambient.Device(key)
```

## Rate Limiting
Ambient Weather API requests are [capped](https://ambientweather.docs.apiary.io/#introduction/rate-limiting) at 1 request per second for each user's apiKey and 3 requests per second for a given applicationKey. When this limit is exceeded, the API will return a 429 response code.
To determine if any of your requests have been rate limited, the ```HTTPResponseCode``` field has been added to response structs to determine the nature of a failed API call.

## Contributing
We would like to cover the entire Ambient Weather API and contributions are of course always welcome.  See [`CONTRIBUTING.md`](CONTRIBUTING.md) for details.

## Versioning
In general, ambient follows [semver](https://semver.org/) as closely as we
can for tagging releases of the package.

## License
This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE)
file.