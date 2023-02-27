package validates

import "testing"

func TestValidates(t *testing.T) {

	err := IsHTML()("<div>ddd</div>")
	if err != nil {
		t.Error(err)
	}

	err = IsURLEncoded()("https://www.baidu.com")
	if err != nil {
		t.Error(err)
	}

	err = IsMac()("f8:b1:56:cb:04:9a")
	if err != nil {
		t.Error(err)
	}

	err = IsLongitude()("-176.332979")
	//err := IsLatitude()("-96.332979")
	if err != nil {
		t.Error(err)
	}

	err = IsIP()("192.168.0.133")
	if err != nil {
		t.Error(err)
	}

	err = IsURL()("http://www.baidu.com")
	if err != nil {
		t.Error(err)
	}

	err = IsE164()("+8618359260323")
	if err != nil {
		t.Error(err)
	}

	err = IsEmail()("33-33@33ddddcom.c")
	if err != nil {
		t.Error(err)
	}

	err = IsBoolean()(true)
	if err != nil {
		t.Error(err)
	}

	err = IsMD5()("1aabac6d068eef6a7bad3fdf50a05cc8")
	if err != nil {
		t.Error(err)
	}

	err = IsSHA256()("82a4a995961e0b083dbfcbae6174605de9cef4568606d25b4c661b0e6c82206c")
	if err != nil {
		t.Error(err)
	}
}
