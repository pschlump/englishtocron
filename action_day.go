package englishtocron

import (
	"fmt"
	"regexp"
	"strings"
)

var reDayMatch = regexp.MustCompile(`(?i)^((days|day)|(((monday|tuesday|wednesday|thursday|friday|saturday|sunday|WEEKEND|MON|TUE|WED|THU|FRI|SAT|SUN)( ?and)?,? ?)+))$`)
var reDayExact = regexp.MustCompile(`(?i)^(day|days)$`)
var reDayWeekdays = regexp.MustCompile(`(?i)(MON|TUE|WED|THU|FRI|SAT|SUN|WEEKEND)`)

var weekDays = []string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}

func dayTryFromToken(s string) bool {
	return reDayMatch.MatchString(s)
}

func dayProcess(token string, cron *Cron) error {
	if reDayExact.MatchString(token) {
		cron.Syntax.DayOfWeek = "?"
		if cron.Syntax.Min == "*" {
			cron.Syntax.Min = "0"
		}
		if cron.Syntax.Hour == "*" {
			cron.Syntax.Hour = "0"
		}
		if last := cron.stackLast(); last != nil {
			if last.Owner == KindFrequencyOnly {
				cron.Syntax.DayOfMonth = fmt.Sprintf("*/%s", last.frequencyToString())
				cron.stackPop()
			} else if last.Owner == KindFrequencyWith {
				cron.Syntax.DayOfMonth = last.frequencyToString()
				cron.stackPop()
			} else {
				cron.Syntax.DayOfMonth = "*"
			}
		} else {
			cron.Syntax.DayOfMonth = "*/1"
		}
	} else {
		matches := reDayWeekdays.FindAllString(token, -1)
		if len(matches) == 0 {
			return errIncorrectValue("day", fmt.Sprintf("value %s is not a weekend format", token))
		}

		cron.Syntax.DayOfWeek = ""
		days := make([]string, len(matches))
		for i, d := range matches {
			days[i] = strings.ToUpper(d)
		}

		if last := cron.stackLast(); last != nil {
			if last.Owner == KindRangeStart {
				var endVal *string
				if last.Day != nil {
					endVal = last.Day.End
				}
				first := days[0]
				last.Day = &StartEndString{Start: &first, End: endVal}
				return nil
			} else if last.Owner == KindRangeEnd {
				var startVal *string
				if last.Day != nil {
					startVal = last.Day.Start
				}
				first := days[0]
				data := StartEndString{Start: startVal, End: &first}
				last.Day = &data
				if data.Start != nil && data.End != nil {
					cron.Syntax.DayOfWeek = fmt.Sprintf("%s-%s", *data.Start, *data.End)
				}
				cron.Syntax.DayOfMonth = "?"
				cron.stackPop()
				return nil
			} else if last.Owner == KindOnlyOn {
				if len(days) == 0 {
					return errIncorrectValue("day", "Expected at least one day in 'only on' syntax but found none")
				}
				cron.Syntax.DayOfWeek = days[0]
				cron.Syntax.DayOfMonth = "?"
				cron.stackPop()
				return nil
			}
			cron.Stack = cron.Stack[:0]
		}

		// Normal day processing
		for _, day := range weekDays {
			contains := false
			for _, d := range days {
				if d == day {
					contains = true
					break
				}
			}
			if contains && !strings.Contains(cron.Syntax.DayOfWeek, day) {
				cron.Syntax.DayOfWeek += day + ","
			}
		}
		// Handle WEEKEND
		for _, d := range days {
			if d == "WEEKEND" {
				for _, wd := range []string{"SAT", "SUN"} {
					if !strings.Contains(cron.Syntax.DayOfWeek, wd) {
						cron.Syntax.DayOfWeek += wd + ","
					}
				}
				break
			}
		}
		cron.Syntax.DayOfWeek = strings.TrimRight(cron.Syntax.DayOfWeek, ",")
		cron.Syntax.DayOfMonth = "?"
	}

	s := newStack(KindDay)
	s.DayOfWeek = strPtr(cron.Syntax.DayOfWeek)
	cron.Stack = append(cron.Stack, s)
	return nil
}
