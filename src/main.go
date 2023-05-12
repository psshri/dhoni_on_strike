package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "https://cricket-live-data.p.rapidapi.com/series"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "cc8d5ba84amshb27935b6f1362f5p1be649jsnfa798a2da04d")
	req.Header.Add("X-RapidAPI-Host", "cricket-live-data.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
