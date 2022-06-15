#!/bin/bash
set -ex
go env -w GOOS=windows
go build -o calculateLoan.exe