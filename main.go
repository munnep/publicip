package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	VERSION string = "0.0.1"
)

func main() {

	jsonPtr := flag.Bool("json", false, "Specify if output should be JSON")
	versionPtr := flag.Bool("version", false, "Show the version")

	flag.Parse()

	if *versionPtr {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	if *jsonPtr {
		fmt.Println(PublicIPJson())
	} else {
		fmt.Println(getPublicIP())
	}

}

func getPublicIP() (publicip string) {

	uris := []string{"https://ifconfig.me", "https://ipinfo.io/ip", "https://api.ipify.org"}

	for _, uri := range uris {
		response, err := http.Get(uri)
		fmt.Println(uri)
		if err != nil {
			continue
		}
		body, _ := io.ReadAll(response.Body)
		defer response.Body.Close() //always have to do this last

		return string(body)

	}

	return string("No internet access")
}

func PublicIPJson() (jsondata string) {
	publicIP := getPublicIP()

	// Define a struct to hold the IP address
	type IPAddress struct {
		IP string `json:"ip"`
	}

	// Create an instance of IPAddress with the public IP
	ip := IPAddress{
		IP: publicIP,
	}

	// Convert the IPAddress struct to JSON
	jsonData, err := json.Marshal(ip)
	if err != nil {
		log.Fatal(err)
	}

	// Output the JSON data as a string
	return string(jsonData)
}
