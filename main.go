package main

import (
	"fmt"
	"io/ioutil"
	rand2 "math/rand"
	"strings"

	"gopkg.in/yaml.v3"
)

type passwordConfig struct {
	Count                int    `yaml:"count"`
	Width                int    `yaml:"width"`
	Letters              string `yaml:"letters"`
	Numbers              string `yaml:"numbers"`
	SpecialCharacters    string `yaml:"specialCharacters"`
	IsRepeatedCharacters bool   `yaml:"isRepeatedCharacters"`
}

func main() {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	var config passwordConfig

	if yaml.Unmarshal(data, &config) != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	var symbols string = config.Letters + config.Numbers + config.SpecialCharacters
	var symbolsCount int = len(symbols)
	uniqueSymbols := make(map[string]int)
	if !config.IsRepeatedCharacters {
		for _, val := range symbols {
			uniqueSymbols[string(val)] = 1
		}
	}

	for i := 0; i < config.Count; i++ {
		password := ""

		for j := 0; j < config.Width; j++ {
			var randIndex int = rand2.Intn(symbolsCount)
			var symbol string = string(symbols[randIndex])

			if !config.IsRepeatedCharacters {
				if len(password) >= len(uniqueSymbols) {
					break
				}
				if strings.Contains(password, symbol) {

					j--
					continue
				}
			}

			password += symbol
		}
		fmt.Println(password)
	}
}
