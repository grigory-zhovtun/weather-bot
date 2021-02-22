package main

type API string

// NewAPI возвращает обернутый в тип API токен
func NewAPI(token string) *API {
	api := API(token)
	return &api
}

// Для поиска координат по названию города используем map
type city struct {
	Name string
	Lat  string
	Lon  string
}

// Документацию см. https://yandex.ru/dev/weather/doc/dg/concepts/forecast-test.html#resp-format__info

type Response struct {
	Now  string `json:"now_dt"`
	Fact Facts  `json:"fact"`
	Info Infos  `json:"info"`
}

type Facts struct {
	Temp      int     `json:"temp"`
	Feels     int     `json:"feels_like"`
	Icon      string  `json:"icon"`
	Condition string  `json:"condition"`
	WindSpeed float64 `json:"wind_speed"`
	WindDir   string  `json:"wind_dir"`
	PrecType  int     `json:"prec_type"`
}

type Infos struct {
	Tzinfos Tzinfo `json:"tzinfo"`
}

type Tzinfo struct {
	Name string `json:"name"`
	Abbr string `json:"abbr"`
}