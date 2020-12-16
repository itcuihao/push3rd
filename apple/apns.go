package apple

import (
	"crypto/tls"
	"errors"
	"fmt"
	"time"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
)

var (
	cert tls.Certificate
	err  error
	defaultPayload = []byte(`{"aps":{"alert":"Bless you.","badge":1,"sound":"bookbookqs_0417.wav"}}`) // See Payload section below
)

func InitApns(path string) error {
	cert, err = certificate.FromP12File(path, "")
	if err != nil {
		return fmt.Errorf("path: %v, err: %v", path, err)
	}
	return nil
}

type Apns struct {
	Topic     string
	IsSandbox bool
	client    *apns2.Client // 发送client
}

func NewApns(topic string, isSandbox bool) *Apns {
	apns := &Apns{Topic: topic, IsSandbox: isSandbox}
	apns.NewClient()
	return apns
}

func (a *Apns) NewClient() {
	client := apns2.NewClient(cert)
	if a.IsSandbox {
		client.Development()
	} else {
		client.Production()
	}
	a.client = client
}

func (a Apns) QuickPush(deviceToken string, payload interface{}) (*apns2.Response, error) {
	notification := &apns2.Notification{}
	notification.DeviceToken = deviceToken
	notification.Topic = a.Topic
	notification.Payload = defaultPayload
	if payload != nil {
		notification.Payload = payload
	}
	res, err := a.client.Push(notification)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a Apns) SendMultiDevice(payload interface{}, deviceTokens ...string) []error {
	notification := &apns2.Notification{}
	notification.Topic = a.Topic
	notification.Payload = []byte(`{"aps":{"alert":"Bless you.","badge":1,"sound":"bookbookqs_0417.wav"}}`) // See Payload section below
	if payload != nil {
		notification.Payload = payload
	}
	var errs []error
	for _, deviceToken := range deviceTokens {
		notification.DeviceToken = deviceToken
		res, err := a.client.Push(notification)
		if err != nil {
			errs = append(errs, fmt.Errorf("device: %s, err: %v", deviceToken, err))
		}
		if res.StatusCode != 200 {
			errs = append(errs, errors.New(fmt.Sprintf("device: %s, err: %+v", deviceToken, res)))
		}
	}
	return errs
}

type (
	Alert struct {
		Title string `json:"title,omitempty"`
		Body  string `json:"body,omitempty"`
	}
	Aps struct {
		Alert Alert  `json:"alert,omitempty"`
		Badge int    `json:"badge,omitempty"`
		Sound string `json:"sound,omitempty"`
	}
	Payload struct {
		Aps  `json:"aps,omitempty"`
		Link string `json:"link,omitempty"`
	}
)

func (a Apns) SendMassDevice(payload interface{}, deviceTokens ...string) []error {
	deviceLen := len(deviceTokens)
	if deviceLen == 0 {
		return nil
	}

	errs := make([]error, 0, deviceLen)
	notifications := make(chan *apns2.Notification, 100)
	responses := make(chan *apns2.Response, deviceLen)
	for i := 0; i < 50; i++ {
		go worker(a.client, notifications, responses)
	}
	for i := 0; i < len(deviceTokens); i++ {
		notification := &apns2.Notification{}
		notification.Topic = a.Topic
		notification.Payload = defaultPayload
		if payload != nil {
			notification.Payload = payload
		}
		notification.DeviceToken = deviceTokens[i]
		notifications <- notification
	}
	for i := 0; i < deviceLen; i++ {
		res := <-responses
		if res.StatusCode != 200 {
			errs = append(errs, fmt.Errorf("apnsId:%v, status: %v, reason: %v", res.ApnsID, res.StatusCode, res.Reason))
		}
	}
	close(notifications)
	close(responses)
	return errs
}

func worker(client *apns2.Client, notifications <-chan *apns2.Notification, responses chan<- *apns2.Response) {
	for n := range notifications {
		res, err := client.Push(n)
		if err != nil {
			timestamp := apns2.Time{Time: time.Now()}
			res = &apns2.Response{
				StatusCode: -1,
				Reason:     err.Error(),
				ApnsID:     n.ApnsID,
				Timestamp:  timestamp,
			}
			responses <- res
		}
		responses <- res
	}
}
