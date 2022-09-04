# benchdb

## Status

[![go test](https://github.com/fbac/benchdb/actions/workflows/gotest.yml/badge.svg)](https://github.com/fbac/benchdb/actions/workflows/gotest.yml)
[![go build](https://github.com/fbac/benchdb/actions/workflows/gobuild.yml/badge.svg)](https://github.com/fbac/benchdb/actions/workflows/gobuild.yml)

## Build

### Requirements

The following binaries are required to run the complete workflow

- psql
- docker

### Steps

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
Usage of bin/benchdb:
  -csv-file string
        path to query csv file

Examples:
bin/benchdb < test/query.csv
bin/benchdb -csv-file test/query.csv
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

- For each worker
  
    1. number of queries run
    2. total processing time of all queries (measure every query processing time)
    3. max query time
    4. min query time
    5. avg query time
    6. median query time

## Technical debt

- Divide packages instead of using package main
- Create doc.go per package
- Use real production-ready folder tree
- Coverage 100%
- Use logrus to log by level
