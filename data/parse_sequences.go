package data

import (
	"bufio"
	"bytes"
	"strings"
)

// ParseSequences parses emoji sequences of the specified type from the specified data.
// Note that prior to version 3.0, type information is not included in the sequence data
// files, and ParseAllSequences should be used instead.
func ParseSequences(seqType SeqType, data []byte) []string {
	var result []string
	typeRegexp := hasTypeRegexp(seqType)
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}
		if !typeRegexp.MatchString(line) {
			continue
		}
		seq, ok := toSeq(line)
		if !ok {
			continue
		}
		result = append(result, seq)
	}
	return result
}

// ParseAllSequences parses all emoji sequences from the specified data, without regard
// to type.
func ParseAllSequences(data []byte) []string {
	var result []string
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}
		seq, ok := toSeq(line)
		if !ok {
			continue
		}
		result = append(result, seq)
	}
	return result
}

func toSeq(line string) (string, bool) {
	seqMatch := SeqRegexp.FindStringSubmatch(line)
	if len(seqMatch) == 1 {
		seq, err := parseSeq(strings.Split(seqMatch[0], " "))
		if err != nil {
			return "" , false
		}
		return seq, true
	}
	return "", false
}

func parseSeq(seq []string) (string, error) {
	var result []rune
	for _, s := range seq {
		val, err := parse32(s)
		if err != nil {
			return "" , err
		}
		result = append(result, rune(*val))
	}
	return string(result), nil
}
