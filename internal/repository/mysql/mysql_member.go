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

func (repo mysqlMemberRepository) Store(ctx context.Context, m *domain.Member) error {
	query := `INSERT  members SET username=? , password=? , nickname=?, privacy_agreed_at=?, policy_agreed_at=?, created_at=?, updated_at=?`
	stmt, err := repo.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, m.Username, m.Password, m.Nickname, m.PrivacyAgreedAt, m.PolicyAgreedAt, m.CreatedAt, m.UpdatedAt)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	m.MemberId = lastID
	return nil
}

func (repo mysqlMemberRepository) createMemberTable() error {

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS members (
		    member_id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			nickname VARCHAR(255),
			privacy_agreed_at DATETIME,
			policy_agreed_at DATETIME,
			created_at DATETIME,
			updated_at DATETIME
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
