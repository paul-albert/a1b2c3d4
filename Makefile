include .env

.SILENT:
	go_version
	os_version
	dependencies
	install
	build
	compile
	run
	build_and_run
	clean
	cleanup
	test
	tests

go_version:
	echo ${GO_VERSION}

os_version:
	echo ${OS}

install: clean
	go mod download
	go get -u ${SRC_DIRECTORY}
	go install ${SRC_DIRECTORY}
	go mod tidy

dependencies: install

build: clean
	GOARCH=${ARCHITECTURE} \
	GOOS=${OS} \
		go build -o ${OUTPUT_DIRECTORY}/${APP_BINARY_FILE}-${OS} ${SRC_DIRECTORY}/${APP_SOURCE}

compile: build

run: clean build
	${OUTPUT_DIRECTORY}/${APP_BINARY_FILE}-${OS}

build_and_run: build run

clean:
	go clean
	rm -f ${OUTPUT_DIRECTORY}/${APP_BINARY_FILE}-${OS}
	go mod tidy

cleanup: clean

test: clean build
	go clean -testcache
	go test -v -count=1 ${TESTS_DIRECTORY}

tests: test
