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

// Api key = 455b79f495c12591e03323c61a869ce5
func main() {
	// Get .env file variables - can be turned into a function
	weatherAPIKey := getWeatherAPIKey()

	// prompt user to put in their zipcode
	fmt.Println("Welcome to Weather Mood!")
	fmt.Print("Please enter the city you are in: ")
	city := getCity()
	// fmt.Println(reflect.TypeOf(zipcode))
	// https://api.openweathermap.org/data/2.5/weather?q=${city}&appid=${apikey}&units=${units}
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

	fmt.Println(string(responseData))

	// after displaying the weather ask the user
	// how do you feel?

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
