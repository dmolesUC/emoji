package data

import (
	"bufio"
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

type ParseSeq func (seqType SeqType, data []byte) []string

// ParseSequences parses emoji sequences of the specified type from the specified data.
// Note that prior to version 3.0, type information is not included in the sequence data
// files, and ParseSequencesLegacy should be used (with the appropriate file) instead.
func ParseSequences(seqType SeqType, data []byte) []string {
	return ParseSequencesMatching(seqTypeRegexp(seqType), data)
}

// ParseSequencesLegacy parses emoji sequences for Emoji versions 1.0 and 2.0. Note
// that in Emoji 1.0, all sequences are in the main data file (filetype Data); for
// Emoji 2.0, all sequences are in the main sequences file (filetype Sequences), with
// no subfiles for variation sequences, ZWJ sequences, etc.
func ParseSequencesLegacy(seqType SeqType, data []byte) []string {
	if re, ok := legacySeqTypeRegexp(seqType); ok {
		return ParseSequencesMatching(re, data)
	}
	return nil
}

// ParseSequencesMatching parses emoji sequences from data lines matching
// the specified regexp.
func ParseSequencesMatching(re *regexp.Regexp, data []byte) []string {
	var result []string
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		if !re.MatchString(line) {
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

// ------------------------------------------------------------
// Unexported symbols

func toSeq(line string) (string, bool) {
	if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
		return "", false
	}
	seqMatch := seqRegexp.FindStringSubmatch(line)
	if len(seqMatch) == 1 {
		seq, err := parseSeq(strings.Split(seqMatch[0], " "))
		if err != nil {
			return "", false
		}
		return seq, true
	}
	return "", false
}

func parseSeq(seq []string) (string, error) {
	var result []rune
	for _, s := range seq {
		val, err := strconv.ParseInt(s, 16, 64)
		if err != nil {
			return "", err
		}
		result = append(result, rune(val))
	}
	return string(result), nil
}
