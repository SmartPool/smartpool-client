package smartpool

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Log struct {
	*log.Logger
	f *os.File
}

func NewLog() *Log {
	f, err := os.OpenFile("smartpool.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return &Log{log.New(f, "", log.Lshortfile|log.LstdFlags), f}
}

func (l *Log) Close() {
	l.f.Close()
}

func (l *Log) Printf(format string, a ...interface{}) (n int, err error) {
	l.Logger.Printf(format, a...)
	t := time.Now()
	ts := t.Format(time.RFC3339)
	fmt.Printf("%v ", ts)
	return fmt.Printf(format, a...)
}
