package emoji

import (
	"bufio"
	"os"
	"regexp"

 //"golang.org/x/text/unicode/runenames"

)

const (
	cp       = "1[0-9A-F]{4}"
	cpSingle = "^(" + cp + ")"
	cpSequence = "^(" + cp + "[[:space:]])+(" + cp + ")"
	cpRange  = "^(" + cp + ")\\.\\.(" + cp + ")"
)

var cpSingleRe = regexp.MustCompile(cpSingle)
var cpSequenceRe = regexp.MustCompile(cpSequence)
var cpRangeRe = regexp.MustCompile(cpRange)

func parseFile(path string) error {
	in, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		if e := in.Close(); e != nil {
			panic(e)
		}
	}()

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		parseLine(line)
		if err := scanner.Err(); err != nil {
			return err
		}
	}
	return nil
}

func parseLine(line string) []string {
	submatch := cpRangeRe.FindStringSubmatch(line)
	if len(submatch) == 0 {

	} else if len(submatch) != 3 {

	}

	single := cpSingleRe.FindString(line)
	if single != "" {

	}
}
