mockgen:
	mockgen -destination=internal/mocks/mock_handler.go -package=mocks github.com/tarkanaciksoz/api-todo-app/internal/todo Handler
	mockgen -destination=internal/mocks/mock_db.go -package=mocks github.com/tarkanaciksoz/api-todo-app/internal/todo DB
	mockgen -destination=internal/mocks/mock_service.go -package=mocks github.com/tarkanaciksoz/api-todo-app/internal/todo Service
test:
	go test ./... -v
run-test:
	docker-compose --env-file ./.env.test up -d
run-prod:
	docker-compose --env-file ./.env.prod up -d
build-test:
	docker-compose --env-file ./.env.test up --build -d
build-prod:
	docker-compose --env-file ./.env.prod up --build -d

install-dev:
	go mod download
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
	go install github.com/golang/mock/mockgen@v1.6.0
dev:
	export APP_ENV=dev && air