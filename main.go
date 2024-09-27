package main

import (
  "log"
  "fmt"
  "database/sql"
  "github.com/romakot321/filestorage/cmd/web"
  dbCon "github.com/romakot321/filestorage/db/sqlc"

  _ "github.com/lib/pq"
  "github.com/spf13/viper"
)

//	@title			Filestorage by http
//	@version		1.0
//	@description	Written in golang

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Token
//	@description				Description for what is this security definition being used

//	@securitydefinitions.oauth2.password	OAuth2Password
//	@tokenUrl								http://localhost:8080/auth/login

type Config struct {
    DbDriver         string `mapstructure:"DB_DRIVER"`
    DbSource         string `mapstructure:"DB_SOURCE"`
    PostgresUser     string `mapstructure:"POSTGRES_USER"`
    PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
    PostgresDb       string `mapstructure:"POSTGRES_DB"`
    ServerAddress    string `mapstructure:"SERVER_ADDRESS"`
}

func loadConfig(path string) (config Config, err error) {
    viper.AddConfigPath(path)
    viper.SetConfigName(".env")
    viper.SetConfigType("env")

    viper.AutomaticEnv()

    err = viper.ReadInConfig()
    if err != nil {
        return
    }

    err = viper.Unmarshal(&config)
    return
}

func main() {
    config, err := loadConfig(".")

    if err != nil {
        log.Fatalf("could not loadconfig: %v", err)
    }

    conn, err := sql.Open(config.DbDriver, config.DbSource)
    if err != nil {
        log.Fatalf("Could not connect to database: %v", err)
    }

    db := dbCon.New(conn)

    fmt.Println("PostgreSql connected successfully...")

    web.Init(db)
}
