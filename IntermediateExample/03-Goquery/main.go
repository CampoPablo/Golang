package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	webPage := "http://yahoo.com"

	getAndPrintTitlePage(webPage)

	printTagContext("example.html", "h1,p")

	filterWords("example.html", "P")

	getLinks(webPage)
}

// getLinks retrieves external links to secured web pages.
func getLinks(page string) {
	webPage := page

	resp, err := http.Get(webPage)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	//In the predicate function we ensure that the link has the https prefix.
	f := func(i int, s *goquery.Selection) bool {
		link, _ := s.Attr("href")
		return strings.HasPrefix(link, "https")
	}

	//We find all the anchor tags, filter them, and then print the filtered links to the console.
	doc.Find("body a").FilterFunction(f).Each(func(_ int, tag *goquery.Selection) {

		link, _ := tag.Attr("href")
		linkText := tag.Text()
		fmt.Printf("%s %s\n", linkText, link)
	})
}

// filterWords retrieve all words starting with param filter in the li tag.
func filterWords(htmlFile string, filter string) {
	//We read the file
	data, err := ioutil.ReadFile(htmlFile)

	if err != nil {
		log.Fatal(err)
	}

	//We generate a new goquery document with NewDocumentFromReader.
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))

	if err != nil {
		log.Fatal(err)
	}

	//We retrieve all words starting with filter param.
	f := func(i int, sel *goquery.Selection) bool {
		return strings.HasPrefix(sel.Text(), filter)
	}

	var words []string

	// We locate the set of matching tags with Find. We filter the set with FilterFunction and go over the filtered results with Each. We add each filtered word to the words slice.
	doc.Find("li").FilterFunction(f).Each(func(_ int, sel *goquery.Selection) {
		words = append(words, sel.Text())
	})

	fmt.Println(words)
}

func printTagContext(htmlFile string, htmlTags string) {
	//We read the file
	data, err := ioutil.ReadFile(htmlFile)

	if err != nil {
		log.Fatal(err)
	}

	//We generate a new goquery document with NewDocumentFromReader.
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))

	if err != nil {
		log.Fatal(err)
	}

	//We get the text contents of two tags: h1 and p.
	text := doc.Find(htmlTags).Text()

	//Using a regular expression, we remove excessive white space.
	re := regexp.MustCompile("\\s{2,}")
	fmt.Println(re.ReplaceAllString(text, "\n"))

}

// getAndPrintTitlePage gets the html of the web page given the parameter, and prints the title of the site.
func getAndPrintTitlePage(webPage string) {
	resp, err := http.Get(webPage)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find("title").Text()
	fmt.Println(title)
}
