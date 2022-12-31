# 设置域名证书代码调用

该项目为调用SetDomainServerCertificate设置指定域名下证书信息。

## 开发环境运行条件

- 在阿里云帐户中获取您的 [凭证](https://usercenter.console.aliyun.com/#/manage/ak)并通过它替换`config/config.yaml`中的 ACCESS_KEY_ID 以及 ACCESS_KEY_SECRET;

- Golang 版本要求 1.13 及以上

## 调试执行
1. 将`config/config.yaml`中的参数替换成实际需要的参数
2. 执行代码:
```sh
export GOPROXY="https://mirrors.aliyun.com/goproxy/,direct" 
go run ./main
```

## 打包可执行文件并运行
```shell
chmod a+x ./scripts/build.sh
chmod a+x ./build/setDomainServerCertificate
./scripts/build.sh
./build/setDomainServerCertificate
```
在生产环境中使用时，可以设置定时任务执行
## 使用的 API

-  SetDomainServerCertificate 调用SetDomainServerCertificate设置指定域名下证书信息及证书功能是否启用。文档示例，可以参考：[文档](https://next.api.aliyun.com/document/Cdn/2018-05-10/SetDomainServerCertificate)



## 返回示例

*实际输出结构可能稍有不同，属于正常返回；下列输出值仅作为参考，以实际调用为准*


- JSON 格式 
```js
{ 
    "RequestId": "0AEDAF20-4DDF-4165-8750-47FF9C1929C9" 
}
```
- XML 格式 
```xml
<SetDomainServerCertificateResponse>
	<RequestId>0AEDAF20-4DDF-4165-8750-47FF9C1929C9</RequestId>
</SetDomainServerCertificateResponse>
```


