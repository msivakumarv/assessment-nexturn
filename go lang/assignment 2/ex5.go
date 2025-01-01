package main

import (
	"errors"
	"fmt"
	"strings"
)

type StateData struct {
	Name        string
	Temperature float64
	Rainfall    float64
}

var stateData = []StateData{
	{"maharashtra", 24.5, 1200},
	{"kerala", 27.0, 3500},
	{"uttar pradesh", 22.0, 800},
	{"tamil nadu", 28.0, 1500},
	{"rajasthan", 30.0, 300},
	{"karnataka", 25.5, 1500},
}

func findHighestTemperature(data []StateData) (StateData, error) {
	if len(data) == 0 {
		return StateData{}, errors.New("no data available")
	}

	highest := data[0]
	for _, state := range data {
		if state.Temperature > highest.Temperature {
			highest = state
		}
	}
	return highest, nil
}

func findLowestTemperature(data []StateData) (StateData, error) {
	if len(data) == 0 {
		return StateData{}, errors.New("no data available")
	}

	lowest := data[0]
	for _, state := range data {
		if state.Temperature < lowest.Temperature {
			lowest = state
		}
	}
	return lowest, nil
}

func calculateAverageRainfall(data []StateData) float64 {
	if len(data) == 0 {
		return 0
	}

	totalRainfall := 0.0
	for _, state := range data {
		totalRainfall += state.Rainfall
	}
	return totalRainfall / float64(len(data))
}

func filterStatesByRainfall(data []StateData, threshold float64) []StateData {
	filtered := []StateData{}
	for _, state := range data {
		if state.Rainfall > threshold {
			filtered = append(filtered, state)
		}
	}
	return filtered
}

func searchStateByName(data []StateData, name string) (StateData, error) {
	for _, state := range data {
		if strings.EqualFold(state.Name, name) {
			return state, nil
		}
	}
	return StateData{}, errors.New("state not found")
}

func main() {
	fmt.Println("Climate Data Analysis - Indian States")

	highest, _ := findHighestTemperature(stateData)
	fmt.Printf("\nState with the highest temperature: %s (%.1f°C)\n", highest.Name, highest.Temperature)

	lowest, _ := findLowestTemperature(stateData)
	fmt.Printf("State with the lowest temperature: %s (%.1f°C)\n", lowest.Name, lowest.Temperature)

	averageRainfall := calculateAverageRainfall(stateData)
	fmt.Printf("\nAverage Rainfall across all states: %.2f mm\n", averageRainfall)

	var threshold float64
	fmt.Print("\nEnter rainfall threshold (mm) to filter states: ")
	_, err := fmt.Scan(&threshold)
	if err != nil {
		fmt.Println("Invalid input for threshold.")
		return
	}

	filteredStates := filterStatesByRainfall(stateData, threshold)
	fmt.Printf("\nStates with rainfall above %.2f mm:\n", threshold)
	for _, state := range filteredStates {
		fmt.Printf("- %s (%.2f mm)\n", state.Name, state.Rainfall)
	}

	var searchName string
	fmt.Print("\nEnter state name to search: ")
	fmt.Scan(&searchName)

	state, err := searchStateByName(stateData, searchName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nState Found: %s\nTemperature: %.1f°C\nRainfall: %.2f mm\n", state.Name, state.Temperature, state.Rainfall)
	}
}
