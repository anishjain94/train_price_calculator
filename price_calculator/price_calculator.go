package price_calculator

import "fmt"

type TrainDetails struct {
	Stations []string
}

type TrainConfigType struct {
	config        map[string]TrainDetails
	PricingConfig PricingConfigType
}

func NewTrainConfig() TrainConfigType {
	return TrainConfigType{
		config:        make(map[string]TrainDetails),
		PricingConfig: DefaultPricing,
	}
}

func (trainConfig *TrainConfigType) GetStations(trainId string) ([]string, error) {
	trainDetails, exists := trainConfig.config[trainId]
	if !exists {
		return nil, fmt.Errorf("train ID %s not found", trainId)
	}
	return trainDetails.Stations, nil
}

func (trainConfig *TrainConfigType) GetPricePerStop(class TicketClass) (int, error) {
	price, exists := DefaultPricing[class]
	if !exists {
		return 0, fmt.Errorf("ticket class %s not found", class)
	}
	return price, nil
}

func (trainConfig *TrainConfigType) AddTrainConfig(trainId string, stations []string) {
	trainConfig.config[trainId] = TrainDetails{Stations: stations}
}

type PassengerType struct {
	Class          TicketClass
	NoOfPassengers int
}

// Strategy Pattern interface
type PricingStrategy interface {
	CalculatePrice(trainId, source, destination string, passengers []PassengerType) (int, error)
}

// Concrete implementation of the PricingStrategy interface
type PriceCalculator struct {
	strategy PricingStrategy
}

func (p *PriceCalculator) CalculatePrice(trainId, source, destination string, passengers []PassengerType) (int, error) {
	return p.strategy.CalculatePrice(trainId, source, destination, passengers)
}
