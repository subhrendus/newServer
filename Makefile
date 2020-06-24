APPNAME := newServer

all: clean build setup-specs run-specs

clean:
	rm -f $(APPNAME)


	
# should be used only once in the local machine
setup-specs:
	go get -u -t github.com/onsi/ginkgo/ginkgo
	go get -u -t github.com/onsi/gomega/...
	mkdir -p /tmp/artifacts # move test coverage

# for running the unit tests in the local machine as well as in the build machine
run-specs:
	ginkgo -r --race --randomizeAllSpecs --randomizeSuites --failOnPending --cover --race -trace
	go tool cover -html=coverage.out
	
test-coverage:
	go build -o $(APPNAME)
	ginkgo  -r --randomizeAllSpecs --randomizeSuites --cover -coverprofile=coverage.out --race --trace
	go tool cover -html=coverage.out

build:
	go build -o $(APPNAME)

run:
	./$(APPNAME)

.PHONY: all