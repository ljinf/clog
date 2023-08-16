package clog

type Formatter interface {
	Format(e *Entry) error
}
