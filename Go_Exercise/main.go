package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://renewing-grackle-95.hasura.app/v1/graphql"
	method := "POST"

	payload := strings.NewReader("{\"query\":\"query { \\r\\n\\ttasks{\\r\\n    id\\r\\n    title\\r\\n    user{\\r\\n      name\\r\\n    }\\r\\n  }\\r\\n}\",\"variables\":{}}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-hasura-admin-secret", "PUHYn6spLASCAVhjp8ZruDJJVcTKM9yKB6yp3yzfXNZaSjrVvD8YMPzbFPqjSUt0")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
