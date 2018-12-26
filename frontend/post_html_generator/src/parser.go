package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

var (
	keys = []string{"title", "slug", "date", "tags"}
)

// Returns yaml frontmatter. Error is not nil if post does not have yaml.
// yaml starts with '\n'
func getYAML(post []byte) (string, error) {
	postString := string(post)
	if strings.Index(postString, "---") == -1 {
		return "", errors.New("Cannot find YAML frontmatter")
	}

	i := strings.Index(postString, "---")
	begin := strings.Index(postString[i:], "\n")
	end := strings.Index(postString[i+begin:], "---")

	if begin >= end {
		return "", errors.New("Cannot find YAML frontmatter")
	}

	yaml := postString[i+begin : i+begin+end]
	return yaml, nil
}

// Return value of key in yaml. Error is not nil if can't find key.
func getValue(key string, yaml string) (string, error) {
	hashIdx := strings.Index(yaml, "\n"+key+":") + 1
	if hashIdx == 0 {
		return "", fmt.Errorf("Can not find key(%s)", key)
	}

	begin := strings.Index(yaml[hashIdx:], ":") + 1
	end := strings.Index(yaml[hashIdx:], "\n")

	value := yaml[hashIdx+begin : hashIdx+end]
	value = strings.TrimSpace(value)
	value = strings.Trim(value, "\"")
	return value, nil
}

// Return yaml key-value map
func buildYAMLMap(yaml string) (map[string]string, error) {
	yamlMap := make(map[string]string)
	for _, key := range keys {
		value, err := getValue(key, yaml)
		if err != nil {
			return nil, err
		}
		yamlMap[key] = value
	}
	return yamlMap, nil
}

// Make JSON key:value pair string. e.g) "key": "value"
func buildJSONString(key string, value string) string {
	if key != "tags" {
		return `"` + key + `": "` + value + `"`
	}

	// Handle value list
	value = strings.Trim(value, "[]")
	tags := strings.Split(value, ",")

	valueJSON := "["

	for i, tag := range tags {
		tags[i] = `"` + strings.TrimSpace(tag) + `"`
	}
	valueJSON += strings.Join(tags, ",")
	valueJSON += "]"

	return `"` + key + `": ` + valueJSON
}

// Get YAML JSON bytes from input
func parse(post []byte) ([]byte, error) {
	yaml, err := getYAML(post)
	if err != nil {
		log.Println("[parse] ", err)
		return nil, errors.New("Fail to parse post")
	}
	yamlMap, err := buildYAMLMap(yaml)
	if err != nil {
		log.Println("[parse] ", err)
		return nil, errors.New("Fail to parse post")
	}

	var jsonStrings []string
	for _, key := range keys {
		value := yamlMap[key]
		jsonStrings = append(jsonStrings, buildJSONString(key, value))
	}
	yamlJSON := "{" + strings.Join(jsonStrings, ",") + "}"

	return []byte(yamlJSON), nil
}
