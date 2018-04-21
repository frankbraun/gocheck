all:
	go install -v github.com/frankbraun/gocheck

.PHONY: test
test:
	gocheck -g -c
