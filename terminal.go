package log

import (
	"fmt"
	"os"
)

func NewTerminal() *Writer {
	queue := NewQueueLoop(QueueSize)
	queue.LoopOnce()

	settingBuilder := &SettingBuilder{}
	setting := settingBuilder.SetCallerShort(true).SetTimeLayout(TimeLayout2).SetFormatter(SimpleFormatter).Build()

	writerBuilder := &WriterBuilder{}
	writer, _ := writerBuilder.SetQueue(queue).SetSetting(setting).SetOut(Info, os.Stdout).SetOut(Error, os.Stderr).Build()

	return writer
}

var terminal = NewTerminal()

func I(tag string, message ...interface{}) {
	_ = terminal.Write(1, "", Info, "", tag, fmt.Sprint(message...), nil, nil, "", nil)
}

func E(tag string, message ...interface{}) {
	_ = terminal.Write(1, "", Error, "", tag, fmt.Sprint(message...), nil, nil, "", nil)
}
