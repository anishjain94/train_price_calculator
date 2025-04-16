package price_calculator

type TicketClass string

const (
	FirstClass   TicketClass = "FirstClass"
	SecondClass  TicketClass = "SecondClass"
	ThirdClass   TicketClass = "ThirdClass"
	GeneralClass TicketClass = "GeneralClass"
)

type PricingConfigType map[TicketClass]int

var DefaultPricing = PricingConfigType{
	FirstClass:   100,
	SecondClass:  80,
	ThirdClass:   50,
	GeneralClass: 10}
