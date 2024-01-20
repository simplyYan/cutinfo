package cutinfo

import (
	"strings"
)

type CutInfo struct{}

func New() *CutInfo {
	return &CutInfo{}
}

func (c *CutInfo) Target(text, startTag, endTag string) string {
	startIndex := strings.Index(text, startTag)
	if startIndex == -1 {
		return ""
	}

	startIndex += len(startTag)
	endIndex := strings.Index(text[startIndex:], endTag)
	if endIndex == -1 {
		return ""
	}

	return text[startIndex : startIndex+endIndex]
}
