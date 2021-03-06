package conv

// CToF converts a Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c - 273.15)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k + 273.15)
}

func FtToM(f Foot) Meter {
	return Meter(f * .3048)
}

func MToFt(m Meter) Foot {
	return Foot(m / .3048)
}
