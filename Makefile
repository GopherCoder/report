BINARY=report

VERSION=1.0.0

BUILD=`date +%FT%T%z`

# Setup the -Idflags options for go build here,interpolate the variable values

LDFLAGS=-ldflags


default:
	go build -o ${BINARY} -tags=jsoniter


install:
	govendor sync -v

dev:
	fresh -c configs/development.conf

prod:
	go build -o ${BINARY} -v ${LDFLAGS} -tags=jsoniter

.PHONY:  default install dev prod