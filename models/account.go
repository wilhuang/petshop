package models

import (
	"petshop/mapper"
	"time"

	"github.com/astaxie/beego"
)

type AccountInfo struct {
	Uid      string
	UserName string
	Role     string
	Email    string
}

func Login(username, password string) (info AccountInfo, err error) {
	tx, err := mapper.GetTx()
	if err != nil {
		return
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewAccountMapper(tx)
	row, err := mappers.FindAccountByUserName(username)
	if err != nil {
		beego.Info(err, "用户名不存在")
		return
	}
	err = mappers.CheckLoginInfo(username, password)
	if err != nil {
		beego.Info(err, "密码错误")
		return
	}
	info = AccountInfo{Uid: row.Uid, UserName: row.UserName, Role: row.Role, Email: row.Email}
	return info, nil
}

//TODO
func Register(username, password, email string) (err error) {
	uid := time.Now().Format("0101150405")
	beego.Info(uid)
	role := "0"
	tx, err := mapper.GetTx()
	if err != nil {
		return
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewAccountMapper(tx)
	_, err = mappers.CreateAccount(uid, username, role, email, password)
	return
}

func ResetEmail(username, email string) error {
	key := "email"
	tx, err := mapper.GetTx()
	if err != nil {
		return err
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewAccountMapper(tx)
	return mappers.ResetAccount(username, key, email)
}

func ResetPwd(username, email string) error {
	key := "password"
	tx, err := mapper.GetTx()
	if err != nil {
		return err
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewAccountMapper(tx)
	return mappers.ResetAccount(username, key, email)
}

func UpdateInfo(uid, username, password, email string) error {
	tx, err := mapper.GetTx()
	if err != nil {
		return err
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewAccountMapper(tx)
	return mappers.ResetUserInfo(uid, username, password, email)
}

func UpdateInfoAdmin(uid, username, email string) error {
	tx, err := mapper.GetTx()
	if err != nil {
		return err
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewAccountMapper(tx)
	return mappers.UpdateUserInfo(uid, username, email)
}

func UnRegister(username string) error {
	tx, err := mapper.GetTx()
	if err != nil {
		return err
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewAccountMapper(tx)
	return mappers.DeleteAccountByUserName(username)
}

func UnRegisterById(uid string) error {
	tx, err := mapper.GetTx()
	if err != nil {
		return err
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewAccountMapper(tx)
	return mappers.DeleteAccountById(uid)
}

func FindUserByUserName(username string) (info AccountInfo, err error) {
	tx, err := mapper.GetTx()
	if err != nil {
		return
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewAccountMapper(tx)
	row, err := mappers.FindAccountByUserName(username)
	if err != nil {
		return
	}
	info = AccountInfo{Uid: row.Uid, UserName: row.UserName, Role: row.Role, Email: row.Email}
	return info, nil
}

func GetAllUser() (num int, accountList []AccountInfo, err error) {
	tx, err := mapper.GetTx()
	if err != nil {
		return
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewAccountMapper(tx)
	rows, err := mappers.FindAccountAll()
	if err != nil {
		return
	}
	i := 0
	for _, row := range rows {
		i++
		info := AccountInfo{}
		info.Uid = row.Uid
		info.UserName = row.UserName
		info.Email = row.Email
		accountList = append(accountList, info)
	}
	return i, accountList, nil
}
