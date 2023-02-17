// Solution to the Exercize 2

// To build and execute (with a sample file) this program run the following 2 commands:
// go build -o ./bin/mostFrequentWord_Ex_2-sol ./exercise-2-solution
// ./bin/mostFrequentWord_Ex_2-sol testdata/count-words/file-with-five-words.txt

package main

import (
	"fmt"

	"github.com/EnricoPicci/go-class-find-most-frequent-word/helpers"
)

func main() {
	file := helpers.ReadFirstCmdLineArg()

	longest := longestWord(file)
	fmt.Printf("The longest word in file \"%v\" is ==>> \"%v\" \n", file, longest)
}
