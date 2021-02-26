package main

import "fmt"

type NotifyPoint int // 通知时间节点

const (
	NPEmpty                       NotifyPoint = 0 // 空
	NPReferrerInstanceSuccOrFail  NotifyPoint = 2 // 发起人-团成功/失败
	NPReferrerBeforeActivityBegin NotifyPoint = 3 // 发起人-活动开始前提醒
	NPReferrerCountDown           NotifyPoint = 4 // 发起人-团倒计时提醒
	NPReferrerTomorrowRemind      NotifyPoint = 5 // 发起人-订阅明日提醒
	NPApplicantInstanceSucc       NotifyPoint = 6 // 参与人-团成功
	NPReferrerBeforeUseCoupon     NotifyPoint = 7 // 参与-优惠券可使用提醒
	NPReferrerBeforeExchange      NotifyPoint = 8 // 参与人商品兑换开始时间
)

func main() {
	fmt.Println(NPApplicantInstanceSucc == NotifyPoint(6))
}
