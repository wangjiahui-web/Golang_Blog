package dao

import (
	"blogProject/pkg/config"
	"blogProject/pkg/zaplogger"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func generateConnectionUrl() string {
	cfg := config.Get().MySqlConfig
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DatabaseName)
}

func New() error {
	url := generateConnectionUrl()
	var err error
	DB ,err = gorm.Open(mysql.Open(url),&gorm.Config{
		Logger:                                   zaplogger.NewGormLogger(zap.L()),
		CreateBatchSize:                          500,

	})
	if err != nil {
		return err
	}
	return nil
}