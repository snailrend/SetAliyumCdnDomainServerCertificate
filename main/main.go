// This file is auto-generated, don't edit it. Thanks.
package main

import (
	"fmt"
	cdn20180510 "github.com/alibabacloud-go/cdn-20180510/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/spf13/viper"
	"io"
	"os"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *cdn20180510.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("cdn.aliyuncs.com")
	_result = &cdn20180510.Client{}
	_result, _err = cdn20180510.NewClient(config)
	return _result, _err
}

/**
* 使用STS鉴权方式初始化账号Client，推荐此方式。本示例默认使用AK&SK方式。
* @param accessKeyId
* @param accessKeySecret
* @param securityToken
* @return Client
* @throws Exception
 */
func CreateClientWithSTS(accessKeyId *string, accessKeySecret *string, securityToken *string) (_result *cdn20180510.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
		// 必填，您的 Security Token
		SecurityToken: securityToken,
		// 必填，表明使用 STS 方式
		Type: tea.String("sts"),
	}
	// 访问的域名
	config.Endpoint = tea.String("cdn.aliyuncs.com")
	_result = &cdn20180510.Client{}
	_result, _err = cdn20180510.NewClient(config)
	return _result, _err
}

func _main(args []*string) (_err error) {
	// 工程代码泄露可能会导致AccessKey泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	client, _err := CreateClient(tea.String(viper.GetString("accessKeyId")), tea.String(viper.GetString("accessKeySecret")))
	privateKeyPath := viper.GetString("privateKeyPath")
	publicKeyPath := viper.GetString("publicKeyPath")
	publicKey, _err := readFileContent(publicKeyPath)
	privateKey, _err := readFileContent(privateKeyPath)
	if _err != nil {
		return _err
	}
	setDomainServerCertificateRequest := &cdn20180510.SetDomainServerCertificateRequest{
		DomainName:              tea.String(viper.GetString("domainName")),
		CertName:                tea.String(viper.GetString("certName")),
		CertType:                tea.String("upload"),
		ServerCertificateStatus: tea.String("on"),
		ForceSet:                tea.String("1"),
		PrivateKey:              tea.String(privateKey),
		ServerCertificate:       tea.String(publicKey),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := client.SetDomainServerCertificateWithOptions(setDomainServerCertificateRequest, runtime)
	if _err != nil {
		return _err
	}

	console.Log(util.ToJSONString(resp))
	return _err
}

func readFileContent(path string) (string, error) {
	file, _err := os.Open(path)
	if _err != nil {
		return "", _err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
	content, _err := io.ReadAll(file)
	if _err != nil {
		return "", _err
	}
	return string(content), nil
}

func main() {
	initConfig()
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		// Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
