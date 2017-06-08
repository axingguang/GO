package dataSyn

import (
	"strconv"
	"strings"

	"net/url"

	"../../common/httpclient"
	logu "../../common/log/util"
	cfg "../../common/property/util"
	zkutil "../../zkManager/util"
)

var c = New()
var client = httpclient.GetHttpclient()

func New() chan *Data {

	channelSizeStr := cfg.GetValue("channel", "channel.pool")
	channelSize := 100
	if !strings.EqualFold(channelSizeStr, "") {
		a, er := strconv.Atoi(channelSizeStr)
		if er == nil {
			channelSize = a
		}
	}

	ch := make(chan *Data, channelSize)
	//ch <- 2
	go func() {
		for a := range ch {
			sendData(a)
		}
	}()
	return ch
}

func SetData(data *Data) {
	c <- data
}

func sendData(data *Data) {
	strArr := zkutil.GetNoteExLocal()
	if strArr != nil && len(strArr) > 0 {
		for _, v := range strArr {
			go func(host string, datas *Data) {
				logu.Info("【数据同步】，host:", host, "dataType:", data.DataType, "key:", data.Key, "value:", data.Val, "ex：", data.Ex)
				if datas.DataType == SetType {
					_, err := client.PostForm("http://"+host+"/cachen/setValue", url.Values{"key": {datas.Key}, "value": {datas.Val}, "ex": {strconv.Itoa(datas.Ex)}})
					if err != nil {
						logu.Error("【同步】设置数据失败,host:", host, "err:", err)
					}
				}
				if datas.DataType == DelType {
					_, err := client.PostForm("http://"+host+"/cachen/delValue", url.Values{"key": {datas.Key}})
					if err != nil {
						logu.Error("【同步】删除数据失败,host:", host, "err:", err)
					}
				}
			}(v, data)
		}
	}
}
