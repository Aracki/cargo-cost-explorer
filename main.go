package main

import (
	"fmt"
	"github.com/DusanKasan/parsemail"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("CarGo_1.eml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	email, err := parsemail.Parse(file) // returns Email struct and error
	if err != nil {
		log.Fatal(err)
	}

	f := findCost(email.HTMLBody)

	fmt.Println(email.Date)
	fmt.Println(f)
}

func findCost(body string) float64 {
	i := strings.Index(body, "rsd")
	s := body[i-7:i-1]
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

