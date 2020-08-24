package main

import (
	"os"

	app "github.com/ganesh-sai/go-curd-tdd"
)

func main() {
	a := app.App{}
	a.Initialize(os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("DB_NAME"))
	a.Run(":8010")
}
