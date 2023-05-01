package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strings"
)

// need a data structure to hold all of the URLs

/*
	Web Crawler
	- Given a URL, parse the page and find any linked web pages, give the option to crawl them as well
	- Grab network connections made by the website
	- Grab cookies made by the website

	- Reference crawling
		- Allow them to select the depth of the crawl?
		- Allow them to go down a path of selections?
		- Allow them to print the HTTP data?

*/

func main() {

	m := make(map[string]int) // reference type

	url := get_user_url()

	max_depth := get_max_depth()

	fmt.Println("Searching", url, "to a depth of", max_depth, "\n###################################################")

	find_refs(url, url, max_depth, m)

	sorted := sort_popular(m)

	fmt.Println("###################################################\nFrequency of referenced websites found:")

	for _, p := range sorted {
		fmt.Println(p.Val, p.Key)
	}
}

// implement sort interface
type Pair struct {
	Key string
	Val int
}
type PairList []Pair

func (pl PairList) Len() int           { return len(pl) }
func (pl PairList) Less(i, j int) bool { return pl[i].Val < pl[j].Val }
func (pl PairList) Swap(i, j int)      { pl[i], pl[j] = pl[j], pl[i] }

func sort_popular(m map[string]int) PairList {
	// given map of urls and counts, sort by count, insertion sort?
	pl := make(PairList, len(m))
	i := 0
	for k, v := range m {
		pl[i] = Pair{k, v}
		i += 1
	}
	sort.Sort(sort.Reverse(pl))

	return pl
}

// gathers initial user input
func get_user_url() string {
	var input string
	fmt.Println("Enter a HTTP URL to crawl")

	fmt.Scanln(&input) // get the URL

	return parse_input(input)
}

// depth to search to
func get_max_depth() int {
	var input int
	fmt.Println("Enter a depth to search to")

	fmt.Scanln(&input)

	if input < 1 {
		log.Fatal(0)
	}
	return input
}

// TODO check user input rigourisly, allow for http or no http, etc.
// parse the user input, making sure it is valid, returning the URL to crawl
func parse_input(url string) string {
	return "http://" + url
}

// TODO create dynamic buffer to take in all data, not just up to buffer size
// fetches HTTP data from URL, returning it as a string
func grab_data(url string) string {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("URL with error:", url)
		return "-1"
	}

	BYTE_LENGTH := 1024 * 10
	bytes := make([]byte, BYTE_LENGTH)
	res.Body.Read(bytes)

	return string(bytes)
}

func find_refs(url string, base string, max_depth int, ref map[string]int) {
	depth := 0 // start at base depth

	ref_search(url, base, depth, max_depth, ref) // recursively searches to a max depth or until complete

}

func ref_search(url string, base string, depth int, max_depth int, ref map[string]int) {

	content := grab_data(url)
	if depth >= max_depth || content == "-1" {
		return
	}
	depth += 1

	re := regexp.MustCompile(`href="(.+)"`) // faster using compiled

	var found []string

	content_list := strings.Split(content, " ")

	for _, word := range content_list {
		result := re.FindStringSubmatch(word)
		if result != nil {
			match := result[1]
			if match[0] == '/' || match[0] == '.' || match[0] == '_' {
				match = base + match
			}
			found = append(found, match)

			_, ok := ref[match] // two-value assignment
			if ok {
				ref[match] += 1
			} else {
				ref[match] = 1
				ref_search(match, base, depth, max_depth, ref) // recursive call, do this with threads, need locks

			}

		}
	}
}