
# CLI TODO App

This project was made on purpose of learning [Go](https://go.dev/). It's probably really unoptimized, any criticizm is welcome.

## Prerequisites
- [SQLite3](https://www.sqlite.org/)

## Installation & usage

  - Windows
```bash
go build
go install
todo
```

  - Linux
```bash
go build
./todo
```
## Packages used
 -  [database/sql - provides a generic interface around SQL (or SQL-like) databases](https://pkg.go.dev/database/sql)
 -	[fmt                    - implements formatted I/O with functions analogous to C's printf and scanf](https://pkg.go.dev/fmt)
 -	[os                     - provides a platform-independent interface to operating system functionality](https://pkg.go.dev/os)
 -	[strconv                - implements conversions to and from string representations of basic data types](https://pkg.go.dev/strconv)
 -	[text/tabwriter         - translates tabbed columns input into properly aligned text](https://pkg.go.dev/text/tabwriter)
 -	[github.com/spf13/cobra - framework for CLI apps](https://cobra.dev/)
 -  [github.com/mattn/go-sqlite3 - sqlite3 driver that conforms to the built-in database/sql interface](https://github.com/mattn/go-sqlite3)
