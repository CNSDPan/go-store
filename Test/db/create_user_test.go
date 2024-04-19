package db

import (
	"github.com/segmentio/ksuid"
	"store/model/db"
	"strconv"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	var user db.Users
	k := ksuid.New()
	for i := 1; i < 6; i++ {

		user = db.Users{
			UserID:    DBModel.Node.Generate().Int64(),
			Token:     k.Next().String(),
			Status:    db.USER_STATUS_1,
			Name:      "用户" + strconv.Itoa(i),
			Fund:      0,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}
		DBModel.DB.Create(user)
	}
}
