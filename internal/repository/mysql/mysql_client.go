package mysql

import (
	"catchreview-api-app/config"
	"database/sql"
	"fmt"
	"log"
)

type Client struct {
	DbConn *sql.DB
	cfg    *config.Config
}

func NewMysqlClient(cfg *config.Config) (*Client, error) {

	jdbcUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", cfg.MySqlInfo.User, cfg.MySqlInfo.Pass, cfg.MySqlInfo.Host, cfg.MySqlInfo.Port, cfg.MySqlInfo.DbName)
	db, err := sql.Open("mysql", jdbcUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Client{DbConn: db, cfg: cfg}, nil
}

func (c *Client) DbClose() {
	if err := c.DbConn.Close(); err != nil {
		log.Println("err : ", err)
	}
}
