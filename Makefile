list:
	@echo "List of options in this Makefile";
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

test:
	go test ./... -v

bench:
	go test --bench . --benchmem

run:
	go run ./cmd/merge

build:
	go build ./cmd/merge
