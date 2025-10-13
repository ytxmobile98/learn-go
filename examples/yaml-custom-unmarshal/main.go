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
	Items  Anchors[Person]      `yaml:"items"`
	Groups map[string][]*Person `yaml:"groups"`
}

type Person struct {
	Name string `yaml:"name"`
	Id   int    `yaml:"id"`
}

type Anchors[T any] map[string]*T

func (a Anchors[T]) PutAnchor(node *yaml.Node) error {
	if node.Anchor == "" {
		return nil
	}

	var value T
	err := node.Decode(&value)
	if err != nil {
		return err
	}

	a[node.Anchor] = &value
	return nil
}

func (a Anchors[T]) ResolveAlias(alias string) *T {
	if a == nil {
		return nil
	}
	return a[alias]
}

func (d *Data) UnmarshalYAML(value *yaml.Node) error {
	if d.Items == nil {
		d.Items = make(Anchors[Person])
	}
	if d.Groups == nil {
		d.Groups = make(map[string][]*Person)
	}

	// process the `items` and `groups` nodes
	var temp struct {
		Items  []yaml.Node          `yaml:"items"`
		Groups map[string]yaml.Node `yaml:"groups"`
	}
	err := value.Decode(&temp)
	if err != nil {
		return err
	}

	// process items, and save them into the anchors map
	for _, node := range temp.Items {
		err := d.Items.PutAnchor(&node)
		if err != nil {
			return err
		}
	}

	// resolve group members from anchors map
	for groupName, nodes := range temp.Groups {
		group := make([]*Person, 0, len(nodes.Content))

		for _, node := range nodes.Content {
			if node.Kind != yaml.AliasNode {
				continue
			}

			person := d.Items.ResolveAlias(node.Value)
			if person == nil {
				return fmt.Errorf("undefined anchor: %s", node.Value)
			}
			group = append(group, person)
		}
		d.Groups[groupName] = group
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
	fmt.Printf("%+v\n", data)
}
