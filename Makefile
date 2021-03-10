.PHONY: start
start: build
	./main

.PHONY: build
build:
	go build src/main/main.go

.PHONY: clean
clean:
	rm ./main

.PHONY: clean/db
clean/db:
	rm test.sqlite3

.PHONY: backup/db
backup/db:
	mv test.sqlite3 test.sqlite3.backup

