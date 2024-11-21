create migrate:
	migrate create -ext sql -dir config/migrations -seq table_name

 migrateup: 
	migrate -database YOUR_DATABASE_URL -path PATH_TO_YOUR_MIGRATIONS up
	