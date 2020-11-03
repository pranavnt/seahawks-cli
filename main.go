package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	url := constructURL(2020, "REG", 8)

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)

	scoreJSON := string(body)

	var score []GameStats
	json.Unmarshal([]byte(scoreJSON), &score)

	for i := 0; i < len(score); i++ {
		homeTeam := score[i].Team
		awayTeam := score[i].Opponent
		homeTeamScore := score[i].Score
		awayTeamScore := score[i].OpponentScore

		if homeTeam == "SEA" {
			fmt.Println("Seahawks Score:")
			fmt.Println(homeTeam + ": " + str(homeTeamScore))
		}
	}

	fmt.Println((score[0].Team))
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
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	key = os.Getenv("API_KEY")
	return
}

type GameStats struct {
	Team          string
	Opponent      string
	Score         int
	OpponentScore int
}
