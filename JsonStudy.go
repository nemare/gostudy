package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type Movie struct {
	Title string
	Year  int  `json:"released"`
	Color bool `json:"color,omitempty"`
	Actor []string
}

type Ext struct {
	DCHN            string `json:"dchn"`
	EntranceChannel string `json:"entrance_channel"`
	XID             string `json:"xid"`
	ActivityID      string `json:"activity_id"`
	InstanceID      string `json:"instance_id"`
}

type SepConf struct {
	Seps     []string `json:"seps"`
	TextConf TextConf `json:"text_conf"`
}

type TextConf struct {
	ButtonText    string `json:"button_text"`
	ShareUserText string `json:"share_user_text"`
}

func main() {

	//movies := []Movie{
	//	{"你是", 2000, false, []string{"李连杰", "成龙"}},
	//	{Title: "我是", Year: 2010, Actor: []string{"李连杰", "成龙"}},
	//	{"他是是", 2020, true, []string{"李连杰", "成龙"}},
	//}

	//ext := Ext{
	//	DCHN: "1",
	//	EntranceChannel: "webx_entrance1",
	//	XID: "123",
	//	ActivityID: "299069675958038",
	//	InstanceID: "131225832",
	//}
	ext := SepConf{
		Seps: []string{"￥", "₠", "₡", "₢", "₣", "₤",
			"₥", "₦", "₧", "₨", "₩", "₪", "₫", "€", "₭", "₮", "₯", "₰", "₱",
			"₲", "₳", "₴", "₵", "₶", "₷", "₸", "₹"},
		TextConf: TextConf{
			ButtonText:    "打开",
			ShareUserText: "给你分享了",
		},
	}

	json_str, err := json.Marshal(ext)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s\n", json_str)

	fmt.Println("-------------")

	var titles []struct {
		Title string
	}

	err1 := json.Unmarshal(json_str, &titles)
	if err1 != nil {
		log.Fatal(err1)
		return
	}

	fmt.Println(titles)
	fmt.Println(reflect.TypeOf(titles))

}
