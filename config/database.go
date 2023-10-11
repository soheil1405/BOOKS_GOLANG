package config

import  (
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	_"github.com/lib/pq" // postgres golang driver
	"crud/helper"	
)


const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "postgres"
	dbName = "test"
)
func DatabaseConnection() *sql.DB{

	sqlInfo  :=fmt.Sprint("host=%s port=%d user=%s password=%s dbName =%s sslmode = disable" , host , port , user , password ,dbName )

	db , err := sql.Open("postgres" , sqlInfo)

	helper.PanicIfErr(err)
	err = db.Ping()
	helper.PanicIfErr(err)


	log.Info().Msg("connected to database")


	return db
}