configPath = /Users/welong/Project/vgo/configs/config.dev.yaml

.PHONY: run
run:
	go run ./cmd/app/ -conf $(configPath)

.PHONY: gen
gen:
	go run ./cmd/gen/ -conf $(configPath)
