package store 

import (
	"github.com/google/uuid" 
	coffeeco "coffeeco/internal"
) 

type Store struct{
	ID uuid.UUID 
	Location string 
	ProductForSale []coffeeco.Product
}