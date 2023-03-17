package purchase

import (
	"time"

	money "github.com/Rhymond/go-money"
	"github.com/google/uuid"

	coffeeco "github.com/irononet/coffeeco/internal"
	"github.com/irononet/coffeeco/internal/store"
)


type Purchase struct{
	id uuid.UUID 
	Store store.Store 
	ProductsToPurchase []coffeeco.Product 
	total 	money.Money 
	PaymentMeans payment.Means 
	timeOfPurchase time.Time 
	CardToken *string 
}