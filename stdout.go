package smartpool

import "fmt"

type StdOut struct{}

func (StdOut) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}
