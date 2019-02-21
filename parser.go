package emoji

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	//"golang.org/x/text/unicode/runenames"

)

const (
	cp         = "1?[0-9A-F]{4}"
	cpSingle   = "^(" + cp + ")"
	cpSequence = "^(" + cp + ")(?: (" + cp + "))+"
	cpRange    = "^(" + cp + ")[.]{2}(" + cp + ")"
)

var cpSingleRe = regexp.MustCompile(cpSingle)
var cpSequenceRe = regexp.MustCompile(cpSequence)
var cpRangeRe = regexp.MustCompile(cpRange)

// TODO: parse bindata bytes, not strings
// TODO: parse to unicode.RangeTable, not strings

func parseFile(path string) (results []string, err error) {
	in, err := os.Open(path)
	if err != nil {
		return
	}
	defer func() {
		if e := in.Close(); e != nil {
			panic(e)
		}
	}()

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if err2 := scanner.Err(); err2 != nil {
			err = err2
			return
		}

		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}
		lineResults, err2 := parseLine(line)
		if err2 != nil {
			err = err2
			return
		}
		results = append(results, lineResults...)
	}
	return
}

func toRune(codepointStr string) (rune, error) {
	cp, err := strconv.ParseInt(codepointStr, 16, 64)
	return rune(cp), err
}

func toSequence(cpStrs []string) (string, error) {
	result := ""
	for i, cpStr := range cpStrs {
		cp, err := strconv.ParseInt(cpStr, 16, 64)
		if err != nil {
			return "", err
		}
		result = result + string(rune(cp))
	}
	return result, nil
}

func toRange(cpStrStart, cpStrEnd string) ([]string, error) {
	runeStart, err := toRune(cpStrStart)
	if err != nil {
		return nil, err
	}
	runeEnd, err := toRune(cpStrEnd)
	if err != nil {
		return nil, err
	}
	var result []string
	for cp := runeStart; cp <= runeEnd; cp++ {
		result = append(result, string(cp))
	}
	return result, nil
}

func parseLine(line string) ([]string, error) {
	rangeMatch := cpRangeRe.FindStringSubmatch(line)
	if len(rangeMatch) > 1 {
		cpStrStart := rangeMatch[1]
		cpStrEnd := rangeMatch[2]
		return toRange(cpStrStart, cpStrEnd)
	}
	seqMatch := cpSequenceRe.FindStringSubmatch(line)
	if len(seqMatch) > 1 {
		seq, err := toSequence(seqMatch[1:])
		if err != nil {
			return nil, err
		}
		return []string{seq}, err
	}
	singleMatch := cpSingleRe.FindStringSubmatch(line)
	if len(singleMatch) > 1 {
		cp, err := toRune(singleMatch[1])
		if err != nil {
			return nil, err
		}
		return []string{string(cp)}, nil
	}
	return nil, nil
}
