package mapper

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestAccountMapper_CreateAccount(t *testing.T) {
	tx, err := GetTx()
	if err != nil {
		panic(err)
	}
	type args struct {
		id          string
		domain      sql.NullString
		username    string
		displayName string
		email       string
		password    string
	}
	tests := []struct {
		name    string
		m       *AccountMapper
		args    args
		want    *AccountRow
		wantErr bool
	}{
		{"case#0", NewAccountMapper(tx),
			args{"001", sql.NullString{"local", true}, "test1", "Test Name 1", "test1@qq.com", "123"},
			&AccountRow{"001", sql.NullString{"local", true}, "test1", "Test Name 1", "test1@qq.com"},
			true,
		},
		{"case#1", NewAccountMapper(tx),
			args{"002", sql.NullString{"local", true}, "test2", "Test Name 2", "test2@qq.com", "123"},
			&AccountRow{"002", sql.NullString{"local", true}, "test2", "Test Name 2", "test2@qq.com"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.CreateAccount(tt.args.id, tt.args.domain, tt.args.username, tt.args.displayName, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountMapper.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountMapper.CreateAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountMapper_FindAccountByID(t *testing.T) {
	tx, err := GetTx()
	if err != nil {
		panic(err)
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		m       *AccountMapper
		args    args
		want    *AccountRow
		wantErr bool
	}{
		{"case#0", NewAccountMapper(tx),
			args{"001"},
			&AccountRow{"001", sql.NullString{"local", true}, "test1", "Test Name 1", "test1@qq.com"},
			false,
		},
		{"case#1", NewAccountMapper(tx),
			args{"002"},
			&AccountRow{"002", sql.NullString{"local", true}, "test2", "Test Name 2", "test2@qq.com"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.FindAccountByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountMapper.FindAccountByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountMapper.FindAccountByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountMapper_SqlCheckLoginInfo(t *testing.T) {
	tx, err := GetTx()
	if err != nil {
		panic(err)
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		m       *AccountMapper
		args    args
		wantErr bool
	}{
		{"case#0", NewAccountMapper(tx), args{"test1", "123"}, false},
		{"case#1", NewAccountMapper(tx), args{"test2", "123"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.SqlCheckLoginInfo(tt.args.username, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("AccountMapper.SqlCheckLoginInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountMapper_FindAccountAll(t *testing.T) {
	tx, err := GetTx()
	if err != nil {
		panic(err)
	}
	tests := []struct {
		name    string
		m       *AccountMapper
		want    []AccountRow
		wantErr bool
	}{
		{"case#0", NewAccountMapper(tx),
			[]AccountRow{{"001", sql.NullString{"local", true}, "test1", "Test Name 1", "test1@qq.com"},
				{"002", sql.NullString{"local", true}, "test2", "Test Name 2", "test2@qq.com"}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.FindAccountAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountMapper.FindAccountAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountMapper.FindAccountAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountMapper_FindAccountCount(t *testing.T) {
	tx, err := GetTx()
	if err != nil {
		panic(err)
	}
	tests := []struct {
		name    string
		m       *AccountMapper
		want    int
		wantErr bool
	}{
		{"case#0", NewAccountMapper(tx), 80, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.FindAccountCount()
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountMapper.FindAccountCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AccountMapper.FindAccountCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountMapper_FindAccountList(t *testing.T) {
	tx, err := GetTx()
	if err != nil {
		panic(err)
	}
	type args struct {
		startIndex int32
		count      int32
	}
	tests := []struct {
		name    string
		m       *AccountMapper
		args    args
		want    []AccountRow
		wantErr bool
	}{
		{"case#0", NewAccountMapper(tx),
			args{1, 2},
			[]AccountRow{{"001", sql.NullString{"local", true}, "test1", "Test Name 1", "test1@qq.com"},
				{"002", sql.NullString{"local", true}, "test2", "Test Name 2", "test2@qq.com"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.FindAccountList(tt.args.startIndex, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountMapper.FindAccountList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountMapper.FindAccountList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountMapper_DeleteAccountByID(t *testing.T) {
	tx, err := GetTx()
	if err != nil {
		panic(err)
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		m       *AccountMapper
		args    args
		wantErr bool
	}{
		{"case#0", NewAccountMapper(tx), args{"001"}, true},
		{"case#1", NewAccountMapper(tx), args{"002"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.DeleteAccountByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AccountMapper.DeleteAccountByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
