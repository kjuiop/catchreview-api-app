package mysql

import (
	"catchreview-api-app/domain"
	"context"
	"database/sql"
	"log"
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

	return repo
}

func (repo mysqlMemberRepository) Store(ctx context.Context, m *domain.Member) {
	//TODO implement me
	panic("implement me")
}

func (repo mysqlMemberRepository) checkExistMemberTable() (bool, error) {
	query := "SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = ?"
	var count int
	if err := repo.Conn.QueryRow(query, repo.tableName).Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}
