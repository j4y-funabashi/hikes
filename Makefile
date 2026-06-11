
build:
	docker build -t hikes-web --file ./apps/api/Dockerfile ./apps/api
	docker build -t hikes-api-test --file ./apps/api/Dockerfile.test ./apps/api

test:
	docker run -it --mount type=bind,src=./apps/api,dst=/api hikes-api-test go test -v ./...