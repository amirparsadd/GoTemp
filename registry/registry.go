package registry

import (
	"gotemp/units"
)

// UnitConverter interface defines the contract for converting to and from Kelvin.
type UnitConverter interface {
	ToKelvin(value float64) float64
	FromKelvin(kelvinValue float64) float64
	Symbol() string
}

// RegisterUnit allows adding support for new temperature units.
func RegisterUnit(converter UnitConverter) {
	UnitRegistry[converter.Symbol()] = converter
}

// Register the base units like Kelvin, Celsius and Fahrenheit
func RegisterBaseUnits() {
	RegisterUnit(units.KelvinConverter{})
	RegisterUnit(units.CelsiusConverter{})
	RegisterUnit(units.FahrenheitConverter{})
}

// GetAllowedUnits returns a slice of the currently supported unit symbols.
func GetAllowedUnits() []string {
	units := make([]string, 0, len(UnitRegistry))
	for symbol := range UnitRegistry {
		units = append(units, symbol)
	}
	return units
}

// UnitRegistry holds the registered UnitConverter implementations.
var UnitRegistry = map[string]UnitConverter{}