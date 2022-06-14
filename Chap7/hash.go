// Hash table construction
package main

import (
	"fmt"
	"hash/fnv" // Fowler–Noll–Vo algorithm
	"strconv"
	"time"
)

const tableSize = 100_000

var length int

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

type WordType struct {
	word string
	list []string
}

// At every index there is a slice of words
type HashTable [tableSize]WordType

func NewTable() HashTable {
	var table HashTable
	for i := 0; i < tableSize; i++ {
		table[i] = WordType{"", []string{}}
	}
	return table
}

// Methods
func (table *HashTable) Insert(word string) {
	index := hash(word) % tableSize // Between 0 and tableSize - 1
	// Search table[index] for word
	if table[index].word == word {
		return // duplicates not allowed
	}
	if len(table[index].list) > 0 {
		for i := 0; i < len(table[index].list); i++ {
			if table[index].list[i] == word {
				return // duplicates not allowed
			}
		}
	}
	if table[index].word == "" {
		table[index].word = word
	} else {
		table[index].list = append(table[index].list, word)
	}
	length += 1
}

func (table HashTable) IsPresent(word string) bool {
	index := hash(word) % tableSize // Between 0 and tableSize - 1
	// Search table[index] for word
	if table[index].word == word {
		return true
	}
	if len(table[index].list) > 0 {
		for i := 0; i < len(table[index].list); i++ {
			if table[index].list[i] == word {
				return true
			}
		}
	}
	return false
}

func main() {
	myTable := NewTable()
	mapCollection := make(map[string]string)

	words := []string{}
	for i := 0; i < 500_000; i++ {
		word := strconv.Itoa(i)
		words = append(words, word)
		myTable.Insert(word)
		mapCollection[word] = ""
	}

	fmt.Println("Benchmark test begins to test words: ", length)
	start := time.Now()
	for i := 0; i < length; i++ {
		if myTable.IsPresent(words[i]) == false {
			fmt.Println("Word not found in table: ", words[i])
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Time to test all words in myTable: ", elapsed)

	start = time.Now()
	for i := 0; i < len(mapCollection); i++ {
		_, present := mapCollection[words[i]]
		if !present {
			fmt.Println("Word not found in mapCollection: ", words[i])
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Time to test words in mapCollection: ", elapsed)
}

/* Output
Benchmark test begins to test words:  500000
Time to test all words in myTable:  1m17.880336666s
Time to test words in mapCollection:  24.405583ms
*/
