package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Data struct {
	Domain       string  `json:""`
	Port         int     `json:"port"`
	StatusCode   int     `json:"status_code"`
	ResponseIp   string  `json:"response_ip"`
	ResponseCode int     `json:"response_code"`
	ResponseTime float64 `json:"response_time"`
}

func Red(string string) string {
	return "\033[0;31m" + string + "\033[0m"
}

func Green(string string) string {
	return "\033[0;32m" + string + "\033[0m"
}

func Yellow(string string) string {
	return "\033[0;33m" + string + "\033[0m"
}

func main() {
	// Read arguments
	if len(os.Args) < 2 {
		fmt.Println("USAGE: is-up <DOMAIN URL>")
		os.Exit(1)
	}
	domain := os.Args[1]

	// Get response
	requestUrl := fmt.Sprintf("https://isitup.org/%s.json", domain)
	res, err := http.Get(requestUrl)
	if err != nil {
		fmt.Println("ERROR: Could not get respnse from server.")
		os.Exit(1)
	}
	defer res.Body.Close()

	// Parse JSON
	var data Data
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Output Result
	responseString := Red("INVALID DOMAIN")
	switch data.StatusCode {
	case 1:
		responseString = Green("UP")
	case 2:
		responseString = Red("DOWN")
	}

	domainString := Yellow(data.Domain)
	fmt.Printf("%s is %s\n", domainString, responseString)
}
