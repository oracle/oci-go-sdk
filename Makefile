DOC_SERVER_URL=http:\/\/lgl-bybliothece-01.virt.lgl.grungy.us
TARGETS = common common/auth identity core objectstorage loadbalancer database audit
TARGETS_WITH_TESTS = common common/auth integtest
TARGETS_BUILD = $(patsubst %,build-%, $(TARGETS))
TARGETS_LINT = $(patsubst %,lint-%, $(TARGETS))
TARGETS_TEST = $(patsubst %,test-%, $(TARGETS_WITH_TESTS))
TARGETS_RELEASE= $(patsubst %,release-%, $(TARGETS))
GOLINT=$(GOPATH)/bin/golint
LINT_FLAGS=-min_confidence 0.9 -set_exit_status


.PHONY: $(TARGETS_BUILD) $(TARGET_TEST)

build: lint $(TARGETS_BUILD)

test: build $(TARGETS_TEST)

lint: $(TARGETS_LINT)

$(TARGETS_LINT): lint-%:%
	@echo "linting and formatting: $<"
	@(cd $< && gofmt -s -w .)
	@if [ \( $< = common \) -o \( $< = common/auth \) ]; then\
		(cd $< && $(GOLINT) -set_exit_status .);\
	else\
		(cd $< && $(GOLINT) $(LINT_FLAGS) .);\
	fi

$(TARGETS_BUILD): build-%:%
	@echo "building: $<"
	@(cd $< && find . -name '*_integ_test.go' | xargs -I{} mv {} ../integtest)
	@(cd $< && go build -v)

$(TARGETS_TEST): test-%:%
	@(cd $< && OCI_GO_SDK_DEBUG=1 go test -v)

clean:
	git clean -dfn
pre-doc:
	find . -name \*.go |xargs sed -i '' 's/{{DOC_SERVER_URL}}/${DOC_SERVER_URL}/g'

gen-version:
	go generate -x

release: gen-version build
