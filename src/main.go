package main

import (
	"src/createDBTable"
	"src/fetchFixtures"
)

func main() {

	fetchFixtures.HitAPI()
	createDBTable.CreateDBTable()

}
