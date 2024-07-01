package main

import (
	"github.com/vsingh58/sharding/db/postgres"
)

func main() {

	config := postgres.Config{
		Host:     "localhost",
		Port:     5432,
		Username: "admin",
		Password: "admin_password",
		Database: "test_conn_pool_db",
	}

	// Benchmark the time taken to execute queries with and without connection pooling
	BenchmarkNonPooledConnection(config, 10)
}
