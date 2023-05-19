package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// const url = "http://services.explorecalifornia.org/json/tours.php"
const url = "https://reqres.in/api/products"

func main() {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close()

	fmt.Println(resp.Status)

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	fmt.Print(content)
}
