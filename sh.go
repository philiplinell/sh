package sh

import "log"

func FailOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
