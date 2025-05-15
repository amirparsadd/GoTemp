package units

// CelsiusConverter implements the UnitConverter interface for Celsius.
type CelsiusConverter struct{}
func (CelsiusConverter) ToKelvin(celsius float64) float64 { return celsius + 273.15 }
func (CelsiusConverter) FromKelvin(kelvin float64) float64 { return kelvin - 273.15 }
func (CelsiusConverter) Symbol() string { return "C" }