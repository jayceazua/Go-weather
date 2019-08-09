package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Api key = 455b79f495c12591e03323c61a869ce5
func main() {

	// prompt user to put in their zipcode
	// deal with .env file for apikey
	var string zipcode
	// apiLink := "https://samples.openweathermap.org/data/2.5/forecast/hourly?zip=94102&appid=455b79f495c12591e03323c61a869ce5"
	response, err := http.Get("https://samples.openweathermap.org/data/2.5/forecast/hourly?zip=33145&appid=455b79f495c12591e03323c61a869ce5")

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
