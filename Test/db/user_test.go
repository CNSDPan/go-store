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

func GetUserById1() (db.UserApi, error) {
	var err error
	user := db.UserApi{
		UserID: 1,
	}
	defer func() {
		if err != nil {
			user.ID = 0
		}
	}()

	err = DBModel.DB.Model(db.Users{}).First(&user).Error
	fmt.Printf("get user by ID:1 err %v", err)
	fmt.Printf("user %v", user)
	return user, err
}

func GetUserById2() (db.UserApi, error) {
	var err error
	user := db.UserApi{
		UserID: 2,
	}
	defer func() {
		if err != nil {
			user.ID = 0
		}
	}()

	err = DBModel.DB.Model(db.Users{}).First(&user).Error
	fmt.Printf("get user by ID:1 err %v", err)
	fmt.Printf("user %v", user)
	return user, err
}

func GetUserById3() (db.UserApi, error) {
	var err error
	user := db.UserApi{
		UserID: 3,
	}
	defer func() {
		if err != nil {
			user.ID = 0
		}
	}()

	err = DBModel.DB.Model(db.Users{}).First(&user).Error
	fmt.Printf("get user by ID:1 err %v", err)
	fmt.Printf("user %v", user)
	return user, err
}

func GetUserById4() (db.UserApi, error) {
	var err error
	user := db.UserApi{
		UserID: 4,
	}
	defer func() {
		if err != nil {
			user.ID = 0
		}
	}()

	err = DBModel.DB.Model(db.Users{}).First(&user).Error
	fmt.Printf("get user by ID:1 err %v", err)
	fmt.Printf("user %v", user)
	return user, err
}

func GetUserById5() (db.UserApi, error) {
	var err error
	user := db.UserApi{
		UserID: 5,
	}
	defer func() {
		if err != nil {
			user.ID = 0
		}
	}()

	err = DBModel.DB.Model(db.Users{}).First(&user).Error
	fmt.Printf("get user by ID:1 err %v", err)
	fmt.Printf("user %v", user)
	return user, err
}
