# Binary file names.
BINARY_NAME=glca

# Build parameters.
DIST_PATH=.dist

### help:					Makefile 帮助
.PHONY: help
help:
	@echo Makefile cmd:
	@echo
	@grep -E '^### [-A-Za-z0-9_]+:' Makefile | sed 's/###/   /'

.PHONY: build
### build:				项目打包
build: build-go

### build-go:			构建 golang 包
build-go:
	rm -rf ${DIST_PATH}
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -trimpath -tags=jsoniter -mod=readonly -v -o $(DIST_PATH)/${BINARY_NAME} main.go
	cp -rf conf $(DIST_PATH)
	cp -rf static $(DIST_PATH)


.PHONY: gen
### init:				初始化数据
gen:
	@gen -f .gen/gen.yml
