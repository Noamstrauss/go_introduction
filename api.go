package main

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"net/url"
	"os"
	"strconv"
	"strings"
)

//type addr struct {
//	House_number int
//	Road         string
//	type       	 string
//	display_name string `json:"total_pages"`
//	Data         []struct {
//		ID        int    `json:"id"`
//		Email     string `json:"email"`
//		FirstName string `json:"first_name"`
//		LastName  string `json:"last_name"`
//		Avatar    string `json:"avatar"`
//	} `json:"data"`
//	Support struct {
//		URL  string `json:"url"`
//		Text string `json:"text"`
//	} `json:"support"`
//}

/*
[,"type":"house","place_rank":30,"importance":9.99999999995449e-06,"addresstype":"place","name":"","display_name":"98, Camp Street, Downtown New Britain, Barrio Latino, New Britain, Capitol Planning Region, Connecticut, 06052, United States","address":{"house_number":"98","road":"Camp Street","neighbourhood":"Downtown New Britain","suburb":"Barrio Latino","city":"New Britain","county":"Capitol Planning Region","state":"Connecticut","ISO3166-2-lvl4":"US-CT","postcode":"06052","country":"United States","country_code":"us"},"boundingbox":["41.6612417","41.6613417","-72.7844727","-72.7843727"]}]
*/

func SearchAddress(addr string, city string) string {
	limit := 5
	url, err := url.Parse("https://nominatim.openstreetmap.org/search?")
	if err != nil {
		log.Fatal(err)
	}

	// Use the Query() method to get the query string params as a url.Values map.
	values := url.Query()

	// Make the changes that you want using the Add(), Set() and Del() methods. If
	// you want to retrieve or check for a specific parameter you can use the Get()
	// and Has() methods respectively.
	if city != "" {
		log.Debug("Adding city to query")
		values.Add("city", city)
		values.Add("street", addr)
	} else {
		values.Set("q", addr+""+city)
	}
	values.Add("addressdetails", "1")
	values.Add("format", "jsonv2")
	values.Add("limit", strconv.Itoa(limit))
	//values.Del("gender")
	//values.Set("city", city)
	// Use the Encode() method to transform the url.Values map into a URL-encoded
	// string (like "age=29&name=alice...") and assign it back to the URL. Note
	// that the encoded values will be sorted alphabetically based on the parameter
	// name.
	url.RawQuery = values.Encode()

	//fmt.Printf("url: %s", url.String())
	return url.String()
	// make the rest api request

}

//func get_address_info(addr string, city string) *request {
//
//	encodedAddress := url.QueryEscape(addr)
//	url := fmt.Sprintf("https://nominatim.openstreetmap.org/search?addressdetails=1&q=%v%2C+%v&format=jsonv2&limit=1", encodedAddress, city)
//	fmt.Println(url)
//	client := &http.Client{}
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		fmt.Println(err)
//	}
//	resp, err := client.Do(req)
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println(string(body))
//	return req
//}

func main() {
	addr, city := GetAddressInput()
	log.Debug("Address = ", addr)
	if city != "" {
		log.Debug("City = ", city)
	}
	log.Info(SearchAddress(addr, city))
}

func GetAddressInput() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	// Prompt the user to enter the address
	log.Info("Enter the address: ")
	scanner.Scan()
	address := scanner.Text()
	// Initialize city variable to handle cases without a comma
	city := ""
	if strings.Contains(address, ",") {
		log.Debug("found city because address contains ','")
		parts := strings.Split(address, ",")  // Split into parts using the comma
		city = strings.TrimSpace(parts[1])    // Extract the city, removing leading/trailing spaces
		address = strings.TrimSpace(parts[0]) // Update address with trimmed first part
	}
	return address, city
}
