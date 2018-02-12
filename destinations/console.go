package destinations

import (
	"fmt"
	"io"
	"os"
)

// Console is a simple console destination that
// logs metrics to stdout
type Console struct {
	out io.Writer
}

// Send sends metrics to destination(s)
func (c *Console) Send(m *MetricValue) error {
	_, err := fmt.Fprintf(c.out, "%v", m.M)
	return err
}

// NewConsole returns new stdout sender
func NewConsole() *Console {
	cons := Console{
		out: os.Stdout,
	}
	return &cons
}

func init() {
	Register("console", func() Sender { return NewConsole() })
}
