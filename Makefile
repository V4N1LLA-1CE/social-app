# load env
include .env
export

# hop into container
pg-exec:
	docker exec -it ${POSTGRES_CONTAINER_NAME} /bin/bash

# db migrate up
pg-migrateup:
	migrate -path cmd/api/migrate -database "${POSTGRES_DSN}" -verbose up ${n}

# db migrate down
pg-migratedown:
	migrate -path cmd/api/migrate -database "${POSTGRES_DSN}" -verbose down ${n}

# force migrate to specified version
pg-fmigrate:
	migrate -path cmd/api/migrate -database "${POSTGRES_DSN}" force ${v}

# run with live reloading
watch:
	air

.PHONY: watch pg-exec pg-migrateup pg-migratedown pg-fmigrate
