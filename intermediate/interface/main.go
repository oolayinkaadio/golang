/**
 * An interface is like a remote control — it doesn’t matter what brand your TV is, as long as it can respond to “power”, “volume up”, and “volume down” buttons!

Imagine you’re at a hotel:

You have a universal power socket that accepts any plug
You don’t care what brand of device you’re plugging in
As long as the plug fits, it will work!
This is exactly how interfaces work in Go:

The socket is your interface
The different devices are your types (structs)
The plug is the method that makes them compatible
Let’s see some real examples with Universal Remote Control 📺:
*/

package main

import (
	"fmt"
)

// INTERFACE 1 TYPE::::::::
	type RemoteControl interface {
		PowerOn() string
		PowerOff() string
 	}

	type TV struct {
		Brand string
	}

	type Radio struct {
		Brand string
	}

		//  making TV work with the remote control:
	func (t TV) PowerOn() string {
		return fmt.Sprintf("%s TV is turning ON", t.Brand)
	}

	func (t TV) PowerOff() string {
		return fmt.Sprintf("%s TV is turning OFF", t.Brand)
	}
		//  making Radio work with the remote control:
	func (r Radio) PowerOn() string {
		return fmt.Sprintf("%s Radio is turning ON", r.Brand)
	}

	func (r Radio) PowerOff() string {
		return fmt.Sprintf("%s Radio is turning OFF", r.Brand)
	}

		// Function that can control any device:
	func ControlDevice(device RemoteControl) {
		fmt.Println((device.PowerOn()))
		fmt.Println((device.PowerOff()))
	}


func execRemoteControl() {
		// Create the devices:
	samsungTv := TV{Brand: "samsung"}
	sonyRadio := Radio{Brand: "sony"}

	// use same remote for different devices:
	fmt.Println("Controlling TV:")
	ControlDevice(samsungTv)

	fmt.Println("Controlling Ravio:")
	ControlDevice(sonyRadio)
	
}
// ----------------------------------------------------------------
// INTERFACE 2 TYPE::::::::
// Interface: defines what makes something a payment method
type PaymentMethod interface {
	ProcessPayment(amount float64) string
	Getbalance() float64
}

// CreditCard type
type CreditCard struct {
	Balance float64
	CardNumber string
	}

	// DebitCard type
	type DebitCard struct {
		Balance float64
		BankName string
	}

	// Making CreditCard work as a PaymentMethod
	func (c *CreditCard) ProcessPayment(amount float64) string {
		if c.Balance >= amount {
			c.Balance -= amount
			return fmt.Sprintf("Paid $%.2f using Credit Card", amount)
		}
		return "Insufficient credit!"
	}

	func (c *CreditCard) Getbalance() float64 {
		return c.Balance
	}

	// Making DebitCard work as a PaymentMethod
	func (d *DebitCard) ProcessPayment(amount float64) string {
		if d.Balance >= amount {
			d.Balance -= amount
			return fmt.Sprintf("Paid $%.2f using Debit Card", amount)
		}
		return "Insufficient funds!"
	}

	func (d *DebitCard) Getbalance() float64 {
		return d.Balance
	}

	// Store function that accepts any payment method
	func ProcessStorePayment(payment PaymentMethod, amount float64) {
		fmt.Println("Processing payment...")
		fmt.Println(payment.ProcessPayment(amount))
		fmt.Printf("Remaining balance: $%.2f\n", payment.Getbalance())
	}


	func payment() {
		// Create our payment methods
		myCredit := &CreditCard{
			Balance: 1000,
			CardNumber: "1234-5678",
		}

		myDebit := &DebitCard{
			Balance: 500,
			BankName: "MyBank",
		}

		// Buy something for $100
		fmt.Println("Buying with Credit Card:")
		ProcessStorePayment(myCredit, 100)

		// Buy something for $100
		fmt.Println("\nBuying with Debit Card:")
		ProcessStorePayment(myDebit, 100)
	}

func main() {
	execRemoteControl()
	fmt.Println("\n\n\n PAYMENT:::::")
	payment()
}