package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	// Third party packages
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// Config a struct where I am able to get .env variables
type config struct {
	APIKey string `env:"weatherApiKey"`
}


func main() {
	// Get .env file variables - can be turned into a function
	weatherAPIKey := getWeatherAPIKey()

	// prompt user to put in their zipcode
	fmt.Println("Welcome to Weather Mood!")
	fmt.Print("Please enter the city you are in: ")
	city := getCity()

	apiLink := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + weatherAPIKey + "&units=Imperial"

	response, err := http.Get(apiLink)
	// display weather
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Well this is the current weather of " + city + " is:\n")
	fmt.Println(string(responseData))

	// after displaying the weather ask the user
	// how do you feel?
	fmt.Print("How do you feel? ")
	mood := getMood()
	fmt.Printf("Well with this weather who can blame you for feeling " + mood + "!\n")

}

func getWeatherAPIKey() string {
	if err := godotenv.Load(); err != nil {
		log.Println("File .env not found, reading configuration from ENV")
	}
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Failed to parse ENV")
	}

	return cfg.APIKey
}

func getCity() string {
	reader := bufio.NewReader(os.Stdin)
	city, _ := reader.ReadString('\n')
	city = strings.Trim(city, "\n")
	return city
}

func getMood() string {
	reader := bufio.NewReader(os.Stdin)
	mood, _ := reader.ReadString('\n')
	mood = strings.Trim(mood, "\n")
	return mood
}
