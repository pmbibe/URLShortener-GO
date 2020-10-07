package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
)

var data string
var dataYAML convertURL

type convertURL struct {
	Convert map[string]string
}

func readFile(file string) string {
	dat, _ := ioutil.ReadFile(file)
	return string(dat)
}

func getKeyValue(data string, dataYAML convertURL) convertURL {
	err := yaml.Unmarshal([]byte(data), &dataYAML.Convert)
	if err != nil {
		fmt.Print("Can not Unmarshal")
	}
	return dataYAML
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
func getLongURL(dataYAML convertURL) []string {
	var listURL []string
	for key := range dataYAML.Convert {
		listURL = append(listURL, key)
	}
	return listURL
}

func getShortURL(longURL string, dataYAML convertURL) string {
	return dataYAML.Convert[longURL]
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
	_, found := Find(getLongURL(getKeyValue(data, dataYAML)), r.URL.Path)
	if !found {
		http.Redirect(w, r, "https://github.com/pmbibe", 301)
	} else {
		http.Redirect(w, r, getShortURL(r.URL.Path, getKeyValue(data, dataYAML)), 301)
	}

}

func main() {

	data = readFile("mapURL")
	getKeyValue(data, dataYAML)
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", nil)

}
