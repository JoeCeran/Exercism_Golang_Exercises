package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	match := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return match.MatchString(text)
}

func SplitLogLine(text string) []string {
	match := regexp.MustCompile(`<[-=~*]*>`)
	return match.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	match := regexp.MustCompile(`(?i)".*password.*"`)
	var i int
	for _, v := range lines {
		if match.MatchString(v) == true {
			i = i + 1
		}
	}
	return i
}

func RemoveEndOfLineText(text string) string {
	match := regexp.MustCompile(`end-of-line\d*`)
	return match.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	match := regexp.MustCompile(`User\s+(\w+)`)
	ret := []string{}
	for _, line := range lines {
		found := match.FindStringSubmatch(line)
		if found != nil {
			line = fmt.Sprintf("[USR] %s %s", found[1], line)
		}
		ret = append(ret, line)
	}
	return ret
}
