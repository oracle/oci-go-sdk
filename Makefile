build: fmt
	go build

fmt:
	gofmt -w .

clean:
	git clean -df
