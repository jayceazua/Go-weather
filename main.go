package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// Third party packages
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)
// Config a struct where I am able to get .env variables
type config struct {
    APIKey string `env:"weatherApiKey"`
}

// Api key = 455b79f495c12591e03323c61a869ce5
func main() {
	// Get .env file variables - can be turned into a function
	weatherAPIKey := getWeatherAPIKey()

	// prompt user to put in their zipcode
	// deal with .env file for apikey
	// var string zipcode
	apiLink := "https://samples.openweathermap.org/data/2.5/forecast/hourly?zip=94102&appid=" + weatherAPIKey
	response, err := http.Get(apiLink)

	// display a nice print on the cli of the weather info
	// City Name
	// description
	// main

	// after displaying the weather ask the user
	// how do you feel?

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
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