package pan

import (
	"strconv"
	"time"
)

type AuthRequest struct {
	PwdID    string `json:"pwd_id"`
	Passcode string `json:"passcode"`
}

// 假设时间为Unix时间戳（秒）
type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {
	ts, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*t = Time(time.Unix(ts, 0))
	return nil
}

type Author struct {
	MemberType string `json:"member_type"`
	AvatarURL  string `json:"avatar_url"`
	NickName   string `json:"nick_name"`
}

type Data struct {
	Subscribed  bool   `json:"subscribed"`
	Stoken      string `json:"stoken"`
	ShareType   int    `json:"share_type"`
	Author      Author `json:"author"`
	ExpiredType int    `json:"expired_type"`
	ExpiredAt   Time   `json:"expired_at"`
	Title       string `json:"title"`
	FileNum     int    `json:"file_num"`
}

type Metadata struct {
	TGroup string `json:"_t_group"`
	GGroup string `json:"_g_group"`
}

type Response struct {
	Status    int      `json:"status"`
	Code      int      `json:"code"`
	Message   string   `json:"message"`
	Timestamp Time     `json:"timestamp"`
	Data      Data     `json:"data"`
	Metadata  Metadata `json:"metadata"`
}
