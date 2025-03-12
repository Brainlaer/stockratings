package connectiondb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func ConnectDB()*sql.DB{
	err:=godotenv.Load()
	if err != nil {
		log.Fatal("Error al obtener las variables de base de datos")
	}

	dsn := fmt.Sprintf(
		"postgresql://%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLLMODE"),
	)
	
	fmt.Println("DSN:", dsn) 
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error al conectarse a la base de datos: %v", err)
	}
	
	
	err = db.Ping()
	if err != nil{
		log.Fatal("No se pudo conectar a la base de datos")
	}
	fmt.Println("Conexión exitosa")

	return db
}