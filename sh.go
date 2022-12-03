package sh

import "log"

// FailOnErr will log with a log.Fatal if there is any errors (meaning it will
// directly exit the routine with a exit code 1).
func FailOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
