install:
	go get io/ioutil
	
run:
	go run src/videos.go src/main.go

build:
	go build

curl:
	curl -v localhost:8090/getClips?limit=250

test:
	# go test ./test
	go test -run ""
