package main

import (
	"fmt"
	"github.com/Ocyss/algorithm/generator/codeforces"
	"github.com/Ocyss/algorithm/generator/leetcode"
	"github.com/Ocyss/algorithm/generator/utils"
	"github.com/Ocyss/algorithm/generator/utils/tool"
	"net/url"
	"os"

	"strings"
)

var conf utils.Conf

func main() {
	conf = utils.InitSettings()
	var selectId int
top:
	fmt.Print(`感谢 灵茶山艾府🎈 的模板生成器 🥰 >
1.LeetCode 周赛
2.LeetCode 双周赛
3.CodeForces
请选择你要参加的比赛> `)
	_, err := fmt.Scan(&selectId)
	if err != nil {
		fmt.Printf("\n错误输入,%v!\n\n", err)
		goto top
	}
	switch selectId {
	case 1, 2:
		genLeetCodeTests(selectId)
	case 3:
		genCodeForcesTests()
	default:
		fmt.Print("\n错误选择!\n\n")
		goto top
	}
}

func genLeetCodeTests(selectId int) {
	var (
		leetcodeUsername = os.Getenv("LEETCODE_USERNAME")
		leetcodePassword = os.Getenv("LEETCODE_PASSWORD")
		testID           int
		contestID        int
		tag              string
		dir              string
	)
	fmt.Print("请输入参赛id (直接回城为本次比赛)：")
	_, _ = fmt.Scanln(&testID)
	if selectId == 1 {
		contestID = leetcode.GetWeeklyContestID(testID) // 自动生成下一场周赛 ID
		tag = leetcode.GetWeeklyContestTag(contestID)
		dir = fmt.Sprintf(utils.Config.Leetcode.Path+"weekly/%d/", contestID) // 自定义生成目录
	} else {
		contestID = leetcode.GetBiweeklyContestID(testID) // 自动生成下一场双周赛 ID
		tag = leetcode.GetBiweeklyContestTag(contestID)
		dir = fmt.Sprintf(utils.Config.Leetcode.Path+"biweekly/%d/", contestID) // 自定义生成目录
	}

	tool.Er(leetcode.GenLeetCodeTests(leetcodeUsername, leetcodePassword, tag, true, dir))
	tool.OpenGoLand(utils.Config.Leetcode.Path, dir+"a/a.go") // 打开GoLand
}

func genCodeForcesTests() {
	var sid int
	fmt.Print("请选择模式 1.单题模板/2.竞赛模板：")
	_, _ = fmt.Scan(&sid)
	if sid == 1 {
		var d string
		fmt.Print("请输入 网址 or ID/题号：")
		_, _ = fmt.Scan(&d)
		_, err := url.ParseRequestURI(d)
		if err != nil {
			d = "https://codeforces.com/problemset/problem/" + strings.ToUpper(d)
		}
		err = codeforces.GenCodeforcesProblemTemplates(d, err != nil)
		tool.Er(err)
	} else if sid == 2 {
		var d string
		fmt.Print("请输入 ID：")
		_, _ = fmt.Scan(&d)
		rootPath := conf.Codeforces.Path + "contests/" + d + "/"
		tool.Er(os.Mkdir(rootPath, os.ModePerm))
		var b int
		fmt.Print("请输入 数量：")
		_, _ = fmt.Scan(&b)
		err := codeforces.GenTemplates(b, rootPath, true)
		tool.Er(err)

	} else {
		panic("模式选择错误")
	}

}
