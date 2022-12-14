package jwt

import "github.com/speps/go-hashids/v2"

type Data struct {
	Uid    string   `json:"uid"`
	Eid    string   `json:"eid"`
	Nick   string   `json:"nick"`
	Avatar string   `json:"avatar"`
	Svc    []string `json:"svc"`
}

func NewDataWithEncoder(encoder *hashids.HashID, uid, eid int, svc ...string) *Data {

	encodeUid, _ := encoder.Encode([]int{uid})
	encodeEid, _ := encoder.Encode([]int{eid})

	return &Data{
		Uid: encodeUid,
		Eid: encodeEid,
		Svc: svc,
	}
}

func (d *Data) WithNick(n string) *Data {
	d.Nick = n
	return d
}

func (d *Data) WithAvatar(a string) *Data {
	d.Avatar = a
	return d
}

func (d *Data) GetUid(encoder *hashids.HashID) (int, error) {
	id, err := encoder.DecodeWithError(d.Uid)
	if err != nil {
		return 0, err
	}
	if len(id) > 0 {
		return id[0], nil
	}
	return 0, nil
}

func (d *Data) GetEid(encoder *hashids.HashID) (int, error) {
	id, err := encoder.DecodeWithError(d.Eid)
	if err != nil {
		return 0, err
	}
	if len(id) > 0 {
		return id[0], nil
	}
	return 0, nil
}
