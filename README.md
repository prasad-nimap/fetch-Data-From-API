# COVID-19 Data Fetcher

This Go program fetches COVID-19 data from the RapidAPI free API and appends it to a JSON file. It allows you to retrieve COVID-19 data for a specific country and store it for further analysis or usage.

## Prerequisites

- Go installed on your machine. You can download and install Go from the [official Go website](https://golang.org/dl/).

## Getting Started

1. Clone this repository to your local machine or download the source code files.

2. Open a terminal or command prompt and navigate to the project's directory.

3. Open the `main.go` file and replace the following placeholders with your actual values:
   - `'YourCountryName'` with the desired country name.
   - `'YourRapidAPIKey'` with your actual RapidAPI key.

4. Save the changes to the `main.go` file.

5. Run the program by executing the following command:

```
go run main.go
```


6. The program will fetch the COVID-19 data for the specified country and append it to a JSON file named `covid_data.json`. If the file doesn't exist, it will be created. If the file already exists, the new data will be appended to it.

7. Once the program finishes running, you can find the JSON file containing the COVID-19 data in the same directory as the `main.go` file.

## Customization

- Fetching data for multiple countries: If you want to fetch COVID-19 data for multiple countries, you can modify the program to accept a list of country names and loop over them to fetch data for each country. You can also adjust the JSON file structure to suit your specific requirements.

- Changing the JSON file name: If you prefer to use a different file name or location to store the COVID-19 data, you can modify the `filename` variable in the `main()` function in the `main.go` file.

## Dependencies

This program uses the following third-party Go package:

- `net/http`: Standard Go library for making HTTP requests.

## License

This project is licensed under the MIT License. Feel free to modify and use it according to your needs.

## Disclaimer

The COVID-19 data is fetched from the RapidAPI free API. The accuracy and reliability of the data depend on the API provider. Please refer to the RapidAPI documentation for more details on the API usage, limitations, and terms of service.

## Support

If you encounter any issues or have any questions or suggestions, please feel free to open an issue or reach out to me.

---

