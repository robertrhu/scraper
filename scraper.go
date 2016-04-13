package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	// Open a file
	file, err := os.Open("elks_urls.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// os.Mkdir("data/elks", 0644)

	//For each line
	scanner := bufio.NewScanner(file)
	var index = 0
	for scanner.Scan() {
		url := scanner.Text()
		fetch(index, url)
		index++
	}

}

//fetch a url
func fetch(index int, url string) {
	//Choose a filename
	filename := fmt.Sprintf("%v.html", index)
	fmt.Println("fetching", url)

	// GET the url
	r, err := http.Get(url)
	defer r.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("OK", url)

	// Read all the body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error Reading Body")
	}

	// Write all the body file
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		fmt.Println("ERROR Writing", filename)
	}

}
