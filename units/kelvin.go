package units

// KelvinConverter implements the UnitConverter interface for Kelvin.
type KelvinConverter struct{}
func (KelvinConverter) ToKelvin(kelvin float64) float64 { return kelvin }
func (KelvinConverter) FromKelvin(kelvin float64) float64 { return kelvin }
func (KelvinConverter) Symbol() string { return "K" }