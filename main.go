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
