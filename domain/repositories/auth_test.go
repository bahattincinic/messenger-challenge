package repositories

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bahattincinic/messenger-challenge/domain/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite

	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository AuthRepository
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.repository = *NewAuthRepo(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestCreateAccessToken() {
	var (
		token   = "121dsdsaaa11"
		userID  = uint(1)
		tokenID = "33"
	)

	var user = models.User{}
	user.ID = userID

	s.mock.MatchExpectationsInOrder(false)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "accesstokens"
		("created_at","updated_at","deleted_at","token","user_id")
		VALUES ($1,$2,$3,$4,$5)
		RETURNING "accesstokens"."id"`,
	)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, token, userID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(tokenID))
	s.mock.ExpectCommit()

	createdToken := s.repository.CreateAccessToken(token, user)

	assert.NotNil(s.T(), createdToken.ID)
	assert.Equal(s.T(), token, createdToken.Token)
	assert.Equal(s.T(), userID, createdToken.UserID)
	assert.Equal(s.T(), tokenID, fmt.Sprint(createdToken.ID))
}

func (s *Suite) TestGetUser() {
	var (
		userID   = 1
		username = "test"
		password = "test123"
		fullName = "Test User"
		now      = time.Now()
	)

	rows := sqlmock.NewRows([]string{
		"id", "created_at", "updated_at", "deleted_at",
		"username", "full_name", "password",
	}).
		AddRow(userID, now, now, nil, username, fullName, password)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"
		WHERE "users"."deleted_at" IS NULL AND
		(("users"."username" = $1) AND ("users"."password" = $2))
		ORDER BY "users"."id" ASC LIMIT 1`,
	)).
		WithArgs(username, password).
		WillReturnRows(rows)

	user, err := s.repository.GetUser(username, password)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), username, user.Username)
	assert.Equal(s.T(), fullName, user.FullName)
	assert.Equal(s.T(), fmt.Sprint(userID), fmt.Sprint(user.ID))
}

func (s *Suite) TestGetUserNotFound() {
	var (
		username = "test"
		password = "test123"
	)

	rows := sqlmock.NewRows([]string{
		"id", "created_at", "updated_at", "deleted_at",
		"username", "full_name", "password",
	})

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"
		WHERE "users"."deleted_at" IS NULL AND
		(("users"."username" = $1) AND ("users"."password" = $2))
		ORDER BY "users"."id" ASC LIMIT 1`,
	)).
		WithArgs(username, password).
		WillReturnRows(rows)

	user, err := s.repository.GetUser(username, password)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), fmt.Sprint(user.ID), "0")
}

func (s *Suite) TestCreateUser() {
	var (
		username = "test"
		password = "passwd"
		fullname = "example"
		userID   = "44"
	)

	s.mock.MatchExpectationsInOrder(false)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users"
		("created_at","updated_at","deleted_at","username","full_name","password")
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING "users"."id"`,
	)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, username, fullname, password).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(userID))
	s.mock.ExpectCommit()

	createdUser := s.repository.CreateUser(username, password, fullname)

	assert.NotNil(s.T(), createdUser.ID)
	assert.Equal(s.T(), username, createdUser.Username)
	assert.Equal(s.T(), fullname, createdUser.FullName)
	assert.Equal(s.T(), password, createdUser.Password)
	assert.Equal(s.T(), userID, fmt.Sprint(createdUser.ID))
}

func (s *Suite) TestCheckAccessToken() {
	var (
		token   = "dsdsdsdsdddd"
		tokenID = 55
		now     = time.Now()
		userID  = uint(3)
	)

	rows := sqlmock.NewRows([]string{
		"id", "created_at", "updated_at", "deleted_at",
		"token", "user_id",
	}).
		AddRow(tokenID, now, now, nil, token, userID)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "accesstokens"
		WHERE "accesstokens"."deleted_at" IS NULL AND
		(("accesstokens"."token" = $1))
		ORDER BY "accesstokens"."id" ASC
		LIMIT 1`,
	)).
		WithArgs(token).
		WillReturnRows(rows)

	obj, err := s.repository.CheckAccessToken(token)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), token, obj.Token)
	assert.Equal(s.T(), userID, obj.UserID)
	assert.Equal(s.T(), fmt.Sprint(tokenID), fmt.Sprint(obj.ID))
}

func (s *Suite) TestCheckAccessTokenNotFound() {
	var token = "1212121"

	rows := sqlmock.NewRows([]string{
		"id", "created_at", "updated_at", "deleted_at",
		"token", "user_id",
	})

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "accesstokens"
		WHERE "accesstokens"."deleted_at" IS NULL AND
		(("accesstokens"."token" = $1))
		ORDER BY "accesstokens"."id" ASC
		LIMIT 1`,
	)).
		WithArgs(token).
		WillReturnRows(rows)

	obj, err := s.repository.CheckAccessToken(token)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), fmt.Sprint(obj.ID), "0")
}
