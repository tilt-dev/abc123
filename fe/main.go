package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	svcLetters = "letters"
	svcNumbers = "numbers"
)

var serviceOwner = flag.String("owner", "", "If specified, abc123 will look for "+
	"hostnames prefixed with this label")
var numbersHost = flag.String("numbers-host", svcNumbers, "The host for the numbers service")
var lettersHost = flag.String("letters-host", svcLetters, "The host for the letters service")
var port = flag.Int("port", 8080, "The port on which to serve http")


type Info struct {
	Letter string
	Number int
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/get_rand", handleTextOnly)

	log.Printf("Serving the frontend on :%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("> Getting a letter and a number!")
	t, err := template.ParseFiles(templatePath("index.tpl"))
	if err != nil {
		log.Fatalf("error parsing template: %v\n", err)
	}

	let, err := getLetter()
	if err != nil {
		log.Printf("error getting letter: %v\n", err)
		http.Error(w, "error getting letter", http.StatusInternalServerError)
		return
	}

	num, err := getNumber()
	if err != nil {
		log.Printf("error getting number: %v\n", err)
		http.Error(w, "error getting number", http.StatusInternalServerError)
		return
	}

	info := Info{
		Letter: let,
		Number: num,
	}

	err = t.Execute(w, info)
	if err != nil {
		log.Fatalf("error executing template: %v\n", err)
	}
}

func templatePath(f string) string {
	dir := os.Getenv("TEMPLATE_DIR")
	if dir == "" {
		dir = "web/templates"
	}

	return filepath.Join(dir, f)
}

func handleTextOnly(w http.ResponseWriter, r *http.Request) {
	let, err := getLetter()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error getting letter: %v\n", err)
		return
	}

	num, err := getNumber()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error getting number: %v\n", err)
		return
	}

	fmt.Fprintf(w, "Random letter \"%s\"; random number \"%d\"\n", let, num)

}

func getLetter() (string, error) {
	resp, err := pingPortForService(*lettersHost)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(resp)), nil
}

func getNumber() (int, error) {
	resp, err := pingPortForService(*numbersHost)
	if err != nil {
		return 0, err
	}

	s := strings.TrimSpace(string(resp))
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func pingPortForService(host string) ([]byte, error) {
	url := fmt.Sprintf("http://%s", host)
	if *serviceOwner != "" {
		url = fmt.Sprintf("http://%s-%s", *serviceOwner, host)
	}
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("request to %s failed: %v", url, err)
	}

	return ioutil.ReadAll(resp.Body)
}
