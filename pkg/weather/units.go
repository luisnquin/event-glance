package weather

type (
	TemperatureUnit   string
	WindspeedUnit     string
	PrecipitationUnit string
)

const (
	Celsius    TemperatureUnit = "celsius"
	Fahrenheit TemperatureUnit = "fahrenheit"
)

const (
	KMH WindspeedUnit = "kmh"
	MS  WindspeedUnit = "ms"
	MPH WindspeedUnit = "mph"
	KN  WindspeedUnit = "kn"
)

const (
	MM   string = "mm"
	INCH string = "inch"
)
