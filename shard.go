package main

import (
	"fmt"
	"strconv"

	"hash/crc32"
)

func main() {

	fmt.Println(getShard(283303518541141, true))

	params := map[string]interface{}{}
	//params["is_order"]z = true

	payRecChannel := false
	payRec, ok := params["is_order"]
	if ok {
		payRecChannel, _ = payRec.(bool)
	}

	fmt.Println(payRecChannel)
	var err error
	//err = errors.New("系统错误")
	fmt.Println(fmt.Sprintf("%v", err.Error()))
}

func getShard(id int64, isUid bool) uint32 {
	var shard uint32
	tableShard := int64(1024)
	if isUid {
		shard = crc32.ChecksumIEEE([]byte(strconv.FormatInt(id, 10))) % uint32(tableShard)
	} else {
		shard = uint32(id % tableShard)
	}
	return shard
}
