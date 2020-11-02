package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You have to provide a command :)")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "score":
		fmt.Println("The seahawks score is...")
	default:
		fmt.Println("command not recognized")
		os.Exit(1)
	}
}

func parseScore() {
	url := "https://therundown-therundown-v1.p.rapidapi.com/affiliates"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "therundown-therundown-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "notAPIKkey")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
