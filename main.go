package main

import (
	"fmt"
	"github.com/DusanKasan/parsemail"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const dirName = "all_emails"
const year = 2019

func main() {
	parseFiles()
}

func parseFiles() {

	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	var sum float64
	sum = 0

	for _, f := range files {
		if !f.IsDir() && strings.Contains(f.Name(), ".eml") {
			sum += parseFile(dirName + "/" + f.Name())
		}
	}

	fmt.Println("Sum in dinars: ", sum)
}

func parseFile(fileName string) float64 {

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	email, err := parsemail.Parse(f) // returns Email struct and error
	if err != nil {
		log.Fatal(err)
	}

	if email.Date.Year() != year {
		return 0
	}
	return findCost(email.HTMLBody)
}

func findCost(body string) float64 {

	i := strings.Index(body, "rsd")
	if i == -1 {
		return 0
	}
	s := body[i-7 : i-1]
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
