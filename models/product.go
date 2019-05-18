package models

import (
	"petshop/mapper"
	"time"

	"github.com/astaxie/beego"
)

type ProductInfo struct {
	Pid    string
	Pname  string
	Oprice string
	Pprice string
	Onhand string
	Pimg1  string
	Pimg2  string
	Pimg3  string
	Pimg4  string
	Pimg5  string
	Pimg6  string
}

//TODO
func AddProcut(username, password, email string) (err error) {
	uid := time.Now().Format("0101150405")
	beego.Info(uid)
	role := "0"
	tx, err := mapper.GetTx()
	if err != nil {
		return
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewProductMapper(tx)
	_, err = mappers.CreateProduct(uid, username, role, email, password)
	return
}

func DelProduct(pid string) error {
	tx, err := mapper.GetTx()
	if err != nil {
		return err
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewProductMapper(tx)
	return mappers.DeleteProductByPid(pid)
}

//TODO
func UpdateDetails(pid, pnmae, oprice, onhand string) error {
	tx, err := mapper.GetTx()
	if err != nil {
		return err
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewProductMapper(tx)
	return mappers.UpdateDetails(pid, pnmae, oprice, onhand)
}

//TODO
// func DelProduct(username string) error {
// 	tx, err := mapper.GetTx()
// 	if err != nil {
// 		return err
// 	}
// 	defer mapper.CloseTx(tx)
// 	mappers := mapper.NewProductMapper(tx)
// 	return mappers.DeleteProductByUserName(username)
// }

func FindProductByPid(pid string) (info ProductInfo, err error) {
	tx, err := mapper.GetTx()
	if err != nil {
		return
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewProductMapper(tx)
	row, err := mappers.FindProductByPid(pid)
	if err != nil {
		return
	}
	info = ProductInfo{Pid: row.Pid, Pname: row.Pname, Oprice: row.Oprice, Pprice: row.Pprice, Onhand: row.Onhand, Pimg1: row.Pimg1, Pimg2: row.Pimg2, Pimg3: row.Pimg3, Pimg4: row.Pimg4, Pimg5: row.Pimg5, Pimg6: row.Pimg6}
	return info, nil
}

func GetAllProduct() (num int, productList []ProductInfo, err error) {
	tx, err := mapper.GetTx()
	if err != nil {
		return
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewProductMapper(tx)
	rows, err := mappers.FindProductAll()
	if err != nil {
		return
	}
	i := 0
	for _, row := range rows {
		i++
		info := ProductInfo{}
		info.Pid = row.Pid
		info.Pname = row.Pname
		info.Oprice = row.Oprice
		info.Pprice = row.Pprice
		info.Onhand = row.Onhand
		info.Pimg1 = row.Pimg1
		info.Pimg2 = row.Pimg2
		info.Pimg3 = row.Pimg3
		info.Pimg4 = row.Pimg4
		info.Pimg5 = row.Pimg5
		info.Pimg6 = row.Pimg6
		productList = append(productList, info)
	}
	return i, productList, nil
}

func GetProductSplit() (productList1, productList2, productList3 []ProductInfo, err error) {
	tx, err := mapper.GetTx()
	if err != nil {
		return
	}
	defer mapper.CloseTx(tx)
	mappers := mapper.NewProductMapper(tx)
	rows, err := mappers.FindProductAll()
	if err != nil {
		return
	}
	i := 0
	for _, row := range rows {
		i++
		info := ProductInfo{}
		info.Pid = row.Pid
		info.Pname = row.Pname
		info.Oprice = row.Oprice
		info.Pprice = row.Pprice
		info.Onhand = row.Onhand
		info.Pimg1 = row.Pimg1
		info.Pimg2 = row.Pimg2
		info.Pimg3 = row.Pimg3
		info.Pimg4 = row.Pimg4
		info.Pimg5 = row.Pimg5
		info.Pimg6 = row.Pimg6
		productList1 = append(productList1, info)
		if i > 10 {
			productList2 = append(productList2, info)
		}
		if i > 14 {
			productList3 = append(productList3, info)
		}
	}
	return productList1, productList2, productList3, nil
}
