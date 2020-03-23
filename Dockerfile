FROM alpine:latest
ADD https://github.com/cloverzrg/file/raw/master/ca-certificates.crt /etc/ssl/certs/
ADD wechat-work-message-push-go /app/wechat-work-message-push-go
EXPOSE 80
ENV Token my_token
ENV DefaultReceiverUserId 13800138000
ENV WechatWorkCorpId ww741038v8sa88hv36d
ENV WechatWorkCorpSecret USVdvsa_ad2k34jk232kjn-asfefeawf_waeasdf-ase
ENV WechatWorkAgentId 1000001
ENV GrafanaWebhookUser admin
ENV GrafanaWebhookPassword admin
CMD ["/app/wechat-work-message-push-go"]