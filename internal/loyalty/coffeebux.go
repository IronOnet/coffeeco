package loyalty 

import (
	"github.com/google/uuid" 

	coffeeco "github.com/irononet/coffeeco/internal" 
	"github.com/irononet/coffeeco/internal/store"
)

type CoffeeBux struct{
	ID		uuid.UUID 
	store store.Store 
	coffeelover coffeeco.CoffeeLover 
	FreeDrinksAvailable int 
	RemainingDrinkPurchasesUntilFreeDrink int 
}

func (c *CoffeeBux) AddStamp(){
	if c.RemainingDrinkPurchasesUntilFreeDrink == 1{
		c.RemainingDrinkPurchasesUntilFreeDrink = 10 
		c.FreeDrinksAvailable += 1 
	} else{
		c.RemainingDrinkPurchasesUntilFreeDrink--
	}
}

func (c *CoffeeBux) Pay(ctx context.Context, purchases []coffeeco.Product) error{
	lp := len(purchases) 
	if lp == 0{
		return errors.New("nothing to buy")
	}

	if c.FreeDrinksAvailable < lp{
		return fmt.Errorf("not enough coffeebux to cover the entire purchase. Have %d, need %d", len(purchases), c.FreeDrinksAvailable) 
	}

	c.FreeDrinksAvailable = c.FreeDrinksAvailable - lp 
	return nil 
}

