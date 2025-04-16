package price_calculator

type StationProvider interface {
	GetStations(trainId string) ([]string, error)
}

type PricingProvider interface {
	GetPricePerStop(class TicketClass) (int, error)
}
