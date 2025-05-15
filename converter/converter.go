package converter

import (
	"errors"
	"gotemp/registry"
)

// Convert performs the temperature conversion.
func Convert(value float64, fromSymbol string, toSymbol string) (float64, error) {
	fromConverter, ok := registry.UnitRegistry[fromSymbol]
	if !ok {
		return 0, errors.New("invalid input unit: " + fromSymbol)
	}

	toConverter, ok := registry.UnitRegistry[toSymbol]
	if !ok {
		return 0, errors.New("invalid output unit: " + toSymbol)
	}

	kelvinValue := fromConverter.ToKelvin(value)
	return toConverter.FromKelvin(kelvinValue), nil
}