package brzao

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
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

type DateOption string

const (
	Today     DateOption = "today"
	Tomorrow  DateOption = "tomorrow"
	Yesterday DateOption = "yesterday"
)

func (d DateOption) IsValid() bool {
	switch d {
	case Today, Tomorrow, Yesterday:
		return true
	}
	return false
}

func (d DateOption) ToTime() time.Time {
	now := time.Now().Local()
	switch d {
	case Tomorrow:
		return now.AddDate(0, 0, 1)
	case Yesterday:
		return now.AddDate(0, 0, -1)
	default:
		return now
	}
}

func Matches(date DateOption) {
	if !date.IsValid() {
		log.Fatalf("Invalid date option. Please use Today, Tomorrow, or Yesterday")
	}

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://v3.football.api-sports.io/fixtures", nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("x-rapidapi-host", "v3.football.api-sports.io")
	req.Header.Add("x-rapidapi-key", os.Getenv("API_KEY"))

	brzao_league_id := 71
	matches_date := date.ToTime().Format("2006-01-02")

	q := req.URL.Query()
	q.Add("league", strconv.Itoa(brzao_league_id))
	q.Add("season", strconv.Itoa(date.ToTime().Year()))
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

	if len(match.Match) == 0 {
		fmt.Println("error: No matches found for the specified date")
		return
	}

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