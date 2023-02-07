package check_error

import "log"

func Check(e error) {
	if e != nil {
		log.Println("Error: " + e.Error())
		panic(e)
	}
}
