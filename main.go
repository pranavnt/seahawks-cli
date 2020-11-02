package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You have to provide a command :)")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "score":
		fmt.Println("getting scores")
		parseScore()

	default:
		fmt.Println("command not recognized")
		os.Exit(1)
	}
}

func parseScore() {
	url := constructURL(2020, "REG", 7)

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	fmt.Println(res)
	fmt.Println(string(body))
}

func constructURL(season int, seasonType string, week int) (url string) {
	url = "https://api.sportsdata.io/v3/nfl/scores/json/TeamGameStats/" + str(season) + seasonType + "/" + str(week) + "?key=" + loadAPI()
	return
}

func str(intStr int) (str string) {
	str = strconv.Itoa(intStr)
	return
}

func loadAPI() (key string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	key = os.Getenv("API_KEY")
	return
}
