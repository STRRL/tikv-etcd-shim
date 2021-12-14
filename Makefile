.PHONY: binary
binary: bin/tikv-etcd-shim

.PHONY: bin/tikv-etcd-shim
bin/tikv-etcd-shim: test
	go build -o bin/tikv-etcd-shim ./cmd/tikv-etcd-shim/main.go

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: test
test:
	go test ./...