package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	s := make([]string, 0)
	s = append(s, "http://www.testingmcafeesites.com/index.html")
	s = append(s, "http://www.testingmcafeesites.com/testcat_ac.html")
	s = append(s, "http://www.testingmcafeesites.com/testcat_al.html")
	s = append(s, "http://www.testingmcafeesites.com/testcat_an.html")
	s = append(s, "http://www.testingmcafeesites.com/testcat_au.html")
	s = append(s, "http://www.testingmcafeesites.com/testcat_be.html")

	for _, value := range s {
		res, err := http.Get(value)
		if err != nil {
			log.Printf("error making http request: %s\n", err)
			os.Exit(1)
		}
		log.Print(res.StatusCode)
	}
}
