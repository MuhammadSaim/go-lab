package main

// WeatherResponse struct to map the API response
type WeatherResponse struct {
	Data struct {
		Time   string `json:"time"`
		Values struct {
			CloudBase                float64 `json:"cloudBase"`
			CloudCeiling             float64 `json:"cloudCeiling"`
			CloudCover               float64 `json:"cloudCover"`
			DewPoint                 float64 `json:"dewPoint"`
			FreezingRainIntensity    float64 `json:"freezingRainIntensity"`
			Humidity                 float64 `json:"humidity"`
			PrecipitationProbability float64 `json:"precipitationProbability"`
			PressureSurfaceLevel     float64 `json:"pressureSurfaceLevel"`
			RainIntensity            float64 `json:"rainIntensity"`
			SleetIntensity           float64 `json:"sleetIntensity"`
			SnowIntensity            float64 `json:"snowIntensity"`
			Temperature              float64 `json:"temperature"`
			TemperatureApparent      float64 `json:"temperatureApparent"`
			UvHealthConcern          float64 `json:"uvHealthConcern"`
			UvIndex                  float64 `json:"uvIndex"`
			Visibility               float64 `json:"visibility"`
			WeatherCode              int     `json:"weatherCode"`
			WindDirection            float64 `json:"windDirection"`
			WindGust                 float64 `json:"windGust"`
			WindSpeed                float64 `json:"windSpeed"`
		} `json:"values"`
	} `json:"data"`
	Location struct {
		Lat  float64 `json:"lat"`
		Lon  float64 `json:"lon"`
		Name string  `json:"name"`
		Type string  `json:"type"`
	} `json:"location"`
}
