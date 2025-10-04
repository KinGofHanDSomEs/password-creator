package main

import (
	"fmt"
	"io/ioutil"
	rand2 "math/rand"

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
		panic(fmt.Sprintf("error: %e", err))
	}

	var config passwordConfig

	if yaml.Unmarshal(data, &config) != nil {
		panic(fmt.Sprintf("error: %e", err))
	}

	var symbols string = config.Letters + config.Numbers + config.SpecialCharacters

	for i := 0; i < config.Count; i++ {
		password, symbolsCopy := "", symbols

		for j := 0; j < config.Width; j++ {

			symbolsCount := len(symbolsCopy)
			if symbolsCount == 0 {
				break
			}

			var randIndex int = rand2.Intn(symbolsCount)
			var symbol string = string(symbolsCopy[randIndex])

			if !config.IsRepeatedCharacters {
				symbolsCopy = symbolsCopy[:randIndex] + symbolsCopy[randIndex+1:]
			}

			password += symbol
		}
		fmt.Println(password)
	}
}
