// +----------------------------------------------------------------------
// | 
// | Created by InteliJ IDE
// +----------------------------------------------------------------------
// | Mohoo Go [ WE CAN DO IT JUST THINK IT ]
// +----------------------------------------------------------------------
// | Copyright (c) 2017 http://www.mohoo.com.cn All rights reserved.
// +----------------------------------------------------------------------
// | 每位工程师都有保持代码优雅的义务
// | Each engineer has a duty to keep the code elegant
// +----------------------------------------------------------------------
// | @Author: Learn <yanzx@infoepoch.com>
// | @Date  : 2018/04/2018/4/17
// | @Time  : 14:00
// +----------------------------------------------------------------------
package uuid

import (
	"testing"
	"fmt"
)

func TestGetUuid(t *testing.T) {
	s, e := GetUUID()
	if e != nil {
		t.Error(e.Error())
	}
	fmt.Println("uuid: ", s)
}
