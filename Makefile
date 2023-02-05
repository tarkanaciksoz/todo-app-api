run:
	APP_ENV=local go run main.go
mockgen:
	mockgen -destination=internal/mocks/mock_handler.go -package=mocks github.com/tarkanaciksoz/api-todo-app/internal/todo Handler
	mockgen -destination=internal/mocks/mock_db.go -package=mocks github.com/tarkanaciksoz/api-todo-app/internal/todo DB
	mockgen -destination=internal/mocks/mock_service.go -package=mocks github.com/tarkanaciksoz/api-todo-app/internal/todo Service
test:
	go test ./... -v

build-local:
	docker-compose --env-file ./.env.local up --build -d
build-test:
	docker-compose --env-file ./.env.test up --build -d
build-prod:
	docker-compose --env-file ./.env.prod up --build -d