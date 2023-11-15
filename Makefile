.PHONY: run
run:
	npx nodemon -e html,go --exec go run main.go --signal SIGTERM
.PHONY: dcup
dcup:
	docker compose up -d
.PHONY: dcdown
dcdown:
	docker compose down