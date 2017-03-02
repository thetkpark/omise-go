package omise

/*
!!! DO NOT MODIFY !!!

autogenerated
 src: gen_list_job.tmpl
 job: &main.GenListJob{Names:[]string{"Account", "Balance", "BankAccount", "Card", "Charge", "Customer", "Deletion", "Dispute", "Document", "Event", "Link", "Recipient", "Refund", "Token", "Transaction", "Transfer"}}
  on: Thu Mar 02 13:31:35 +0700 2017
  by: chakrit
*/

import (
	"bytes"
)

// AccountList represents the list structure returned by Omise's REST API that contains
// Account struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type AccountList struct {
	List
	Data []*Account `json:"data"`
}

// Find finds and returns Account with the given id. Returns nil if not found.
func (list *AccountList) Find(id string) *Account {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func (list *AccountList) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}

// BalanceList represents the list structure returned by Omise's REST API that contains
// Balance struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type BalanceList struct {
	List
	Data []*Balance `json:"data"`
}

// Find finds and returns Balance with the given id. Returns nil if not found.
func (list *BalanceList) Find(id string) *Balance {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func (list *BalanceList) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}

// BankAccountList represents the list structure returned by Omise's REST API that contains
// BankAccount struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type BankAccountList struct {
	List
	Data []*BankAccount `json:"data"`
}

// Find finds and returns BankAccount with the given id. Returns nil if not found.
func (list *BankAccountList) Find(id string) *BankAccount {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func (list *BankAccountList) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}

// CardList represents the list structure returned by Omise's REST API that contains
// Card struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type CardList struct {
	List
	Data []*Card `json:"data"`
}

// Find finds and returns Card with the given id. Returns nil if not found.
func (list *CardList) Find(id string) *Card {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func (list *CardList) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}

// ChargeList represents the list structure returned by Omise's REST API that contains
// Charge struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type ChargeList struct {
	List
	Data []*Charge `json:"data"`
}

// Find finds and returns Charge with the given id. Returns nil if not found.
func (list *ChargeList) Find(id string) *Charge {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func (list *ChargeList) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}

// CustomerList represents the list structure returned by Omise's REST API that contains
// Customer struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type CustomerList struct {
	List
	Data []*Customer `json:"data"`
}

// Find finds and returns Customer with the given id. Returns nil if not found.
func (list *CustomerList) Find(id string) *Customer {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func (list *CustomerList) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}

// DeletionList represents the list structure returned by Omise's REST API that contains
// Deletion struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type DeletionList struct {
	List
	Data []*Deletion `json:"data"`
}

// Find finds and returns Deletion with the given id. Returns nil if not found.
func (list *DeletionList) Find(id string) *Deletion {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func (list *DeletionList) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}

// DisputeList represents the list structure returned by Omise's REST API that contains
// Dispute struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type DisputeList struct {
	List
	Data []*Dispute `json:"data"`
}

// Find finds and returns Dispute with the given id. Returns nil if not found.
func (list *DisputeList) Find(id string) *Dispute {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func (list *DisputeList) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}

// DocumentList represents the list structure returned by Omise's REST API that contains
// Document struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type DocumentList struct {
	List
	Data []*Document `json:"data"`
}

// Find finds and returns Document with the given id. Returns nil if not found.
func (list *DocumentList) Find(id string) *Document {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func (list *DocumentList) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}

// EventList represents the list structure returned by Omise's REST API that contains
// Event struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type EventList struct {
	List
	Data []*Event `json:"data"`
}

// Find finds and returns Event with the given id. Returns nil if not found.
func (list *EventList) Find(id string) *Event {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func (list *EventList) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}

// LinkList represents the list structure returned by Omise's REST API that contains
// Link struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type LinkList struct {
	List
	Data []*Link `json:"data"`
}

// Find finds and returns Link with the given id. Returns nil if not found.
func (list *LinkList) Find(id string) *Link {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func (list *LinkList) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}

// RecipientList represents the list structure returned by Omise's REST API that contains
// Recipient struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type RecipientList struct {
	List
	Data []*Recipient `json:"data"`
}

// Find finds and returns Recipient with the given id. Returns nil if not found.
func (list *RecipientList) Find(id string) *Recipient {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func (list *RecipientList) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}

// RefundList represents the list structure returned by Omise's REST API that contains
// Refund struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type RefundList struct {
	List
	Data []*Refund `json:"data"`
}

// Find finds and returns Refund with the given id. Returns nil if not found.
func (list *RefundList) Find(id string) *Refund {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func (list *RefundList) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}

// TokenList represents the list structure returned by Omise's REST API that contains
// Token struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type TokenList struct {
	List
	Data []*Token `json:"data"`
}

// Find finds and returns Token with the given id. Returns nil if not found.
func (list *TokenList) Find(id string) *Token {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func (list *TokenList) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}

// TransactionList represents the list structure returned by Omise's REST API that contains
// Transaction struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type TransactionList struct {
	List
	Data []*Transaction `json:"data"`
}

// Find finds and returns Transaction with the given id. Returns nil if not found.
func (list *TransactionList) Find(id string) *Transaction {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func (list *TransactionList) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}

// TransferList represents the list structure returned by Omise's REST API that contains
// Transfer struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type TransferList struct {
	List
	Data []*Transfer `json:"data"`
}

// Find finds and returns Transfer with the given id. Returns nil if not found.
func (list *TransferList) Find(id string) *Transfer {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func (list *TransferList) String() string {
	buf := &bytes.Buffer{}
	for _, item := range list.Data {
		buf.WriteString(item.String() + "\n")
	}
	return buf.String()
}
