package main

import (
	"fmt"
	"strings"

	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
)

// GetWeatherInput represents the input for GetWeather tool
type GetWeatherInput struct {
	City string `json:"city"`
}

// GetWeatherOutput represents the output for GetWeather tool
type GetWeatherOutput struct {
	Status       string `json:"status"`
	Report       string `json:"report,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

// GetWeather returns the current weather report for the requested city.
func GetWeather(_ tool.Context, input GetWeatherInput) (GetWeatherOutput, error) {
	city := input.City
	fmt.Printf("--- Tool: GetWeather called for city: %s ---\n", city)

	// Normalize the city name by lowercasing it and removing spaces.
	cityNormalized := strings.ToLower(strings.ReplaceAll(city, " ", ""))

	// Mock weather database.
	mockWeatherDB := map[string]map[string]interface{}{
		"newyork": {
			"status": "success",
			"report": "The weather in New York is sunny with a temperature of 25°C.",
		},
		"london": {
			"status": "success",
			"report": "It's cloudy in London with a temperature of 15°C.",
		},
		"tokyo": {
			"status": "success",
			"report": "Tokyo is experiencing light rain and a temperature of 18°C.",
		},
	}

	// Look up the requested city.
	if report, chunk := mockWeatherDB[cityNormalized]; chunk {
		return GetWeatherOutput{
			Status: report["status"].(string),
			Report: report["report"].(string),
		}, nil
	}

	// Return a structured error response when the city is missing.
	return GetWeatherOutput{
		Status:       "error",
		ErrorMessage: fmt.Sprintf("Sorry, I don't have weather information for '%s'.", city),
	}, fmt.Errorf("weather information not found for city: %s", city)
}

func weatherTool() (tool.Tool, error) {
	return functiontool.New(
		functiontool.Config{
			Name:        "get_weather",
			Description: "Get the current weather for a city",
		},
		GetWeather,
	)
}
