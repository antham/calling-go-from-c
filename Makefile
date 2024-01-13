build: build-go-lib move-go-lib build-binary move-binary

build-go-lib:
	@cd go; \
	CGO_ENABLED=1 go build -buildmode=c-archive -o libhttp.a

move-go-lib:
	@mv go/lib* c

build-binary:
	@cd c; \
	gcc main.c -o main -L. -framework CoreFoundation -framework Security -lhttp

move-binary:
	@mv c/main bin;

run:
	@./bin/main

clean:
	@git clean -f -x
