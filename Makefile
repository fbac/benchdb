BINARY = benchdb
BIN_FOLDER = bin
TEST_FOLDER = test
PSQL_BIN = $(which psql)
PG_IS_READY_BIN = $(pg_isready)

.PHONY: build clean test docs

#######################
# benchdb cmd targets #
#######################

build:
	go build -o ${BIN_FOLDER}

clean:
	rm -f ${BIN_FOLDER}/${BINARY}

test:
	go test -v

coverage:
	go test -cover -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	
doc:
	godoc -http localhost:8888

#######################
#    pgsql targets    #
#######################

db: dbclean dbrun dbwait dbcheck dbcfg

dbrun:
	@echo -e "# starting pgsql\n"
	@mkdir -p /tmp/timescaledb/pgsql
	@docker run -d --name timescaledb -p 5432:5432 -v /tmp/timescaledb/pgsql:/var/lib/postgresql/data -e POSTGRES_PASSWORD=postgres timescale/timescaledb-ha:pg14-latest

dbcfg:
	@echo -e "# configure pgsql\n"
	@psql -v -w -U postgres -h127.0.0.1 -p5432 < assets/cpu_usage.sql
	@psql -v -w -U postgres -h127.0.0.1 -p5432 -d homework -c "\COPY cpu_usage FROM assets/cpu_usage.csv CSV HEADER"

dbcheck:
	@echo -e "# checking pgsql readiness\n"
	@docker exec -it timescaledb pg_isready

dbclean:
	@echo -e "# cleaning pgsql container\n"
	@docker stop timescaledb
	@docker rm -v timescaledb
	@rm -rf /tmp/timescaledb/

dbwait:
	@echo -e "# waiting for pgsql to boot\n"
	@sleep 30