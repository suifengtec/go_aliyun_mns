package mns

//Queue struct
type Queue struct {
	*Client
	QueueName string
	Base64    bool
}

//Message struct
type Message struct {
	MessageBody string `xml:"MessageBody"`
}

//MsgSend struct
type MsgSend struct {
	MessageId      string `xml:"MessageId"`
	MessageBodyMD5 string `xml:"MessageBodyMD5"`
}

//MsgReceive struct
type MsgReceive struct {
	MessageId       string `xml:"MessageId"`
	MessageBodyMD5  string `xml:"MessageBodyMD5"`
	MessageBody     string `xml:"MessageBody"`
	ReceiptHandle   string `xml:"ReceiptHandle"`
	EnqueueTime     int64  `xml:"EnqueueTime"`
	NextVisibleTime int64  `xml:"NextVisibleTime"`
	DequeueCount    int    `xml:"DequeueCount"`
	Priority        int    `xml:"Priority"`
}

//Error struct ...
type Error struct {
	StatusCode int
	Code       string `xml:"Code"`
	Message    string `xml:"Message"`
}

//Error() ...
func (err *Error) Error() string {
	return fmt.Sprintf("Aliyun MNS API Error: Status Code: %d Code: %s Message: %s", err.StatusCode, err.Code, err.Message)
}
