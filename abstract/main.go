package main

import "fmt"

type account struct {
	accountNo string
	pwd       string
	balance   float64
}

func NewAccount(accountNo string, pwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号的长度不对")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("密码长度不对")
		return nil
	}
	if balance < 20 {
		fmt.Println("余额金额不对")
		return nil
	}
	return &account{
		accountNo: accountNo,
		pwd:       pwd,
		balance:   balance,
	}
}

func (acc *account) Deposite(money float64, pwd string) {
	if pwd != acc.pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}
	acc.balance += money
	fmt.Println("存款成功")
}

func (acc *account) WithDraw(money float64, pwd string) bool {
	if pwd != acc.pwd {
		fmt.Println("你输入的密码不正确")
	}
	if money > acc.balance || money <= 0 {
		fmt.Println("你输入的金额不正确")
	}
	acc.balance -= money
	fmt.Println("取款成功")
	return true
}

func (acc *account) Query(pwd string) {
	if pwd != acc.pwd {
		fmt.Println("你输入的密码不正确")
	}
	fmt.Printf("你的账号为%v 余额为%v \n ", acc.accountNo, acc.balance)
}

func main() {
	account := &account{
		accountNo: "1111",
		balance:   7732.23,
		pwd:       "12323",
	}
	account.Query("323")
}
