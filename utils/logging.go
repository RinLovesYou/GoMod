package utils

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type LogMessage struct {
	TimeStamp string
	Prefix    string
	Message   string
}

var LogBuffer = make([]LogMessage, 0)

func pushLog(message LogMessage) {
	//if the buffer is bigger than 100, remove the first element
	if len(LogBuffer) > 100 {
		LogBuffer = LogBuffer[1:]
	}

	LogBuffer = append(LogBuffer, message)
}

func Log(message string, format ...any) {
	fmt.Printf("%s %s %s\n", timeStamp(false), prefix(false, "GoMod"), fmt.Sprintf(message, format...))
	pushLog(LogMessage{
		TimeStamp: timeStampLog(),
		Prefix:    "GoMod",
		Message:   fmt.Sprintf(message, format...),
	})
}

func Error(message string, format ...any) {
	fmt.Printf("%s %s %s\n", timeStamp(true), prefix(true, "GoMod"), fmt.Sprintf(message, format...))
	pushLog(LogMessage{
		TimeStamp: timeStampLog(),
		Prefix:    "GoMod",
		Message:   fmt.Sprintf(message, format...),
	})
}

func timeStamp(err bool) string {
	if err {
		return fmt.Sprintf("%s[%s]", Color(Red), time.Now().Format("15:04:05.000"))
	}

	return fmt.Sprintf("%s[%s%s%s]", Color(Gray), Color(Green), time.Now().Format("15:04:05.000"), Color(Gray))
}

func timeStampLog() string {
	return time.Now().Format("15:04:05.000")
}

func prefix(err bool, prefix string) string {
	if err {
		return fmt.Sprintf("%s[%s]", Color(Red), prefix)
	}
	return fmt.Sprintf("%s[%s%s%s]", Color(Gray), Color(Pink), prefix, Color(Gray))
}

func prefixLog(prefix string) string {
	return fmt.Sprintf("[%s]", prefix)
}

func Color(color ConsoleColor) string {
	if col, ok := ConsoleAnsiiDict[color]; ok {
		return col
	} else {
		return "\x1b[97m" //white
	}
}

var (
	ConsoleAnsiiDict = map[ConsoleColor]string{
		Black:       "\x1b[30m",
		DarkBlue:    "\x1b[34m",
		DarkGreen:   "\x1b[32m",
		DarkCyan:    "\x1b[36m",
		DarkRed:     "\x1b[31m",
		DarkMagenta: "\x1b[35m",
		DarkYellow:  "\x1b[33m",
		Gray:        "\x1b[37m",
		DarkGray:    "\x1b[90m",
		Blue:        "\x1b[94m",
		Green:       "\x1b[92m",
		Cyan:        "\x1b[96m",
		Red:         "\x1b[91m",
		Magenta:     "\x1b[95m",
		Yellow:      "\x1b[93m",
		White:       "\x1b[97m",
		Pink:        "\033[38;5;206m",
	}
)

type ConsoleColor int

const (
	Black ConsoleColor = iota
	DarkBlue
	DarkGreen
	DarkCyan
	DarkRed
	DarkMagenta
	DarkYellow
	Gray
	DarkGray
	Blue
	Green
	Cyan
	Red
	Magenta
	Yellow
	White
	Pink
)

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return false
	}
}
