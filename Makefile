# Ensure DB_CONN_STRING is provided
check-db-conn:
	@if [ -z "$(DB_CONN_STRING)" ]; then \
		echo "Error: DB_CONN_STRING is not set"; \
		exit 1; \
	fi

migration-up: check-db-conn
	GOOSE_MIGRATION_DIR="./migration/" goose mysql "$(DB_CONN_STRING)" up

migration-down: check-db-conn
	GOOSE_MIGRATION_DIR="./migration/" goose mysql "$(DB_CONN_STRING)" down
