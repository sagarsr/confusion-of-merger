package models

import (
	"context"

	"github.com/jackc/pgx/v4"
)

//Bank table contains bank details
type Bank struct {
	BankName string `json:"bank_name"`
	BankID   int64
}

//Branch table contains details of Branch
type Branch struct {
	Ifsc        string `json:"ifsc"`
	BankDetails Bank
	Branch      string `json:"branch"`
	Address     string `json:"address"`
	City        string `json:"city"`
	District    string `json:"district"`
	State       string `json:"state"`
}

//FetchBankDetailsWithIFSC is to obtain bank details with IFSC code
func (b *Branch) FetchBankDetailsWithIFSC(conn *pgx.Conn, ifsc string) (err error) {
	err = conn.QueryRow(context.Background(), "SELECT ifsc, name, branch, address, city, district, state FROM branches INNER JOIN banks ON branches.bank_id = banks.id where ifsc = $1", ifsc).Scan(
		&b.Ifsc, &b.BankDetails.BankName, &b.Branch, &b.Address, &b.City, &b.District, &b.State)
	return
}

//FetchBranches lists all details of the banks when city and the bank name is given
func (b *Branch) FetchBranches(conn *pgx.Conn, bankname string, city string, limit int, offset int) (BranchList []Branch, count int, err error) {
	err = conn.QueryRow(context.Background(), "SELECT COUNT(*) FROM branches INNER JOIN banks ON branches.bank_id = banks.id where banks.name=$1 and city=$2;", bankname, city).Scan(&count)
	rows, err := conn.Query(context.Background(), "SELECT ifsc, name, branch, address, city, district, state from branches INNER JOIN banks ON branches.bank_id = banks.id where banks.name=$1 and city=$2 LIMIT $3 OFFSET $4;", bankname, city, limit, offset)
	for rows.Next() {
		err = rows.Scan(&b.Ifsc, &b.BankDetails.BankName, &b.Branch, &b.Address, &b.City, &b.District, &b.State)
		BranchList = append(BranchList, *b)
	}
	return
}
