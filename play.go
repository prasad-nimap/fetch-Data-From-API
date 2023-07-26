package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type APIResponse struct {
	Confirmed int    `json:"confirmed"`
	Deaths    int    `json:"deaths"`
	Recovered int    `json:"recovered"`
	Location  string `json:"location"`
}

// Append data to file
func appendToJSONFile(filepath string, apiResponse APIResponse) {
	// Read existing data from the file, if any
	existingData, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error while reading file: %s\n", err)
		return
	}

	// Unmarshal the existing JSON data into a slice of APIResponse
	var existingDataSlice []APIResponse
	if err := json.Unmarshal(existingData, &existingDataSlice); err != nil {
		fmt.Printf("Error while unmarshaling existing data: %s\n", err)
		return
	}

	// Append the new APIResponse to the existing slice
	existingDataSlice = append(existingDataSlice, apiResponse)

	// Marshal the updated data back to JSON
	updatedJSONData, err := json.Marshal(existingDataSlice)
	if err != nil {
		fmt.Printf("Error while marshaling updated data: %s\n", err)
		return
	}

	// Update the JSON file
	err = ioutil.WriteFile(filepath, updatedJSONData, 0644)
	if err != nil {
		fmt.Printf("Error while writing to file: %s\n", err)
		return
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the output file path as an argument.")
		return
	}
	filePath := os.Args[1]

	url := "https://covid-19-coronavirus-statistics.p.rapidapi.com/v1/country"
	country := "Canada"

	req, err := http.NewRequest("GET", url+"?country="+country, nil)
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return
	}

	req.Header.Add("X-RapidAPI-Key", "4d2c003e3emsh82512a559ce6b89p16ac42jsnc0e7dc1d61ac")
	req.Header.Add("X-RapidAPI-Host", "covid-19-coronavirus-statistics.p.rapidapi.com")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error sending the request: %s\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API request failed with status code: %d\n", resp.StatusCode)
		return
	}

	var apiResponse struct {
		Data struct {
			Covid19Stats []APIResponse `json:"covid19Stats"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		fmt.Printf("Error decoding JSON: %s\n", err)
		return
	}

	if len(apiResponse.Data.Covid19Stats) == 0 {
		fmt.Println("Country data not found.")
		return
	}

	latestData := apiResponse.Data.Covid19Stats[len(apiResponse.Data.Covid19Stats)-1]
	appendToJSONFile(filePath, latestData)

	fmt.Println("Data has been successfully saved to", filePath)
} 	