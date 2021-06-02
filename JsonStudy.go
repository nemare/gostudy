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

func (a Ext) String() string {

	return "hello " + a.DCHN
}

type SepConf struct {
	Seps     []string    `json:"seps"`
	TextConf []*TextConf `json:"text_conf"`
}

type TextConf struct {
	ButtonText    string `json:"button_text"`
	ShareUserText string `json:"share_user_text"`
}

func main() {

	s, _ := json.Marshal(Ext{
		DCHN:            "2",
		EntranceChannel: "1",
		XID:             "3",
		ActivityID:      "4",
		InstanceID:      "5",
	})
	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", Ext{
		DCHN:            "2",
		EntranceChannel: "1",
		XID:             "3",
		ActivityID:      "4",
		InstanceID:      "5",
	})
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

	err1 := json.Unmarshal([]byte(""), &titles)
	if err1 != nil {
		log.Fatal(err1)
		return
	}

	fmt.Println(titles)
	fmt.Println(reflect.TypeOf(titles))

}
