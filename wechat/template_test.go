package wechat

import (
	"testing"
)

func TestM(t *testing.T) {
	token := "30_EN51TrSEfCrboPg26FmnjkunBPrFxkEGUyiR59vQhY0WkY064uoU0Lv8FdpCsp4uawWszWYI3HieT10BMWJqr3AOrih_xYjOhDe6U2x7hjQTD1Of_-0zvgxL80wuMxxD2n6k88geRGyjKXwbEWXfAFAUMN"
	user := "oFqEd6HUL3u0grQ6Zsa27SJ2w558"
	tid := "vy-_mZpvXezlnGeT3Qipyz_zNdy56bdi0smvrUyvfJ4"
	url := UrlTemplateMessageSend
	result := &TemplateSendResp{}
	data := struct {
		Name TemplateDataItem `json:"name"`
		Card TemplateDataItem `json:"card"`
	}{
		Name: TemplateDataItem{Value: "平安度疫", Color: "#ff0000"},
		Card: TemplateDataItem{Value: "年卡", Color: "#ff0000"},
	}
	template := NewTemplate().
		SetToUser(user).
		SetTemplateId(tid).
		SetData(data)
	err := PushReceptor(PushTypeTemplate, token).
		SetTemplate(template).
		Push(url, result)
	t.Log(err)
	t.Log(result)
}
