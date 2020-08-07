FROM harbor.nj.com/ops/centos:official

WORKDIR /opt/app
COPY logtest .
RUN set -ex \
    && chmod +x logtest \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

ENTRYPOINT ["/opt/app/logtest", "-f", "/opt/app/template.log"]
