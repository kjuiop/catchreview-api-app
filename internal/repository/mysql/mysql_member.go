package mysql

import (
	"catchreview-api-app/domain"
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlMemberRepository struct {
	Conn      *sql.DB
	tableName string
}

func NewMysqlMemberRepository(conn *sql.DB) domain.MemberRepository {

	repo := &mysqlMemberRepository{Conn: conn, tableName: "Member"}
	exist, err := repo.checkExistMemberTable()
	if err != nil {
		log.Fatalln("checkExistMemberTable err : ", err)
	}
	if exist {
		return repo
	}

	if err := repo.createMemberTable(); err != nil {
		log.Fatalln("createMemberTable err : ", err)
	}

	return repo
}

func (repo mysqlMemberRepository) Store(ctx context.Context, m *domain.Member) {
	//TODO implement me
	panic("implement me")
}

func (repo mysqlMemberRepository) createMemberTable() error {

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS Members (
			Username VARCHAR(255) PRIMARY KEY,
			Password VARCHAR(255) NOT NULL,
			Nickname VARCHAR(255),
			PrivacyAgreedAt DATETIME,
			PolicyAgreedAt DATETIME,
			CreatedAt DATETIME,
			UpdatedAt DATETIME
		);
	`

	_, err := repo.Conn.Exec(createTableQuery)
	if err != nil {
		return err
	}

	return nil
}

func (repo mysqlMemberRepository) checkExistMemberTable() (bool, error) {
	query := "SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = ?"
	var count int
	if err := repo.Conn.QueryRow(query, repo.tableName).Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}
