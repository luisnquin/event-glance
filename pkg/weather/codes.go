package weather

type Code int

const (
	ClearSky                     Code = 0
	MainlyClear                  Code = 1
	PartlyCloudy                 Code = 2
	Overcast                     Code = 3
	Fog                          Code = 45
	DepositingRimeFog            Code = 48
	DrizzleLight                 Code = 51
	DrizzleModerate              Code = 53
	DrizzleDense                 Code = 55
	FreezingDrizzleLight         Code = 56
	FreezingDrizzleDense         Code = 57
	RainSlight                   Code = 61
	RainModerate                 Code = 63
	RainHeavy                    Code = 65
	FreezingRainLight            Code = 66
	FreezingRainHeavy            Code = 67
	SnowFallSlight               Code = 71
	SnowFallModerate             Code = 73
	SnowFallHeavy                Code = 75
	SnowGrains                   Code = 77
	RainShowersSlight            Code = 80
	RainShowersModerate          Code = 81
	RainShowersViolent           Code = 82
	SnowShowersSlight            Code = 85
	SnowShowersHeavy             Code = 86
	ThunderstormSlightOrModerate Code = 95
	ThunderstormSlightHail       Code = 96
	ThunderstormHeavyHail        Code = 99
)

func (c Code) String() string {
	switch c {
	case ClearSky:
		return "Clear sky"
	case MainlyClear:
		return "Mainly clear"
	case PartlyCloudy:
		return "Partly cloudy"
	case Overcast:
		return "Overcast"
	case Fog:
		return "Fog"
	case DepositingRimeFog:
		return "Depositing rime fog"
	case DrizzleLight:
		return "Drizzle light"
	case DrizzleModerate:
		return "Drizzle moderate"
	case DrizzleDense:
		return "Drizzle dense"
	case FreezingDrizzleLight:
		return "Freezing drizzle light"
	case FreezingDrizzleDense:
		return "Freezing drizzle dense"
	case RainSlight:
		return "Rain slight"
	case RainModerate:
		return "Rain moderate"
	case RainHeavy:
		return "Rain heavy"
	case FreezingRainLight:
		return "Freezing rain light"
	case FreezingRainHeavy:
		return "Freezing rain heavy"
	case SnowFallSlight:
		return "Snow fall slight"
	case SnowFallModerate:
		return "Snow fall moderate"
	case SnowFallHeavy:
		return "Snow fall heavy"
	case SnowGrains:
		return "Snow grains"
	case RainShowersSlight:
		return "Rain showers slight"
	case RainShowersModerate:
		return "Rain showers moderate"
	case RainShowersViolent:
		return "Rain showers violent"
	case SnowShowersSlight:
		return "Snow showers slight"
	case SnowShowersHeavy:
		return "Snow showers heavy"
	case ThunderstormSlightOrModerate:
		return "Thunderstorm slight or moderate"
	case ThunderstormSlightHail:
		return "Thunderstorm with slight hail"
	case ThunderstormHeavyHail:
		return "Thunderstorm with heavy hail"
	}

	return ""
}
