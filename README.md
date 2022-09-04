# benchdb

## Status

[![go test](https://github.com/fbac/benchdb/actions/workflows/gotest.yml/badge.svg)](https://github.com/fbac/benchdb/actions/workflows/gotest.yml)
[![go build](https://github.com/fbac/benchdb/actions/workflows/gobuild.yml/badge.svg)](https://github.com/fbac/benchdb/actions/workflows/gobuild.yml)

## Build

### Requirements

The following binaries must be installed in your system

- psql
- docker

### Steps

NOTE: During the db config process, the password will be asked via prompt.

It is "postgres" and has to be introduced manually.

#### Manual procedure

- Create, configure and populate database

```bash
make db
```

- Create bin/benchdb

```bash
make build
```

- Run bin/benchdb
  Usage as follows

```bash
$ bin/benchdb -help
  -csv-file string
        path to query csv file (default "data-file.csv")
  -max-threads int
        max threads to process csv (default 1)

Example:
bin/benchdb -csv-file test/query.csv
```

#### Automatic procedure

- Alternatively a test-run is available
  This will set the db up and running, and run a query test suite

```bash
make test-run
```

### Additional Makefile targets

- Clean binary

```bash
make clean
```

- Run test suite

```bash
make test
```

- Run coverage report

```bash
make coverage
```

- View docs in browser

```bash
make docs
```

## Statistics gathered

- number of queries run
- total processing time of all queries (measure every query processing time)
- max query time
- min query time
- avg query time
- median query time

## Technical debt

- Divide packages instead of using only package main
- Use real production-ready folder tree, inside internal/ and divided by pkg/ and cmd/
- Create doc.go per package
- Test coverage 100%
- Use logrus to log by level
- Modify timescaledb entrypoint to accept scripts from stdin
- Introduce password automatically during workflow
