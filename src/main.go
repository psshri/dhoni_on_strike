package main

import (
	"src/createDBTable"
	"src/fixtures"
)

func main() {

	fixtures.HitAPI()
	createDBTable.CreateDBTable()

}
