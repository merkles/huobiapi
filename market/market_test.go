package market

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/stretchr/testify/assert"
)

func TestNewMarket(t *testing.T) {
	m, err := NewMarket(Endpoint, false, "", "")
	assert.NoError(t, err)

	// 订阅
	//err = m.Subscribe("market.eosusdt.kline.1min", func(topic string, json *simplejson.Json) {
	//	fmt.Println(topic, json)
	//})
	//assert.NoError(t, err)
	err = m.Subscribe("market.eosusdt.trade.detail", func(topic string, json *simplejson.Json) {
		fmt.Println(time.Now(), json)
		//fmt.Println(topic, json)
	})
	assert.NoError(t, err)

	//// 请求
	//rep, err := m.Request("market.eosusdt.detail")
	//assert.NoError(t, err)
	//fmt.Println(rep)

	// 阻塞事件循环
	fmt.Println(m)
	go func() {
		time.Sleep(time.Second * 10)
		m.Close()
	}()
	m.Loop()

	fmt.Println(strings.Repeat("-------------------\n", 10))

	// 重新连接
	m.ReConnect()
	go func() {
		time.Sleep(time.Second * 12)
		m.Close()
	}()
	go func() {
		for {
			time.Sleep(time.Second * 2)
			rep, err := m.Request("market.eosusdt.detail")
			fmt.Println(err, rep)
		}
	}()
	m.Loop()

	fmt.Println(m)
}

func TestMarketAlive(t *testing.T) {
	m, err := NewMarket(Endpoint, false, "", "")
	assert.NoError(t, err)
	err = m.Subscribe("market.eosusdt.kline.1min", func(topic string, json *simplejson.Json) {
		fmt.Println(topic, json)
	})
	err = m.Subscribe("market.eosusdt.kline.1min", func(topic string, json *simplejson.Json) {
		fmt.Println(topic, json)
	})
	assert.NoError(t, err)
	go func() {
		time.Sleep(time.Minute * 2)
		m.Close()
	}()
	m.Loop()
}

//var key_secret_huobi = "31609e5e-cfcf0540-f7d6f1e0-c2947"
//var key_access_huobi = "f6248841-df15ce66-2a354233-1qdmpe4rty"

var key_secret_huobi = "7af726f1-59da0020-1d1c936f-c285e"
var key_access_huobi = "33dd3194-mn8ikls4qg-41282b7f-0b64f"

func TestMarketAuth(t *testing.T) {
	m, err := NewMarket(AssetEndPoint, true, key_access_huobi, key_secret_huobi)
	assert.NoError(t, err)
	//m.sendAuth()
	//订阅账户
	err = m.Subscribe("accounts", func(topic string, json *simplejson.Json) {
		fmt.Println("callback:", topic, json)
	})

	//go func() {
	//	err = m.Subscribe("accounts", func(topic string, json *simplejson.Json) {
	//		fmt.Println(topic, json)
	//	})
	//
	//}()

	err = m.Subscribe("orders.bttusdt.update", func(topic string, json *simplejson.Json) {
		fmt.Println("callback:", topic, json)
	})

	//go func() {
	//	err = m.Subscribe("orders.bttusdt.update", func(topic string, json *simplejson.Json) {
	//		fmt.Println(topic, json)
	//	})
	//
	//}()
	//
	//go func() {
	//	err = m.Subscribe("orders.bttusdt", func(topic string, json *simplejson.Json) {
	//		fmt.Println(topic, json)
	//	})
	//
	//}()

	//err = m.Subscribe("market.eosusdt.kline.1min", func(topic string, json *simplejson.Json) {
	//	fmt.Println(topic, json)
	//})
	//go func() {
	//	time.Sleep(time.Minute * 2)
	//	m.Close()
	//}()
	m.Loop()
}
