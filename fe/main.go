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
	"strings"
)

const (
	svcLetters = "letters"
	svcNumbers = "numbers"
	svcDBWriter = "db-writer"
)

var serviceOwner = flag.String("owner", "", "If specified, abc123 will look for "+
	"hostnames prefixed with this label")
var numbersHost = flag.String("numbers-host", svcNumbers, "The host for the numbers service")
var lettersHost = flag.String("letters-host", svcLetters, "The host for the letters service")
var dbWriterHost = flag.String("db-writer-host", svcDBWriter, "The host for the db_writer service")
var port = flag.Int("port", 8080, "The port on which to serve http")


type Info struct {
	Letter string
	Number string
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/get_rand", handleTextOnly)

	log.Printf("Serving the frontend on :%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		return
	}

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
	resp, err := getRequest(*lettersHost)
	if err != nil {
		return "", err
	}
	letter := strings.TrimSpace(string(resp))
	postRequest(*dbWriterHost, "letter", letter)
	return letter, nil
}

func getNumber() (string, error) {
	resp, err := getRequest(*numbersHost)
	if err != nil {
		return "", err
	}

	num := strings.TrimSpace(string(resp))
	postRequest(*dbWriterHost, "number", num)
	return num, nil
}

func getRequest(host string) ([]byte, error) {
	url := urlForService(host)
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("request to %s failed: %v", url, err)
	}

	return ioutil.ReadAll(resp.Body)
}

func postRequest(host, path, body string) {
	url := urlForService(host)
	url = url + "/" + path

	http.Post(url, "text/plain", strings.NewReader(body))
}

func urlForService(host string) string {
	url := fmt.Sprintf("http://%s", host)
	if *serviceOwner != "" {
		url = fmt.Sprintf("http://%s-%s", *serviceOwner, host)
	}
	return url
}
