package config

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"gopkg.in/gcfg.v1"
)

func LoadConfig() (*Config, error) {
	filename := "config.toml"

	cfg := &Config{}

	err := gcfg.ReadFileInto(cfg, filename)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func InitDatabase(cfg *Config) (*sqlx.DB, error) {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
	db, err := sqlx.Connect("mysql", source)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.Database.MaxIdleConnection)
	db.SetMaxOpenConns(cfg.Database.MaxOpenConnection)
	db.SetConnMaxIdleTime(time.Duration(cfg.Database.MaxIdletimeConnection) * time.Second)
	db.SetConnMaxLifetime(time.Duration(cfg.Database.MaxLifetimeConnection) * time.Second)

	return db, nil
}

func InitGrpcClient(cfg Config) error {
	//Create connection to gRPC server
	conn, err := grpc.Dial(cfg.GRPC.Host+":"+cfg.GRPC.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to service: %v", err)
		return err
	}
	defer conn.Close()

	//Create new userService client
	// _ = sensorpbA.NewSensorServicesClient(conn)

	return nil
}
