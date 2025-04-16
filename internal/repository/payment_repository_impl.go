package repository

import (
	"payment-options/internal/models"
	"time"
)

type paymentRepo struct{}

func NewPaymentRepo() PaymentRepository {
	return &paymentRepo{}
}

func (r *paymentRepo) CallOVO() models.PaymentMethod {
    time.Sleep(3 * time.Second)  // Add 2 second delay
    return models.PaymentMethod{
        Account: "6288xx",
        Status:  "Active",
        Balance: "10000",
        Icon:    "https://sampleurl.com/ovo.jpg",
    }
}

func (r *paymentRepo) CallDANA() models.PaymentMethod {
    time.Sleep(1 * time.Second)  // Add 1 second delay
    return models.PaymentMethod{
        Account: "6288xx",
        Status:  "Active",
        Balance: "10000",
        Icon:    "https://sampleurl.com/dana.jpg",
    }
}

func (r *paymentRepo) CallGoPay() models.PaymentMethod {
    time.Sleep(1500 * time.Millisecond)  // Add 1.5 second delay
    return models.PaymentMethod{
        Account: "6288xx",
        Status:  "Active",
        Balance: "10000",
        Icon:    "https://sampleurl.com/gopay.jpg",
    }
}

func (r *paymentRepo) CallShopee() models.PaymentMethod {
    time.Sleep(2 * time.Second)  // Add 2 second delay
    return models.PaymentMethod{
        Account: "6288xx",
        Status:  "Active",
        Balance: "10000",
        Icon:    "https://sampleurl.com/shopee.jpg",
    }
}

func (r *paymentRepo) CallOneKlik() models.PaymentMethod {
    time.Sleep(1 * time.Second)  // Add 1 second delay
    return models.PaymentMethod{
        Account: "6288xx",
        Status:  "Active",
        Balance: "10000",
        Icon:    "https://sampleurl.com/oneklik.jpg",
    }
}

func (r *paymentRepo) CallBRIDD() models.PaymentMethod {
    time.Sleep(1500 * time.Millisecond)  // Add 1.5 second delay
    return models.PaymentMethod{
        Account: "6288xx",
        Status:  "Active",
        Balance: "10000",
        Icon:    "https://sampleurl.com/bridd.jpg",
    }
}

func (r *paymentRepo) CallLinkAja() models.PaymentMethod {
    time.Sleep(2 * time.Second)  // Add 2 second delay
    return models.PaymentMethod{
        Account: "0812xx", 
        Status:  "Active", 
        Balance: "150000", 
        Icon:    "https://sampleurl.com/linkaja.jpg",
    }
}



