package ambient

import (
	"encoding/json"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func getValidDeviceRecord() *DeviceRecord {
	result := &DeviceRecord{
		Macaddress: faker.MacAddress(),
		Info:       DeviceInfo{},
		LastData:   Record{},
		LastDataFields: map[string]interface{}{
			"one": float64(1),
			"two": "two",
		},
	}

	faker.FakeData(&result.Info)
	faker.FakeData(&result.LastData)
	overrideFakeBatteryData(&result.LastData)

	return result
}

func getValidRecordSlice(count int) []*Record {
	result := make([]*Record, 0)

	for i := 0; i < count; i++ {
		data := &Record{}
		faker.FakeData(data)
		overrideFakeBatteryData(data)
		result = append(result, data)
	}

	return result
}

func overrideFakeBatteryData(data *Record) {
	data.Battin = "1"
	data.Battout = "2"
	data.Batt_co2 = "3"
	data.Batt_lightning = "4"
	data.Batt1 = "5"
	data.Batt2 = "6"
	data.Batt3 = "7"
	data.Batt4 = "8"
	data.Batt5 = "9"
	data.Batt6 = "10"
	data.Batt7 = "11"
	data.Batt8 = "12"
	data.Batt9 = "13"
	data.Batt10 = "14"
}

func getMockHttpGetRequest(statusCode int, responseData interface{}) func(url string) (*http.Response, error) {
	return func(url string) (*http.Response, error) {
		response := &http.Response{}
		response.StatusCode = statusCode
		response.Body = io.NopCloser(strings.NewReader(mapToJson(responseData)))

		return response, nil
	}
}

func mapToJson(toMap interface{}) string {
	result, err := json.Marshal(toMap)
	if err != nil {
		return "{}"
	}

	return string(result)
}

func requireDeviceRecordsEqualValues(t *testing.T, expected []*DeviceRecord, actual []DeviceRecord) {
	require.Len(t, actual, len(expected))

	actualByMacAddress := deviceRecordByMacAddress(actual)

	for _, expectedRecord := range expected {
		actualRecord := actualByMacAddress[expectedRecord.Macaddress]
		requireDeviceRecordEqualValues(t, actualRecord, expectedRecord)
	}
}

func requireDeviceRecordEqualValues(t *testing.T, actualRecord DeviceRecord, expectedRecord *DeviceRecord) {
	require.NotNil(t, actualRecord)
	require.Equal(t, expectedRecord.Macaddress, actualRecord.Macaddress)
	require.Equal(t, expectedRecord.Info.Name, actualRecord.Info.Name)
	require.Equal(t, expectedRecord.Info.Location, actualRecord.Info.Location)
	require.EqualValues(t, expectedRecord.LastDataFields, actualRecord.LastDataFields)
	requireRecordEqualValues(t, expectedRecord.LastData, actualRecord.LastData)
}

func deviceRecordByMacAddress(data []DeviceRecord) map[string]DeviceRecord {
	result := make(map[string]DeviceRecord)

	for _, item := range data {
		result[item.Macaddress] = item
	}

	return result
}

func requireRecordsEqualValues(t *testing.T, expected []*Record, actual []Record) {
	require.Equal(t, len(expected), len(actual))

	actualByDate := recordByDate(actual)

	for _, expectedRecord := range expected {
		actualRecord := actualByDate[expectedRecord.Date.UTC()]

		require.NotNil(t, actualRecord)
		requireRecordEqualValues(t, *expectedRecord, actualRecord)
	}
}

func recordByDate(data []Record) map[time.Time]Record {
	result := make(map[time.Time]Record)

	for _, item := range data {
		result[item.Date.UTC()] = item
	}

	return result
}

func requireRecordEqualValues(t *testing.T, expected Record, actual Record) {
	require.EqualValues(t, expected.Date.UTC(), actual.Date.UTC())
	require.EqualValues(t, expected.Baromabsin, actual.Baromabsin)
	require.EqualValues(t, expected.Baromrelin, actual.Baromrelin)
	require.EqualValues(t, toFloat64(expected.Battin), toFloat64(actual.Battin))
	require.EqualValues(t, toFloat64(expected.Battout), toFloat64(actual.Battout))
	require.EqualValues(t, toFloat64(expected.Batt1), toFloat64(actual.Batt1))
	require.EqualValues(t, toFloat64(expected.Batt2), toFloat64(actual.Batt2))
	require.EqualValues(t, toFloat64(expected.Batt3), toFloat64(actual.Batt3))
	require.EqualValues(t, toFloat64(expected.Batt4), toFloat64(actual.Batt4))
	require.EqualValues(t, toFloat64(expected.Batt5), toFloat64(actual.Batt5))
	require.EqualValues(t, toFloat64(expected.Batt6), toFloat64(actual.Batt6))
	require.EqualValues(t, toFloat64(expected.Batt7), toFloat64(actual.Batt7))
	require.EqualValues(t, toFloat64(expected.Batt8), toFloat64(actual.Batt8))
	require.EqualValues(t, toFloat64(expected.Batt9), toFloat64(actual.Batt9))
	require.EqualValues(t, toFloat64(expected.Batt10), toFloat64(actual.Batt10))
	require.EqualValues(t, toFloat64(expected.Batt_co2), toFloat64(actual.Batt_co2))
	require.EqualValues(t, toFloat64(expected.Batt_lightning), toFloat64(actual.Batt_lightning))
	require.EqualValues(t, expected.Co2, actual.Co2)
	require.EqualValues(t, expected.Dailyrainin, actual.Dailyrainin)
	require.EqualValues(t, expected.Dewpoint, actual.Dewpoint)
	require.EqualValues(t, expected.Dewpoint1, actual.Dewpoint1)
	require.EqualValues(t, expected.Dewpoint2, actual.Dewpoint2)
	require.EqualValues(t, expected.Dewpoint3, actual.Dewpoint3)
	require.EqualValues(t, expected.Dewpoint4, actual.Dewpoint4)
	require.EqualValues(t, expected.Dewpoint5, actual.Dewpoint5)
	require.EqualValues(t, expected.Dewpoint6, actual.Dewpoint6)
	require.EqualValues(t, expected.Dewpoint7, actual.Dewpoint7)
	require.EqualValues(t, expected.Dewpoint8, actual.Dewpoint8)
	require.EqualValues(t, expected.Dewpoint9, actual.Dewpoint9)
	require.EqualValues(t, expected.Dewpoint10, actual.Dewpoint10)
	require.EqualValues(t, expected.Dewpointin, actual.Dewpointin)
	require.EqualValues(t, expected.Eventrainin, actual.Eventrainin)
	require.EqualValues(t, expected.Feelslike, actual.Feelslike)
	require.EqualValues(t, expected.Feelslike1, actual.Feelslike1)
	require.EqualValues(t, expected.Feelslike2, actual.Feelslike2)
	require.EqualValues(t, expected.Feelslike3, actual.Feelslike3)
	require.EqualValues(t, expected.Feelslike4, actual.Feelslike4)
	require.EqualValues(t, expected.Feelslike5, actual.Feelslike5)
	require.EqualValues(t, expected.Feelslike6, actual.Feelslike6)
	require.EqualValues(t, expected.Feelslike7, actual.Feelslike7)
	require.EqualValues(t, expected.Feelslike8, actual.Feelslike8)
	require.EqualValues(t, expected.Feelslike9, actual.Feelslike9)
	require.EqualValues(t, expected.Feelslike10, actual.Feelslike10)
	require.EqualValues(t, expected.Feelslikein, actual.Feelslikein)
	require.EqualValues(t, expected.Hourlyrainin, actual.Hourlyrainin)
	require.EqualValues(t, expected.Humidity, actual.Humidity)
	require.EqualValues(t, expected.Humidity1, actual.Humidity1)
	require.EqualValues(t, expected.Humidity2, actual.Humidity2)
	require.EqualValues(t, expected.Humidity3, actual.Humidity3)
	require.EqualValues(t, expected.Humidity4, actual.Humidity4)
	require.EqualValues(t, expected.Humidity5, actual.Humidity5)
	require.EqualValues(t, expected.Humidity6, actual.Humidity6)
	require.EqualValues(t, expected.Humidity7, actual.Humidity7)
	require.EqualValues(t, expected.Humidity8, actual.Humidity8)
	require.EqualValues(t, expected.Humidity9, actual.Humidity9)
	require.EqualValues(t, expected.Humidity10, actual.Humidity10)
	require.EqualValues(t, expected.Humidityin, actual.Humidityin)
	require.EqualValues(t, expected.LastRain.UTC(), actual.LastRain.UTC())
	require.EqualValues(t, expected.Maxdailygust, actual.Maxdailygust)
	require.EqualValues(t, expected.Lightning_day, actual.Lightning_day)
	require.EqualValues(t, expected.Lightning_distance, actual.Lightning_distance)
	require.EqualValues(t, expected.Lightning_hour, actual.Lightning_hour)
	require.EqualValues(t, expected.Lightning_time.UTC(), actual.Lightning_time.UTC())
	require.EqualValues(t, expected.Pm25, actual.Pm25)
	require.EqualValues(t, expected.Pm25_24h, actual.Pm25_24h)
	require.EqualValues(t, expected.Relay1, actual.Relay1)
	require.EqualValues(t, expected.Relay2, actual.Relay2)
	require.EqualValues(t, expected.Relay3, actual.Relay3)
	require.EqualValues(t, expected.Relay4, actual.Relay4)
	require.EqualValues(t, expected.Relay5, actual.Relay5)
	require.EqualValues(t, expected.Relay6, actual.Relay6)
	require.EqualValues(t, expected.Relay7, actual.Relay7)
	require.EqualValues(t, expected.Relay8, actual.Relay8)
	require.EqualValues(t, expected.Relay9, actual.Relay9)
	require.EqualValues(t, expected.Relay10, actual.Relay10)
	require.EqualValues(t, expected.Monthlyrainin, actual.Monthlyrainin)
	require.EqualValues(t, expected.Soiltemp1f, actual.Soiltemp1f)
	require.EqualValues(t, expected.Soiltemp2f, actual.Soiltemp2f)
	require.EqualValues(t, expected.Soiltemp3f, actual.Soiltemp3f)
	require.EqualValues(t, expected.Soiltemp4f, actual.Soiltemp4f)
	require.EqualValues(t, expected.Soiltemp5f, actual.Soiltemp5f)
	require.EqualValues(t, expected.Soiltemp6f, actual.Soiltemp6f)
	require.EqualValues(t, expected.Soiltemp7f, actual.Soiltemp7f)
	require.EqualValues(t, expected.Soiltemp8f, actual.Soiltemp8f)
	require.EqualValues(t, expected.Soiltemp9f, actual.Soiltemp9f)
	require.EqualValues(t, expected.Soiltemp10f, actual.Soiltemp10f)
	require.EqualValues(t, expected.Soilhum1, actual.Soilhum1)
	require.EqualValues(t, expected.Soilhum2, actual.Soilhum2)
	require.EqualValues(t, expected.Soilhum3, actual.Soilhum3)
	require.EqualValues(t, expected.Soilhum4, actual.Soilhum4)
	require.EqualValues(t, expected.Soilhum5, actual.Soilhum5)
	require.EqualValues(t, expected.Soilhum6, actual.Soilhum6)
	require.EqualValues(t, expected.Soilhum7, actual.Soilhum7)
	require.EqualValues(t, expected.Soilhum8, actual.Soilhum8)
	require.EqualValues(t, expected.Soilhum9, actual.Soilhum9)
	require.EqualValues(t, expected.Soilhum10, actual.Soilhum10)
	require.EqualValues(t, expected.Solarradiation, actual.Solarradiation)
	require.EqualValues(t, expected.Tempf, actual.Tempf)
	require.EqualValues(t, expected.Temp1f, actual.Temp1f)
	require.EqualValues(t, expected.Temp2f, actual.Temp2f)
	require.EqualValues(t, expected.Temp3f, actual.Temp3f)
	require.EqualValues(t, expected.Temp4f, actual.Temp4f)
	require.EqualValues(t, expected.Temp5f, actual.Temp5f)
	require.EqualValues(t, expected.Temp6f, actual.Temp6f)
	require.EqualValues(t, expected.Temp7f, actual.Temp7f)
	require.EqualValues(t, expected.Temp8f, actual.Temp8f)
	require.EqualValues(t, expected.Temp9f, actual.Temp9f)
	require.EqualValues(t, expected.Temp10f, actual.Temp10f)
	require.EqualValues(t, expected.Tempinf, actual.Tempinf)
	require.EqualValues(t, expected.Totalrainin, actual.Totalrainin)
	require.EqualValues(t, expected.Uv, actual.Uv)
	require.EqualValues(t, expected.Weeklyrainin, actual.Weeklyrainin)
	require.EqualValues(t, expected.Winddir, actual.Winddir)
	require.EqualValues(t, expected.Windgustmph, actual.Windgustmph)
	require.EqualValues(t, expected.Windgustdir, actual.Windgustdir)
	require.EqualValues(t, expected.Windspeedmph, actual.Windspeedmph)
	require.EqualValues(t, expected.Winddir_avg2m, actual.Winddir_avg2m)
	require.EqualValues(t, expected.Windspdmph_avg2m, actual.Windspdmph_avg2m)
	require.EqualValues(t, expected.Winddir_avg10m, actual.Winddir_avg10m)
	require.EqualValues(t, expected.Windspdmph_avg10m, actual.Windspdmph_avg10m)
	require.EqualValues(t, expected.Yearlyrainin, actual.Yearlyrainin)
	require.EqualValues(t, expected.TZ, actual.TZ)
	require.EqualValues(t, expected.Aqi_pm25_in, actual.Aqi_pm25_in)
	require.EqualValues(t, expected.Aqi_pm25_in_24h, actual.Aqi_pm25_in_24h)
}

func toFloat64(data json.Number) float64 {
	result, err := data.Float64()
	if err != nil {
		return float64(0)
	}

	return result
}
