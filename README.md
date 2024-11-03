# Advent of Code Go
> Combined solutions from multiple Advent of Code challenges

# Usage
Starter directories for each puzzle can be generated with the included `newday` program.

```
$ ./bin/newday -h
Usage of ./bin/newday:
  -day int
        day number to create [0-99]
  -dir string
        directory in which to generate the new day files (default ".")
  -year int
        year to create day in (default 2024)

$ ./bin/newday -day 1
$ ./bin/newday -year 2015 -day 1
```

The included `Makefile` can be used to quickly execute all puzzle tests.

```
# The default target both executes all tests and builds the newday program.
$ make
```

A specific year or day can be tested by manually invoking the Go CLI.

```
$ go test ./years/2020/...
$ go test ./years/2020/day16/...
```

# Background
This repo consolidates my previous Advent of Code projects into a single location.  This is being done both to reduce repo-sprawl and to encourage increased investment in a single utility library.
