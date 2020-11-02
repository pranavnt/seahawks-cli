package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
	url := "https://therundown-therundown-v1.p.rapidapi.com/affiliates"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "therundown-therundown-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", loadAPI())

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func loadAPI() (key string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	key = os.Getenv("API_KEY")
	return
}
