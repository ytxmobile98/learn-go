package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"example.com/utils"
)

type (
	LogLevel    uint8
	LogLevelStr string
)

func (l LogLevel) String() string {
	if s, ok := logLevels.GetKey(l); ok {
		return string(s)
	}
	return ""
}

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

var logLevels = utils.BuildBidiMap(
	[]utils.BidiMapItem[LogLevelStr, LogLevel]{
		{Key: "debug", Value: LogLevelDebug},
		{Key: "info", Value: LogLevelInfo},
		{Key: "warn", Value: LogLevelWarn},
		{Key: "error", Value: LogLevelError},
		{Key: "fatal", Value: LogLevelFatal},
	},
)

func (l *LogLevel) UnmarshalJSON(data []byte) error {
	// first try to unmarshal as its usual type
	var v any
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	// then check type and do the conversion
	switch level := v.(type) {
	case float64:
		logLevel := LogLevel(level)
		if _, ok := logLevels.GetKey(logLevel); ok {
			*l = logLevel
			return nil
		}
	case string:
		logLevelStr := LogLevelStr(strings.ToLower(level))
		if logLevel, ok := logLevels.GetValue(logLevelStr); ok {
			*l = logLevel
			return nil
		}
	}
	return fmt.Errorf("invalid log level: %v", v)
}

func (l *LogLevel) MarshalJSON() ([]byte, error) {
	if l == nil {
		return json.Marshal(nil)
	}

	if logLevelStr, ok := logLevels.GetKey(*l); ok {
		return json.Marshal(logLevelStr)
	}
	return nil, fmt.Errorf("invalid log level: %v", *l)
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
			[]byte(`{"logLevel":1}`),      // Should output: {LogLevel:info}
			[]byte(`{"logLevel":"warn"}`), // Should output: {LogLevel:warn}
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
			{LogLevel: LogLevelWarn},  // Should output: {"logLevel":"warn"}
			{LogLevel: LogLevelFatal}, // Should output: {"logLevel":"fatal"}
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
