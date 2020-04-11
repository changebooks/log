package log

import "errors"

func NewStream(p *Profile) (*Stream, error) {
	if p == nil {
		return nil, errors.New("profile can't be nil")
	}

	channel := p.GetChannel()
	directory := p.GetDirectory()
	level := p.GetLevel()
	timeLayout := p.GetTimeLayout()
	queueSize := p.GetQueueSize()
	enableCaller := p.GetEnableCaller()
	callerShort := p.GetCallerShort()
	enableStacktrace := p.GetEnableStacktrace()
	perm := p.GetPerm()
	pathJoiner := p.GetPathJoiner()
	formatter := p.GetFormatter()
	idGenerator := p.GetIdGenerator()
	timeProcessor := p.GetTimeProcessor()

	fileBuilder := &FileBuilder{}
	fileBuilder.
		SetDirectory(directory).
		SetChannel(channel).
		SetPathJoiner(pathJoiner).
		SetPerm(perm)

	settingBuilder := &SettingBuilder{}
	setting := settingBuilder.
		SetEnableCaller(enableCaller).
		SetCallerShort(callerShort).
		SetEnableStacktrace(enableStacktrace).
		SetTimeLayout(timeLayout).
		SetFormatter(formatter).
		SetIdGenerator(idGenerator).
		SetTimeProcessor(timeProcessor).
		Build()

	queue := NewQueueLoop(queueSize)
	queue.LoopOnce()

	streamBuilder := &StreamBuilder{}
	streamBuilder.
		SetChannel(channel).
		SetQueue(queue).
		SetSetting(setting).
		SetFileBuilder(fileBuilder).
		SetLevel(level)

	return streamBuilder.Build()
}
