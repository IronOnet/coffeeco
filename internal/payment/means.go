package payment 

type Means string 

const (
	MEANS_CARD = "CARD" 
	MEANS_CASH = "CASH" 
	MEANS_COFFEEBUX = "coffeebux" 
)

type CardDetails struct{
	cardToken string 
}