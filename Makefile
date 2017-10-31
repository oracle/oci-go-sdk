DOC_SERVER_URL=http:\/\/lgl-bybliothece-01.virt.lgl.grungy.us

build: fmt
	go build

fmt:
	gofmt -s -w .

clean:
	git clean -dfn
pre-doc:
	find . -name \*.go |xargs sed -i '' 's/{{DOC_SERVER_URL}}/${DOC_SERVER_URL}/g'
