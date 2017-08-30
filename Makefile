build: fmt
	go build

build-gen:
	go install ./cmd/modelmerger

gen: build-gen
	go generate ./identity

fmt:
	gofmt -s -w .

clean:
	git clean -dfn
