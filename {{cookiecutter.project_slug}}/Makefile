PROJECT_NAME={{cookiecutter.project_name}}

e ?= local.env

ENV_FILE=~/Projects/envs/${PROJECT_NAME}/${e}

include ${ENV_FILE}
export $(shell sed 's/=.*//' ${ENV_FILE})

MIGRATION_DIR=store/migrations


MIGRATION_TABLE_PATH=$(MIGRATION_DIR)/tables
MIGRATION_TABLE_TABLE=schema_migrations

MIGRATION_FUNCTION_PATH=$(MIGRATION_DIR)/functions
MIGRATION_FUNCTION_TABLE=function_migrations

runserver:
	go run cmd/server/main.go

migrate_create:
	migrate create -ext sql -dir ${MIGRATION_TABLE_PATH} -seq ${seq}

migrate_up:
	migrate -database "${DATABASE_URL}?sslmode=disable&x-migrations-table=${MIGRATION_TABLE_TABLE}" -path ${MIGRATION_TABLE_PATH} up

migrate_down:
	migrate -database "${DATABASE_URL}?sslmode=disable&x-migrations-table=${MIGRATION_TABLE_TABLE}" -path ${MIGRATION_TABLE_PATH} down


function_create:
	migrate create -ext sql -dir ${MIGRATION_FUNCTION_PATH} -seq ${seq}

function_up:
	migrate -database "${DATABASE_URL}?sslmode=disable&x-migrations-table=${MIGRATION_FUNCTION_TABLE}" -path ${MIGRATION_FUNCTION_PATH} up

function_down:
	migrate -database "${DATABASE_URL}?sslmode=disable&x-migrations-table=${MIGRATION_FUNCTION_TABLE}" -path ${MIGRATION_FUNCTION_PATH} down

format_models:
	gofmt -w store/models

seed:
	go build -o bin/seed cmd/seed/main.go && ./bin/seed -seed=${seed}

gen_models:
	sqlc generate
