package englishtocron

import "regexp"

var reRangeStartMatch = regexp.MustCompile(`(?i)(between|starting|start)`)
var reRangeStartBetween = regexp.MustCompile(`(?i)(between)`)

func rangeStartTryFromToken(s string) bool {
	return reRangeStartMatch.MatchString(s)
}

func rangeStartProcess(token string, cron *Cron) {
	s := newStack(KindRangeStart)
	s.IsBetweenRange = reRangeStartBetween.MatchString(token)
	cron.Stack = append(cron.Stack, s)
}
