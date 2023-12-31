package helper

import (
	"bufio"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

// map to store first names
var firstnames = make(map[string]string)

// map to store last names
var lastnames = make(map[string]string)

// list to store first names
var firstNameslist = []string{}

// list to store last names
var lastNameslist = []string{}

// slice to store full names
var fullnames = []string{}

// slice to store user data entries
var enteris = make([]UserData, 0)

// minimum length required for a name
var name_min_length = int(3)

// UserData represents information about a user.
type UserData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FullName  string `json:"fullName"`
	City      string `json:"city"`
}

// Greeting_msg displays a welcome message.
func Greeting_msg(phonebookmaxsize uint, appname string) {
	log.Infof("Welcome to the %s application\n", appname)
	log.Debugf("%d remaining places available in the %s", phonebookmaxsize, appname)
}

// Get_info retrieves user information from the user.
func Get_info(delimiter string) (string, string) {
	for {
		// setting up reader (stdin)
		reader := bufio.NewReader(os.Stdin)
		// asking for full name from user input
		log.Info("Enter your full name:")
		// reading fullname from stdin
		fullname, _ := reader.ReadString('\n')
		log.Debugf("Your fullname is: %s\n", fullname)
		// adding fullname to fullnames slice
		fullnames = append(fullnames, fullname)
		// splitting fullname into first and last name
		var names = strings.Fields(fullname)

		// setting firstName from names index 0
		firstName := names[0]

		// setting lastName from names index 1
		lastName := names[1]

		// checking if the length of firstName or lastName is less than name_min_length
		if len(firstName) < name_min_length {
			log.Errorf("First Name must be more than %d characters, please try again", name_min_length)
			continue
		}
		if len(lastName) < name_min_length {
			log.Errorf("Last Name must be more than %d characters, please try again", name_min_length)
			continue
		}
		log.Debugf("First Name is: %s \n", firstName)
		log.Debug(delimiter)
		log.Debugf("Last Name is: %s \n", lastName)
		log.Debug(delimiter)
		return firstName, lastName
	}
}

// Get_address retrieves user's address information from the user.
func Get_address(firstName string) (string, string) {
	// setting up reader (stdin)
	reader := bufio.NewReader(os.Stdin)
	log.Infof("Hey %s what is your address? (e.x 67 brown st)", firstName)
	address, _ := reader.ReadString('\n')
	log.Infof("City? (e.x Boston)")
	city, _ := reader.ReadString('\n')
	get_address_info := UserData{FirstName: firstName, LastName: "", FullName: ""}
	//city, _ := reader.ReadString('\n')
	//// capitalizing & trimming spaces using strings.TrimSpace & strings.Title functions
	//city = strings.Title(strings.TrimSpace(city))
	//switch city {
	//// checking if city is in the list of cities
	//case "New York", "Boston", "Los Angeles":
	//	log.Infof("%s is the best place in the US!\n", city)
	//	fallthrough
	//case "London":
	//	log.Infof("%s is pretty nice\n", city)
	//case "Tel Aviv", "Jerusalem":
	//	log.Infof("%s is the city of the Jews!\n", city)
	//case "Paris":
	//	log.Infof("%s has great food\n", city)
	//default:
	//	log.Infof("I don't know where %s is, but it sounds fun\n", city)
	//}

}

//func Get_country(city string) string {
//	// checking for country
//	ninjaapikey := os.Getenv("NINJA_API_KEY")
//	log.Debugf(ninjaapikey)
//	ninja_url := fmt.Sprintf("https://api.api-ninjas.com/v1/city?name=%v", city)
//	log.Debugf(ninja_url)
//
//	client := &http.Client{}
//	req, err := http.NewRequest("GET", ninja_url, nil)
//	if err != nil {
//		log.Error(err)
//	}
//	req.Header.Add("X-Api-Key", ninjaapikey)
//
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Error(err)
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Error(err)
//	}
//
//	// Parse JSON response
//	var cityInfo []map[string]interface{}
//	if err := json.NewDecoder(resp.Body).Decode(&cityInfo); err != nil {
//		log.Error(err)
//	}
//
//	// Extract country from JSON response
//	var country string
//	if len(cityInfo) > 0 {
//		countryValue, ok := cityInfo[0]["country"].(string)
//		if ok {
//			country = countryValue
//		}
//	}
//	log.Debugf(string(body))
//	log.Infof("Country: %s\n", country)
//	return country
//}
//
//func Get_zipcode(country string) (string, string) {
//	//checking for zip code
//	if country == "us" {
//		state_name := "MA"
//		log.Infof("Country is %s, setting state %s", country, state_name)
//	} else {
//		state_name := "null"
//		log.Infof("Country is %s, setting state %s", country, state_name)
//	}
//	zipcodebaseapikey := os.Getenv("apikey")
//	zipcodebase_url := fmt.Sprintf("https://app.zipcodebase.com/api/v1/code/city?apikey=%s&city=%s&state_code=NY&country=us: %d", zipcodebaseapikey, city)
//	resp, err := http.Get(zipcodebase_url)
//
//	if err != nil {
//		log.Errorf("Request Failed: %s", err)
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Errorf("Reading body failed: %s", err)
//	}
//	// Log the request body
//	bodyString := string(body)
//	log.Print(bodyString)
//	return zipcode
//}

// Store_info stores user information and prints JSON data.
func Store_info(firstName string, lastName string, city string) ([]string, []string) {
	// adding firstName to firstNameslist slice
	firstNameslist = append(firstNameslist, firstName)

	// adding lastName to lastNameslist slice
	lastNameslist = append(lastNameslist, lastName)

	// adding firstName to map firstnames
	firstnames["firstName"] = firstName
	lastnames["lastName"] = lastName

	// creating an instance of UserData struct
	userData := UserData{
		LastName:  lastName,
		FirstName: firstName,
		FullName:  firstName + " " + lastName,
		City:      city,
	}

	// converting userData struct to JSON
	jsonData, err := json.Marshal(userData)
	if err != nil {
		log.Errorf("Error: %v", err)
	}
	// print the JSON data
	log.Debugln("userData struct instance in json: ", string(jsonData))
	// append userData to enteris slice
	enteris = append(enteris, userData)

	// print maps
	log.Debugf("firstnames map is %s", firstnames)
	log.Debugf("lastnames map is %s", lastnames)
	// print slices
	log.Debugf("firstNameslist slice is %s", firstNameslist)
	log.Debugf("lastNameslist slice is %s", lastNameslist)
	log.Debugf("List of entries is %v\n", enteris)

	return firstNameslist, lastNameslist
}
