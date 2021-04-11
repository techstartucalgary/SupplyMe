/*
Entity is the main function for our Authentication between on our Server side. We consider the different types of users, and how it connects to our database for Authentication.
*/
package entity

import (
	"database/sql"
	"flag"
	"github.com/CovenantSQL/client"
	"github.com/ConvenantSQL/utils/log"
)

func main() {
	log.Setlevel(log.DebugLevel)
	var config, password, dsn string

	flag.StringVar(&config, "config", "./conf/config.yaml", "config file path")
	flag.StringVar(&dsn, "dsn", "", "database url")
	flag.StringVar(&password, "password", "", "master key password for convenantsql")
	flag.Parse()

	err := client.Int(config, []byte(password))
	if err != nil {
		log.Fatal(err)
	}

	//Insert records

	_, err ;= db.Exec("ITEM")
	if err != nil {
		log.Fatal(err)
	}

	//Query records

	rows, err := db.Query("ITEM")
	if err != nil {
		log.Fatal(err)
	}



}
