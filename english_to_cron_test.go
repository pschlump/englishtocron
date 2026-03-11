package englishtocron

import (
	"testing"
)

func TestStrCronSyntax(t *testing.T) {
	cases := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		// Seconds
		{"Run second", "* * * * * ? *", false},
		{"every 5 second", "0/5 * * * * ? *", false},
		{"every 5 second on september", "0/5 * * * SEP ? *", false},
		{"every 5 second on 9 month", "0/5 * * * 9 ? *", false},
		{"Every 2 seconds, only on thursday", "0/2 * * ? * THU *", false},
		{"Run every 2 second on the 12th day", "0/2 0 0 12 * ? *", false},
		{"Run every 2 second on Monday thursday", "0/2 * * ? * MON,THU *", false},
		{"Run every 10 seconds Monday through thursday between 6:00 am and 8:00 pm", "0/10 * 6-20 ? * MON-THU *", false},
		// Minutes
		{"Run every minute", "0 * * * * ? *", false},
		{"Run every 15 minutes", "0 0/15 * * * ? *", false},
		{"every minutes on thursday", "0 * * ? * THU *", false},
		{"every 2 minutes on Thursday", "0 0/2 * ? * THU *", false},
		{"Run every 10 minutes Monday through Friday every month", "0 0/10 * ? * MON-FRI *", false},
		{"Run every 1 minutes Monday through Thursday between 6:00 am and 9:00 pm", "0 0/1 6-21 ? * MON-THU *", false},
		{"Run every 5 minutes Monday through Thursday between 6:00 am and 9:00 am", "0 0/5 6-9 ? * MON-THU *", false},
		{"Every 5 minutes, only on Friday", "0 0/5 * ? * FRI *", false},
		// Hours
		{"Run every 3 hours", "0 0 0/3 * * ? *", false},
		{"Run every 6 hours, starting at 1:00 pm on day Monday", "0 0 0/6 ? * MON *", false},
		{"Run every 1 hour only on weekends", "0 0 0/1 ? * SAT,SUN *", false},
		{"Run every hour only on weekends", "0 0 * ? * SAT,SUN *", false},
		{"2pm on Tuesday, Wednesday and Thursday", "0 0 14 ? * TUE,WED,THU *", false},
		// Days
		{"Run every day", "0 0 0 */1 * ? *", false},
		{"Run every 4 days", "0 0 0 */4 * ? *", false},
		{"every day at 4:00 pm", "0 0 16 */1 * ? *", false},
		{"every 2 day at 4:00 pm", "0 0 16 */2 * ? *", false},
		{"every 5 day at 4:30 pm", "0 30 16 */5 * ? *", false},
		{"every 5 day at 4:30 pm only in September", "0 30 16 */5 SEP ? *", false},
		{"every 5 day at 4:30 pm Monday through Thursday", "0 30 16 ? * MON-THU *", false},
		{"Run every day from January to March", "0 0 0 */1 JAN-MAR ? *", false},
		{"Run every 3 days at noon", "0 0 12 */3 * ? *", false},
		{"Run every 2nd day of the month", "0 0 0 2 * ? *", false},
		// Month
		{"Run every sec from January to March", "* * * * JAN-MAR ? *", false},
		{"Run every minute from January to March", "0 * * * JAN-MAR ? *", false},
		{"Run every hours from January to March", "0 0 * * JAN-MAR ? *", false},
		// Year
		{"every 2 day from January to August in 2020 and 2024", "0 0 0 */2 JAN-AUG ? 2020,2024", false},
		// Specific Times (AM/PM)
		{"Run at 10:00 am", "0 0 10 * * ? *", false},
		{"Run at 12:15 pm", "0 15 12 * * ? *", false},
		{"Run at 6:00 pm every Monday through Friday", "0 0 18 ? * MON-FRI *", false},
		{"Run at noon every Sunday", "0 0 12 ? * SUN *", false},
		{"Run at midnight on the 1st and 15th of the month", "0 0 0 1,15 * ? *", false},
		{"midnight on Tuesdays", "0 0 0 ? * TUE *", false},
		{"Run at 5:15am every Tuesday", "0 15 5 ? * TUE *", false},
		{"7pm every Thursday", "0 0 19 ? * THU *", false},
		{"2pm and 6pm", "0 0 14,18 * * ? *", false},
		{"5am, 10am and 3pm", "0 0 5,10,15 * * ? *", false},
		{"Run every hour only on Monday", "0 0 * ? * MON *", false},
		{"Run every 30 seconds only on weekends", "0/30 * * ? * SAT,SUN *", false},
		{"4pm, 5pm and 7pm", "0 0 16,17,19 * * ? *", false},
		{"4pm, 5pm, and 7pm", "0 0 16,17,19 * * ? *", false},
		{"4pm, 5pm, 7pm", "0 0 16,17,19 * * ? *", false},
		{"4pm and 5pm and 7pm", "0 0 16,17,19 * * ? *", false},
		// Invalid
		{"invalid input", "", true},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := StrCronSyntax(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Errorf("expected error for %q, got %q", tc.input, got)
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error for %q: %v", tc.input, err)
				return
			}
			if got != tc.expected {
				t.Errorf("input %q\n  got:  %q\n  want: %q", tc.input, got, tc.expected)
			}
		})
	}
}
