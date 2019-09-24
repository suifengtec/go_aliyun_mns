# go_aliyun_mns
go_aliyun_mns


## 阿里云 MNS 的 Golang 封装

```go

package main

import (
	"encoding/xml"
	"fmt"
	"github.com/suifengtec/go_aliyun_mns/mns"
)

var (
	AliYunAccessKeyId     = "LT***bw"
	AliYunAccessKeySecret = "gF3***z5tAe"
	AliYunMnsEndpoint     = "1***37.mns.cn-shanghai.aliyuncs.com"
	AliYunMnsQueueName    = "n***ForNewOrder"
)

func send() {

	client := mns.NewClient(AliYunAccessKeyId, AliYunAccessKeySecret, AliYunMnsEndpoint)

	queue := mns.Queue{
		Client:    client,
		QueueName: AliYunMnsQueueName,
		Base64:    false,
	}

	msg := mns.Message{MessageBody: "世界,你好"}
	data, err := xml.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}

	msgId, err2 := queue.Send(mns.GetCurrentUnixMicro(), data)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("send message=>%v", msgId)
		fmt.Println()
	}

}

func receive() {

	client := mns.NewClient(AliYunAccessKeyId, AliYunAccessKeySecret, AliYunMnsEndpoint)
	queue := mns.Queue{
		Client:    client,
		QueueName: AliYunMnsQueueName,
		Base64:    false,
	}
	respChan := make(chan mns.MsgReceive)
	errChan := make(chan error)
	end := make(chan int)
	receiptHandle := ""

	go func() {
		select {
		case resp := <-respChan:
			{
				fmt.Printf("receive message=>%v", resp)
				receiptHandle = resp.ReceiptHandle
				end <- 1
			}
		case err := <-errChan:
			{
				fmt.Println(err)
				end <- 0
			}
		}
	}()

	queue.Receive(respChan, errChan)
	received := <-end

	if received == 1 {
		msgDelete(receiptHandle)
	}

}

func msgDelete(receiptHandle string) {

	client := mns.NewClient(AliYunAccessKeyId, AliYunAccessKeySecret, AliYunMnsEndpoint)
	queue := mns.Queue{
		Client:    client,
		QueueName: AliYunMnsQueueName,
		Base64:    false,
	}
	errChan := make(chan error)
	end := make(chan int)
	go func() {
		select {
		case err := <-errChan:
			{
				if err != nil {

					fmt.Println(err)

					end <- 0
				} else {

					fmt.Println()

					fmt.Println("deletesuccess=>" + receiptHandle)
					end <- 1
				}
			}
		}
	}()

	queue.Delete(receiptHandle, errChan)
	<-end
}

func main() {

	fmt.Println("Aliyun MNS")

	//send()
	receive()

}


```
