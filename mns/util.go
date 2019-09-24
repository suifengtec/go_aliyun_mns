package mns

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"time"
)

//getGMTime ...
func getGMTime() string {
	return time.Now().UTC().Format(http.TimeFormat)
}

//GetCurrentMillisecond 获取当前时间戳(毫秒)
func GetCurrentMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

//GetCurrentUnixMicro 获取当前时间戳(秒)
func GetCurrentUnixMicro() int64 {
	return time.Now().Unix() * 1000
}

//Sha1 对字符串进行sha1 计算
func Sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//Md5 对数据进行md5计算
func Md5(byteMessage []byte) string {
	h := md5.New()
	h.Write(byteMessage)
	return hex.EncodeToString(h.Sum(nil))
}

//HamSha1 ...
func HamSha1(data string, key []byte) string {
	hmac := hmac.New(sha1.New, key)
	hmac.Write([]byte(data))

	return base64.StdEncoding.EncodeToString(hmac.Sum(nil))
}

/*
func dump(data interface{}) {
	fmt.Println(data)
}
*/
