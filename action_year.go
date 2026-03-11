package englishtocron

import (
	"fmt"
	"regexp"
	"strings"
)

var reYearMatch = regexp.MustCompile(`(?i)((years|year)|([0-9]{4}[0-9]*(( ?and)?,? ?))+)`)
var reYearExact = regexp.MustCompile(`(?i)^(years|year)$`)
var reYearNumeric = regexp.MustCompile(`[0-9]+`)
var reYearFormat = regexp.MustCompile(`^[0-9]{4}$`)

func yearTryFromToken(s string) bool {
	return reYearMatch.MatchString(s)
}

func yearProcess(token string, cron *Cron) error {
	if reYearExact.MatchString(token) {
		cron.Syntax.Year = "?"
		if last := cron.stackLast(); last != nil {
			if last.Owner == KindFrequencyOnly {
				cron.Syntax.Year = fmt.Sprintf("0/%s", last.frequencyToString())
				cron.stackPop()
			} else if last.Owner == KindFrequencyWith {
				cron.Syntax.Year = last.frequencyToString()
			} else {
				cron.Syntax.Year = "*"
			}
		}
	} else {
		rawMatches := reYearNumeric.FindAllString(token, -1)
		var years []int
		for _, m := range rawMatches {
			if reYearFormat.MatchString(m) {
				v, err := parseInt(m, "year")
				if err == nil {
					years = append(years, v)
				}
			}
		}

		if last := cron.stackLast(); last != nil {
			if last.Owner == KindRangeStart {
				var endVal *int
				if last.Year != nil {
					endVal = last.Year.End
				}
				if len(years) > 0 {
					last.Year = &StartEnd{Start: intPtr(years[0]), End: endVal}
				}
				return nil
			} else if last.Owner == KindRangeEnd {
				var startVal *int
				if last.Year != nil {
					startVal = last.Year.Start
				}
				var endVal *int
				if len(years) > 0 {
					endVal = intPtr(years[0])
				}
				s := 0
				e := 0
				if startVal != nil {
					s = *startVal
				}
				if endVal != nil {
					e = *endVal
				}
				cron.Syntax.Year = fmt.Sprintf("%d-%d", s, e)
				cron.stackPop()
				return nil
			}
		}

		if len(years) == 0 {
			return errIncorrectValue("year", fmt.Sprintf("value %s is not a year format", token))
		}

		parts := make([]string, len(years))
		for i, y := range years {
			parts[i] = fmt.Sprintf("%d", y)
		}
		cron.Syntax.Year = strings.Join(parts, ",")
	}

	cron.Stack = append(cron.Stack, newStack(KindYear))
	return nil
}
