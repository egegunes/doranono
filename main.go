package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func main() {
	resp, err := http.Get("http://nonolive.com/dora")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	re := regexp.MustCompile(`"anchor_intro":"([^"]*)"`)
	match := re.FindStringSubmatch(string(body))

	fmt.Println(match[1])
}
