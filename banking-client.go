package main

import (
	"ethos/altEthos"
	"ethos/syscall"
	"ethos/myRpc"
	"fmt"
	"log"
)

func init() {

	myRpc.SetupMyRpcIncrementReply(incrementReply)
	myRpc.SetupMyRpcGetBalanceReply(getBalanceReply)
	myRpc.SetupMyRpcMakeDepositReply(getMakeDepositReply)
	myRpc.SetupMyRpcWithDrawCashReply(getWithdrawCashReply)
	myRpc.SetupMyRpcTransferMoneyReply(getTransferReply)


}

func incrementReply(count uint64) (myRpc.MyRpcProcedure) {
	fmt.Printf("Received Increment Reply: %v\n", count)
	log.Printf("Received Increment Reply: %v\n", count)
	return nil
}

func getBalanceReply(amount uint64, err string)(myRpc.MyRpcProcedure) {
	if err != "ErrorNone" {
		log.Printf("There was some problem getting Balance for %s and the problem was %s", altEthos.GetUser(), err)	
	}
	fmt.Printf("Received balance Reply: amount = %d , error = %s\n",amount, err)
	log.Printf("Received balance Reply: amount = %d , error = %s\n",amount, err)
	return nil
}

func getMakeDepositReply(err string)(myRpc.MyRpcProcedure) {
	if err != "ErrorNone" {
		log.Printf("There was some problem making deposit for %s and the problem was %s", altEthos.GetUser(), err)
		return nil	
	}
	fmt.Printf("Received make deposit Reply: %s\n", err)
	log.Printf("Received make deposit Reply: %s\n", err)
	return nil
}

func getWithdrawCashReply(err string)(myRpc.MyRpcProcedure) {
	if err != "ErrorNone" {
		log.Printf("There was some problem withdrawing cash for %s nd the problem was %s ", altEthos.GetUser(), err)	
		return nil
	}
	fmt.Printf("Received withdraw cash Reply: %s\n", err)
	log.Printf("Received withdraw cash Reply: %s\n", err)
	return nil
}

func getTransferReply(err string)(myRpc.MyRpcProcedure) {
	if err != "ErrorNone" {
		log.Printf("There was some problem transfering cash from account name ->  %s and the problem was %s", altEthos.GetUser(), err)
		return nil
	}
	fmt.Printf("Received transfer Reply: %s\n", err)
	log.Printf("Received transfer Reply: %s\n", err)
	return nil
}


func main () {

	altEthos.LogToDirectory("test/syncClient")
	
	log.Println("before call123")

	log.Printf("User calling is -> %s", altEthos.GetUser())
	userName := altEthos.GetUser()
	

	fd, status := altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status)
		altEthos.Exit(status)
	}
/*
	call := myRpc.MyRpcIncrement{}
	status = altEthos.ClientCall(fd, &call)
	if status != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", status)
		altEthos.Exit(status)
	}
*/
	call1 := myRpc.MyRpcGetBalance{userName}
	status = altEthos.ClientCall(fd, &call1)
	
	if status != syscall.StatusOk {
		log.Printf("clientCall get balance failed: %v\n", status)
		altEthos.Exit(status)
	} 

	fd, status = altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status)
		altEthos.Exit(status)
	}

	call2 := myRpc.MyRpcMakeDeposit{userName, 200}
	status = altEthos.ClientCall(fd, &call2)
	if status != syscall.StatusOk {
		log.Printf("clientCall make deposit failed: %v\n", status)
		altEthos.Exit(status)
	}

	fd, status = altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status)
		altEthos.Exit(status)
	}

	call3 := myRpc.MyRpcWithDrawCash{userName, 191}
	status = altEthos.ClientCall(fd, &call3)
	if status != syscall.StatusOk {
		log.Printf("clientCall withdraw cashfailed: %v\n", status)
		altEthos.Exit(status)
	}

	fd, status = altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status)
		altEthos.Exit(status)
	}

	call4 := myRpc.MyRpcTransferMoney{userName, "yl", 450}
	status = altEthos.ClientCall(fd, &call4)
	if status != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", status)
		altEthos.Exit(status)
	}

	log.Println("Client Program Exited")
}

