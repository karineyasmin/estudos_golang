package main

import "fmt"

type fahrenheit float64
type celsius float64
type kelvin float64

// receiver fica entre o func e o nome da função
func (f fahrenheit) toCelsius() celsius {
	return celsius(f-32) * 5 / 9
}

func (c celsius) toFahrenheit() fahrenheit {
	return fahrenheit(c*9/5) + 32
}

func (k kelvin) toCelsius() celsius {
	return celsius(k - 273.15)
}

func (k kelvin) toFahrenheit() fahrenheit {
	return k.toCelsius().toFahrenheit()
}
func main() {
	var tempAtualF fahrenheit = 32
	var tempAtualC celsius = 0
	var tempAtualK kelvin = 273.15
	fmt.Printf("%.2f°C\n", tempAtualF.toCelsius())
	fmt.Printf("%.2f°F\n", tempAtualC.toFahrenheit())
	fmt.Printf("%.2f°F\n", tempAtualK.toCelsius())
	fmt.Printf("%.2f°F\n", tempAtualK.toFahrenheit())

}
