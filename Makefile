TEST?=$$(go list ./... | grep -v vendor)
VETARGS?=-all
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: test vet

clean:
	rm -Rf $(CURDIR)/bin/*

build: clean vet
	govendor build -o $(CURDIR)/bin/sacloud-upload-archive $(CURDIR)/main.go

build-x: clean vet
	sh -c "'$(CURDIR)/scripts/build.sh'"

test: vet
	govendor test $(TEST) $(TESTARGS) -timeout=30s -parallel=4

vet: fmt
	@echo "go tool vet $(VETARGS) ."
	@go tool vet $(VETARGS) $$(ls -d */ | grep -v vendor) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

docker-test: 
	sh -c "'$(CURDIR)/scripts/build_on_docker.sh' 'test'"

docker-build: clean 
	sh -c "'$(CURDIR)/scripts/build_on_docker.sh' 'build-x'"


.PHONY: default test vet testacc fmt fmtcheck
