package englishtocron

import (
	"fmt"
	"regexp"
)

var reMinuteMatch = regexp.MustCompile(`(?i)(minutes|minute|mins|min)`)
var reMinuteExact = regexp.MustCompile(`(?i)^(minutes|minute|mins|min)$`)

func minuteTryFromToken(s string) bool {
	return reMinuteMatch.MatchString(s)
}

func minuteProcess(token string, cron *Cron) {
	if !reMinuteExact.MatchString(token) {
		return
	}

	var minutes *StartEnd
	if last := cron.stackLast(); last != nil {
		if last.Owner == KindFrequencyOnly {
			minutes = &StartEnd{Start: last.Frequency}
			cron.Syntax.Min = fmt.Sprintf("0/%s", last.frequencyToString())
			cron.stackPop()
		} else if last.Owner == KindFrequencyWith {
			minutes = &StartEnd{Start: last.Frequency}
			cron.Syntax.Min = last.frequencyToString()
			cron.stackPop()
		} else if last.Owner == KindRangeStart {
			last.Min = &StartEnd{Start: last.FrequencyStart}
			return
		} else if last.Owner == KindRangeEnd {
			last.Min = &StartEnd{Start: last.FrequencyStart, End: last.FrequencyEnd}
			last.FrequencyEnd = nil
			if last.FrequencyStart != nil && last.Min.End != nil {
				cron.Syntax.Min = fmt.Sprintf("%d-%d", *last.FrequencyStart, *last.Min.End)
			}
			return
		}
	}

	if minutes != nil {
		s := newStack(KindMinute)
		s.Min = minutes
		cron.Stack = append(cron.Stack, s)
	}
}
