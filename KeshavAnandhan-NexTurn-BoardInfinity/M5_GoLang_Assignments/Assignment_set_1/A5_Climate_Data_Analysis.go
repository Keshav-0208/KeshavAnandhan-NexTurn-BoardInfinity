package main

import (
	"fmt"
	"errors"
)

type City struct {
	Name           string
	Temperature    float64
	Rainfall       float64
}

var cities = []City{
	{"Chennai", 32.0, 1200.0}, // Chennai, Tamil Nadu
	{"Coimbatore", 28.0, 800.0},
	{"Madurai", 33.0, 700.0},
	{"Trichy", 30.0, 500.0},
}

func highestTemp() string {
	highest := cities[0]
	for _, city := range cities {
		if city.Temperature > highest.Temperature {
			highest = city
		}
	}
	return highest.Name
}

func lowestTemp() string {
	lowest := cities[0]
	for _, city := range cities {
		if city.Temperature < lowest.Temperature {
			lowest = city
		}
	}
	return lowest.Name
}

func avgRainfall() float64 {
	total := 0.0
	for _, city := range cities {
		total += city.Rainfall
	}
	return total / float64(len(cities))
}

func citiesAboveRainfall(threshold float64) []string {
	var above []string
	for _, city := range cities {
		if city.Rainfall > threshold {
			above = append(above, city.Name)
		}
	}
	return above
}

func searchCity(name string) (City, error) {
	for _, city := range cities {
		if city.Name == name {
			return city, nil
		}
	}
	return City{}, errors.New("City not found")
}

func main() {
	fmt.Println("Highest Temp:", highestTemp())
	fmt.Println("Lowest Temp:", lowestTemp())
	fmt.Println("Average Rainfall:", avgRainfall())

	var threshold float64
	fmt.Print("Enter rainfall threshold: ")
	_, err := fmt.Scan(&threshold)
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}

	above := citiesAboveRainfall(threshold)
	if len(above) > 0 {
		fmt.Println("Cities with rainfall above threshold:")
		for _, city := range above {
			fmt.Println(city)
		}
	} else {
		fmt.Println("No cities have rainfall above the threshold.")
	}

	var searchName string
	fmt.Print("Enter city name to search: ")
	fmt.Scan(&searchName)

	city, err := searchCity(searchName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("City %s has %.2f°C temperature and %.2f mm rainfall.\n", city.Name, city.Temperature, city.Rainfall)
	}
}
