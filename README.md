# ambient
[![License](https://img.shields.io/badge/License-BSD%202--Clause-orange.svg)](https://opensource.org/licenses/BSD-2-Clause)

AmbientWeather.net API Helper

Example program in https://github.com/lrosenman/ambient/example

Official Doc:

https://ambientweather.docs.apiary.io/

```
package ambient // import "github.com/lrosenman/ambient"

Package ambient provides helper functions and Go types for accessing
ambientweather.net's API which is documented at

    https://ambientweather.docs.apiary.io/

const ApiEP = "https://api.ambientweather.net/" + ApiVer
const ApiVer = "v1"
type AmbientRecord struct{ ... }
type ApiDeviceMacResponse struct{ ... }
    func DeviceMac(key Key, macaddr string, endtime time.Time, limit int64) (ApiDeviceMacResponse, error)
type ApiDeviceResponse struct{ ... }
    func Device(key Key) (ApiDeviceResponse, error)
type DeviceInfo struct{ ... }
type DeviceRecord struct{ ... }
type Key struct{ ... }
    func NewKey(applicationKey string, apiKey string) Key
```
