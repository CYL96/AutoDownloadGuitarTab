package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	. "TEST/爬取吉他谱/run"
)

var nowPage int
var exit int

const (
	firstPage_num    = 0
	search_num       = 3
	searchResult_num = 4
)

func main() {
	for {
		switch nowPage {
		case 0:
			firstPage()
			if exit == 1 {
				return
			}
		case 1:
			newGT()
		case 2:
			hotGT()
		case 3:
			search()
		}
	}

	firstPage()
}
func Clear() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func firstPage() {
	nowPage = firstPage_num
	Clear()
	fmt.Println("[主页]")
	fmt.Println("1：获取当前最新")
	fmt.Println("2：获取当前最多下载")
	fmt.Println("3：搜索")
	fmt.Println("0:退出")
	a := 0
	fmt.Scan(&a)
	switch a {
	case 1:
		newGT()
	case 2:
		hotGT()
	case 3:
		search()
	case 0:
		fmt.Println("0:退出")
		time.Sleep(1 * time.Second)
		exit = 1
		return
	default:
		fmt.Println("选择错误：")
		time.Sleep(1 * time.Second)
		firstPage()
	}
}
func newGT() {
	searchResult(1, "", 1)

}
func hotGT() {
	searchResult(2, "", 1)

}
func search() {
	Clear()
	fmt.Println("0:返回上一级")
	fmt.Println("请输入搜索内容:")
	str := ""
	fmt.Scan(&str)
	switch str {
	case "0":
		nowPage = firstPage_num
	default:
		searchResult(3, str, 1)
	}
}

func searchResult(qType int, str string, page int) {
	Clear()

	var (
		result []SearchResultExt
		err    error
	)
	switch qType {
	case 1:
		//最新
		fmt.Println("正在搜索.....")
		result, err = GetNewGT(strconv.Itoa(page))
	case 2:
		//最多下载
		fmt.Println("正在搜索.....")
		result, err = GetHotGT(strconv.Itoa(page))
	case 3:
		//搜索
		fmt.Println("正在搜索.....")
		result, err = GetSearchResult(str, strconv.Itoa(page))
	}
	if err != nil {
		Clear()
		fmt.Println("搜错出错：", err.Error())
		fmt.Println("输入任何数返回上一级")
		ctrl := 0
		fmt.Scan(&ctrl)
		switch ctrl {
		case 0:
			if qType == 3 {
				nowPage = search_num
			} else {
				nowPage = firstPage_num
			}
			return
		}
	}
	Clear()
	fmt.Println("搜索结果：")
	fmt.Println("-----------------------------------------------------------------------------")
	for i := range result {
		fmt.Println(i+1, result[i].Name)
	}
	fmt.Println("-----------------------------------------------------------------------------")
	fmt.Println("（l:上一页 n:下一页 ）")
	fmt.Println("0:返回上一级")
	fmt.Println("输入对应编号进行下载(-1:下载全部)")
	fmt.Println("[第", page, "页]")
	ctrl := ""
	fmt.Scan(&ctrl)
	switch ctrl {
	case "l":
		if page-1 >= 1 {
			page--
		}
		searchResult(qType, str, page)
	case "n":
		searchResult(qType, str, page+1)
	case "0":
		if qType == 3 {
			nowPage = search_num
		} else {
			nowPage = firstPage_num
		}
		return

	default:
		down(result, ctrl)
		searchResult(qType, str, page)
	}
}

func down(result []SearchResultExt, ids string) {
	Clear()
	fmt.Println("【下载】")
	if ids == "-1" {
		// 下载全部
		ids = ""
		for i := 1; i <= len(result); i++ {
			if i != 1 {
				ids += ","
			}
			ids += strconv.Itoa(i)
		}
	}
	idSlice := strings.Split(ids, ",")
	for _, id := range idSlice {
		num, _ := strconv.Atoi(id)
		fmt.Println("-----------------------------------------------------------------------------")
		if num > 0 && num <= len(result) {
			fmt.Println("下载：", num, result[num-1].Name)
			err := DownGTPic("GT/"+result[num-1].Name, result[num-1].Id)
			if err != nil {
				fmt.Println("下载：", num, result[num-1].Name, "失败", err)
			} else {
				fmt.Println("下载：", num, result[num-1].Name, "成功，保存至：", "GT/"+result[num-1].Name)
			}
		}
		fmt.Println("-----------------------------------------------------------------------------")
	}
	fmt.Println("下载完成")
	fmt.Println("输入任何数返回")
	ctrl := ""
	fmt.Scan(&ctrl)
}
