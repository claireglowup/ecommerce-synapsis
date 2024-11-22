create migrate:
	migrate create -ext sql -dir config/migrations -seq init_table

 migrateup: 
	migrate -database postgresql://root:secret@localhost:5432/synapsis?sslmode=disable -path config/migrations up

run: 
	go run main.go

sqlc: 
	sqlc generate