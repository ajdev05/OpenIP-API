package main

import (
  "net/http"
  "io/ioutil"
	"encoding/json"
	"fmt"	
)

type OpenIPAPI_INFO struct {
	IP      string `json:"ip"`
	Country string `json:"country"`
	City    string `json:"city"`
}

func main() {
	ipAddress := "1.1.1.1"
	url := fmt.Sprintf("https://open-ip-api.vercel.app/api?ip=%s", ipAddress)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to make GET request:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	var OpenIPAPI_INFO OpenIPAPI_INFO
	err = json.Unmarshal(body, &OpenIPAPI_INFO)
	if err != nil {
		fmt.Println("Failed to decode JSON response:", err)
		return
	}

  fmt.Println("IP:", OpenIPAPI_INFO.IP)
	fmt.Println("Country:", OpenIPAPI_INFO.Country)
	fmt.Println("City:", OpenIPAPI_INFO.City)
}
