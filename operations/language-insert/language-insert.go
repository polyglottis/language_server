// Package language-insert contains the language-insert executable.
package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/polyglottis/language_server/operations"
	"github.com/polyglottis/platform/config"
	"github.com/polyglottis/platform/language"
)

var inputFile = flag.String("in", "", "Input file (tab-delimited, columns: [code, iso639-1, iso639-3, wikidata, comment])")

func main() {
	flag.Parse()

	conf := config.Get()

	if *inputFile == "" {
		flag.Usage()
		log.Fatalln("Input file is mandatory")
	}

	lines, err := readLines(*inputFile)
	if err != nil {
		log.Fatalln(err)
	}

	filtered := filterComments(lines)

	c, err := operations.NewClient(conf.LanguageOp)
	if err != nil {
		log.Fatalln(err)
	}

	for _, l := range filtered {
		split := strings.Split(l.content, "\t")
		if len(split) != 5 {
			log.Printf("Line %3d: Expecting 5, but got %d: %v", l.num, len(split), l.content)
			continue
		}

		language := &language.Language{
			Code:      language.Code(split[0]),
			ISO_639_1: split[1],
			ISO_639_3: split[2],
			WikiData: split[3],
		}

		err = c.Insert(split[4], language)
		if err != nil {
			log.Printf("Line %3d: Error inserting (%v)", l.num, err)
		}
	}
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

type line struct {
	num     int
	content string
}

func filterComments(lines []string) []*line {
	filtered := make([]*line, 0, len(lines))
	for i, l := range lines {
		n := strings.TrimSpace(l)
		if len(n) > 0 && !strings.HasPrefix(n, "#") {
			filtered = append(filtered, &line{
				num:     i,
				content: n,
			})
		}
	}
	return filtered
}
