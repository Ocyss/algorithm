package main

import (
	"fmt"
	codeforces "github.com/Ocyss/algorithm/generator/cf"
	"github.com/Ocyss/algorithm/generator/leetcode"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	LeetcodeWeeklypath   = "../周赛/%d/"
	LeetcodeBiweeklyPath = "../双周赛/%d/"
	CFPath               = "../CF/contest/<id>/"
	Comment              = "https://github.com/Ocyss"
	GoLandPath           = "D:\\IDE\\GoLand\\bin\\goland64.exe"
)

var (
	leetcodeUsername = os.Getenv("LEETCODE_USERNAME")
	leetcodePassword = os.Getenv("LEETCODE_PASSWORD")
)

func main() {
	var selectId int
	fmt.Print(
		`感谢 灵茶山艾府🎈 的代码生成器 🥰 >
1.LeetCode 周赛
2.LeetCode 双周赛
3.CodeForces
请选择你要参加的比赛> `)
	_, err := fmt.Scanln(&selectId)
	Er(err)
	switch selectId {
	case 1, 2:
		var testid int
		fmt.Print("请输入参赛id (直接回城为本次比赛)：")
		_, err = fmt.Scanln(&testid)
		Er(err)
		err = genLeetCodeTests(selectId == 1, testid)
		Er(err)
	case 3:
		var sid int
		fmt.Print("请选择模式 1.单题模板/2.竞赛模板：")
		_, _ = fmt.Scanln(&sid)
		if sid == 1 {
			var d string
			fmt.Print("请输入 网址 or ID/题号：")
			_, _ = fmt.Scanln(&d)
			_, err := url.ParseRequestURI(d)
			if err != nil {
				d = "https://codeforces.com/problemset/problem/" + strings.ToUpper(d)
			}
			genCodeForcesSOLOTests(d, err != nil)
		} else if sid == 2 {
			fmt.Println("竞赛模板")
		} else {
			panic("模式选择错误")
		}
	default:
		panic("错误的选择")
	}
}

func genLeetCodeTests(weekly bool, testID int) error {
	var tag, dir string
	if weekly {
		contestID := leetcode.GetWeeklyContestID(testID) // 自动生成下一场周赛 ID
		tag = leetcode.GetWeeklyContestTag(contestID)
		dir = fmt.Sprintf(LeetcodeWeeklypath, contestID) // 自定义生成目录
	} else {
		contestID := leetcode.GetBiweeklyContestID(testID) // 自动生成下一场双周赛 ID
		tag = leetcode.GetBiweeklyContestTag(contestID)
		dir = fmt.Sprintf(LeetcodeBiweeklyPath, contestID) // 自定义生成目录
	}
	dirAbs, _ := filepath.Abs(dir)
	err := os.MkdirAll(dirAbs, os.ModePerm)
	if err == nil {
		cmd := exec.Command(GoLandPath, dirAbs)
		_, _ = cmd.CombinedOutput()
	}

	return leetcode.GenLeetCodeTests(leetcodeUsername, leetcodePassword, tag, true, dir, Comment)

}
func genCodeForcesSOLOTests(problemURL string, openWebsite bool) {
	err := codeforces.GenCodeforcesProblemTemplates(problemURL, openWebsite)
	Er(err)
}

func Er(e error) {
	if e != nil {
		panic(e)
	}
}
