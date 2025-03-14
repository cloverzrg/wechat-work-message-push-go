# wechat-work-message-push-go
[![Build Status](http://drone.jeongen.com/api/badges/cloverzrg/wechat-work-message-push-go/status.svg)](http://drone.jeongen.com/cloverzrg/wechat-work-message-push-go)
[![](https://img.shields.io/microbadger/image-size/cloverzrg/wechat-work-message-push-go.svg)](https://hub.docker.com/r/cloverzrg/wechat-work-message-push-go/)

### 部署
- 镜像(支持 AMD64、ARM)  
  https://hub.docker.com/r/cloverzrg/wechat-work-message-push-go
  
- docker compose  
```
services:
  wechat-work-message-push-go:
    image: cloverzrg/wechat-work-message-push-go
    container_name: wechat-work-message-push-go
    environment:
      TZ: Asia/Shanghai
      Token:
      DefaultReceiverUserId:
      WechatWorkCorpId:
      WechatWorkCorpSecret:
      WechatWorkAgentId:
    ports:
      - 51234:80
    restart: always
    network_mode: bridge
```

### 使用：
1.创建企业号（200人以下不需要认证），获取企业ID `WechatWorkCorpId` [https://work.weixin.qq.com/wework_admin/frame#profile](https://work.weixin.qq.com/wework_admin/frame#profile)

2.创建自建应用，用用管理页面获取 `WechatWorkAgentId` 和 `WechatWorkCorpSecret`

3.到通讯录查看自己的账号，获取 `DefaultReceiverUserId` 

3.复制 docker-compose.yaml 到本地，编辑文件，补充以上环境变量和随机字符串token,然后执行
```shell
docker-compose up -d
```

4.到 企业微信->我的企业->微信插件，扫二维码关注企业微信 https://work.weixin.qq.com/wework_admin/frame#profile/wxPlugin

5.发送以下请求
```shell
curl -X POST \
  http://127.0.0.1:60009/push \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'token: your_token' \
  -d 'message=1234'
```

6.刚才在微信上的关注的微工作台应收到第五步发送的消息
![](https://github.com/cloverzrg/wechat-work-message-push-go/raw/master/image/IMG_8017.jpg)


7.grafana 报警通知功能
设置GrafanaWebhookUser和GrafanaWebhookPassword两个环境变量就可以用了
![](https://github.com/cloverzrg/wechat-work-message-push-go/raw/master/image/grafana_webhook.png)


8. 代替Telegram的通知连接  
   比如nezha面板的通知，可以设置为
   GET https://xx/push/push?token=xxx&message=#NEZHA#

### 开发

#### 使用 drone 构建镜像

https://drone.io/

