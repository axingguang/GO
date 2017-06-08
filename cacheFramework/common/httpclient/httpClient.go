package httpclient

import (
	//	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	cfg "../property/util"
)

var HttpClient = New()

func New() *http.Client {
	timeoutStr := cfg.GetValue("httpclient", "client.timeout")
	timeOutI := 2
	if !strings.EqualFold(timeoutStr, "") {
		timeOUtAtoi, err := strconv.Atoi(timeoutStr)
		if err == nil {
			timeOutI = timeOUtAtoi
		}
	}
	/*
		client := &http.Client{
			Transport: &http.Transport{
				Dial: func(netw, addr string) (net.Conn, error) {
					conn, err := net.DialTimeout(netw, addr, time.Second*time.Duration(timeOutI))
					if err != nil {
						return nil, err
					}
					//conn.SetDeadline(time.Now().Add(time.Second * time.Duration(timeOutI)))
					return conn, nil
				},
				ResponseHeaderTimeout: time.Second * time.Duration(timeOutI),
			},
		}*/
	client := &http.Client{
		Timeout: time.Second * time.Duration(timeOutI),
	}
	return client
}
func GetHttpclient() *http.Client {
	return HttpClient
}
