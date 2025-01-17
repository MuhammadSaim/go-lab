package main

type WeatherData struct {
	Data     Data     `json:"data"`
	Location Location `json:"location"`
}

type Data struct {
	Time   string `json:"time"`
	Values Values `json:"values"`
}

type Values struct {
	CloudBase                *float64 `json:"cloudBase"`
	CloudCeiling             *float64 `json:"cloudCeiling"`
	CloudCover               *int     `json:"cloudCover"`
	DewPoint                 *float64 `json:"dewPoint"`
	FreezingRainIntensity    *float64 `json:"freezingRainIntensity"`
	HailProbability          *float64 `json:"hailProbability"`
	HailSize                 *float64 `json:"hailSize"`
	Humidity                 *int     `json:"humidity"`
	PrecipitationProbability *int     `json:"precipitationProbability"`
	PressureSurfaceLevel     *float64 `json:"pressureSurfaceLevel"`
	RainIntensity            *float64 `json:"rainIntensity"`
	SleetIntensity           *float64 `json:"sleetIntensity"`
	SnowIntensity            *float64 `json:"snowIntensity"`
	Temperature              *float64 `json:"temperature"`
	TemperatureApparent      *float64 `json:"temperatureApparent"`
	UVHealthConcern          *int     `json:"uvHealthConcern"`
	UVIndex                  *int     `json:"uvIndex"`
	Visibility               *float64 `json:"visibility"`
	WeatherCode              *int     `json:"weatherCode"`
	WindDirection            *int     `json:"windDirection"`
	WindGust                 *float64 `json:"windGust"`
	WindSpeed                *float64 `json:"windSpeed"`
}

type Location struct {
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Name string  `json:"name"`
	Type string  `json:"type"`
}
