BINDIR := $(CURDIR)/bin

test:
	go mod tidy -modfile=go_test.mod
	go test ./... -modfile go_test.mod -shuffle=on -race

lint:
	go vet -modfile=go_test.mod ./...

test.cover:
	go mod tidy -modfile=go_test.mod
	go test -modfile=go_test.mod -race -shuffle=on -coverprofile=coverage.txt -covermode=atomic ./...

mod.clean:
	rm -f go.mod go.sum
	cat go.mod.bk > go.mod

BENCH_COUNT := 30

bench: bin/benchstat
	perl -pi -e 's|(github.com/qawatake/ctxvls/internal/ctxvls)(\d+)|$$1 . (1)|ge' ctxvls.go
	go test -modfile=go_test.mod -bench=. -benchmem -run=^$$ . -count $(BENCH_COUNT) | tee bench1.txt
	perl -pi -e 's|(github.com/qawatake/ctxvls/internal/ctxvls)(\d+)|$$1 . ($$2+1)|ge' ctxvls.go
	go test -modfile=go_test.mod -bench=. -benchmem -run=^$$ . -count $(BENCH_COUNT) | tee bench2.txt
	perl -pi -e 's|(github.com/qawatake/ctxvls/internal/ctxvls)(\d+)|$$1 . ($$2+1)|ge' ctxvls.go
	go test -modfile=go_test.mod -bench=. -benchmem -run=^$$ . -count $(BENCH_COUNT) | tee bench3.txt
	$(BINDIR)/benchstat bench2.txt bench3.txt bench1.txt
	perl -pi -e 's|(github.com/qawatake/ctxvls/internal/ctxvls)(\d+)|$$1 . (3)|ge' ctxvls.go

bin/benchstat:
	GOBIN=$(BINDIR) go install golang.org/x/perf/cmd/benchstat@latest
