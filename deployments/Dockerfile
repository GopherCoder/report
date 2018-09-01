# 官方Golang Image

FROM golang:1.9.4

LABEL maintainer="xiewei(1156143589@qq.com)"

WORKDIR /go/src/report


# 对外开放接口8080，纯文字提示这里并不真正开放接口！
EXPOSE 8080
EXPOSE 8082

# 定义工作环境位置

# 时区设定
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone
RUN dpkg-reconfigure -f noninteractive tzdata



# 运行环境初始化
RUN go get -u github.com/kardianos/govendor

COPY vendor/vendor.json vendor/

COPY Makefile .

RUN make install

COPY . .


# 当container运行时，默认执行编译后文件
CMD [nginx]
