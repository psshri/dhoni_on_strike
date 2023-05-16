// this code parses the fixtures.json file and pushes the data to MySQL db/table

package main

import (
	"fmt"
	"os"
)

func main() {
	// open the deDetails.json file
	fixtures, err := os.Open("../fetchFixtures/fixtures.json")
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of dbDetails file so that we can parse it later on
	defer fixtures.Close()
}
