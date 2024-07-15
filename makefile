GO_MODULE := github.com/dwprz/prasorganic-proto

.PHONY: clean
clean: 
	rm --force --recursive ./protogen 
	mkdir -p ./protogen

.PHONY: protoc-go
protoc-go:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	--go-grpc_opt=module=${GO_MODULE} --go-grpc_out=. \
	./proto/email/*.proto  ./proto/email/type/*.proto \
	./proto/user/*.proto ./proto/user/type/*.proto \

.PHONY: build
build: clean protoc-go