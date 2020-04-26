// IRC quotes handling.
// Author:
//    Andrea Cervesato <andrea.cervesato@mailbox.org>

package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

// Quotes file rapresentation.
type IRCQuotes struct {
	Location string
	Lines    []string
}

// Read a random quote inside quotes file.
func (quotes IRCQuotes) RandQuote() string {
	rand.Seed(time.Now().UnixNano())

	count := len(quotes.Lines)
	randline := rand.Intn(count)
	quote := quotes.Lines[randline]

	return quote
}

// Create a new quotes file object.
func NewIRCQuotes(location string) *IRCQuotes {
	quotes := IRCQuotes{
		Location: location,
		Lines:    make([]string, 0),
	}

	// open file
	file, err := os.Open(location)
	if err != nil {
		log.Fatal(err)
	}

	// always close quotes file
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// read all quotes
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		quotes.Lines = append(quotes.Lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &quotes
}
