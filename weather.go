package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	host = "https://api.weather.yandex.ru/v2/forecast" // пути лучше выносить в константы, чтобы их можно было быстро поменять (при смене версии, например)
)


// Request метод API, обертка над функцией request
func (api *API) Request(cityName string) (*Response, error) {

	findCity, err := coordinates(cityName)
	if err != nil {
		return nil, err
	}

	return request(string(*api), findCity.Lat, findCity.Lon)
}

// это функция, реализующая сам запрос
func request(token, lat, lon string) (*Response, error) {
	param := make(url.Values)
	param.Add("lat", lat)
	param.Add("lon", lon)
	param.Add("lang", "ru_RU")
	param.Add("extra", "true")

	u, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	u.RawQuery = param.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Yandex-API-Key", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := new(Response)

	if err := json.NewDecoder(resp.Body).Decode(r); err != nil {
		return nil, err
	}

	return r, nil
}

// вместо чтения файла при каждом запросе мы сохраняем все данные в map
func coordinates(name string) (*city, error) {
	find, ok := Cities[name]
	if !ok {
		return nil, fmt.Errorf("%s не найден", name)
	}

	return &find, nil
}