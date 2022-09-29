build:
	docker build -t devsecopshub .

run:
	docker run --name=devsecopshub -p 8080:8080 devsecopshub

migrate:
	migrate -path ./migrations -database 'postgres://app:qwer1234@127.0.0.1:5432/test?sslmode=disable' up


 docker run -d --name=devsecopshub -p 8040:8080 devsecopshub:latest
