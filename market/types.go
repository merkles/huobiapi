package market

import (
	"github.com/bitly/go-simplejson"
	"time"
)

type pongData struct {
	Pong int64 `json:"pong"`
}

type pingData struct {
	Ping int64 `json:"ping"`
}

type subData struct {
	Sub string `json:"sub"`
	ID  string `json:"id"`
}

type unSubData struct {
	Unsub string `json:"unsub"`
	ID    string `json:"id"`
}

type reqData struct {
	Req string `json:"req"`
	ID  string `json:"id"`
}

type jsonChan = chan *simplejson.Json

//资产,订单websocket原型结构

type authPingPongData struct {
	OP string `json:"op"`
	TS int64  `json:"ts"`
}

type authReqData struct {
	OP    string `json:"op"`
	Topic string `json:"topic"`
	CID   string `json:"cid"`
}

type authReplyData struct {
	OP        string `json:"op"`
	Topic     string `json:"topic"`
	CID       string `json:"cid"`
	ErrorCode int    `json:"error_code"`
	TS        int64  `json:"ts"`
}

//{
//"op": "isMustAuth",
//"AccessKeyId": "e2xxxxxx-99xxxxxx-84xxxxxx-7xxxx",
//"SignatureMethod": "HmacSHA256",
//"SignatureVersion": "2",
//"Timestamp": "2017-05-11T15:19:30",
//"Signature": "4F65x5A2bLyMWVQj3Aqp+B4w+ivaA7n5Oi2SuYtCJ9o=",
//}

type authSub struct {
	OP               string `json:"op"`
	CID              string `json:"cid"`
	AccessKeyId      string `json:"AccessKeyId"`
	SignatureMethod  string `json:"SignatureMethod"`
	SignatureVersion string `json:"SignatureVersion"`
	Timestamp        string `json:"Timestamp"`
	Signature        string `json:"Signature"`
}

func NowSec() int {
	return int(time.Now().UnixNano() / 1000000000)
}
