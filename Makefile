postgres:
	docker run --name postgresmain -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres
createdb":
	docker exec -it postgresmain createdb --username=root --owner=root simplebank

dropdb:
		docker exec -it postgresmain createdb simplebank
sqlc:
	docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc generate

.PHONY: postgres createdb dropdb sqlc