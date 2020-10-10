package main

import (
	"fmt"
	//"strings"
)

const PATTERN = "%*=command%*="

func main() {

	//str := "5.0复制内容%*=command%*=到 淘tao寳 或点击链接 https://m.tb.cn/h.VwXRVJA?sm=cf1819 至流览器【爱玩具 正版ACTOYS猫铃铛少女心盲盒 超可爱网红猫咪手办公仔摆件】\""
	//r, _ := regexp.Compile(`￥([^￥]*)￥`)
	//r,_:=regexp.Compile(`(?<='1').*(?='2')`)
	//r, _ := regexp.Compile(`%\*=()%\*=`)

	//result := strings.Replace(str,PATTERN ,"woshi", 1)

	var a string
	a = "18810575856"
	fmt.Println(len(a))
	fmt.Println(getTailNumber(a))

	//result := r.FindString(str)

	//fmt.Println(result)

}

func getTailNumber(phone string) string {
	return phone[len(phone)-4:]
}