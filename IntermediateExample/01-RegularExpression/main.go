package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	exampleMatch()
	fmt.Println(strings.Repeat("-", 50))
	exampleCompile()
	fmt.Println(strings.Repeat("-", 50))
	exampleFindAllStrings()
	fmt.Println(strings.Repeat("-", 50))
	exampleFindAllStringIndex()
	fmt.Println(strings.Repeat("-", 50))
	exampleSplit()
	fmt.Println(strings.Repeat("-", 50))
	exampleGroup()
	fmt.Println(strings.Repeat("-", 50))
	exampleReplaceAllString()
	fmt.Println(strings.Repeat("-", 50))
	exampleReplaceAllStringFunc()
}

//The ReplaceAllStringFunc returns a copy of a string in which all matches of the regular expression have been replaced by the return value of the specified function.
func exampleReplaceAllStringFunc() {
	content := "an old eagle"

	re := regexp.MustCompile(`[^aeiou]`)

	//we apply the strings.ToUpper function on all wovels of a string.
	fmt.Println(re.ReplaceAllStringFunc(content, strings.ToUpper))
}

//It is possible to replace strings with ReplaceAllString. The method returns the modified string.
func exampleReplaceAllString() {

	//The example reads HTML data of a web page and strips its HTML tags using a regular expression.
	resp, err := http.Get("http://webcode.me")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		log.Fatal(err)
	}

	content := string(body)

	//We read the body of the response object.
	re := regexp.MustCompile("<[^>]*>")
	//This pattern defines a regular expression that matches HTML tags.
	replaced := re.ReplaceAllString(content, "")

	fmt.Println(strings.TrimSpace(replaced))
}

//Round brackets () are used to create capturing groups. This allows us to apply a quantifier to the entire group or to restrict alternation to a part of the regular expression.
//To find capturing groups (Go uses the term subexpressions), we use the FindStringSubmatch function.
func exampleGroup() {
	websites := [...]string{"webcode.me", "zetcode.com", "freebsd.org", "netbsd.org", ""}

	//we divide the domain names into two parts by using groups.
	re := regexp.MustCompile("(\\w+)\\.(\\w+)")

	for _, website := range websites {

		//We define two groups with parentheses. The FindStringSubmatch returns a slice of strings holding the matches, including those from the capturing groups.
		parts := re.FindStringSubmatch(website)

		for i := range parts {
			fmt.Println(parts[i])
		}

	}
}

//The Split function cuts a string into substrings separated by the defined regular expression. It returns a slice of the substrings between those expression matches.
func exampleSplit() {
	var data = `22, 1, 3, 4, 5, 17, 4, 3, 21, 4, 5, 1, 48, 9, 42, 1`

	sum := 0

	re := regexp.MustCompile(",\\s*")

	vals := re.Split(data, -1)

	for _, val := range vals {

		n, err := strconv.Atoi(val)

		sum += n

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(sum)
}

//The FindAllStringIndex returns a slice of all successive indexes of matches of the expression.
func exampleFindAllStringIndex() {
	var content = `Foxes are omnivorous mammals belonging to several genera 
of the family Canidae. Foxes have a flattened skull, upright triangular ears, 
a pointed, slightly upturned snout, and a long bushy tail. Foxes live on every 
continent except Antarctica. By far the most common and widespread species of 
fox is the red fox.`

	re := regexp.MustCompile("(?i)fox(es)?")

	idx := re.FindAllStringIndex(content, -1)

	for _, j := range idx {
		match := content[j[0]:j[1]]
		fmt.Printf("%s at %d:%d\n", match, j[0], j[1])
	}
}

//The FindAllString function returns a slice of all successive matches of the regular expression.
func exampleFindAllStrings() {
	var content = `Foxes are omnivorous mammals belonging to several genera 
of the family Canidae. Foxes have a flattened skull, upright triangular ears, 
a pointed, slightly upturned snout, and a long bushy tail. Foxes live on every 
continent except Antarctica. By far the most common and widespread species of 
fox is the red fox.`

	//we find all occurrences of the word fox, including its plural form.
	//With the (?i) syntax, the regular expression is case insensitive. The (es)? indicates that "es" characters might be included zero times or once.
	re := regexp.MustCompile("(?i)fox(es)?")

	//We look for all occurrences of the defined regular expression with FindAllString. The second parameter is the maximum matches to look for;
	//-1 means search for all possible matches.
	found := re.FindAllString(content, -1)

	fmt.Printf("%q\n", found)

	if found == nil {
		fmt.Printf("no match found\n")
		os.Exit(1)
	}

	for _, word := range found {
		fmt.Printf("%s\n", word)
	}
}

//The Compile function parses a regular expression and returns, if successful, a Regexp object that can be used to match against text.
//Compiled regular expressions yield faster code.
//The MustCompile function is a convenience function which compiles a regular expression and panics if the expression cannot be parsed.
func exampleCompile() {
	words := [...]string{"Seven", "even", "Maven", "Amen", "eleven", "Pablo", "Mark"}

	re, err := regexp.Compile(".even")

	if err != nil {
		log.Fatal(err)
	}

	for _, word := range words {

		found := re.MatchString(word)

		if found {

			fmt.Printf("%s matches\n", word)
		} else {

			fmt.Printf("%s does not match\n", word)
		}
	}
}

//The MatchString function reports whether a string contains any match of the regular expression pattern.
func exampleMatch() {
	words := [...]string{"Seven", "even", "Maven", "Amen", "eleven", "Mark", "Paul"}

	for _, word := range words {

		found, err := regexp.MatchString(".even", word)

		if err != nil {
			log.Fatal(err)
		}

		if found {

			fmt.Printf("%s matches\n", word)
		} else {

			fmt.Printf("%s does not match\n", word)
		}
	}
}
