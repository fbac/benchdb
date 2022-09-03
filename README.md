# benchdb

## Status

[![go test](https://github.com/fbac/benchdb/actions/workflows/gotest.yml/badge.svg)](https://github.com/fbac/benchdb/actions/workflows/gotest.yml)
[![go build](https://github.com/fbac/benchdb/actions/workflows/gobuild.yml/badge.svg)](https://github.com/fbac/benchdb/actions/workflows/gobuild.yml)

## Build

> Requirements
>
> psql client
>
> docker

- Create bin/benchdb

```bash
make build
```

- Clean binary

```bash
make clean
```

## Usage

## Statistics

- For each worker
  
    1. number of queries run
    2. total processing time of all queries (measure every query processing time)
    3. max query time
    4. min query time
    5. avg query time
    6. median query time

## Technical debt

- Create doc.go per package
- Coverage 100%
