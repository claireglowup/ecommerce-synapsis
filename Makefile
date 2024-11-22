create migrate:
	migrate create -ext sql -dir config/migrations -seq init_table

 migrateup: 
	migrate -database {DATABASE_URL} -path config/migrations up
	