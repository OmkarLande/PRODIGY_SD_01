package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r, "static/index.html") // Remove "/static/"
		} else if r.Method == "POST" {
			tempStr := r.FormValue("temperature")
			unit := r.FormValue("unit")

			temp, err := strconv.ParseFloat(tempStr, 64)
			if err != nil {
				http.Error(w, "Invalid temperature value", http.StatusBadRequest)
				return
			}

			unit = strings.ToLower(unit)

			var result string
			switch unit {
			case "c":
				result = fmt.Sprintf("Fahrenheit: %.2f\nKelvin: %.2f\n", celsiusToFahrenheit(temp), celsiusToKelvin(temp))
			case "f":
				result = fmt.Sprintf("Celsius: %.2f\nKelvin: %.2f\n", fahrenheitToCelsius(temp), fahrenheitToKelvin(temp))
			case "k":
				result = fmt.Sprintf("Celsius: %.2f\nFahrenheit: %.2f\n", kelvinToCelsius(temp), kelvinToFahrenheit(temp))
			default:
				http.Error(w, "Invalid unit", http.StatusBadRequest)
				return
			}

			w.Write([]byte(result))
		}
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// Celsius to Fahrenheit
func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

// Celsius to Kelvin
func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

// Fahrenheit to Celsius
func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) / 1.8
}

// Fahrenheit to Kelvin
func fahrenheitToKelvin(fahrenheit float64) float64 {
	return (fahrenheit-32)/1.8 + 273.15
}

// Kelvin to Celsius
func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

// Kelvin to Fahrenheit
func kelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin-273.15)*1.8 + 32
}
