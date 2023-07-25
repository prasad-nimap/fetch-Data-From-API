package main

import (
	// "encoding/json"
	"fmt"
	"io"
	// "io/ioutil"
	"net/http"
)

/* func main() {

	// NewRequest function of the http package to specify our API call
	// because it allows us to set custom headers.
	url := "https://covid-193.p.rapidapi.com/countries"

	req, err := http.NewRequest("GET", url, nil)
	checkerror(err)

	// req.Header.Add("x-rapidapi-key", "YOU_API_KEY")
	// use the header.Add the function to add headers in API
	req.Header.Add("X-rapidapi-key", "4d2c003e3emsh82512a559ce6b89p16ac42jsnc0e7dc1d61ac")

	req.Header.Add("X-rapidapi-Host", "covid-193.p.rapidapi.com")

	// send request to receive the response from the api
	res, err := http.DefaultClient.Do(req)
	checkerror(err)

	// response
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// check error
func checkerror(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}
*/

/*
	func fetchCOVIDData(country string) ([]COVIDData, error) {
		url := "https://covid-19-data.p.rapidapi.com/country"
		queryString := "?name=" + "USA" // Replace "YourCountryName" with the desired country name

		req, err := http.NewRequest("GET", url+queryString, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Add("X-RapidAPI-Key", "4d2c003e3emsh82512a559ce6b89p16ac42jsnc0e7dc1d61ac") // Replace "YourRapidAPIKey" with your actual RapidAPI key

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var data []COVIDData
		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, err
		}

		return data, nil
	}

	func appendToJSONFile(data []COVIDData, filename string) error {
		var jsonData []COVIDData

		// Check if the file already exists
		file, err := os.Open(filename)
		if err != nil {
			// If the file doesn't exist, create an empty array
			jsonData = []COVIDData{}
		} else {
			// If the file exists, read its content
			defer file.Close()

			fileContent, err := ioutil.ReadAll(file)
			if err != nil {
				return err
			}

			// Unmarshal the existing JSON data
			err = json.Unmarshal(fileContent, &jsonData)
			if err != nil {
				return err
			}
		}

		// Append the new data to the existing JSON array
		jsonData = append(jsonData, data...)

		// Marshal the updated JSON data
		updatedData, err := json.Marshal(jsonData)
		if err != nil {
			return err
		}

		// Write the updated data to the file
		err = ioutil.WriteFile(filename, updatedData, 0644)
		if err != nil {
			return err
		}

		return nil
	}
*/

type COVIDData struct {
	Location  string `json:"location"`
	Confirmed int    `json:"confirmed"`
	Deaths    int    `json:"deaths"`
}

func play() {
	/*
		country := "USA" // Replace "YourCountryName" with the desired country name

		// Fetch COVID-19 data
		covidData, err := fetchCOVIDData(country)
		if err != nil {
			fmt.Println("Error fetching COVID-19 data:", err)
			return
		}

		// Append data to the JSON file
		filename := "covid_data.json"
		err = appendToJSONFile(covidData, filename)
		if err != nil {
			fmt.Println("Error appending data to JSON file:", err)
			return
		}

		fmt.Println("Data appended to", filename)
	*/
	url := "https://covid-19-coronavirus-statistics.p.rapidapi.com/v1/total?country=India"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "4d2c003e3emsh82512a559ce6b89p16ac42jsnc0e7dc1d61ac")
	req.Header.Add("X-RapidAPI-Host", "covid-19-coronavirus-statistics.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
