package main

import (
	"fmt"

	"github.com/go-zoox/fetch"
	"github.com/tidwall/gjson"
)

func GetCppResponses(code string) string {
	// https://api-berita-indonesia.vercel.app/cnn/terbaru/
	response, err := fetch.Post(fmt.Sprintf("https://onecompiler.com/api/code/exec"), &fetch.Config{
		Body: map[string]interface{}{
			"name":         "C++",
			"title":        "C++ Hello World",
			"version":      "latest",
			"mode":         "c_cpp",
			"description":  "",
			"extension":    "cpp",
			"languageType": "programming",
			"active":       true,
			"visibility":   "public",
			"properties": map[string]interface{}{
				"language":    "cpp",
				"docs":        true,
				"tutorials":   true,
				"cheatsheets": true,
				"files": []map[string]interface{}{
					{
						"name":    "Main.cpp",
						"content": code,
					},
				},
			},
		},
	})

	json := string(response.Body)

	if err != nil || !gjson.Valid(json) {
		Error("Cant get cpp response : %s", response.Error().Error())
		return "Gak bisa jalanin kodemu, fixkan dulu"
	}

	s := ""
	stdout := gjson.Get(json, "stdout").String()
	stderr := gjson.Get(json, "stderr").String()

	if stdout != "" {
		s += fmt.Sprintf("```%s```", stdout)
	} else if stderr != "" {
		s += fmt.Sprintf("```Aaa tidak!\n%s```", stderr)
	}

	return s
}

func GetJsResponses(code string) string {
	// https://api-berita-indonesia.vercel.app/cnn/terbaru/
	response, err := fetch.Post(fmt.Sprintf("https://onecompiler.com/api/code/exec"), &fetch.Config{
		Body: map[string]interface{}{
			"name":         "JavaScript",
			"title":        "3zrhvv9ny",
			"version":      "ES6",
			"mode":         "javascript",
			"description":  "",
			"extension":    "js",
			"languageType": "programming",
			"active":       true,
			"visibility":   "public",
			"properties": map[string]interface{}{
				"language":       "javascript",
				"docs":           true,
				"tutorials":      true,
				"cheatsheets":    true,
				"filesEditable":  true,
				"filesDeletable": true,
				"files": []map[string]interface{}{
					{
						"name":    "index.js",
						"content": code,
					},
				},
			},
		},
	})

	json := string(response.Body)

	if err != nil || !gjson.Valid(json) {
		Error("Cant get cpp response : %s", response.Error().Error())
		return "Gak bisa jalanin kodemu, fixkan dulu"
	}

	s := ""
	stdout := gjson.Get(json, "stdout").String()
	stderr := gjson.Get(json, "stderr").String()

	if stdout != "" {
		s += fmt.Sprintf("```%s```", stdout)
	} else if stderr != "" {
		s += fmt.Sprintf("```Aaa tidak!\n%s```", stderr)
	}

	return s
}

func GetKotlinResponses(code string) string {
	// https://api-berita-indonesia.vercel.app/cnn/terbaru/
	response, err := fetch.Post(fmt.Sprintf("https://onecompiler.com/api/code/exec"), &fetch.Config{
		Body: map[string]interface{}{
			"name":         "Kotlin",
			"title":        "Kotlin Hello World!",
			"mode":         "groovy",
			"description":  "",
			"extension":    "kt",
			"languageType": "programming",
			"active":       true,
			"visibility":   "public",
			"properties": map[string]interface{}{
				"language":    "kotlin",
				"docs":        false,
				"tutorials":   false,
				"cheatsheets": false,
				"files": []map[string]interface{}{
					{
						"name":    "HelloWorld.kt",
						"content": code,
					},
				},
			},
		},
	})

	json := string(response.Body)

	if err != nil || !gjson.Valid(json) {
		Error("Cant get cpp response : %s", response.Error().Error())
		return "Gak bisa jalanin kodemu, fixkan dulu"
	}

	s := ""
	stdout := gjson.Get(json, "stdout").String()
	stderr := gjson.Get(json, "stderr").String()

	if stdout != "" {
		s += fmt.Sprintf("```%s```", stdout)
	} else if stderr != "" {
		s += fmt.Sprintf("```Aaa tidak!\n%s```", stderr)
	}

	return s
}
