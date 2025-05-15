package converter

import "errors"

// UnitConverter interface defines the contract for converting to and from Kelvin.
type UnitConverter interface {
	ToKelvin(value float64) float64
	FromKelvin(kelvinValue float64) float64
	Symbol() string
}

// celsiusConverter implements the UnitConverter interface for Celsius.
type celsiusConverter struct{}
func (celsiusConverter) ToKelvin(celsius float64) float64 { return celsius + 273.15 }
func (celsiusConverter) FromKelvin(kelvin float64) float64 { return kelvin - 273.15 }
func (celsiusConverter) Symbol() string { return "C" }

// fahrenheitConverter implements the UnitConverter interface for Fahrenheit.
type fahrenheitConverter struct{}
func (fahrenheitConverter) ToKelvin(fahrenheit float64) float64 { return (fahrenheit - 32) * 5 / 9 + 273.15 }
func (fahrenheitConverter) FromKelvin(kelvin float64) float64 { return (kelvin - 273.15) * 9 / 5 + 32 }
func (fahrenheitConverter) Symbol() string { return "F" }

// kelvinConverter implements the UnitConverter interface for Kelvin.
type kelvinConverter struct{}
func (kelvinConverter) ToKelvin(kelvin float64) float64 { return kelvin }
func (kelvinConverter) FromKelvin(kelvin float64) float64 { return kelvin }
func (kelvinConverter) Symbol() string { return "K" }

// UnitRegistry holds the registered UnitConverter implementations.
var UnitRegistry = map[string]UnitConverter{}

// RegisterUnit allows adding support for new temperature units.
func RegisterUnit(converter UnitConverter) {
	UnitRegistry[converter.Symbol()] = converter
}

func RegisterBaseUnits() {
	RegisterUnit(kelvinConverter{})
	RegisterUnit(celsiusConverter{})
	RegisterUnit(fahrenheitConverter{})
}

// Convert performs the temperature conversion.
func Convert(value float64, fromSymbol string, toSymbol string) (float64, error) {
	fromConverter, ok := UnitRegistry[fromSymbol]
	if !ok {
		return 0, errors.New("invalid input unit: " + fromSymbol)
	}

	toConverter, ok := UnitRegistry[toSymbol]
	if !ok {
		return 0, errors.New("invalid output unit: " + toSymbol)
	}

	kelvinValue := fromConverter.ToKelvin(value)
	return toConverter.FromKelvin(kelvinValue), nil
}

// GetAllowedUnits returns a slice of the currently supported unit symbols.
func GetAllowedUnits() []string {
	units := make([]string, 0, len(UnitRegistry))
	for symbol := range UnitRegistry {
		units = append(units, symbol)
	}
	return units
}