package englishtocron

import (
	"fmt"
	"regexp"
	"strings"
)

var reMonthMatch = regexp.MustCompile(`(?i)^((months|month)|(((january|february|march|april|may|june|july|august|september|october|november|december|JAN|FEB|MAR|APR|MAY|JUN|JUL|AUG|SEPT|OCT|NOV|DEC)( ?and)?,? ?)+))$`)
var reMonthExact = regexp.MustCompile(`(?i)^(month|months)$`)
var reMonthAbbrev = regexp.MustCompile(`(?i)(JAN|FEB|MAR|APR|MAY|JUN|JUL|AUG|SEP|OCT|NOV|DEC)`)

var months = []string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}

func monthTryFromToken(s string) bool {
	return reMonthMatch.MatchString(s)
}

func monthProcess(token string, cron *Cron) error {
	if reMonthExact.MatchString(token) {
		if last := cron.stackLast(); last != nil {
			if last.Owner == KindFrequencyOnly {
				cron.Syntax.Month = last.frequencyToString()
				cron.stackPop()
			} else if last.Owner == KindFrequencyWith {
				cron.Syntax.Month = last.frequencyToString()
				cron.stackPop()
			} else if last.Owner == KindRangeEnd {
				fs := 0
				fe := 0
				if last.FrequencyStart != nil {
					fs = *last.FrequencyStart
				}
				if last.FrequencyEnd != nil {
					fe = *last.FrequencyEnd
				}
				cron.Syntax.DayOfMonth = fmt.Sprintf("%d,%d", fs, fe)
			} else {
				cron.Syntax.Month = "*"
			}
		} else {
			cron.Syntax.Month = "*"
		}
	} else {
		matches := reMonthAbbrev.FindAllString(token, -1)
		if len(matches) == 0 {
			return errIncorrectValue("month", fmt.Sprintf("value %s is not a month format", token))
		}

		cron.Syntax.Month = ""
		monthList := make([]string, len(matches))
		for i, m := range matches {
			monthList[i] = strings.ToUpper(m)
		}

		if last := cron.stackLast(); last != nil {
			if last.Owner == KindFrequencyOnly || last.Owner == KindFrequencyWith {
				cron.Syntax.DayOfMonth = last.frequencyToString()
				cron.stackPop()
			} else if last.Owner == KindRangeStart {
				var endVal *string
				if last.Month != nil {
					endVal = last.Month.End
				}
				first := monthList[0]
				last.Month = &StartEndString{Start: &first, End: endVal}
				cron.stackPop()
				return nil
			} else if last.Owner == KindRangeEnd {
				if last.FrequencyEnd != nil {
					cron.Syntax.DayOfWeek = "?"
					if last.FrequencyStart != nil {
						cron.Syntax.DayOfMonth = fmt.Sprintf("%d-%d", *last.FrequencyStart, *last.FrequencyEnd)
					}
				}
				var startVal *string
				if last.Month != nil {
					startVal = last.Month.Start
				}
				first := monthList[0]
				data := StartEndString{Start: startVal, End: &first}
				last.Month = &data
				if data.Start != nil && data.End != nil {
					cron.Syntax.Month = fmt.Sprintf("%s-%s", *data.Start, *data.End)
				}
				cron.stackPop()
				return nil
			} else {
				cron.stackPop()
			}
		}

		for _, m := range months {
			for _, ml := range monthList {
				if ml == m && !strings.Contains(cron.Syntax.Month, m) {
					cron.Syntax.Month += m + ","
					break
				}
			}
		}
		cron.Syntax.Month = strings.TrimRight(cron.Syntax.Month, ",")
	}

	s := newStack(KindMonth)
	start := cron.Syntax.Month
	s.Month = &StartEndString{Start: &start}
	cron.Stack = append(cron.Stack, s)
	return nil
}
