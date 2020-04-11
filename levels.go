package log

var levels = map[int]*Level{
	Verbose: {
		name:    "Verbose",
		num:     Verbose,
		lower:   "verbose",
		upper:   "VERBOSE",
		initial: "V",
	},
	Debug: {
		name:    "Debug",
		num:     Debug,
		lower:   "debug",
		upper:   "DEBUG",
		initial: "D",
	},
	Info: {
		name:    "Info",
		num:     Info,
		lower:   "info",
		upper:   "INFO",
		initial: "I",
	},
	Notice: {
		name:    "Notice",
		num:     Notice,
		lower:   "notice",
		upper:   "NOTICE",
		initial: "N",
	},
	Warning: {
		name:    "Warning",
		num:     Warning,
		lower:   "warning",
		upper:   "WARNING",
		initial: "W",
	},
	Error: {
		name:    "Error",
		num:     Error,
		lower:   "error",
		upper:   "ERROR",
		initial: "E",
	},
	Critical: {
		name:    "Critical",
		num:     Critical,
		lower:   "critical",
		upper:   "CRITICAL",
		initial: "C",
	},
	Alert: {
		name:    "Alert",
		num:     Alert,
		lower:   "alert",
		upper:   "ALERT",
		initial: "A",
	},
	Emergency: {
		name:    "Emergency",
		num:     Emergency,
		lower:   "emergency",
		upper:   "EMERGENCY",
		initial: "M",
	},
	Silent: {
		name:    "Silent",
		num:     Silent,
		lower:   "silent",
		upper:   "SILENT",
		initial: "S",
	},
}

func GetLevels() map[int]*Level {
	return levels
}

func GetLevel(num int) *Level {
	if r, ok := levels[num]; ok {
		return r
	}
	return nil
}

func IsLevelExist(num int) bool {
	_, ok := levels[num]
	return ok
}
