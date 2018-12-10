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

var serviceOwner = flag.String("owner", "", "If specified, abc123 will look for "+
	"service names prefixed with this label")

const (
	svcLetters = "letters"
	svcNumbers = "numbers"
)

type Info struct {
	Letter string
	Number int
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/get_rand", handleTextOnly)

	log.Println("Serving the fontend on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(templatePath("index.tpl"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error parsing template: %v\n", err)
		return
	}

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

	info := Info{
		Letter: let,
		Number: num,
	}

	err = t.Execute(w, info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error executing template: %v\n", err)
		return
	}
}

func templatePath(f string) string {
	dir := os.Getenv("TEMPLATE_DIR")
	if dir == "" {
		dir = "fortune/web/templates"
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
	resp, err := pingPortForService(svcLetters)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(resp)), nil
}

func getNumber() (int, error) {
	resp, err := pingPortForService(svcNumbers)
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

func pingPortForService(svcName string) ([]byte, error) {
	url := fmt.Sprintf("http://%s", svcName)
	if *serviceOwner != "" {
		url = fmt.Sprintf("http://%s-%s", *serviceOwner, svcName)
	}
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("request to %s failed: %v", url, err)
	}

	return ioutil.ReadAll(resp.Body)
}
