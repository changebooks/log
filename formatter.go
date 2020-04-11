package log

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// 格式化日志函数：标准格式
func NormalFormatter(s *Schema) []byte {
	if s == nil {
		return []byte{EOL}
	}

	var r bytes.Buffer

	r.WriteString(fmt.Sprintf("%s %d %s %s %s %v", s.Time, s.ProcessId, s.Caller, s.Level, s.Tag, s.Message))
	r.WriteRune(EOL)

	if s.Error != "" {
		r.WriteString(LabelError)
		r.WriteRune(EOL)
		r.WriteString(s.Error)
		r.WriteRune(EOL)
	}

	if s.Stacktrace != "" {
		r.WriteString(LabelStacktrace)
		r.WriteRune(EOL)
		r.WriteString(s.Stacktrace)
		r.WriteRune(EOL)
	}

	return r.Bytes()
}

// 格式化日志函数：json
func JsonFormatter(s *Schema) []byte {
	if s == nil {
		return []byte{EOL}
	}

	r, err := json.Marshal(s)
	if err != nil {
		r = []byte(err.Error())
	}

	return append(r, EOL)
}

// 格式化日志函数：简要描述
func SimpleFormatter(s *Schema) []byte {
	if s == nil {
		return []byte{EOL}
	}

	r := fmt.Sprintf("%s %d %s %s %v %s %s %c", s.Time, s.ProcessId, s.Initial, s.Tag, s.Message, s.Error, s.Stacktrace, EOL)
	return []byte(r)
}
