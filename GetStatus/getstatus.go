package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {
	//t := time.Now()
	data := make(url.Values)
	data["userid"] = []string{"45"}
	username := "13512121211"
	pwd := "123456L"
	timestamp := time.Now().Format("20060102150405")
	data["timestamp"] = []string{timestamp}
	signstring := username + pwd + timestamp
	sign := MD5(signstring)
	data["sign"] = []string{sign}
	data["action"] = []string{"query"}
	res, err := http.PostForm("http://119.23.247.150:8888/v2statusApi.aspx", data)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println("post send success")
	//fmt.Printf("type is  %T\n", body)
	fmt.Println(string(body))
}

func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}
