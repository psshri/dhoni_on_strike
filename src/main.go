package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
)

var apiKey string = "cc8d5ba84amshb27935b6f1362f5p1be649jsnfa798a2da04d"
var apiHost string = "cricket-live-data.p.rapidapi.com"

func hitAPI(url string) []uint8 {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", apiHost)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}

func main() {

	url := "https://cricket-live-data.p.rapidapi.com/series"

	body := hitAPI(url)

	// find out the data type of body variable
	fmt.Println(reflect.TypeOf(body))

	// fmt.Println(res)
	// fmt.Println(string(body))

	// create a file output.txt
	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}
	defer file.Close()

	// // write body to the file output.txt
	// _, err = file.Write(body)
	// if err != nil {
	// 	fmt.Println("Error writing to file: ", err)
	// 	return
	// }

	// fmt.Println("Output written to file successfully!")

	// explore the data type
	// body_string := string(body)
	var data_json interface{}
	err2 := json.Unmarshal(body, &data_json)
	if err2 != nil {
		fmt.Println("Error unmarshaling body JSON: ", err2)
		return
	}

	formattedData, _ := json.MarshalIndent(data_json, "", "    ")
	// fmt.Println(string(formattedData))

	// writing formattedData to file
	_, err = file.Write(formattedData)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}

	fmt.Println("Output written to file successfully!")

	var data_map map[string]interface{}
	err1 := json.Unmarshal(formattedData, &data_map)
	if err1 != nil {
		fmt.Println("Error unmarshaling formattedData JSON: ", err1)
		return
	}

	for key := range data_map {
		fmt.Println(key)
	}

	// meta := data_map.(map[string]interface{})["meta"]
	// fmt.Println(meta)

	fmt.Println(data_map["meta"])

	// // used to find the type of data in golang
	// fmt.Println(reflect.TypeOf(data_map["meta"].(map[string]interface{})))

	// meta := data_map["meta"].(map[string]interface{})
	// fmt.Println(meta)

}
