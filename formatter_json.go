package log

import (
	"fmt"
	"github.com/bournex/ordered_container"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"time"
)

type JsonFormatter struct {
	IgnoreBasicFields bool
}

func (f *JsonFormatter) Format(e *Entry) error {
	if !f.IgnoreBasicFields {
		e.Map.Values[0] = ordered_container.OrderedValue{Key: "level", Value: LevelNameMapping[e.Level]}
		e.Map.Values[1] = ordered_container.OrderedValue{Key: "time", Value: e.Time.Format(time.RFC3339)}
		e.Map.Values[2] = ordered_container.OrderedValue{Key: "file", Value: e.File + ":" + strconv.Itoa(e.Line)}
		e.Map.Values[3] = ordered_container.OrderedValue{Key: "func", Value: e.Func}

		switch e.Format {
		case FmtEmptySeparate:
			e.Map.Values[4] = ordered_container.OrderedValue{Key: "message", Value: fmt.Sprint(e.Args...)}
		default:
			e.Map.Values[4] = ordered_container.OrderedValue{Key: "message", Value: fmt.Sprintf(e.Format, e.Args...)}
		}

		return jsoniter.NewEncoder(e.Buffer).Encode(e.Map)
	}

	switch e.Format {
	case FmtEmptySeparate:
		for _, arg := range e.Args {
			if err := jsoniter.NewEncoder(e.Buffer).Encode(arg); err != nil {
				return err
			}
		}
	default:
		e.Buffer.WriteString(fmt.Sprintf(e.Format, e.Args...))
	}
	return nil
}
