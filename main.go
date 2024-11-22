package main

import (
	"log"
	config "synapsis-ecommerce/config/postgre"
	"synapsis-ecommerce/config/util"
)

func main() {

	env, err := util.LoadConfig(".env")
	if err != nil {
		log.Fatal("cannot load env", err)
	}

	pg := config.Database(env.DBDriver,env.DatabaseURL)
	pg.InitPostgre()
	pg.RunDBMigration(env.MigrationURL)
	defer pg.DB.Close()



	// server := src.InitServer(pg, env)

	// wg := sync.WaitGroup{}

	// wg.Add(1)
	// go func() {

	// 	defer wg.Done()
	// 	server.Start()
	// }()

	// wg.Wait()
}