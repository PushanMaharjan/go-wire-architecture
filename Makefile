include .env
export

MIGRATE=sql-migrate

create:
	@read -p  "What is the name of migration?" NAME; \
	${MIGRATE} new -config=config/dbconfig.yml $$NAME

.PHONY: create