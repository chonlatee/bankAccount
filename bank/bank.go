package bank

import (
	"strconv"
	"sync"
)

type currency struct {
	sync.Mutex
	amount float64
	code   string
}

// Balance export
var Balance = &currency{amount: 0, code: "TH"}

func (c *currency) Deposit(amount float64) {
	c.Lock()
	c.amount += amount
	c.Unlock()
}

func (c *currency) Withdraw(amount float64) {
	c.Lock()
	c.amount -= amount
	c.Unlock()
}

func (c *currency) Display() string {
	c.Lock()
	defer c.Unlock()
	return strconv.FormatFloat(c.amount, 'f', 2, 64) + " " + c.code
}
