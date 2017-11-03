DOC_SERVER_URL=http:\/\/lgl-bybliothece-01.virt.lgl.grungy.us
TARGETS = common identity
TARGETS_BUILD = $(patsubst %,build-%, $(TARGETS))
TARGETS_TEST = $(patsubst %,test-%, $(TARGETS))


.PHONY: $(TARGETS_BUILD) $(TARGET_TEST)

build: $(TARGETS_BUILD)

test: $(TARGETS_TEST)

$(TARGETS_BUILD): build-%:%
	@echo "\nbuilding: $<"
	@(cd $< && gofmt -s -w .)
	@(cd $< && go build -v)

$(TARGETS_TEST): test-%:%
	@(cd $< && go test)

clean:
	git clean -dfn
pre-doc:
	find . -name \*.go |xargs sed -i '' 's/{{DOC_SERVER_URL}}/${DOC_SERVER_URL}/g'
