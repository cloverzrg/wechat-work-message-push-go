FROM alpine:latest
ADD ca-certificates.crt /etc/ssl/certs/
ADD wechat-work-message-push-go /app/wechat-work-message-push-go
CMD ["/app/wechat-work-message-push-go","-c","/app/config/config.json"]