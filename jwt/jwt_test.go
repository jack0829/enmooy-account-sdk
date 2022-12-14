package jwt

import "testing"

func TestJWT(t *testing.T) {

	key := []byte("your key")
	salt := "your salt"

	var s string

	// 生成 JWT
	{
		uid := 123
		eid := 456
		svc := []string{"oss", "slb"}
		nick := "abc"
		avatar := "xx.jpg"

		data := NewDataWithEncoder(NewEncoder(salt), uid, eid, svc...).WithNick(nick).WithAvatar(avatar)
		j1 := New(key, salt).WithData(data)

		s = j1.Encode()
		t.Log(s)
	}

	// 解析 JWT
	{
		j2 := New(key, salt)
		if err := j2.Decode(s); err != nil {
			t.Error(err)
		}

		data := j2.Data()
		t.Logf("%#+v", data)
		if uid, err := data.GetUid(j2.GetEncoder()); err == nil {
			t.Logf("UID: %d", uid)
		}
		if eid, err := data.GetEid(j2.GetEncoder()); err == nil {
			t.Logf("EID: %d", eid)
		}
	}
}
