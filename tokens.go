package englishtocron

import (
	"regexp"
	"strings"
)

var reTokens = regexp.MustCompile(
	`(?i)(?:seconds|second|secs|sec)|(?:hours?|hrs?)|(?:minutes?|mins?|min)|(?:months?|(?:january|february|march|april|may|june|july|august|september|october|november|december|jan|feb|mar|apr|may|jun|jul|aug|sept|oct|nov|dec)(?: ?and)?,? ?)+|[0-9]+(?:th|nd|rd|st)|(?:[0-9]+:)?[0-9]+ ?(?:am|pm)|[0-9]+:[0-9]+|(?:noon|midnight)|(?:days?|(?:monday|tuesday|wednesday|thursday|friday|saturday|sunday|weekend|mon|tue|wed|thu|fri|sat|sun)(?: ?and)?,? ?)+|(?:[0-9]{4}[0-9]*(?: ?and)?,? ?)+|[0-9]+|(?:only on)|(?:to|through|ending|end|and)|(?:between|starting|start)`,
)

type tokenizer struct{}

func newTokenizer() *tokenizer { return &tokenizer{} }

func (t *tokenizer) run(input string) []string {
	processed := strings.ReplaceAll(input, ", ", " and ")
	if strings.Contains(processed, "only on") {
		processed = strings.ReplaceAll(processed, " and only on", " only on")
	}
	matches := reTokens.FindAllString(processed, -1)
	result := make([]string, 0, len(matches))
	for _, m := range matches {
		result = append(result, strings.TrimSpace(m))
	}
	return result
}
