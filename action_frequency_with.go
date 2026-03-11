package englishtocron

import "regexp"

var reFrequencyWith = regexp.MustCompile(`^[0-9]+(th|nd|rd|st)$`)
var reFrequencyWithNumPrefix = regexp.MustCompile(`^[0-9]+`)

func frequencyWithTryFromToken(s string) bool {
	return reFrequencyWith.MatchString(s)
}

func frequencyWithProcess(token string, cron *Cron) error {
	loc := reFrequencyWithNumPrefix.FindString(token)
	if loc == "" {
		return errCapture("frequency_with", token)
	}
	frequency, err := parseInt(loc, "frequency_with")
	if err != nil {
		return err
	}

	if last := cron.stackLast(); last != nil {
		if last.Owner == KindRangeEnd {
			last.FrequencyEnd = intPtr(frequency)
			return nil
		} else if last.Owner == KindRangeStart {
			last.FrequencyStart = intPtr(frequency)
			return nil
		}
	}

	s := newStack(KindFrequencyWith)
	s.Frequency = intPtr(frequency)
	s.DayOfWeek = strPtr(cron.Syntax.DayOfWeek)
	cron.Stack = append(cron.Stack, s)
	return nil
}
