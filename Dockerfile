FROM --platform=$BUILDPLATFORM golang:alpine as builder
ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH
ARG DRONE_TAG
ENV CGO_ENABLED=0
ENV GOOS=$TARGETOS
ENV GOARCH=$TARGETARCH
RUN echo "I am running on $BUILDPLATFORM, building for $TARGETPLATFORM, GOOS $GOOS, GOARCH $GOARCH"
ADD . /go/src/wechat-work-message-push-go
WORKDIR /go/src/wechat-work-message-push-go
RUN pwd && ls -lah
RUN apk update && apk add --no-cache git build-base make
RUN buildflags="-X 'main.BuildTime=`date`' -X 'main.GitHead=`git rev-parse --short HEAD`' -X 'main.GoVersion=$(go version)'" && go build -trimpath -ldflags "-s -w $buildflags" -o wechat-work-message-push-go

FROM --platform=$TARGETPLATFORM alpine:latest
RUN apk update && apk add --no-cache ca-certificates tzdata
COPY --from=builder /go/src/wechat-work-message-push-go/wechat-work-message-push-go /app/wechat-work-message-push-go
ENV TZ=Asia/Shanghai
ENV Token my_token
ENV DefaultReceiverUserId 13800138000
ENV WechatWorkCorpId ww741038v8sa88hv36d
ENV WechatWorkCorpSecret USVdvsa_ad2k34jk232kjn-asfefeawf_waeasdf-ase
ENV WechatWorkAgentId 1000001
ENV GrafanaWebhookUser admin
ENV GrafanaWebhookPassword admin
EXPOSE 80
WORKDIR /app
CMD ["/app/wechat-work-message-push-go"]