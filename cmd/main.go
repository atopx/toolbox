package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CaiYunApi struct {
	url string
}

type CaiYunData struct {
	Status     string    `json:"status"`
	ApiVersion string    `json:"api_version"`
	ApiStatus  string    `json:"api_status"`
	Lang       string    `json:"lang"`
	Unit       string    `json:"unit"`
	Tzshift    int       `json:"tzshift"`
	Timezone   string    `json:"timezone"`
	ServerTime int       `json:"server_time"`
	Location   []float64 `json:"location"`
	Result     struct {
		Hourly struct {
			Status        string `json:"status"`
			Description   string `json:"description"`
			Precipitation []struct {
				Datetime    string `json:"datetime"`
				Value       int    `json:"value"`
				Probability int    `json:"probability"`
			} `json:"precipitation"`
			Temperature []struct {
				Datetime string `json:"datetime"`
				Value    int    `json:"value"`
			} `json:"temperature"`
			ApparentTemperature []struct {
				Datetime string  `json:"datetime"`
				Value    float64 `json:"value"`
			} `json:"apparent_temperature"`
			Wind []struct {
				Datetime  string `json:"datetime"`
				Speed     int    `json:"speed"`
				Direction int    `json:"direction"`
			} `json:"wind"`
			Humidity []struct {
				Datetime string  `json:"datetime"`
				Value    float64 `json:"value"`
			} `json:"humidity"`
			Cloudrate []struct {
				Datetime string `json:"datetime"`
				Value    int    `json:"value"`
			} `json:"cloudrate"`
			Skycon []struct {
				Datetime string `json:"datetime"`
				Value    string `json:"value"`
			} `json:"skycon"`
			Pressure []struct {
				Datetime string  `json:"datetime"`
				Value    float64 `json:"value"`
			} `json:"pressure"`
			Visibility []struct {
				Datetime string `json:"datetime"`
				Value    int    `json:"value"`
			} `json:"visibility"`
			Dswrf []struct {
				Datetime string  `json:"datetime"`
				Value    float64 `json:"value"`
			} `json:"dswrf"`
			AirQuality struct {
				Aqi []struct {
					Datetime string `json:"datetime"`
					Value    struct {
						Chn int `json:"chn"`
						Usa int `json:"usa"`
					} `json:"value"`
				} `json:"aqi"`
				Pm25 []struct {
					Datetime string `json:"datetime"`
					Value    int    `json:"value"`
				} `json:"pm25"`
			} `json:"air_quality"`
		} `json:"hourly"`
		Primary          int    `json:"primary"`
		ForecastKeypoint string `json:"forecast_keypoint"`
	} `json:"result"`
}

// 116.493088,39.989839
const url = "http://api.caiyunapp.com/v2.6/d7NtAjbCooopUEyf/116.493088,39.989839/hourly"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var data CaiYunData
	_ = json.Unmarshal(body, &data)
	fmt.Println(data)
}
