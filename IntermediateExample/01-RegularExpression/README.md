#Go regular expressions

Go regular expressions tutorial shows how to parse text in Go using regular expressions.

##Regular expressions
Regular expressions are used for text searching and more advanced text manipulation. Regular expressions are built into tools including grep and sed, text editors including vi and emacs, programming languages including Go, Java, and Python.

Go has built-in API for working with regular expressions; it is located in regexp package.

A regular expression defines a search pattern for strings. It is used to match text, replace text, or split text. A regular expression may be compiled for better performance. The Go syntax of the regular expressions accepted is the same general syntax used by Perl, Python, and other languages.

##Regex examples
The following table shows a couple of regular expression strings.

| **Regex** | **Description**                                      |
|-----------|------------------------------------------------------|
| .         | Match any single character.                          |
| ?         | Match preceing element once or not at all.           |
| +         | Matches the preceding element once or more times.    |
| *         | Matches the preceding element zero or more times.    |
| ^         | Matches the starting position within the string.     |
| $         | Matches the ending position within the string.       |
| \|        | Alternation operator.                                |
| [abc]     | matches a or b or c.                                 |
| [a-c]     | Range, matches a or b or c.                          |
| [^abc]    | Negation, matches everything, except a or b or c.    |
| \s        | Matches white space character.                       |
| \w        | Matches  a word character; equivalent to [a-zA-Z_0-9]|

