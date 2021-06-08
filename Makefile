list:
	@echo "List of options in this Makefile"
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

test:
	go test ./... -v

bench:
	go test --bench . --benchmem

run:
	go run ./cmd/merge

build:
	go build ./cmd/merge

depgraph:
	godepgraph github.com/tehsphinx/tss | dot -Tpng -o depgraph.png

doc:
	@echo "Documentation URL: http://localhost:6060/pkg/github.com/tehsphinx/tss/"
	@echo "(Press CTRL+C to stop)"
	@godoc -http=:6060 >/dev/null 2>&1
