package main
import (
 "fmt"
 "github.com/hackebrot/turtle"
)

// Using third-party packages

func main() {
	emoji, ok := turtle.Emojis["smiley"]
	if !ok {
		fmt.Println("No emoji found.")
	} else {
		fmt.Println(emoji.Char)
	}
}