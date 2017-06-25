package file

import (
	"flashCoder/utils"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type FlashFile struct {
	FileType string
}

func (f *FlashFile) DeleteLines(contents []string, lines []int) []string {
	if len(contents) == 0 || len(lines) == 0 || len(contents) <= len(lines) {
		return nil
	}
	sort.Ints(lines)
	res := make([]string, len(contents)-len(lines))
	lastLine, lastNum := 0, 0
	for k, line := range lines {
		if k > 0 {
			copy(res[lastNum:], contents[lastLine:line-1])
			lastNum += len(contents[lastLine : line-1])
		} else {
			if line > 1 {
				lastNum = len(contents[0 : line-1])
				copy(res, contents[0:line-1])
			} else {
				copy(res, contents[line:])
			}
		}
		lastLine = line
	}
	copy(res[lastNum:], contents[lastLine:len(contents)])
	return res
}

func (f *FlashFile) AddLines(contents []string, line int, content string, isPre bool, newLine bool) []string {
	if line <= 0 {
		return nil
	}
	addContents := strings.Split(content, "\n")
	for k, v := range addContents {
		addContents[k] = "\t\t" + v
	}
	var res []string
	if newLine {
		res = make([]string, len(contents)+len(addContents))
	} else {
		res = make([]string, len(contents)+len(addContents)-1)
	}
	if len(contents) >= line {
		if isPre && newLine {
			copy(res, contents[:line-1])
			total := len(contents[:line-1])
			copy(res[total:], addContents)
			total += len(addContents)
			copy(res[total:], contents[line-1:])
		} else if isPre && !newLine {
			copy(res, contents[:line-1])
			total := len(contents[:line-1])
			addContents[len(addContents)-1] += contents[line-1]
			copy(res[total:], addContents)
			total += len(addContents)
			if line < len(contents) {
				copy(res[total:], contents[line:])
			}
		} else if !isPre && newLine {
			copy(res, contents[:line])
			total := len(contents[:line])
			copy(res[total:], addContents)
			total += len(addContents)
			if line < len(contents) {
				copy(res[total:], contents[line:])
			}
		} else { // next, !newline
			copy(res, contents[:line])
			total := len(contents[:line])
			res[total-1] += addContents[0]
			addNewContents := make([]string, len(addContents)-1)
			copy(addNewContents, addContents[1:])
			copy(res[total:], addNewContents)
			total += len(addNewContents)
			if line < len(contents) {
				copy(res[total:], contents[line:])
			}
		}
	} else {
		if newLine {
			copy(res, contents)
			total := len(contents)
			copy(res[total:], addContents)
		} else {
			copy(res, contents)
			total := len(contents)
			res[total-1] += addContents[0]
			addNewContents := make([]string, len(addContents)-1)
			copy(addNewContents, addContents[1:])
			copy(res[total:], addNewContents)
		}

	}

	return res

}

func (f *FlashFile) FindFuncLine(contents []string, funcName string, funcRule string) (int, int) {
	lineBegin, lineEnd := 0, 0
	beginTag, endTag := "{", "}"
	lfunc := strings.ToLower(funcName)
	for k, v := range contents {
		lv := strings.ToLower(v)
		match, err := regexp.Match(funcRule, []byte(lv))
		if err != nil {
			fmt.Println(err)
			continue
		}
		if match && strings.Contains(lv, lfunc) {
			if strings.Contains(lv, beginTag) {
				lineBegin = k + 1
			} else {
				tmp := contents[k+1:]
				for k1, v1 := range tmp {
					lv1 := strings.ToLower(v1)
					if strings.Contains(lv1, beginTag) {
						lineBegin += (k + k1 + 2)
						break
					}
				}
			}
			break
		}
	}

	stat := 0
	if lineBegin > 0 {
		tmp := contents[lineBegin-1:]
		for k2, v2 := range tmp {
			lv2 := strings.ToLower(v2)
			if strings.Contains(lv2, beginTag) {
				stat += 1
			}
			if strings.Contains(lv2, endTag) {
				stat -= 1
			}

			if stat == 0 {
				lineEnd += (lineBegin + k2)
				break
			}
		}
	}
	return lineBegin, lineEnd
}

func (f *FlashFile) AddFuncContent(contents []string, funcName string, content string, isBegin bool, offset int) []string {
	funcRule := f.getFuncRule()
	lineBegin, lineEnd := f.FindFuncLine(contents, funcName, funcRule)
	if lineBegin > 0 && lineEnd > 0 {
		if isBegin {
			return f.AddLines(contents, lineBegin+offset, content, false, true)
		} else {
			return f.AddLines(contents, lineEnd-offset, content, true, true)
		}
	}
	return nil
}

func (f *FlashFile) getFuncRule() string {
	fileType := strings.ToLower(f.FileType)
	switch fileType {
	case ".php":
		return "function\\s+[a-zA-Z_]+\\s*\\("
	}
	utils.CheckError("err", "no explicit file type")
	return ""
}
