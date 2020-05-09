package repositories

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UserRepoSuite struct {
	suite.Suite

	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository UserRepository
}

func (s *UserRepoSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.repository = *NewUserRepo(s.DB)
}

func (s *UserRepoSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestUserRepoInit(t *testing.T) {
	suite.Run(t, new(UserRepoSuite))
}

func (s *UserRepoSuite) TestFetchUsers() {
	var now = time.Now()

	rows := sqlmock.NewRows([]string{
		"id", "created_at", "updated_at", "deleted_at",
		"username", "full_name", "password",
	}).
		AddRow(1, now, now, nil, "user1", "User 1", "123456").
		AddRow(2, now, now, nil, "user2", "User 2", "123456").
		AddRow(3, now, now, nil, "user3", "User 3", "123456").
		AddRow(4, now, now, nil, "user4", "User 4", "123456")

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"
		WHERE "users"."deleted_at" IS NULL`,
	)).
		WillReturnRows(rows)

	users := s.repository.FetchUsers()

	assert.Equal(s.T(), len(users), 4)
}

func (s *UserRepoSuite) TestFetchUserByUsername() {
	var (
		userID   = 33
		now      = time.Now()
		username = "user33"
		password = "12121"
		fullName = "Test Test"
	)

	rows := sqlmock.NewRows([]string{
		"id", "created_at", "updated_at", "deleted_at",
		"username", "full_name", "password",
	}).
		AddRow(userID, now, now, nil, username, fullName, password)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"
		 WHERE "users"."deleted_at" IS NULL AND
		 (("users"."username" = $1))
		 ORDER BY "users"."id" ASC
		 LIMIT 1`,
	)).
		WithArgs(username).
		WillReturnRows(rows)

	user, err := s.repository.FetchUserByUsername(username)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), username, user.Username)
	assert.Equal(s.T(), fmt.Sprint(userID), fmt.Sprint(user.ID))
}

func (s *UserRepoSuite) TestFetchUserByID() {
	var (
		userID   = uint(35)
		now      = time.Now()
		username = "user35"
		password = "12125"
		fullName = "Test3 Test"
	)

	rows := sqlmock.NewRows([]string{
		"id", "created_at", "updated_at", "deleted_at",
		"username", "full_name", "password",
	}).
		AddRow(userID, now, now, nil, username, fullName, password)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"
		 WHERE "users"."deleted_at" IS NULL AND
		 ((id = $1))
		 ORDER BY "users"."id" ASC
		 LIMIT 1`,
	)).
		WithArgs(userID).
		WillReturnRows(rows)

	user, err := s.repository.FetchUserByID(uint(userID))

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), username, user.Username)
	assert.Equal(s.T(), fmt.Sprint(userID), fmt.Sprint(user.ID))
}
