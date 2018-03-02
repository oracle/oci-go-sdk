INTEGTEST_DIR = integtest
TEST_HELPERS = test_helpers.go test_service_deps_helpers.go
TEST_FILES = $(notdir $(wildcard $(INTEGTEST_DIR)/*_integ_test.go))
TARGETS = $(patsubst %_client_integ_test.go, %, $(TEST_FILES))
TEST_TARGETS = $(patsubst %_client_integ_test.go, test-%, $(TEST_FILES))

.PHONY: list-test

list-test:
	@echo $(TEST_TARGETS)

test-all: $(TEST_TARGETS)

$(TEST_TARGETS): test-%:%
	@echo Testing $(INTEGTEST_DIR)/$<_client_integ_test.go
	@(cd $(INTEGTEST_DIR) && go test -v $(TEST_HELPERS) $<_client_integ_test.go)

$(TARGETS): %:integtest/%_client_integ_test.go
