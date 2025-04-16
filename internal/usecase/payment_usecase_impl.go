package usecase

import (
	"payment-options/internal/models"
	"payment-options/internal/repository"
	"sync"
)

type paymentUsecase struct {
	repo repository.PaymentRepository
}

func NewPaymentUsecase(r repository.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{repo: r}
}

func (u *paymentUsecase) GetPaymentOptions() (map[string]models.PaymentMethod, error) {
    result := make(map[string]models.PaymentMethod)

    result["ovo"] = u.repo.CallOVO()
    result["dana"] = u.repo.CallDANA()
    result["gopay"] = u.repo.CallGoPay()
    result["shopee"] = u.repo.CallShopee()
    result["oneklik"] = u.repo.CallOneKlik()
    result["bridd"] = u.repo.CallBRIDD()
    result["linkaja"] = u.repo.CallLinkAja()

    return result, nil
}


