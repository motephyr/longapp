package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DatabaseDriver struct {
	Driver      string `yaml:"driver" env:"DB_DRIVER"`
	Host        string `yaml:"host" env:"DB_HOST"`
	Username    string `yaml:"username" env:"DB_USER"`
	Password    string `yaml:"password" env:"DB_PASS"`
	DBName      string `yaml:"db_name" env:"DB_NAME"`
	Port        int    `yaml:"port" env:"DB_PORT"`
	Connections int    `yaml:"connections" env:"DB_CONNECTIONS"`
}

type DatabaseConfig struct {
	*sql.DB
	Drivers map[string]DatabaseDriver `yaml:"drivers"`
	Default DatabaseDriver            `yaml:"default" env:"DEFAULT_DB_DRIVER"`
}

func (d *DatabaseConfig) Setup() error {
	//nolint:wsl,lll
	var err error //nolint:wsl
	connectionString := ""

	if d.DB != nil {
		return nil
	}
	// gormLogger := gorm2.Logger{
	// 	Log: &log.DefaultLogger,
	// }
	// newLogger := gormLogger.LogMode(logger.Info)
	switch d.Default.Driver {
	case "postgres":
		connectionString = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", d.Default.Host, d.Default.Port, d.Default.Username, d.Default.DBName, d.Default.Password)
		d.DB, err = sql.Open("postgres", connectionString)

	default:
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", d.Default.Username, d.Default.Password, d.Default.Host, d.Default.Port, d.Default.DBName)
		d.DB, err = sql.Open(d.Default.Driver, connectionString)

	}
	if err != nil {
		fmt.Println(d.Default)
		panic(err)
	}
	// boil.SetDB(d.DB)

	// d.DB.Use(
	// 	dbresolver.Register(dbresolver.Config{}).
	// 		SetConnMaxLifetime(24 * time.Hour).
	// 		SetMaxIdleConns(100).
	// 		SetMaxOpenConns(100),
	// )
	return nil //nolint:wsl
}
