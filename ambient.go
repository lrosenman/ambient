// SPDX-License-Identifier: BSD-2-Clause
//
// Copyright 2018 Larry Rosenman, LERCTR Consulting, larryrtx@gmail.com
//

// Package ambient provides helper functions and Go types
// for accessing ambientweather.net's API which is documented at
//      https://ambientweather.docs.apiary.io/
package ambient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

// APIVer is the current version of the API.
const APIVer = "v1"

// APIEP is the endpoint to be called.
const APIEP = "https://api.ambientweather.net/" + APIVer

// Record maps the data for a specific time
// as returned by the API.
type Record struct {
	Date              time.Time
	Baromabsin        float64
	Baromrelin        float64
	Dailyrainin       float64
	Dewpoint          float64
	Feelslike         float64
	Hourlyrainin      float64
	Humidity          int
	Humidityin        int
	// Note: LastRain only comes in with
	// the Device API call, not the DeviceMac 
	// API
	LastRain          time.Time
	Maxdailygust      float64
	Monthlyrainin     float64
	Solarradiation    float64
	Tempf             float64
	Tempinf           float64
	Uv                int
	Weeklyrainin      float64
	Winddir           int
	Windgustmph       float64
	Windspeedmph      float64
	Winddir_avg2m     int
	Windspdmph_avg2m  float64
	Winddir_avg10m    int
	Windspdmph_avg10m float64
	Yearlyrainin      float64
}

// DeviceInfo maps the info portion of the /devices API.
type DeviceInfo struct {
	Name     string
	Location string
}

//DeviceRecord maps one record of the /devices API.
type DeviceRecord struct {
	Macaddress string
	Info       DeviceInfo
	LastData   Record
}

// APIDeviceMacResponse returns the data from
// /devices/macaddr API.
type APIDeviceMacResponse struct {
	Record           []Record
	JSONResponse     []byte
	HTTPResponseCode int
	ResponseTime     time.Duration
}

// APIDeviceResponse returns the data from
// /devices API.
type APIDeviceResponse struct {
	DeviceRecord     []DeviceRecord
	JSONResponse     []byte
	HTTPResponseCode int
	ResponseTime     time.Duration
}

// Key holds the keys.
type Key struct {
	applicationKey string
	apiKey         string
}

// NewKey returns Key stucture to be used.
func NewKey(applicationKey string, apiKey string) Key {
	return Key{applicationKey: applicationKey, apiKey: apiKey}
}

// APIKey returns the currently set APIKey.
func (Key Key) APIKey() string {
	return Key.apiKey
}

// ApplicationKey returns the currently set ApplicationKey.
func (Key Key) ApplicationKey() string {
	return Key.applicationKey
}

// SetApplicationKey sets the applicationKey
func (Key Key) SetApplicationKey(applicationKey string) {
	Key.applicationKey = applicationKey
}

// SetAPIKey sets the APIKey
func (Key Key) SetAPIKey(apiKey string) {
	Key.apiKey = apiKey
}

// Device issues a /devices call.
func Device(key Key) (APIDeviceResponse, error) {
	var ar APIDeviceResponse

	url := APIEP + "/devices?applicationKey=" + key.applicationKey +
		"&apiKey=" + key.apiKey
	startTime := time.Now()
	resp, err := http.Get(url)
	ar.ResponseTime = time.Since(startTime)
	if err != nil {
		return ar, err
	}
	ar.HTTPResponseCode = resp.StatusCode
	switch resp.StatusCode {
	case 200:
	case 503, 429:
		{
			return ar, nil
		}
	default:
		{
			fmt.Fprintf(os.Stderr, "ambient.Device: HTTPResponseCode=%d\nFull Response:\n%+v",
				resp.StatusCode, resp)
			return ar, errors.New("Bad non-200/429/503 Response Code")
		}
	}
	ar.JSONResponse, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return ar, err
	}
	err = json.Unmarshal(ar.JSONResponse, &ar.DeviceRecord)
	if err != nil {
		return ar, err
	}
	return ar, nil

}

// DeviceMac issues a /devices/macaddr call.
func DeviceMac(key Key, macaddr string, endtime time.Time, limit int64) (APIDeviceMacResponse, error) {
	var ar APIDeviceMacResponse
	url := APIEP + "/devices/" + macaddr + "?endDate=" + url.QueryEscape(endtime.Format(time.RFC3339)) +
		"&limit=" + fmt.Sprintf("%d", limit) + "&applicationKey=" + key.applicationKey +
		"&apiKey=" + key.apiKey
	startTime := time.Now()
	resp, err := http.Get(url)
	ar.ResponseTime = time.Since(startTime)
	if err != nil {
		return ar, err
	}
	ar.HTTPResponseCode = resp.StatusCode
	switch resp.StatusCode {
	case 200:
	case 503, 429:
		{
			return ar, nil
		}
	default:
		{
			fmt.Fprintf(os.Stderr,
				"ambient.DeviceMac: HTTPResponseCode=%d\n"+
					"Full Response:\n%+v",
				resp.StatusCode, resp)
			return ar, errors.New("Bad non-200/429/503 Response Code")
		}
	}
	ar.JSONResponse, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return ar, err
	}
	err = json.Unmarshal(ar.JSONResponse, &ar.Record)
	if err != nil {
		return ar, err
	}
	return ar, nil
}
