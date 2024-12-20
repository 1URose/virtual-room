@echo off
:: Define Goose migration directory and PostgreSQL connection string
set "GOOSE_DIR=."
set "CONN_STR=postgres://db-user:db-password@localhost:5432/base-db?sslmode=disable"

:: Run Goose migrations
goose -dir %GOOSE_DIR% postgres "%CONN_STR%" up

:: Pause to view output
pause
