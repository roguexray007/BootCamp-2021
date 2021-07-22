package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)


type Account struct{
	currentBalance int
	mutex sync.Mutex
}
func (account *Account) deposit(depositAmount int,waitGroup *sync.WaitGroup){
	defer waitGroup.Done()
	account.mutex.Lock()
	account.currentBalance += depositAmount
	fmt.Println("$",depositAmount,"credited in the account. Current Balance is: $",account.currentBalance)
	account.mutex.Unlock()
}
func (account *Account) withdraw(withdrawalNo int,withdrawalAmount int, waitGroup *sync.WaitGroup,errorChannel chan<- error){
	defer waitGroup.Done()
	account.mutex.Lock()
	if withdrawalAmount > account.currentBalance{
		errorChannel <- errors.New(strconv.Itoa(withdrawalNo) + "failed. Insufficient Balance. Maximum Withdrawal Anount: $"+strconv.Itoa(account.currentBalance))
	} else {
		account.currentBalance -= withdrawalAmount
		fmt.Println("$",withdrawalAmount,"was debited. Current Balance is: $",account.currentBalance)
	}
	account.mutex.Unlock()
}
func errorMessage(errors chan error,worker chan bool){
	for err := range errors{
		fmt.Println(err)
	}
	worker <- true
}
func main(){
	errorChannel := make(chan error,5)
	var waitGroup sync.WaitGroup
	accountDetails := Account{}
	fmt.Println("Enter the Starting Balance for the account: ")
	fmt.Scanln(&accountDetails.currentBalance)
	var DepositWithdrawalRange int
	fmt.Println("Enter the maximum amount that we can deposit and withdrawal:")
	fmt.Scanln(&DepositWithdrawalRange)
	//we are going for 10 deposits and 10 withdrawals
	for depositId := 1; depositId <= 10; depositId++{
		waitGroup.Add(1)
		go accountDetails.deposit(rand.Intn(DepositWithdrawalRange),&waitGroup)
	}
	for withdrawalNo := 1; withdrawalNo <= 10; withdrawalNo++{
		waitGroup.Add(1)
		go accountDetails.withdraw(withdrawalNo,rand.Intn(DepositWithdrawalRange),&waitGroup,errorChannel)
	}
	WorkerChannel := make(chan bool)
	go errorMessage(errorChannel,WorkerChannel)
	waitGroup.Wait()
	fmt.Println("Current Account Balance is: $",accountDetails.currentBalance)
	close(errorChannel)
	<-WorkerChannel

}
