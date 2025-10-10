package main

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"gopkg.in/yaml.v3"
)

func getDataFile() string {
	_, file, _, _ := runtime.Caller(0)
	dir := path.Dir(file)
	return path.Join(dir, "data.yaml")
}

func unmarshalYaml(filename string) (any, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var a yaml.Unmarshaler = &CustomData{}
	err = yaml.Unmarshal(bytes, a)
	return a, err
}

type CustomData map[any]any

func (c *CustomData) UnmarshalYAML(value *yaml.Node) error {
	fmt.Printf("Value: %+v\n", value)
	return nil
}

func main() {
	yamlFile := getDataFile()
	data, err := unmarshalYaml(yamlFile)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
