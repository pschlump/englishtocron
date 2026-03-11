package englishtocron

import "fmt"

// Syntax holds the seven fields of a cron expression.
type Syntax struct {
	Seconds    string
	Min        string
	Hour       string
	DayOfMonth string
	DayOfWeek  string
	Month      string
	Year       string
}

func defaultSyntax() Syntax {
	return Syntax{
		Seconds:    "0",
		Min:        "*",
		Hour:       "*",
		DayOfMonth: "*",
		DayOfWeek:  "?",
		Month:      "*",
		Year:       "*",
	}
}

// Cron holds the parsed cron expression state.
type Cron struct {
	Syntax Syntax
	Stack  []*Stack
}

func newCron() *Cron {
	return &Cron{Syntax: defaultSyntax()}
}

// String formats the cron as a 7-field expression.
func (c *Cron) String() string {
	return fmt.Sprintf("%s %s %s %s %s %s %s",
		c.Syntax.Seconds,
		c.Syntax.Min,
		c.Syntax.Hour,
		c.Syntax.DayOfMonth,
		c.Syntax.Month,
		c.Syntax.DayOfWeek,
		c.Syntax.Year,
	)
}

// New parses an English schedule description into a Cron.
func New(text string) (*Cron, error) {
	tokenizer := newTokenizer()
	tokens := tokenizer.run(text)

	if len(tokens) == 0 {
		return nil, errInvalidInput()
	}

	cron := newCron()
	for _, token := range tokens {
		if k, ok := tryFromToken(token); ok {
			if err := k.process(token, cron); err != nil {
				return nil, err
			}
		}
	}
	return cron, nil
}

// StrCronSyntax converts an English schedule description to cron syntax.
func StrCronSyntax(input string) (string, error) {
	cron, err := New(input)
	if err != nil {
		return "", err
	}
	return cron.String(), nil
}

// stackLast returns the last element of the stack, or nil.
func (c *Cron) stackLast() *Stack {
	if len(c.Stack) == 0 {
		return nil
	}
	return c.Stack[len(c.Stack)-1]
}

// stackPop removes the last element from the stack.
func (c *Cron) stackPop() {
	if len(c.Stack) > 0 {
		c.Stack = c.Stack[:len(c.Stack)-1]
	}
}
