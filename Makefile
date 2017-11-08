DOC_SERVER_URL=http:\/\/lgl-bybliothece-01.virt.lgl.grungy.us
TARGETS = common identity core objectstorage loadbalancer database
TARGETS_WITH_TESTS = common integtest
TARGETS_BUILD = $(patsubst %,build-%, $(TARGETS))
TARGETS_TEST = $(patsubst %,test-%, $(TARGETS_WITH_TESTS))


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
