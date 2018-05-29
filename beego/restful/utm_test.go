package restful

import (
	"net/url"
	"fmt"
	"testing"
)

// 测试URL地址基础构建
func TestUrlValue(t *testing.T) {
	c := url.Values{"method": {"get"}, "id": {"1"}}
	fmt.Println(c.Encode())
	fmt.Println(c.Get("method"))

	c.Set("method", "post")
	fmt.Println(c.Encode())
	fmt.Println(c.Get("method"))

	c.Del("method")
	fmt.Println(c.Encode())
	fmt.Println(c.Get("method"))

	c.Add("new", "hi")
	fmt.Println(c.Encode())
}
