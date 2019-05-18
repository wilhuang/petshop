package models

import (
	"petshop/mapper"
	"strconv"

	"github.com/astaxie/beego"
)

type CartInfo struct {
	Uid  string
	Pid  string
	Pnum string
}

type ShowInfo struct {
	Pid    string
	Pname  string
	Oprice string
	Pimg1  string
	Pnum   string
	Total  int
}

//TODO
func AddCart(uid, pid string) (err error) {
	tx, err := mapper.GetTx()
	if err != nil {
		return
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewCartMapper(tx)
	row, err := mappers.CheckCart(uid, pid)
	if err == nil {
		pnum, err := strconv.Atoi(row.Pnum)
		if err != nil {
			return err
		}
		return mappers.ResetPnum(uid, pid, strconv.FormatInt(int64(pnum+1), 10))
	}
	_, err = mappers.AddCart(uid, pid, strconv.FormatInt(1, 10))
	return
}

func CheckCart(uid, pid string) (info CartInfo, err error) {
	tx, err := mapper.GetTx()
	if err != nil {
		return
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewCartMapper(tx)
	row, err := mappers.CheckCart(uid, pid)
	if err != nil {
		return
	}
	info.Pnum = row.Pnum
	info.Uid = row.Uid
	info.Pid = row.Pid
	return
}

func UpdatePnum(uid, pid, pnum string) error {
	tx, err := mapper.GetTx()
	if err != nil {
		return err
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewCartMapper(tx)
	return mappers.ResetPnum(uid, pid, pnum)
}

func DeleteCart(uid, pid string) error {
	tx, err := mapper.GetTx()
	if err != nil {
		return err
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewCartMapper(tx)
	return mappers.DeleteCart(uid, pid)
}

func FindCart(uid string) (cartList []CartInfo, err error) {
	tx, err := mapper.GetTx()
	if err != nil {
		return
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewCartMapper(tx)
	rows, err := mappers.FindCartByUid(uid)
	if err != nil {
		return
	}
	for _, row := range rows {
		info := CartInfo{}
		info.Uid = row.Uid
		info.Pid = row.Pid
		info.Pnum = row.Pnum
		cartList = append(cartList, info)
	}
	return cartList, nil
}
func CartShow(uid string) (allnum, allprice int, cartList []ShowInfo, err error) {
	carts, err := FindCart(uid)
	if err != nil {
		return 0, 0, nil, err
	}
	beego.Info(carts)
	for _, row := range carts {
		info, err := FindProductByPid(row.Pid)
		if err != nil {
			return 0, 0, nil, err
		}
		infos := ShowInfo{}
		infos.Pid = row.Pid
		infos.Pname = info.Pname
		infos.Oprice = info.Oprice
		infos.Pimg1 = info.Pimg1
		infos.Pnum = row.Pnum
		num, _ := strconv.Atoi(row.Pnum)
		price, _ := strconv.Atoi(info.Oprice)
		infos.Total = num * price
		cartList = append(cartList, infos)
		allnum = allnum + num
		allprice = allprice + num*price
	}
	return allnum, allprice, cartList, nil
}

func GetAllCart() (cartList []CartInfo, err error) {
	tx, err := mapper.GetTx()
	if err != nil {
		return
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewCartMapper(tx)
	rows, err := mappers.FindCartAll()
	if err != nil {
		return
	}
	for _, row := range rows {
		info := CartInfo{}
		info.Uid = row.Uid
		info.Pid = row.Pid
		info.Pnum = row.Pnum
		cartList = append(cartList, info)
	}
	return cartList, nil
}
