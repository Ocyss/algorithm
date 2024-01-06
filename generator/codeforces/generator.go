package codeforces

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Ocyss/OI/generator/utils"
	"github.com/Ocyss/OI/generator/utils/tool"
	"github.com/skratchdot/open-golang/open"
)

const (
	CmdCodeforces = "cf"  // https://github.com/xalanq/cf-tool
	CmdAtcoder    = "atc" // https://github.com/sempr/cf-tool rename
)

// GenCodeforcesContestTemplates 生成 CF 比赛模板（需要先 cf race，以确认题目数量）
func GenCodeforcesContestTemplates(cmdName, rootPath, contestID string, overwrite bool) error {
	if contestID == "" {
		fmt.Println("contest ID is empty")
		return nil
	}

	// openedOneFile := false
	openedOneFile := true

	return filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == rootPath || !info.IsDir() {
			return nil
		}

		parentName := filepath.Base(path)
		for _, srcFileName := range []string{"main.go", "main_test.go"} {
			// 为了便于区分，把 main 替换成所在目录的名字
			dstFileName := strings.Replace(srcFileName, "main", parentName, 1)
			dstFilePath := filepath.Join(path, dstFileName)
			if !overwrite {
				if _, err := os.Stat(dstFilePath); !os.IsNotExist(err) {
					continue
				}
			}
			if err := copyFile(dstFilePath, srcFileName); err != nil {
				return err
			}
			if !openedOneFile {
				openedOneFile = true
				open.Run(absPath(dstFilePath))
			}
		}
		cmd := fmt.Sprintf("%s submit contest %s %s -f %s.go", cmdName, contestID, parentName, parentName)
		if err := os.WriteFile(filepath.Join(path, parentName+".bat"), []byte(cmd), 0644); err != nil {
			return err
		}
		return nil
	})
}

// GenCodeforcesProblemTemplates 生成单道题目的模板（Codeforces）
func GenCodeforcesProblemTemplates(problemURL string, openWebsite bool) error {
	urlObj, err := url.Parse(problemURL)
	if err != nil {
		return err
	}

	contestID, problemID, isGYM := parseCodeforcesProblemURL(problemURL)
	if _, err := strconv.Atoi(contestID); err != nil {
		return fmt.Errorf("invalid URL: %v", err)
	}

	luoguURL := fmt.Sprintf("https://www.luogu.com.cn/problem/CF%s%s", contestID, problemID)

	var statusURL string
	if isGYM {
		statusURL = fmt.Sprintf("https://%s/gym/%s/status/%s", urlObj.Host, contestID, problemID)
	} else {
		statusURL = fmt.Sprintf("https://%s/problemset/status/%s/problem/%s", urlObj.Host, contestID, problemID)
	}

	if openWebsite {
		//_ = open.Run(luoguURL)
		_ = open.Run(statusURL)
		//_ = open.Run(problemURL)
	}

	example, err := parseExamples(luoguURL)
	if err != nil {
		return err
	}

	exampleStr := ""
	for _, p := range example {
		in := strings.TrimSpace(p[0])
		out := strings.TrimSpace(p[1])
		exampleStr += "\n\t\t{\n"
		exampleStr += "\t\t\t`" + in + "`,\n"
		exampleStr += "\t\t\t`" + out + "`,\n"
		exampleStr += "\t\t},"
	}

	if !isGYM {
		problemID = contestID + problemID
	}
	mainStr := fmt.Sprintf(`package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CF%[1]s(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
}

func main() { CF%[1]s(os.Stdin, os.Stdout) }
`, problemID)
	mainTestStr := fmt.Sprintf(`package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// %s
// %s
func TestCF%[3]s(t *testing.T) {
	testCases := [][2]string{%s
}
	testutil.AssertEqualStringCase(t, testCases, 0, CF%[3]s)
}
`, problemURL, statusURL, problemID, exampleStr)

	var dir string
	if isGYM {
		dir = utils.Config.Codeforces.Path + "gym/" + contestID + "/"
	} else {
		dir = utils.Config.Codeforces.Path + genDirName(contestID) + "/" + problemID + "/"
	}
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	mainFilePath := dir + problemID + ".go"
	if _, err := os.Stat(mainFilePath); !os.IsNotExist(err) {
		open.Run(absPath(mainFilePath))
		return fmt.Errorf("文件已存在！")
	}
	if err := os.WriteFile(mainFilePath, []byte(mainStr), 0o644); err != nil {
		return err
	}
	//_ = open.Run(absPath(mainFilePath))

	testFilePath := dir + problemID + "_test.go"
	if err := os.WriteFile(testFilePath, []byte(mainTestStr), 0o644); err != nil {
		return err
	}
	tool.OpenGoLand(utils.Config.Codeforces.Path, mainFilePath, testFilePath) // 打开GoLand
	//_ = open.Run(absPath(testFilePath))
	return nil
}

// GenTemplates 在某一路径下批量生成模板
func GenTemplates(problemNum int, rootPath string, overwrite bool) error {
	for i := 'a'; i < 'a'+int32(problemNum); i++ {
		dir := rootPath + string(i) + "/"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
		for j, fileName := range []string{"main.go", "main_test.go"} {
			goFilePath := dir + strings.Replace(fileName, "main", string(i), 1)
			if !overwrite {
				if _, err := os.Stat(goFilePath); !os.IsNotExist(err) {
					continue
				}
			}
			if err := copyFile(goFilePath, "codeforces/templates/"+fileName); err != nil {
				return err
			}
			if i == 'a' && j == 0 {
				//_ = open.Run(absPath(goFilePath))
				tool.OpenGoLand(utils.Config.Codeforces.Path, goFilePath)
			}
		}
	}
	return nil
}
