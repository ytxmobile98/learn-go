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

func unmarshalYaml[T any](filename string) (*T, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var a T
	err = yaml.Unmarshal(bytes, &a)
	return &a, err
}

type Data struct {
	Items  Anchors `yaml:"items"`
	Groups Refs    `yaml:"groups"`
}

type Person struct {
	Name string `yaml:"name"`
	Id   int    `yaml:"id"`
}
type Anchors map[string]*Person
type Refs map[string][]*Person

func (a *Anchors) UnmarshalYAML(value *yaml.Node) error {
	*a = make(Anchors)

	for _, c := range value.Content {
		anchorName := c.Anchor

		var person Person
		err := c.Decode(&person)
		if err != nil {
			return err
		}

		(*a)[anchorName] = &person
	}
	return nil
}

func main() {
	yamlFile := getDataFile()

	var data *Data
	data, err := unmarshalYaml[Data](yamlFile)
	if err != nil {
		panic(err)
	}

	for anchor, item := range data.Items {
		fmt.Printf("%s: (%p) %+v\n", anchor, item, *item)
	}
	for group, items := range data.Groups {
		fmt.Println(group, items)
		for _, item := range items {
			fmt.Println("  ", item)
		}
	}
}
