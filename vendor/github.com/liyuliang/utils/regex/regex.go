package regex

import (
	"regexp"
)

func Get(content string, expr string) string {
	result := ""
	rule, _ := regexp.Compile(expr)
	allMatch := rule.FindStringSubmatch(content)
	if 2 == len(allMatch) {
		result = allMatch[1]
	}
	return result
}

func GetAll(content string, expr string) (results []string) {
	rule := regexp.MustCompile(expr)
	allMatch := rule.FindAllStringSubmatch(content,-1)
	for _, value := range allMatch {
		if len(value) > 1 {
			results = append(results,value[1])
		}else{
			results = append(results,value[0])
		}
	}
	return results
}

func Replace(content string, expr string, replacement string) string {
	rule, _ := regexp.Compile(expr)
	return rule.ReplaceAllString(content, replacement)
}
