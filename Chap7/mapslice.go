// We compare dictionary lookup using slice versus map
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

var mapCollection map[string]string

var sliceCollection []string

func IsPresent(word string, sliceCollection []string) bool {
	for i := 0; i < len(sliceCollection); i++ {
		if sliceCollection[i] == word {
			return true
		}
	}
	return false
}

func IsPresentBinarySearch(word string, sliceCollection []string) bool {
	low := 0
	high := len(sliceCollection) - 1

	for low <= high {
		median := (low + high) / 2

		if sliceCollection[median] < word {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low == len(sliceCollection) || sliceCollection[low] != word {
		return false
	}
	return true
}

func main() {
	file, err := os.Open("words.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}

	// Fill mapCollection and sliceConnection with words
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	mapCollection = make(map[string]string)
	sliceCollection = make([]string, 1)

	var words []string

	for scanner.Scan() {
		word := scanner.Text()
		words = append(words, word)
		mapCollection[word] = word
		sliceCollection = append(sliceCollection, word)
	}

	// Benchmark time to test for presence of each word in mapCollection
	start := time.Now()
	for i := 0; i < len(words); i++ {
		_, present := mapCollection[words[i]]
		if !present {
			fmt.Println("Word not found in mapCollectio0n")
		}
	}
	elapsed := time.Since(start)

	fmt.Println("Number of words in mapCollection: ", len(mapCollection))
	fmt.Println("\nTime to test words in mapCollection: ", elapsed)

	sort.Strings(sliceCollection)

	// Benchmark time to test for presence of each word in sliceCollection
	start = time.Now()
	for i := 0; i < len(sliceCollection); i++ {
		if !IsPresent(sliceCollection[i], sliceCollection) {
			fmt.Println("Word not found in mapCollectio0n")
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Time to test words in sliceCollection: ", elapsed)

	// Benchmark time to test for presence of each word in sorted sliceCollection
	start = time.Now()
	for i := 0; i < len(sliceCollection); i++ {
		if !IsPresentBinarySearch(sliceCollection[i], sliceCollection) {
			fmt.Println("Word not found in mapCollectio0n")
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Time to test words in sorted sliceCollection: ", elapsed)

}

/* Output
Number of words in mapCollection:  466468

Time to test words in mapCollection:  29.022542ms
Time to test words in sliceCollection:  2m20.874580833s
Time to test words in sorted sliceCollection:  51.836708ms
*/
