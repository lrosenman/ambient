// Ambient Weather API stuff
package ambient

import (
	"time"
)

const ApiVer = "v1"
const ApiEP = "https://api.ambientweather.net/" + ApiVer

type AmbientRecord struct {
	Date           time.Time
	Baromabsin     float64
	Baromrelin     float64
	Dailyrainin    float64
	Dewpoint       float64
	Feelslike      float64
	Hourlyrainin   float64
	Humidity       int
	Humidityin     int
	LastRain       time.Time
	Maxdailygust   float64
	Monthlyrainin  float64
	Solarradiation float64
	Tempf          float64
	Tempinf        float64
	Uv             int
	Weeklyrainin   float64
	Winddir        int
	Windgustmph    float64
	Windspeedmph   float64
	Yearlyrainin   float64
}

type DeviceInfo struct {
	Name     string
	Location string
}

type DeviceRecord struct {
	Macaddr    string
	DeviceInfo DeviceInfo
	LastData   AmbientRecord
}

type ApiDeviceMacResponse struct {
	Ar               []AmbientRecord
	JSONResponse     string
	HTTPResponseCode int
}

type ApiDeviceResponse struct {
	DeviceRecord     []DeviceRecord
	JSONResponse     string
	HTTPResponseCode int
}
