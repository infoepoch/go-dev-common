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
// | @Time  : 13:57
// +----------------------------------------------------------------------
package uuid

import "errors"

// GetUUID 获取 uuid
func GetUUID() (string, error) {
	u, e := uuid.NewV4()
	if e != nil {
		return "", errors.New("new uuid v4 error")
	}
	return u.String(), nil
}
