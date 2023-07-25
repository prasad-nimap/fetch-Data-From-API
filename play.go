package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CountryData struct {
	Data struct {
		Covid19Stats []struct {
			Confirmed int `json:"confirmed"`
			Deaths    int `json:"deaths"`
		} `json:"covid19Stats"`
	} `json:"data"`
}

func main() {
	country := "India" // Replace this with the desired country name

	url := fmt.Sprintf("https://covid-19-coronavirus-statistics.p.rapidapi.com/countries?country=%s", country)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return
	}

	req.Header.Add("X-RapidAPI-Key", "4d2c003e3emsh82512a559ce6b89p16ac42jsnc0e7dc1d61ac")
	req.Header.Add("X-RapidAPI-Host", "covid-19-coronavirus-statistics.p.rapidapi.com")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %s\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API request failed with status code: %d\n", resp.StatusCode)
		return
	}

	var data CountryData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding JSON: %s\n", err)
		return
	}

/* 	if len(data.Data.Covid19Stats) == 0 {
		fmt.Println("Country data not found.")
		return
	} */

	countryData := data.Data.Covid19Stats[0]
	fmt.Printf("Country: %s\nConfirmed cases: %d\nDeaths: %d\n", country, countryData.Confirmed, countryData.Deaths)
}
