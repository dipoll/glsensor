package destinations

import (
	"io"
	"os"
)

/*
Console is a simple console destination that
logs metrics to stdout
*/
type Console struct {
	out io.Writer
}

/*
Send sends metrics to destination(s)
*/
func (c *Console) Send(m *MetricValue) error {
	c.out.Write([]byte(m.M.Info.Name))
	return nil
}

/*
NewConsole returns new stdout sender
*/
func NewConsole() *Console {
	cons := Console{
		out: os.Stdout,
	}
	return &cons
}

func init() {
	Register("console", func() Sender { return NewConsole() })
}
