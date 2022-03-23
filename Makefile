.PHONY: test

help:   ## Show help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

tests:  ## Run unit tests.
	for d in ./*/ ; do \
	 	(cd "$$d" && go get -v -t -d ./... && go test -v -failfast --race ./... -timeout 5000ms); \
	done
