# go-tg-store

[![codecov](https://codecov.io/gh/mr-linch/go-tg-store/branch/main/graph/badge.svg?token=os3QgtAZ3k)](https://codecov.io/gh/mr-linch/go-tg-store)

That repository contains diffrent implementations of [go-tg](https://github.com/mr-linch/go-tg) session store.
Each implementation is a separate module.

## sessionbolt [![Go Reference](https://pkg.go.dev/badge/github.com/mr-linch/go-tg-store/sessionbolt.svg)](https://pkg.go.dev/github.com/mr-linch/go-tg-store/sessionbolt)

This is a [BoltDB](https://pkg.go.dev/go.etcd.io/bbolt) implementation of [session store](https://pkg.go.dev/github.com/mr-linch/go-tg@latest/tgb/session#Store).
It uses a latest version.

### Installation

```bash
go get github.com/mr-linch/go-tg-store/sessionbolt
```

## sessionredis [![Go Reference](https://pkg.go.dev/badge/github.com/mr-linch/go-tg-store/sessionredis.svg)](https://pkg.go.dev/github.com/mr-linch/go-tg-store/sessionredis)

This is a [go-redis](https://github.com/go-redis/redis) implementation of [session store](https://pkg.go.dev/github.com/mr-linch/go-tg/tgb/session#Store).

### Installation

```bash
go get github.com/mr-linch/go-tg-store/sessionredis
```

## sessionsql [![Go Reference](https://pkg.go.dev/badge/github.com/mr-linch/go-tg-store/sessionsql.svg)](https://pkg.go.dev/github.com/mr-linch/go-tg-store/sessionsql)

This generic SQL implementation of [session store](https://pkg.go.dev/github.com/mr-linch/go-tg/tgb/session#Store). It uses [database/sql](https://pkg.go.dev/database/sql) package as backend.

### Installation

```bash
go get github.com/mr-linch/go-tg-store/sessionsql
```
