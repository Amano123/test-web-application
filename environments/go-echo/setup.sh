#!/bin/sh

go mod tidy
go mod download
air init 
air -c .air.toml