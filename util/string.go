package util

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
)

const (
	Black = iota + 30
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	White
)

func Formatter(src string, params ...any) string {
	if len(params) == 0 || !strings.Contains(src, "{}") {
		return src
	}

	for _, item := range params {
		src = strings.Replace(src, "{}", anyConvertToString(item), 1)
	}

	return src
}

func anyConvertToString(v any) string {
	switch v.(type) {
	case string:
		return v.(string)
	default:
		kind := reflect.TypeOf(v).Kind()
		if kind == reflect.Struct {
			bytes, err := json.Marshal(v)
			if err != nil {
				log.Printf("json marshal %v error: %s", v, err)
			}
			return string(bytes)
		}
		return fmt.Sprint(v)
	}
}

func ColorText(color int, src string) string {
	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, src)
}
