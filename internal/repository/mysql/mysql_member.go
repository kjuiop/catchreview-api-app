package mysql

import (
	"catchreview-api-app/domain"
	"context"
	"database/sql"
)

type mysqlMemberRepository struct {
	Conn *sql.DB
}

func NewMysqlMemberRepository(conn *sql.DB) domain.MemberRepository {
	return &mysqlMemberRepository{
		Conn: conn,
	}
}

func (repo mysqlMemberRepository) Store(ctx context.Context, m *domain.Member) {
	//TODO implement me
	panic("implement me")
}
