package main 

import (
	"context" 
	"log"

	"github.com/Rhymond/go-money" 
	"github.com/google/uuid"

	coffeeco "github.com/irononet/coffeeco/internal" 
	"github.com/irononet/coffeeco/internal/payment" 
	"github.com/irononet/coffeeco/internal/purchase" 
	"github.com/irononet/coffeeco/internal/store"
)

func main(){
	
	ctx := context.Background() 

	// This is the test key from stripes documentation 
	stripeTestKey := "sk_test_4eC39HQlYJwDarjtT1zdp7dc" 

	cardToken := "tok_visa" 

	mongoConString := "mongodb://root:example@localhost:27017" 
	csvc, err := payment.NewStripeService(stripeTestKey) 
	if err != nil{
		log.Fatal(err)
	}

	prepo, err := purchase.NewMongoRepo(ctx, mongoConString)
	if err != nil{
		log.Fatal(err)
	}
	if err := prepo.Ping(ctx); err != nil{
		log.Fatal(err)
	}

	sRepo, err := store.NewMongoRepo(ctx, mongoConString) 
	if err != nil{
		log.Fatal(err)
	}
	if err := sRepo.Ping(ctx); err != nil{
		log.Fatal(err)
	}

	sSvc := store.NewService(sRepo) 

	svc := purchase.NewService(csvc, prepo, sSvc)

	storeID := uuid.New() 

	pur := &purchase.Purchase{
		CardToken: &cardToken,
		Store: store.Store{
			ID: storeID, 
		}, 
		ProductsToPurchase: []coffeeco.Product{{
			ItemName: "item1", 
			BasePrice: *money.New(3300, "USD"),
		}},
		PaymentMeans: payment.MEANS_CARD,
	}
	if err := svc.CompletePurchase(ctx, storeID, pur, nil); err != nil{
		log.Fatal(err)
	}

	log.Println("purchase was successful")
}