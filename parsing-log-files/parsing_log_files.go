package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`\<(\~|\*|\=|\-)*\>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)\".*password.*\"`)
	count := 0

	for _, line := range lines {
		count += len(re.FindAllString(line, -1))
	}

	return count

}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]*`)
	return re.ReplaceAllLiteralString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User +[a-zA-Z0-9]* `)
	multipleSpacesRe := regexp.MustCompile(` +`)
	spaceRe := regexp.MustCompile(` `)

	var resultLines []string

	for _, v := range lines {
		result := re.FindStringSubmatch(v)
		if nil == result {
			resultLines = append(resultLines, v)
			continue
		}
		resultMultipleSpaces := multipleSpacesRe.ReplaceAllLiteralString(result[0], " ")
		resultSpace := spaceRe.Split(resultMultipleSpaces, -1)
		line := "[USR] " + resultSpace[1] + " " + v
		resultLines = append(resultLines, line)
	}

	return resultLines
}
