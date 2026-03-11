package englishtocron

import (
	"fmt"
	"regexp"
)

var reSecondsMatch = regexp.MustCompile(`(?i)(seconds|second|sec|secs)`)
var reSecondsExact = regexp.MustCompile(`(?i)^(seconds|second|sec|secs)$`)

func secondsTryFromToken(s string) bool {
	return reSecondsMatch.MatchString(s)
}

func secondsProcess(token string, cron *Cron) {
	if !reSecondsExact.MatchString(token) {
		return
	}
	if last := cron.stackLast(); last != nil {
		if last.Owner == KindFrequencyOnly {
			cron.Syntax.Seconds = fmt.Sprintf("0/%s", last.frequencyToString())
			cron.stackPop()
		} else if last.Owner == KindFrequencyWith {
			cron.Syntax.Seconds = last.frequencyToString()
			cron.stackPop()
		}
	} else {
		cron.Syntax.Seconds = "*"
	}
	cron.Stack = append(cron.Stack, newStack(KindSecund))
}
