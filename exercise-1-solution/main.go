package main

import (
	"fmt"

	"github.com/EnricoPicci/go-class-find-most-frequent-word/helpers"
)

func main() {
	file := helpers.ReadFirstCmdLineArg()

	mostFrequent := mostFrequentWordWithCount(file)

	fmt.Printf("The most frequent word is ==>> \"%v\" \n", mostFrequent.Word)
	fmt.Printf("It is found %v times in the file %v \n", mostFrequent.Occurrencies, file)
}
