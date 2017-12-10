package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/tkajder/adventofcode17/fileutils"
)

func verifyPassphraseNoDuplicates(passphrase []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range passphrase {
		wordSet[word] = true
	}

	return len(wordSet) == len(passphrase)
}

func verifyPassphraseNoAnagrams(passphrase []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range passphrase {
		runes := []rune(word)
		sort.Slice(runes, func(i int, j int) bool {
			return runes[i] < runes[j]
		})

		wordSet[string(runes)] = true
	}

	fmt.Println(wordSet)

	return len(wordSet) == len(passphrase)
}

func main() {
	fileNamePtr := flag.String("file", "", "Required: file containing inverse captcha")
	flag.Parse()

	// Check if file provided
	if *fileNamePtr == "" {
		flag.Usage()
		os.Exit(1)
	}

	var numValidNoDuplicates uint64
	var numValidNoAnagrams uint64

	lines, errc := fileutils.ByLine(*fileNamePtr)
	for line := range lines {
		passphrase := strings.Split(line, " ")

		if verifyPassphraseNoDuplicates(passphrase) {
			numValidNoDuplicates++
		}

		if verifyPassphraseNoAnagrams(passphrase) {
			numValidNoAnagrams++
		}
	}
	if err := <-errc; err != nil {
		log.Fatal(err)
	}

	fmt.Println(numValidNoDuplicates)
	fmt.Println(numValidNoAnagrams)
}
