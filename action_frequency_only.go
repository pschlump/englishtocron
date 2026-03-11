package englishtocron

import "regexp"

var reFrequencyOnly = regexp.MustCompile(`^[0-9]+$`)

func frequencyOnlyTryFromToken(s string) bool {
	return reFrequencyOnly.MatchString(s)
}

func frequencyOnlyProcess(frequency int, cron *Cron) {
	if last := cron.stackLast(); last != nil {
		if last.Owner == KindRangeEnd {
			last.FrequencyEnd = intPtr(frequency)
			return
		} else if last.Owner == KindRangeStart {
			last.FrequencyStart = intPtr(frequency)
			return
		}
	}
	s := newStack(KindFrequencyOnly)
	s.Frequency = intPtr(frequency)
	cron.Stack = append(cron.Stack, s)
}
