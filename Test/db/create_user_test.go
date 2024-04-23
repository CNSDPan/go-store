package db

import (
	"fmt"
	"github.com/segmentio/ksuid"
	"store/common"
	"store/model/db"
	"strconv"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	users := []db.Users{}
	k := ksuid.New()
	for i := 1; i < 6; i++ {
		users = append(users, db.Users{
			UserID:    DBModel.Node.Generate().Int64(),
			Token:     k.Next().String(),
			Status:    db.USER_STATUS_1,
			Name:      "用户" + strconv.Itoa(i),
			Fund:      common.EnterExchange(1000),
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		})

	}
	fmt.Printf("%v \r\n", users)
	res := DBModel.DB.Create(&users)
	fmt.Printf("创建 %v", res.Error)
}
