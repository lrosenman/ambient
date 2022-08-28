# ambient
[![License](https://img.shields.io/badge/License-BSD%202--Clause-orange.svg)](https://opensource.org/licenses/BSD-2-Clause)
[![GoDoc](https://godoc.org/github.com/lrosenman/ambient?status.svg)](https://godoc.org/github.com/lrosenman/ambient)

AmbientWeather.net API Helper

Example program in https://github.com/lrosenman/ambient/tree/master/example

Official Doc:

https://ambientweather.docs.apiary.io/

Pull Requests, bug reports, code improvements, etc. very much welcome.


```
package ambient // import "github.com/lrosenman/ambient"

Package ambient provides helper functions and Go types for accessing
ambientweather.net's API which is documented at

    https://ambientweather.docs.apiary.io/

const APIEP = "https://api.ambientweather.net/" + APIVer
const APIVer = "v1"
type APIDeviceMacResponse struct{ ... }
    func DeviceMac(key Key, macaddr string, endtime time.Time, limit int64) (APIDeviceMacResponse, error)
type APIDeviceResponse struct{ ... }
    func Device(key Key) (APIDeviceResponse, error)
type DeviceInfo struct{ ... }
type DeviceRecord struct{ ... }
type Key struct{ ... }
    func NewKey(applicationKey string, apiKey string) Key
type Record struct{ ... }
```
