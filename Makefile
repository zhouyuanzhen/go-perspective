all:
	@echo TBD

build:
	@go build .

clean:
	@rm -f go-perspective

latest-inst:
	@go install github.com/zhouyuanzhen/go-perspective@latest
