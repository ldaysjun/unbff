package main

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

func main() {
	tmpl, err := template.New("tmpl").Parse("{\n\t\"user\": {\n\t\t\"firstName\": \"{{.id}}\",\n\t\t\"lastName\": \"{{.id}}\"\n\t}\n}")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	var tmplBytes bytes.Buffer
	if err := tmpl.Execute(&tmplBytes, map[string]string{
		"id": "id",
	}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tmplBytes.String())
}

func jsonPath(m map[string]interface{}) interface{} {
	path := "data.results"
	pathSplit := strings.Split(path, ".")
	temp := m
	for index, p := range pathSplit {
		v, ok := temp[p]
		if !ok {
			break
		}
		if index == len(pathSplit)-1 {
			return v
		}
		switch value := v.(type) {
		case nil:
			break
		case string:
			break
		case int:
			break
		case float64:
			break
		case []interface{}:
			break
		case map[string]interface{}:
			temp = value
		default:
			fmt.Println(p, "is unknown type", fmt.Sprintf("%T", v))
		}
	}
	return nil
}
