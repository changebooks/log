package log

type Level struct {
	name    string // 名称，如：Info
	num     int    // 数字，如：3
	lower   string // 全小写，如：info
	upper   string // 全大写，如：INFO
	initial string // 缩写，如：I
}

func (x *Level) GetName() string {
	return x.name
}

func (x *Level) GetNum() int {
	return x.num
}

func (x *Level) GetLower() string {
	return x.lower
}

func (x *Level) GetUpper() string {
	return x.upper
}

func (x *Level) GetInitial() string {
	return x.initial
}
