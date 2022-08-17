package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionRepository interface {
	SaveTransaction(transaction Transaction, creditCard CreditCard) error
	GetCreditCard(creditCard CreditCard) (CreditCard, error)
	CreateCreditCard(creditCard CreditCard) error
}

type CreditCard struct {
	ID string
	Name string
	Number string
	ExpirationMonth int32
	ExpirationYear int32
	CVV int32
	Balance float64
	Limit float64
	CreatedAt time.Time
}

func NewCreditCard() *CreditCard {
	c := &CreditCard{}
	c.ID = uuid.NewV4().String()
	c.CreatedAt = time.Now()
	return c
}

func (t *Transaction) ProcessAndValidate(CreditCard *CreditCard) {
	if t.Amount + CreditCard.Balance > CreditCard.Limit {
		t.Status = "rejected"
	} else {
		t.Status = "approved"
		CreditCard.Balance = CreditCard.Balance + t.Amount
	}
}