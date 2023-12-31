package main

import (
	log "github.com/sirupsen/logrus"
	"main/helper"
)

// setting up global variables
var appname = "phonebook"
var phonebookmaxsize = uint(2)
var delimiter = "---------------"

type msg struct {
	a, b string
}

func (m msg) Greeting_msg(phone uint, app string) {
	m.a = "aa"
	helper.Greeting_msg(phone, app)
}

type msgI interface {
	Greeting_msg(uint, string)
}

func main() {
	// setting log level to Info
	log.SetLevel(log.DebugLevel)
	var myMsgI msgI
	myMsgI := msg2{}

	mm := msgI{}
	myMsgI.Greeting_msg()
	for {
		// only running if phonebookmaxsize is bigger than 0
		if phonebookmaxsize > 0 {
			log.Infof(delimiter)
			// greeting message
			helper.Greeting_msg(phonebookmaxsize, appname)
			// getting first, last name from user
			firstName, lastName := helper.Get_info(delimiter)
			// getting city from user
			city := helper.Get_city(firstName)
			// storing info in phonebook
			helper.Store_info(firstName, lastName, city)
			//get_gender(get_firstname(firstNameslist, lastNameslist, fullnames, delimiter))

		} else { // if phonebookmaxsize is 0 telling the user that there are no spots avilable
			log.Warnf("The is %d spots avilable, try again later", phonebookmaxsize)
			break
		}
		// subtracting 1 from phonebookmaxsize variable
		phonebookmaxsize = phonebookmaxsize - 1
		log.Infof("Your number of avilable spots in the %s are: %d \n", appname, phonebookmaxsize)
	}
}
