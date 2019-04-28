package models

import (
	"petshop/mapper"
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
		return
	}
	err = mappers.CheckLoginInfo(username, password)
	if err != nil {
		return
	}
	info = AccountInfo{Uid: row.Uid, UserName: row.UserName, Role: row.Role, Email: row.Email}
	return info, nil
}

//TODO
func Register(username, password, email string) (err error) {
	id := "a"
	role := "0"
	tx, err := mapper.GetTx()
	if err != nil {
		return
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewAccountMapper(tx)
	_, err = mappers.CreateAccount(id, username, role, email, password)
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

func UnRegister(username string) error {
	tx, err := mapper.GetTx()
	if err != nil {
		return err
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewAccountMapper(tx)
	return mappers.DeleteAccountByUserName(username)
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

func GetAllUser() (accountList []AccountInfo, err error) {
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
	for _, row := range rows {
		info := AccountInfo{}
		info.Uid = row.Uid
		info.UserName = row.UserName
		info.Email = row.Email
		accountList = append(accountList, info)
	}
	return accountList, nil
}
