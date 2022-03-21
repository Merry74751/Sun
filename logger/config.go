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
	writeFileLevel  int
	path            string
	fileName        string
	level           int
}
