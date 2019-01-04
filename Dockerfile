FROM alpine:latest
ADD https://github.com/cloverzrg/file/raw/master/ca-certificates.crt /etc/ssl/certs/
ADD wechat-work-message-push-go /app/wechat-work-message-push-go
CMD ["/app/wechat-work-message-push-go","-c","/app/config/config.json"]