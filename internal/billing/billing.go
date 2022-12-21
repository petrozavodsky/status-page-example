package billing

import (
	"status_page/internal/utils"
)

type BillingData struct {
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
}

// InstancBilling - Функция инициации экземпляра
func InstancBilling() *BillingData {
	res := utils.ReadFileCsv(utils.ConfigData.BillingDataPath)[0][0]

	var newBillingData BillingData

	newBillingData.CreateCustomer = utils.ByteToBool(res[0])
	newBillingData.Purchase = utils.ByteToBool(res[1])
	newBillingData.Payout = utils.ByteToBool(res[2])
	newBillingData.Recurring = utils.ByteToBool(res[3])
	newBillingData.FraudControl = utils.ByteToBool(res[4])
	newBillingData.CheckoutPage = utils.ByteToBool(res[5])

	return &newBillingData
}
