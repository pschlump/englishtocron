package englishtocron

import "regexp"

var reRangeEndMatch = regexp.MustCompile(`(?i)(to|through|ending|end|and)`)
var reRangeEndAnd = regexp.MustCompile(`(?i)(and)`)

func rangeEndTryFromToken(s string) bool {
	return reRangeEndMatch.MatchString(s)
}

func rangeEndProcess(token string, cron *Cron) {
	isAnd := reRangeEndAnd.MatchString(token)

	last := cron.stackLast()
	if last == nil {
		return
	}

	last.IsAndConnector = isAnd

	switch last.Owner {
	case KindFrequencyWith, KindFrequencyOnly:
		last.FrequencyStart = last.Frequency
	case KindDay:
		if last.Day != nil {
			last.Day = &StartEndString{Start: last.DayOfWeek, End: last.Day.End}
		} else {
			last.Day = &StartEndString{Start: last.DayOfWeek}
		}
	case KindMonth:
		last.Owner = KindRangeEnd
	case KindRangeStart:
		last.Owner = KindRangeEnd
	}
	last.Owner = KindRangeEnd
}
