dev:
	go run main.go
docker-build:
	docker-compose pull || true
	docker-compose build || true
docker: docker-down docker-build
	docker-compose up -d
	docker-compose ps
	docker-compose logs -f
docker-down:
	docker-compose down
	docker rmi -f niels58/go-rest-gin-k8s:v1 || true