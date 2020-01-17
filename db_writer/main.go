package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

var port = flag.Int("port", 8080, "The port on which to serve http")

var isLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func main() {
	flag.Parse()
	http.HandleFunc("/number", handleNumber)
	http.HandleFunc("/letter", handleLetter)

	log.Printf("Serving the frontend on :%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func handleNumber(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error reading request body: %v\n", err)
		return
	}
	num, err := strconv.Atoi(string(body))
	if err != nil {
		fmt.Println("Fatal error!")
		os.Exit(0)
	}

	fmt.Printf("Received number: %d\n", num)
}

func handleLetter(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error reading request body: %v\n", err)
		return
	}
	letter := string(body)
	if len(letter) != 1 || !isLetter(letter){
		fmt.Println("Fatal error!")
		os.Exit(0)
	}

	fmt.Printf("Received letter: %s\n", letter)
}
