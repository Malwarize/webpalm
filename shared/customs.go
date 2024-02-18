package shared

import (
	"fmt"
	"strings"
)

type RegexFlag struct {
	Regexes map[string]string
}

func (r *RegexFlag) String() string {
	return fmt.Sprintf("%v", r.Regexes)
}

func (r *RegexFlag) Set(value string) error {
	var err error
	r.Regexes, err = ParseRegexes(value)
	if err != nil {
		return err
	}
	return nil
}

func (r *RegexFlag) Type() string {
	return "regex"
}

func (r *RegexFlag) Value() map[string]string {
	return r.Regexes
}

func isBetween(s string, a string, b string, index int) bool {
	return strings.Count(s[:index], a) == strings.Count(s[index:], b) && strings.Count(s[:index], b) != 0
}

func ParseRegexes(stream string) (map[string]string, error) {
	result := make(map[string]string)
	var key, value string
	prev := 0
	for i, c := range stream {
		if c == ',' && !isBetween(stream, "{", "}", i) {
			keyValue := strings.Split(stream[prev:i], "=")
			if len(keyValue) != 2 {
				return nil, fmt.Errorf("invalid regexes format")
			}
			key = keyValue[0]
			value = keyValue[1]
			result[key] = value
			prev = i + 1
		}
	}
	keyValue := strings.Split(stream[prev:], "=")
	if len(keyValue) != 2 {
		return nil, fmt.Errorf("invalid regexes format")
	}
	key = keyValue[0]
	value = keyValue[1]
	result[key] = value
	return result, nil
}
