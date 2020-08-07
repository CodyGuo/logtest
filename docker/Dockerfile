FROM alpine

WORKDIR /opt/app
COPY logtest .
RUN set -ex \
    && chmod +x logtest \
    && apk add --no-cache tzdata\
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

ENTRYPOINT ["/opt/app/logtest", "-f", "/opt/app/template.log"]
