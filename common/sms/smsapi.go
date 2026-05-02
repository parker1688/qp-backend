package sms

import (
	"bootpkg/common/global"
	"encoding/json"
	"fmt"
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v5/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

func CreateClient() (_result *dysmsapi20170525.Client, _err error) {
	//conf := global.CONFIG.Sms
	config := &openapi.Config{
		// 您的AccessKey ID

		AccessKeyId: tea.String(global.CONFIG.Sms.AccessKeyId),
		//AccessKeyId: tea.String(conf.ALiSMSAppId),
		// 您的AccessKey Secret
		AccessKeySecret: tea.String(global.CONFIG.Sms.AccessKeySecret),
		//AccessKeySecret: tea.String(conf.AliSMSAppKey),
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String(global.CONFIG.Sms.Endpoint)
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func Handle(to string, code string) (err error) {
	client, _err := CreateClient()
	if _err != nil {
		return _err
	}

	tmcode := "{code:" + code + "}"
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(to),
		SignName:      tea.String(global.CONFIG.Sms.SignName),
		TemplateCode:  tea.String(global.CONFIG.Sms.TemplateCode),
		TemplateParam: tea.String(tmcode),
	}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.SendSmsWithOptions(sendSmsRequest, &util.RuntimeOptions{})
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		error := &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		// 错误 message
		fmt.Println(tea.StringValue(error.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		err1 := d.Decode(&data)
		if err1 != nil {
			return err1
		}
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}
