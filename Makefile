all:
	CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"'

with-docker:
	docker run -it --rm -v $(PWD):/go-cobra-practice --workdir="/go-cobra-practice" golang:1.14-alpine ./build.sh


