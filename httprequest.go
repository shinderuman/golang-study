package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer resp.Body.Close()
	contents, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Printf("%s", err2)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents))
}
