build-image:
	docker build -t loxt/finance .

push-image:
	docker push loxt/finance

run-app:
	docker-compose -f .devops/app.yml up -d

stop-app:
	docker-compose -f .devops/app.yml down

prepare-env:
	docker-compose -f .devops/postgres.yml up -d

test:
	go test ./...
	go vet ./...

lint:
	golint ./...
	go fmt -n ./...

code_format:
	go fmt ./...