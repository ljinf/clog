package clog

import (
	"bytes"
	"github.com/bournex/ordered_container"
	"runtime"
	"strings"
	"time"
)

type Entry struct {
	logger *logger
	Buffer *bytes.Buffer
	Map    ordered_container.OrderedMap
	Level  Level
	Time   time.Time
	File   string
	Line   int
	Func   string
	Format string
	Args   []interface{}
}

func entry(logger *logger) *Entry {
	return &Entry{
		logger: logger,
		Buffer: new(bytes.Buffer),
		Map:    ordered_container.OrderedMap{Values: make([]ordered_container.OrderedValue, 5)},
	}
}

func (e *Entry) write(level Level, format string, args ...interface{}) {
	if e.logger.opt.level > level {
		return
	}

	e.Time = time.Now()
	e.Level = level
	e.Format = format
	e.Args = args
	if !e.logger.opt.disableCaller {
		if pc, file, line, ok := runtime.Caller(2); !ok {
			e.File = "???"
			e.Func = "???"
		} else {
			e.File, e.Line, e.Func = file, line, runtime.FuncForPC(pc).Name()
			e.Func = e.Func[strings.LastIndex(e.Func, "/")+1:]
		}
		e.format()
		e.writer()
		e.release()
	}
}

func (e *Entry) format() {
	_ = e.logger.opt.formatter.Format(e)
}

func (e *Entry) writer() {
	e.logger.mu.Lock()
	defer e.logger.mu.Unlock()
	for _, w := range e.logger.opt.output {
		_, _ = w.Write(e.Buffer.Bytes())
	}
}

func (e *Entry) release() {
	e.Args, e.Line, e.File, e.Format, e.Func = nil, 0, "", "", ""
	e.Buffer.Reset()
	e.logger.entryPool.Put(e)
}
