package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	s := make([]string, 0)
	results := make([]map[string]string, 0)
	s = append(s, "http://www.testingmcafeesites.com/index.html")
	s = append(s, "http://www.neverssl.com/")
	s = append(s, "http://www.testingmcafeesites.com/testcat_al.html")
	s = append(s, "http://www.testingmcafeesites.com/testcat_an.html")
	s = append(s, "http://www.testingmcafeesites.com/testcat_au.html")
	s = append(s, "http://www.testingmcafeesites.com/testcat_be.html")

	for _, value := range s {
		start := time.Now()
		res, err := http.Get(value)
		delta := time.Since(start)

		result := make(map[string]string)
		result["url"] = value
		result["status"] = fmt.Sprintf("%d", res.StatusCode)
		result["time"] = delta.String()

		if err != nil {
			result["error"] = err.Error()
			log.Fatalf("error: %s url: %s\n", err, value)
		}

		results = append(results, result)

		defer res.Body.Close()
	}
	for _, r := range results {
		log.Println(r)
	}
}
