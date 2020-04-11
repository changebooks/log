package log

import (
	"errors"
	"io"
	"os"
	"strings"
	"sync"
)

type File struct {
	io.WriteCloser
	mu         sync.RWMutex                                                // ensures atomic writes; protects the following fields
	file       *os.File                                                    // 文件资源
	path       string                                                      // 最近一次日志路径
	perm       os.FileMode                                                 // 文件权限，缺省：define.FilePerm
	directory  string                                                      // 文件夹
	channel    string                                                      // 频道
	level      string                                                      // 日志级别
	pathJoiner func(directory string, channel string, level string) string // 拼接日志路径函数
	listeners  []func(data []byte, n int, err error)                       // 写日志后，回调接口列表
}

func (x *File) Write(b []byte) (n int, err error) {
	x.mu.Lock()
	defer x.mu.Unlock()

	path := x.pathJoiner(x.directory, x.channel, x.level)

	if x.path != path {
		// 文件路径改变
		// 关闭文件资源，写日志时，重新打开
		if x.file != nil {
			if err1 := x.file.Close(); err1 != nil {
				return 0, err1
			}

			x.file = nil
		}

		x.path = path
	}

	if IsNotExist(x.path) {
		x.file = nil
	}

	if x.file == nil {
		// 写日志时，才打开文件资源
		fd, err2 := os.OpenFile(x.path, FileFlag, x.perm)
		if err2 != nil {
			return 0, err2
		}

		if err3 := os.Chmod(x.path, x.perm); err3 != nil {
			_ = fd.Close()
			return 0, err3
		}

		x.file = fd
	}

	n, err = x.file.Write(b)

	if x.listeners != nil {
		for _, l := range x.listeners {
			if l != nil {
				l(b, n, err)
			}
		}
	}

	return
}

func (x *File) Close() error {
	x.mu.Lock()
	defer x.mu.Unlock()

	if x.file == nil {
		// 防范重复关闭
		// 同一个File给多个Writer使用时，该函数会被调用多次
		return nil
	}

	if err := x.file.Close(); err != nil {
		return err
	}

	x.file = nil
	return nil
}

func (x *File) GetFile() *os.File {
	return x.file
}

func (x *File) GetPath() string {
	return x.path
}

func (x *File) GetPerm() os.FileMode {
	return x.perm
}

func (x *File) GetDirectory() string {
	return x.directory
}

func (x *File) GetChannel() string {
	return x.channel
}

func (x *File) GetLevel() string {
	return x.level
}

func (x *File) GetPathJoiner() func(directory string, channel string, level string) string {
	return x.pathJoiner
}

func (x *File) GetListeners() []func(data []byte, n int, err error) {
	return x.listeners
}

type FileBuilder struct {
	mu         sync.Mutex // ensures atomic writes; protects the following fields
	perm       os.FileMode
	directory  string
	channel    string
	level      string
	pathJoiner func(directory string, channel string, level string) string
	listeners  []func(data []byte, n int, err error)
}

func (x *FileBuilder) Build() (*File, error) {
	if x.directory == "" {
		return nil, errors.New("directory can't be empty")
	}

	if x.channel == "" {
		return nil, errors.New("channel can't be empty")
	}

	if x.level == "" {
		return nil, errors.New("level can't be empty")
	}

	if err := Mkdir(x.directory); err != nil {
		return nil, err
	}

	perm := x.perm
	if perm <= 0 {
		perm = FilePerm
	}

	pathJoiner := x.pathJoiner
	if pathJoiner == nil {
		pathJoiner = PathJoiner
	}

	return &File{
		perm:       perm,
		directory:  x.directory,
		channel:    x.channel,
		level:      x.level,
		pathJoiner: pathJoiner,
		listeners:  x.listeners,
	}, nil
}

func (x *FileBuilder) SetPerm(p os.FileMode) *FileBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.perm = p
	return x
}

func (x *FileBuilder) SetDirectory(s string) *FileBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.directory = s
	return x
}

func (x *FileBuilder) SetChannel(s string) *FileBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.channel = s
	return x
}

func (x *FileBuilder) SetLevel(s string) *FileBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.level = s
	return x
}

func (x *FileBuilder) SetPathJoiner(f func(directory string, channel string, level string) string) *FileBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.pathJoiner = f
	return x
}

func (x *FileBuilder) AddListener(f func(data []byte, n int, err error)) *FileBuilder {
	if f == nil {
		return x
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	x.listeners = append(x.listeners, f)
	return x
}

func IsNotExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}
