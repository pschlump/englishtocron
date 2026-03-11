package englishtocron

import (
	"fmt"
	"regexp"
)

var reHourMatch = regexp.MustCompile(`(?i)(hour|hrs|hours)`)
var reHourExact = regexp.MustCompile(`(?i)^(hour|hrs|hours)$`)

func hourTryFromToken(s string) bool {
	return reHourMatch.MatchString(s)
}

func hourProcess(token string, cron *Cron) {
	if !reHourExact.MatchString(token) {
		return
	}

	var hour *StartEnd
	if last := cron.stackLast(); last != nil {
		if last.Owner == KindFrequencyOnly {
			hour = &StartEnd{Start: last.Frequency}
			cron.Syntax.Hour = fmt.Sprintf("0/%s", last.frequencyToString())
			cron.Syntax.Min = "0"
			cron.stackPop()
		} else if last.Owner == KindFrequencyWith {
			hour = &StartEnd{Start: last.Frequency}
			cron.Syntax.Hour = last.frequencyToString()
			cron.Syntax.Min = "0"
			cron.stackPop()
		} else if last.Owner == KindRangeStart {
			last.Min = &StartEnd{Start: last.FrequencyStart}
			return
		} else if last.Owner == KindRangeEnd {
			last.Min = &StartEnd{Start: last.FrequencyStart, End: last.FrequencyEnd}
			last.FrequencyEnd = nil
			if last.FrequencyStart != nil && last.Min.End != nil {
				cron.Syntax.Hour = fmt.Sprintf("%d-%d", *last.FrequencyStart, *last.Min.End)
				cron.Syntax.Min = "0"
			}
			return
		}
	}
	cron.Syntax.Min = "0"

	if hour != nil {
		s := newStack(KindMinute)
		s.Hour = hour
		cron.Stack = append(cron.Stack, s)
	}
}
