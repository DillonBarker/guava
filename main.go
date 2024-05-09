package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	accessToken := os.Getenv("ACCESS_TOKEN")

	url := "https://www.strava.com/api/v3/athlete"

	var bearer = "Bearer " + accessToken

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("Error on request.\n[ERROR] -", err)
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	b, _ := pretty(body)
	fmt.Printf("%s", b)
}

func pretty(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}
