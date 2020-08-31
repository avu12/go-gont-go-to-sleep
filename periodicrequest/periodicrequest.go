package periodicrequest

import (
	"fmt"
	"net/http"
	"time"
)

//CreateRequest send a request with the given method to the targeted url in the given interval (in minutes )
func CreateRequest(interval int, url string, method string) {
	for range time.NewTicker(time.Minute * time.Duration(interval)).C {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			fmt.Println("ERROR in request!")
		}
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("ERROR in self resp!")
		}
		resp.Body.Close()
	}
}
