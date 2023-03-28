MyRpc interface {
	Increment() (count uint64)
	GetBalance(user string) (balance uint64, error string)
	MakeDeposit(user string, amount uint64) (error string)
	WithDrawCash(user string,amount uint64) (error string)
	TransferMoney(user1 string, user2 string,amount uint64) (error string)
}

