# load env
include .env
export

# hop into container
exec:
	docker exec -it ${POSTGRES_CONTAINER_NAME} /bin/bash

# db migrate up
up:
	migrate -path migrations -database "${POSTGRES_DSN}" -verbose up ${n}

# db migrate down
down:
	migrate -path migrations -database "${POSTGRES_DSN}" -verbose down ${n}

# force migrate to specified version
fmigrate:
	migrate -path migrations -database "${POSTGRES_DSN}" force ${v}

# run with live reloading
watch:
	air

.PHONY: watch exec up down fmigrate
