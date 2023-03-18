package coffeeco

import (
	"errors"
	"time"

	money "github.com/Rhymond/go-money"
	"github.com/google/uuid"
) 

type Product struct{
	ItemName string 
	BasePrice money.Money  
}


