ARTIFACT_NAME := bloomfilter-tool

build:
	@go build -o bin/${ARTIFACT_NAME} main.go bloom_filter.go

run:
	@go run main.go bloom_filter.go

clean:
	rm -f bin/${ARTIFACT_NAME}
