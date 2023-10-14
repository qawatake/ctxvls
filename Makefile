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

bench: bin/benchstat
	go test -modfile=go_test.mod -bench=. -benchmem -run=^$ . -count 10 | tee bench1.txt
	perl -pi -e 's|github.com/qawatake/ctxvls/internal/ctxvls|github.com/qawatake/ctxvls/internal/ctxvls2|g' ctxvls.go
	go test -modfile=go_test.mod -bench=. -benchmem -run=^$ . -count 10 | tee bench2.txt
	perl -pi -e 's|github.com/qawatake/ctxvls/internal/ctxvls2|github.com/qawatake/ctxvls/internal/ctxvls|g' ctxvls.go
	benchstat bench1.txt bench2.txt

bench.s: bin/benchstat
	go test -modfile=go_test.mod -bench=. -benchmem -run=^$ . -count 10 > bench1.txt
	perl -pi -e 's|github.com/qawatake/ctxvls/internal/ctxvls|github.com/qawatake/ctxvls/internal/ctxvls2|g' ctxvls.go
	go test -modfile=go_test.mod -bench=. -benchmem -run=^$ . -count 10 > bench2.txt
	perl -pi -e 's|github.com/qawatake/ctxvls/internal/ctxvls2|github.com/qawatake/ctxvls/internal/ctxvls|g' ctxvls.go
	$(BINDIR)/benchstat bench1.txt bench2.txt

bin/benchstat:
	go install golang.org/x/perf/cmd/benchstat@latest
