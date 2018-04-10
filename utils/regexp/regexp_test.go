package regexp

import "testing"

func TestIsOpenid(t *testing.T) {
	if IsOpenid("oKXUCjyIgojbWpg1osglQSVAnlC8") {
		t.Log("oKXUCjyIgojbWpg1osglQSVAnlC8 is the openid")
	}
	if !IsOpenid("oKXUCjyIgojbWpg1oSVAnlC8") {
		t.Log("oKXUCjyIgojbWpg1oSVAnlC8 is not openid")
	}
}

func TestIsMobile(t *testing.T) {
	if IsMobile("11111111111") {
		t.Log("11111111111 is the mobile")
	}
	if !IsMobile("11111111111") {
		t.Log("11111111111 is not mobile")
	}
}

func TestIsGH(t *testing.T) {
	if IsGH("20000000") {
		t.Log("20000000 is the GH")
	}
	if !IsGH("2000000") {
		t.Log("20000000 is not GH")
	}
}

func TestIsKD(t *testing.T) {
	if IsKD("AD1111111111") {
		t.Log("AD1111111111 is the KD")
	}
	if !IsKD("AD1111111111") {
		t.Log("AD1111111111 is not KD")
	}
}

func TestIsZHJT(t *testing.T) {
	if IsZHJT("ZH11111111") {
		t.Log("ZH11111111 is the ZHJT")
	}
	if !IsZHJT("ZH11111111") {
		t.Log("ZH11111111 is not ZHJT")
	}
}
