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

type MessagesRepoSuite struct {
	suite.Suite

	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository MessageRepository
}

func (s *MessagesRepoSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.repository = *NewMessageRepo(s.DB)
}

func (s *MessagesRepoSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestMessageRepoInit(t *testing.T) {
	suite.Run(t, new(MessagesRepoSuite))
}

func (s *MessagesRepoSuite) TestCreateMessage() {
	var fromUser = models.User{
		Model:    gorm.Model{ID: 11},
		Username: "user11",
		FullName: "User 11",
		Password: "123456",
	}
	var toUser = models.User{
		Model:    gorm.Model{ID: 12},
		Username: "user12",
		FullName: "User 12",
		Password: "123456",
	}
	var message = models.MessageCreate{
		Message: "Hello",
	}
	var messageID = "44"

	s.mock.MatchExpectationsInOrder(false)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "messages"
		("created_at","updated_at","deleted_at","from_id","to_id","message")
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING "messages"."id"`,
	)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil,
			fromUser.ID, toUser.ID, message.Message).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(messageID))
	s.mock.ExpectCommit()

	createdMessage := s.repository.CreateMessage(
		fromUser, toUser, message,
	)

	assert.NotNil(s.T(), createdMessage.ID)
	assert.Equal(s.T(), fmt.Sprint(fromUser.ID), fmt.Sprint(createdMessage.FromID))
	assert.Equal(s.T(), fmt.Sprint(toUser.ID), fmt.Sprint(createdMessage.ToID))
	assert.Equal(s.T(), messageID, fmt.Sprint(createdMessage.ID))
	assert.Equal(s.T(), message.Message, createdMessage.Message)
}

func (s *MessagesRepoSuite) TestFetchUsersMessages() {
	var now = time.Now()
	var fromUser = models.User{
		Model:    gorm.Model{ID: 11},
		Username: "user11",
		FullName: "User 11",
		Password: "123456",
	}
	var toUser = models.User{
		Model:    gorm.Model{ID: 12},
		Username: "user12",
		FullName: "User 12",
		Password: "123456",
	}

	rows := sqlmock.NewRows([]string{
		"id", "created_at", "updated_at", "deleted_at",
		"from_id", "to_id", "message",
	}).
		AddRow(1, now, now, nil, fromUser.ID, toUser.ID, "Hello 1").
		AddRow(2, now, now, nil, fromUser.ID, toUser.ID, "hello").
		AddRow(3, now, now, nil, fromUser.ID, toUser.ID, "Hi").
		AddRow(4, now, now, nil, fromUser.ID, toUser.ID, "Say Hello")

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT messages.message,
				messages.created_at,
				from_user.username as from_user,
				to_user.username as to_user
		FROM "messages"
		JOIN users as from_user ON from_user.id = messages.from_id
		JOIN users as to_user ON to_user.id = messages.to_id
		WHERE ((from_id = $1 AND to_id = $2)) OR ((from_id = $3 AND to_id = $4))`,
	)).
		WithArgs(fromUser.ID, toUser.ID, toUser.ID, fromUser.ID).
		WillReturnRows(rows)

	messages := s.repository.FetchUsersMessages(fromUser, toUser)

	assert.Equal(s.T(), len(messages), 4)
}
