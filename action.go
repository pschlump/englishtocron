package englishtocron

import (
	"fmt"
	"strings"
)

// Kind enumerates the token types.
type Kind int

const (
	KindFrequencyWith Kind = iota
	KindFrequencyOnly
	KindClockTime
	KindDay
	KindSecund
	KindMinute
	KindHour
	KindMonth
	KindYear
	KindRangeStart
	KindRangeEnd
	KindOnlyOn
)

var kindOrder = []Kind{
	KindFrequencyWith,
	KindFrequencyOnly,
	KindClockTime,
	KindDay,
	KindSecund,
	KindMinute,
	KindHour,
	KindMonth,
	KindYear,
	KindRangeStart,
	KindRangeEnd,
	KindOnlyOn,
}

// tryFromToken returns the Kind that matches the token, or -1 if none.
func tryFromToken(token string) (Kind, bool) {
	for _, k := range kindOrder {
		var matched bool
		switch k {
		case KindFrequencyWith:
			matched = frequencyWithTryFromToken(token)
		case KindFrequencyOnly:
			matched = frequencyOnlyTryFromToken(token)
		case KindClockTime:
			matched = clockTimeTryFromToken(token)
		case KindDay:
			matched = dayTryFromToken(token)
		case KindSecund:
			matched = secondsTryFromToken(token)
		case KindMinute:
			matched = minuteTryFromToken(token)
		case KindHour:
			matched = hourTryFromToken(token)
		case KindMonth:
			matched = monthTryFromToken(token)
		case KindYear:
			matched = yearTryFromToken(token)
		case KindRangeStart:
			matched = rangeStartTryFromToken(token)
		case KindRangeEnd:
			matched = rangeEndTryFromToken(token)
		case KindOnlyOn:
			matched = strings.ToLower(token) == "only on"
		}
		if matched {
			return k, true
		}
	}
	return -1, false
}

// process dispatches token processing to the appropriate handler.
func (k Kind) process(token string, cron *Cron) error {
	switch k {
	case KindFrequencyWith:
		return frequencyWithProcess(token, cron)
	case KindFrequencyOnly:
		freq, err := parseInt(token, "frequency_only")
		if err != nil {
			return err
		}
		frequencyOnlyProcess(freq, cron)
	case KindClockTime:
		return clockTimeProcess(token, cron)
	case KindDay:
		return dayProcess(token, cron)
	case KindSecund:
		secondsProcess(token, cron)
	case KindMinute:
		minuteProcess(token, cron)
	case KindHour:
		hourProcess(token, cron)
	case KindMonth:
		return monthProcess(token, cron)
	case KindYear:
		return yearProcess(token, cron)
	case KindRangeStart:
		rangeStartProcess(token, cron)
	case KindRangeEnd:
		rangeEndProcess(token, cron)
	case KindOnlyOn:
		// nothing to do; next token (a day) handles it
	}
	return nil
}

func parseInt(s, state string) (int, error) {
	var v int
	_, err := fmt.Sscanf(s, "%d", &v)
	if err != nil {
		return 0, errParseToNumber(state, s)
	}
	return v, nil
}
