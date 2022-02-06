package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]

queries:
	for _, q := range query {
	search:
		for i, w := range words {
			switch q {
			case "and", "or", "the":
				break search
			}
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				continue queries
			}
		}
	}

}

// for loop with break statement
// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// const corpus = "lazy cat jumps again and again and again"

// func main() {
// 	words := strings.Fields(corpus)
// 	query := os.Args[1:]

// 	for _, q := range query {
// 		for i, w := range words {
// 			if q == w {
// 				fmt.Printf("#%-2d: %q\n", i+1, w)
// 				break
// 			}
// 		}
// 	}

// }
