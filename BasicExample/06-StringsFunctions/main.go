package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {

	exampleCompareStrings()
	exampleReplaceString()
	exampleIndexString()
	exampleCountingStrings()
	exampleUpperTitleCase()
	examplePrefixSubfix()
	exampleContainContainsRune()
	exampleSplit()
	exampleJoins()
}

// The Join function concatenates the elements of the slice argument to create a single string.
// We join words of a slice by a space character.
func exampleJoins() {
	words := []string{"an", "old", "falcon", "in", "the", "sky"}

	msg := strings.Join(words, " ")
	fmt.Println(msg)
}

// The Split function splits a slice into all substrings separated by the given separator and returns a slice of the substrings between those separators.
// We have a string of integer values, separated with comma. The string is cut into parts by the comma.
// The string parts are converted to integers with strconv.Atoi and the integers are summed.
func exampleSplit() {
	msg := "3,4,5,6,7,8,9,10,11"

	data := strings.Split(msg, ",")
	fmt.Printf("%v\n", data)

	var sum = 0
	for _, e := range data {

		val, err := strconv.Atoi(e)

		if err != nil {
			log.Fatal(err)
		}

		sum += val
	}

	fmt.Println(sum)
}

// The Contains function checks if the given substring is present in the string.
// The ConstainsRune checks if the Unicode code point is in the strings.
func exampleContainContainsRune() {
	msg := "a blue 🐋"
	r := '🐋'

	if strings.ContainsRune(msg, r) {

		fmt.Println("yes")
	} else {

		fmt.Println("no")
	}

	fmt.Println("-----------------")

	if strings.Contains(msg, "🐋") {

		fmt.Println("yes")
	} else {

		fmt.Println("no")
	}
}

// The HasPrefix function checks whether the string begins with the given prefix.
// The HasSufffix function checks whether the string ends with the given suffix.
func examplePrefixSubfix() {
	words := []string{"sky", "lot", "car", "wood", "cloud",
		"cup", "war", "wind", "near", "tell", "cheer", "coin", "book"}

	for _, word := range words {

		if strings.HasPrefix(word, "c") {

			fmt.Println(word)
		}
	}

	fmt.Println("----------------------")

	for _, word := range words {

		if strings.HasSuffix(word, "r") {

			fmt.Println(word)
		}
	}
}

// The ToLower function returns a lowercased copy of the string, while the ToUpper returns an uppercased string.
// The Title function returns a titlecased copy of the given string (only the first letter is uppercased).
func exampleUpperTitleCase() {
	msg := "and old falcon"
	msg2 := "čerešňa"

	fmt.Println("Title of", msg, "is", strings.Title(msg))
	fmt.Println("Upper of", msg, "is", strings.ToUpper(msg))
	fmt.Println("Upper of", msg2, "is", strings.ToUpper(msg2))
	fmt.Println("Title of", msg2, "is", strings.Title(msg2))
}

// The Count function counts the number of substrings found in a string.
func exampleCountingStrings() {
	word := "wood"

	c1 := "o"
	c2 := "w"

	n1 := strings.Count(word, c1)
	fmt.Printf("# of %s in %s: %d\n", c1, word, n1)

	n2 := strings.Count(word, c2)
	fmt.Printf("# of %s in %s: %d\n", c2, word, n2)
}

// The Index function returns the index of the first substring found, while the LastIndex finds the last index.
func exampleIndexString() {
	msg := "I saw a ox in the forest. The fox had brown fur. I like foxes."

	idx1 := strings.Index(msg, "fox")
	fmt.Println("First index string", idx1)

	idx2 := strings.LastIndex(msg, "fox")
	fmt.Println("Last index string", idx2)
}

// The Replace function returns a copy of the string with the first n occurrences of the string replaced, while the ReplaceAll returns a copy where all occurrences are replaced.
func exampleReplaceString() {
	msg := "I saw a fox in the forest. The fox had brown fur."

	output := strings.Replace(msg, "fox", "wolf", 2)
	fmt.Println(output)

	output2 := strings.ReplaceAll(msg, "fox", "wolf")
	fmt.Println(output2)
}

// The Compare function compare two strings lexicographically. To compare two strings in a case-insensitive manner, we use the EqualFold function.
func exampleCompareStrings() {
	w1 := "pablo"
	w2 := "Pablo"

	if strings.Compare(w1, w2) == 0 {

		fmt.Println("The words are equal")
	} else {

		fmt.Println("The words are not equal")
	}

	if strings.EqualFold(w1, w2) {

		fmt.Println("The words are equal")
	} else {

		fmt.Println("The words are not equal")
	}
}
