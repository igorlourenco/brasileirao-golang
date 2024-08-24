package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Match struct {
	Fixture struct {
		Date string `json:"date"`
	} `json:"fixture"`

	Teams struct {
		Home struct {
			Name string `json:"name"`
		} `json:"home"`
		Away struct {
			Name string `json:"name"`
		} `json:"away"`
	} `json:"teams"`
}

type Response struct {
	Match []Match `json:"response"`
}

func reverseArray(arr []Match) {
	start := 0
	end := len(arr) - 1

	for start < end {
		// Swap the elements
		arr[start], arr[end] = arr[end], arr[start]
		// Move towards the middle
		start++
		end--
	}
}

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://v3.football.api-sports.io/fixtures", nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("x-rapidapi-host", "v3.football.api-sports.io")
	req.Header.Add("x-rapidapi-key", "3506bef834adf16f3a3d9897fa385f67")

	brzao_league_id := 71
	current_time := time.Now().Local()
	matches_date := current_time.Format("2006-01-02")

	q := req.URL.Query()
	q.Add("league", strconv.Itoa(brzao_league_id))
	q.Add("season", strconv.Itoa(current_time.Year()))
	q.Add("date", matches_date)
	q.Add("timezone", "America/Araguaina")
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatal(res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var match Response

	err = json.Unmarshal(body, &match)

	if err != nil {
		log.Fatal(err)
	}

	reverseArray(match.Match)

	parsedTime, err := time.Parse(time.RFC3339, match.Match[0].Fixture.Date)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parsedTime.Format("2006-01-02"))
	fmt.Println("--------------------------------")
	for _, match := range match.Match {

		parsedTime, err := time.Parse(time.RFC3339, match.Fixture.Date)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(parsedTime.Format("15:04"))
		fmt.Printf("%s x %s\n", match.Teams.Home.Name, match.Teams.Away.Name)
		fmt.Println("--------------------------------")
	}
}
