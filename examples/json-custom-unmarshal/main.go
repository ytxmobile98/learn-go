package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type (
	LogLevel    uint8
	LogLevelStr string
)

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

const (
	LogLevelDebugStr LogLevelStr = "debug"
	LogLevelInfoStr  LogLevelStr = "info"
	LogLevelWarnStr  LogLevelStr = "warn"
	LogLevelErrorStr LogLevelStr = "error"
	LogLevelFatalStr LogLevelStr = "fatal"
)

func (l *LogLevel) UnmarshalJSON(data []byte) error {
	// first try to unmarshal as its usual type
	var level any
	err := json.Unmarshal(data, &level)
	if err != nil {
		return err
	}

	// then check type and do the conversion
	switch logLevel := level.(type) {
	case float64:
		*l = LogLevel(logLevel)
	case string:
		logLevelStr := LogLevelStr(strings.ToLower(logLevel))
		switch logLevelStr {
		case LogLevelDebugStr:
			*l = LogLevelDebug
		case LogLevelInfoStr:
			*l = LogLevelInfo
		case LogLevelWarnStr:
			*l = LogLevelWarn
		case LogLevelErrorStr:
			*l = LogLevelError
		case LogLevelFatalStr:
			*l = LogLevelFatal
		default:
			return fmt.Errorf("invalid log level: %s", level)
		}
	default:
		return fmt.Errorf("invalid log level: %v", level)
	}
	return nil
}

func (l *LogLevel) MarshalJSON() ([]byte, error) {
	if l == nil {
		return json.Marshal(nil)
	}

	switch *l {
	case LogLevelDebug:
		return json.Marshal(LogLevelDebugStr)
	case LogLevelInfo:
		return json.Marshal(LogLevelInfoStr)
	case LogLevelWarn:
		return json.Marshal(LogLevelWarnStr)
	case LogLevelError:
		return json.Marshal(LogLevelErrorStr)
	case LogLevelFatal:
		return json.Marshal(LogLevelFatalStr)
	default:
		return nil, fmt.Errorf("invalid log level: %v", *l)
	}
}

type Test struct {
	LogLevel LogLevel `json:"logLevel"`
}

func unmarshal(data []byte) (test Test, err error) {
	err = json.Unmarshal(data, &test)
	return
}

func marshal(test *Test) ([]byte, error) {
	return json.Marshal(test)
}

func main() {
	{ // test unmarshal data
		fmt.Println("========== unmarshal data ==========")
		for _, data := range [][]byte{
			[]byte(`{"logLevel":1}`),      // Should output: {LogLevel:1}
			[]byte(`{"logLevel":"warn"}`), // Should output: {LogLevel:2}
		} {
			test, err := unmarshal(data)
			if err != nil {
				fmt.Println("error:", err)
			} else {
				fmt.Printf("%+v\n", test)
			}
		}
	}

	fmt.Println()

	{ // test marshal data
		fmt.Println("========== marshal data ==========")
		for _, test := range []Test{
			{LogLevel: LogLevelWarn},
			{LogLevel: LogLevelFatal},
		} {
			data, err := marshal(&test)
			if err != nil {
				fmt.Println("error:", err)
			} else {
				fmt.Printf("%s\n", data)
			}
		}
	}
}
