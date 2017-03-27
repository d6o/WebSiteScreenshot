package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	api := flag.String("api", "", "API KEY from thumbnail.ws")
	website := flag.String("website", "", "Website URL")

	flag.Parse()

	config := NewConfig(*api, *website)
	err := config.Validate()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	app := NewApp(config)

	err = app.takeScreenShot()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Download finished!")
}
