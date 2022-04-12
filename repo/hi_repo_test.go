package repo

import (
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//real test with db which we do not want.
/* func TestHiRepo(t *testing.T){
	h := NewRepo()
	got:= h.SayHi()
	want:= dto.MyMessage{Message:"hi from repo" }
	if got!=want {
        t.Errorf("got %q, wanted %q", got, want)
    }
} */
func NewMock() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		panic(err) // Error here
	}

	return gormDB, mock
}

//TestWithMockDB tests the sql queries generated using mock
func TestWithMockDB(t *testing.T) {
	gormDB, mock := NewMock()
	repo := HiRepo{gormDB}
	query := `SELECT * FROM "message" WHERE id=$1 ORDER BY "message"."message_text" LIMIT 1`
	rows := sqlmock.NewRows([]string{"id", "message_text"}).
		AddRow("11e41e52-0333-47c2-af3b-791a68ba6ad0", "hi from postgres db")
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("11e41e52-0333-47c2-af3b-791a68ba6ad0").WillReturnRows(rows)
	res := repo.SayHi()

	assert.NotNil(t, res)
}
