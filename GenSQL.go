package main

import "fmt"

func main() {

	sql := "select * from recommend_activity_reward_%d\t \nwhere scene = 103 and create_time >= \"2020-12-24 00:00:00\" \nand activity_id = 488\nand extra -> '$.city_id' in (116, 318, 321)\nand action_type = 'send_wj_zero_goods_item' union"
	for i := 700; i <= 999; i++ {
		fmt.Println(fmt.Sprintf(sql, i))
	}
}
