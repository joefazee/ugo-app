BINARY_NAME=ladiwork
APP_NAME=LadiWork

build:
	@go mod vendor
	@echo "Building ${APP_NAME}"
	@go build -o bin/${BINARY_NAME} .
	@echo "Done!"


run: build
	@echo "Starting ${APP_NAME}"
	@./bin/${BINARY_NAME}
	@echo "${APP_NAME} started!"

clean: stop_compose
	@echo "Cleaning.."
	@go clean
	@rm -rf bin
	@echo "Cleaned!"

test:
	@echo "Testing.."
	@go -v test ./...
	@echo "Done!"

start:  run

stop:
	@echo "Stopping ${APP_NAME}"
	@-pkill -SIGTERM -f "./bin/${BINARY_NAME}"
	@echo "Stopped ${APP_NAME}"

restart: stop start

start_compose:
	@docker compose up -d 

stop_compose:
	@docker compose down

start_mail:
	@mailhog
