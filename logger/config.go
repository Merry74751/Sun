package logger

var config *sumConfig

type sumConfig struct {
	enableWriteFile bool
	path            string
	fileName        string
	level           string
}
