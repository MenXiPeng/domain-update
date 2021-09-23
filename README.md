# domain-update
动态更新阿里域名 ip 解析   
提供邮件的方式通知 ip 变更

# 打包方式
set GOARCH=amd64
set GOOS=linux
go build main.go

# 首次执行
go install
go run main.go -k '' -s '' -d '' -r ''

--- 
    k : 阿里云 AccessKey
    s : 阿里云 AccessKey Secret
    d : 根域名
    r : 二级域名
---

## 阿里云申请地址

https://usercenter.console.aliyun.com/#/manage/ak
