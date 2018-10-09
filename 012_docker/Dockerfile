FROM alpine:3.8

RUN apk update && \
    apk upgrade && \
    apk add --no-cache curl && \
    apk add --no-cache musl-dev && \
    apk add --no-cache go

RUN mkdir /root/code
ADD src /root/code/

RUN cd /root/code && go build -o app .

ENTRYPOINT [ "/root/code/app" ]
