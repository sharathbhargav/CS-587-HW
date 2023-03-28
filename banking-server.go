package main

import (
	"ethos/altEthos"
	"ethos/myRpc"
	"ethos/syscall"
	"fmt"
	"log"
)

var myRpc_increment_counter uint64 = 0

var Bank map[string]uint64
// Make a map of users and their account balance

func init() {
	myRpc.SetupMyRpcIncrement(increment)
	myRpc.SetupMyRpcGetBalance(getBalance)
	myRpc.SetupMyRpcMakeDeposit(makeDeposit)
	myRpc.SetupMyRpcWithDrawCash(withDrawCash)
	myRpc.SetupMyRpcTransferMoney(transferMoney)
}

func transferMoney(user1 string, user2 string, amount uint64) (myRpc.MyRpcProcedure){
	fmt.Printf("Server : transfer money : %s, to %s\n", user1,user2)
	log.Printf("Server : transfer money : %s, to %s\n", user1,user2)
	
	if  _,ok := Bank[user1] ; ok {
		log.Printf("Amount inside %s is %d", user1, Bank[user1])
		if Bank[user1] < amount {
			return &myRpc.MyRpcTransferMoneyReply{user1+" has insufficient funds"}
		}

		Bank[user1] -= amount
		Bank[user2] += amount
		log.Printf("Amount transfered successful from  %s to %s is %d", user1, user2, amount)
		log.Printf("New Balance of %s is %d", user1, Bank[user1])
		log.Printf("New Balance of %s is %d", user2, Bank[user2])
		return &myRpc.MyRpcTransferMoneyReply{"ErrorNone"}
	}
	
	return &myRpc.MyRpcTransferMoneyReply{"User not found"}

	
}

func withDrawCash(user string,amount uint64) (myRpc.MyRpcProcedure){
	fmt.Printf("Server : withdraw cash  : %s of %d\n", user, amount)
	log.Printf("Server : withdraw cash : %s of %d\n", user, amount)

	if _,ok := Bank[user]; ok {
		log.Printf("Found Bank Balance for %s with bank balance %d for withdrawl", user, Bank[user])
		if amount > Bank[user] {
			return &myRpc.MyRpcWithDrawCashReply{user+" has insufficient funds to withdraw"}
		}
	
		Bank[user] -= amount
		log.Printf("Server : Successfully withdraw cash : %s of %d\n", user, amount)
		return &myRpc.MyRpcWithDrawCashReply{"ErrorNone"}
	}

	return &myRpc.MyRpcWithDrawCashReply{"User not found"}
	
}

func makeDeposit(user string,amount uint64) (myRpc.MyRpcProcedure){
	fmt.Printf("Server : make deposit cash  : %s, %d\n", user, amount)
	log.Printf("Server : make deposit cash : %s, %d\n", user, amount)

	if _,ok := Bank[user]; ok {
		log.Printf("Found Bank Balance for making a deposit %s", user)
		Bank[user] += amount
		log.Printf("Server : Successfully made deposit cash : %s, %d\n", user, amount)
		log.Printf("Server : New balance is -> %d", Bank[user])
		return &myRpc.MyRpcMakeDepositReply{"ErrorNone"}
	}

	return &myRpc.MyRpcMakeDepositReply{"User not found"}
	
}

func getBalance(user string) (myRpc.MyRpcProcedure){
	fmt.Printf("Server : balance cash  : %s\n", user)
	log.Printf("Server : balance cash : %s\n", user)
	if amt,ok := Bank[user]; ok {
		log.Printf("Server : Balance for user : %s is %d \n", user, Bank[user])
		s:= fmt.Sprintf("%s has balance %d",user,amt)

		return &myRpc.MyRpcGetBalanceReply{amt, s}
	}

	return &myRpc.MyRpcGetBalanceReply{0, "Error in getting Balance"}
	//return &myRpc.MyRpcTransferMoneyReply{"User not found"}
	//return &myRpc.MyRpcGetBalance{"User not found"}
}


func increment() (myRpc.MyRpcProcedure) {
	log.Println("called increment123")
	myRpc_increment_counter++
	return &myRpc.MyRpcIncrementReply{myRpc_increment_counter}
}


func main () {

	altEthos.LogToDirectory("test/syncService")

	Bank = make(map[string]uint64)
	
	Bank["me"] = 10000
	Bank["pat"] = 1000
	Bank["yl"] = 1000
	Bank["nobody"] = 1000

	listeningFd, status := altEthos.Advertise("myRpc")
	if status != syscall.StatusOk {
		log.Println("Advertising service failed: ", status)
		altEthos.Exit(status)
	}

	for {
		_, fd, status := altEthos.Import(listeningFd)
		if status != syscall.StatusOk {
			log.Printf("Error calling Import: %v\n", status)
			altEthos.Exit(status)
		}

		log.Println("new connection accepted123")

		t := myRpc.MyRpc{}
		altEthos.Handle(fd, &t)
	}
}