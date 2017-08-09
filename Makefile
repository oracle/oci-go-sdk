build: fmt
	go build

build-gen:
	go install ./cmd/modelmerger

gen: build-gen
	go generate ./identity

fmt:
	gofmt -w .

clean:
	git clean -df --exclude=cmd/
