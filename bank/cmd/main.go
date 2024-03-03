package main

import (
	"bank"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var accounts = map[string]*bank.Account{}

func main() {
	accounts["1111-2222-3333-4444"] = &bank.Account{
		Customer: bank.Customer{
			Name: "John Client",
		},
		Number:     100,
		CardNumber: "1111-2222-3333-4444",
		Balance:    100,
	}

	accounts["1234-5678-1234-5678"] = &bank.Account{
		Customer: bank.Customer{
			Name: "John Merchant",
		},
		Number:     100,
		CardNumber: "1234-5678-1234-5678",
		Balance:    100,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the bank!")
	})
	http.HandleFunc("/transfer", transferHandler)
	http.HandleFunc("/deposit", depositHandler)
	http.HandleFunc("/withdraw", withdrawHandler)
	fmt.Println("Server running on 0.0.0.0:3000")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}

func depositHandler(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("accountNumber")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		bank.JsonResponse(w, http.StatusBadRequest, "Invalid amount number!")
	} else {
		account, ok := accounts[numberqs]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", numberqs)
		} else {
			err := account.Deposit(amount)
			if err != nil {
				println(err.Error())
				bank.JsonResponse(w, http.StatusBadRequest, err.Error())
			}
			bank.JsonResponse(w, http.StatusOK, "Deposito exitoso", account.Balance)
		}
	}
}

func withdrawHandler(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("accountNumber")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[numberqs]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", numberqs)
		} else {
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				bank.JsonResponse(w, http.StatusOK, "Retiro exitoso", account.Balance)
			}
		}
	}
}

func transferHandler(w http.ResponseWriter, req *http.Request) {
	originAccount := req.URL.Query().Get("originAccountNumber")
	amountFromRequest := req.URL.Query().Get("amount")
	destinationAccount := req.URL.Query().Get("destinationAccountNumber")

	if originAccount == "" {
		//fmt.Fprintf(w, "Account number is missing!")
		bank.JsonResponse(w, http.StatusBadRequest, "originAccount number is missing!")
		return
	}

	if destinationAccount == "" {
		//fmt.Fprintf(w, "Account number is missing!")
		bank.JsonResponse(w, http.StatusBadRequest, "destinationAccount number is missing!")
		return
	}

	if amount, err := strconv.ParseFloat(amountFromRequest, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		if origin, ok := accounts[originAccount]; !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", originAccount)
		} else if destination, ok := accounts[destinationAccount]; !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", destinationAccount)
		} else {
			err := origin.Transfer(amount, destination)
			if err != nil {
				fmt.Println("Error en la transferencia: ", err.Error())
				bank.JsonResponse(w, http.StatusBadRequest, err.Error())
			} else {
				fmt.Println("Transferencia exitosa de ", amount, " de ", originAccount, " a ", destinationAccount)
				bank.JsonResponse(w, http.StatusOK, "Transferencia exitosa")
			}
		}
	}
}
