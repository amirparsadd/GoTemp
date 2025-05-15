package units

// FahrenheitConverter implements the UnitConverter interface for Fahrenheit.
type FahrenheitConverter struct{}
func (FahrenheitConverter) ToKelvin(fahrenheit float64) float64 { return (fahrenheit - 32) * 5 / 9 + 273.15 }
func (FahrenheitConverter) FromKelvin(kelvin float64) float64 { return (kelvin - 273.15) * 9 / 5 + 32 }
func (FahrenheitConverter) Symbol() string { return "F" }