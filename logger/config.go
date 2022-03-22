package logger

var config *sumConfig

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
)

type sumConfig struct {
	enableWriteFile bool
	path            string
	fileName        string
	writeFileLevel  int
	level           int
}
