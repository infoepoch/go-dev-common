package regexp

import "regexp"

// IsOpenid 匹配ASCII值从0-127的字符，如数字，英文字符，半角符号，以及某些控制字符。即非中文28位
func IsOpenid(openid string) bool {
	var reg = regexp.MustCompile(`^[\x00-\x7F]{28}$`)
	return reg.MatchString(openid)
}

// IsMobile 校验入参是否是手机号码
func IsMobile(mobile string) bool {
	var reg = regexp.MustCompile(`^[1][123456789]\d{9}$`)
	return reg.MatchString(mobile)
}

// IsGH 校验入参是否是固话
func IsGH(deviceno string) bool {
	var reg = regexp.MustCompile(`^(2|3|5|6)[0-9]{7}$`)
	return reg.MatchString(deviceno)
}

// IsKD 校验入参是否是宽带
func IsKD(deviceno string) bool {
	var reg = regexp.MustCompile(`^(AD|KD|LN|FH)[0-9]{10}$`)
	return reg.MatchString(deviceno)
}

// IsZHJT 校验入参是否是智慧家庭
func IsZHJT(deviceno string) bool {
	var reg = regexp.MustCompile(`^(ZH)[0-9]{8}$`)
	return reg.MatchString(deviceno)
}
