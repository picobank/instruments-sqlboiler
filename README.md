# SQLBoiler experimentation

This repository expects the database to have been setup in accordance
with https://github.com/picobank/instruments

# Generating SQLBoiler code

    go generate -v

# Testing the generated code

    go test -v ./models/

# Running the sample main program

In the repository directory run:

    go run main.go

Configuration for connecting to the database is read from `sqlboiler.toml`.

# Running the benchmarks

In the repository directory run:

    go test -bench=.

Configuration for connecting to the database is read from `sqlboiler.toml`.
    