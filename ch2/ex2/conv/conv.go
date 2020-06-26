// Package conv performs conversions.
package conv

import "fmt"

type Meters float64
type Feet float64

type Celsius float64
type Fahrenheit float64

type Kilograms float64
type Pounds float64

func (f Feet) ToMeters() Meters {
	return Meters(f * 0.3048)
}

func (m Meters) ToFeet() Feet {
	return Feet(m / 0.3048)
}

func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (p Pounds) ToKilograms() Kilograms {
	return Kilograms(p * 0.45359237)
}

func (k Kilograms) ToPounds() Pounds {
	return Pounds(k / 0.45359237)
}
func (m Meters) String() string     { return fmt.Sprintf("%.3fm", m) }
func (f Feet) String() string       { return fmt.Sprintf("%.3f\"", f) }
func (c Celsius) String() string    { return fmt.Sprintf("%.3f°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.3f°F", f) }
func (k Kilograms) String() string  { return fmt.Sprintf("%.3fkg", k) }
func (p Pounds) String() string     { return fmt.Sprintf("%.3flbs", p) }
