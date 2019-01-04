# wechat-work-message-push-go



### 使用：
1.创建企业号（200人以下不需要认证），获取 `WechatWorkCorpId`

2.创建自建应用，获取 `WechatWorkAgentId` 和 `WechatWorkCorpSecret`

3.到通讯录查看自己的账号，设置 `DefaultReceiverUserId`

3.复制 docker-compose.yaml 到本地，补充以上环境变量和随机字符串token,然后执行
```$xslt
docker-compose up -d
```

### 测试
```shell
curl -X POST \
  http://127.0.0.1/push/ \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'token: your_token' \
  -F message=1234 \
  -F to_user=XiaoMing
```


### 开发

## 使用 drone ^1.0 构建镜像

https://drone.io/