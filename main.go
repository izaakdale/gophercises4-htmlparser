package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/izaakdale/htmlParser/links"
)

func main() {

	file, err := os.Open("ex4.html")
	if err != nil {
		log.Fatal(err.Error())
	}

	links, err := links.Parse(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	var buf bytes.Buffer
	je := json.NewEncoder(&buf)

	je.Encode(links)
	fmt.Print(&buf)
}
