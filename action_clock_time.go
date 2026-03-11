package englishtocron

import (
	"fmt"
	"regexp"
	"strings"
)

var reClockTimeMatch = regexp.MustCompile(`(?i)^([0-9]+:)?[0-9]+ *(AM|PM)$|^([0-9]+:[0-9]+)$|(noon|midnight)`)
var reClockHour = regexp.MustCompile(`^[0-9]+`)
var reClockMinute = regexp.MustCompile(`:[0-9]+`)
var reNoonMidnight = regexp.MustCompile(`(?i)(noon|midnight)`)

func clockTimeTryFromToken(s string) bool {
	return reClockTimeMatch.MatchString(s)
}

func clockTimeProcess(token string, cron *Cron) error {
	hour := 0
	minute := 0

	if loc := reClockHour.FindString(token); loc != "" {
		v, err := parseInt(loc, "clock_time")
		if err != nil {
			return err
		}
		hour = v
	}

	if loc := reClockMinute.FindString(token); loc != "" {
		parts := strings.SplitN(loc, ":", 2)
		if len(parts) == 2 {
			v, err := parseInt(parts[1], "clock_time")
			if err != nil {
				return err
			}
			if v >= 60 {
				return errIncorrectValue("clock_time", fmt.Sprintf("minute %d should be lower or equal to 60", v))
			}
			minute = v
		}
	}

	lower := strings.ToLower(token)
	if strings.Contains(lower, "pm") {
		if hour < 12 {
			hour += 12
		} else if hour > 12 {
			return errIncorrectValue("clock_time", fmt.Sprintf("please correct the time before PM. value: %d", hour))
		}
	} else if strings.Contains(lower, "am") {
		if hour == 12 {
			hour = 0
		} else if hour > 12 {
			return errIncorrectValue("clock_time", fmt.Sprintf("please correct the time before AM. value: %d", hour))
		}
	}

	if reNoonMidnight.MatchString(token) {
		if strings.ToLower(token) == "noon" {
			hour = 12
		} else {
			hour = 0
		}
		minute = 0
	}

	if last := cron.stackLast(); last != nil {
		if last.Owner == KindRangeStart {
			last.Hour = &StartEnd{Start: intPtr(hour)}
			return nil
		} else if last.Owner == KindRangeEnd {
			if last.Hour != nil {
				if last.Hour.Start != nil && *last.Hour.Start == hour {
					last.Min = &StartEnd{Start: intPtr(hour), End: intPtr(hour)}
					cron.Syntax.Hour = fmt.Sprintf("%d-%d", hour, hour)
				} else {
					last.Hour.End = intPtr(hour)
					startHour := 0
					if last.Hour.Start != nil {
						startHour = *last.Hour.Start
					}
					if last.IsAndConnector && !last.IsBetweenRange {
						if strings.Contains(cron.Syntax.Hour, ",") {
							cron.Syntax.Hour = fmt.Sprintf("%s,%d", cron.Syntax.Hour, hour)
						} else {
							cron.Syntax.Hour = fmt.Sprintf("%d,%d", startHour, hour)
						}
					} else {
						cron.Syntax.Hour = fmt.Sprintf("%d-%d", startHour, hour)
					}
				}
			}
			return nil
		}
	}

	cron.Syntax.Min = fmt.Sprintf("%d", minute)
	cron.Syntax.Hour = fmt.Sprintf("%d", hour)

	s := newStack(KindClockTime)
	s.Hour = &StartEnd{Start: intPtr(hour)}
	s.Min = &StartEnd{Start: intPtr(minute)}
	cron.Stack = append(cron.Stack, s)
	return nil
}
