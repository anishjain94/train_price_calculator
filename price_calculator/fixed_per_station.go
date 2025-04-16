package price_calculator

import "fmt"

type FixedPerStationStrategy struct {
	stationProvider StationProvider
	pricingProvider PricingProvider
}

func NewFixedPerStationStrategy(sp StationProvider, pp PricingProvider) *FixedPerStationStrategy {
	return &FixedPerStationStrategy{
		stationProvider: sp,
		pricingProvider: pp,
	}
}

func (f *FixedPerStationStrategy) CalculatePrice(trainId, source, destination string, passengersInfo []PassengerType) (int, error) {
	stations, err := f.stationProvider.GetStations(trainId)
	if err != nil {
		return 0, fmt.Errorf("train ID %s not found", trainId)
	}

	sourceIndex := -1
	destinationIndex := -1
	for i, station := range stations {
		if station == source {
			sourceIndex = i
		}
		if station == destination {
			destinationIndex = i
		}
	}

	if sourceIndex == -1 {
		return 0, fmt.Errorf("source station %s not found", source)
	}
	if destinationIndex == -1 {
		return 0, fmt.Errorf("destination station %s not found", destination)
	}
	if sourceIndex >= destinationIndex {
		return 0, fmt.Errorf("source station %s is after destination station %s", source, destination)
	}

	stationsBetween := destinationIndex - sourceIndex
	totalPrice := 0
	for _, passenger := range passengersInfo {
		classPrice, err := f.pricingProvider.GetPricePerStop(passenger.Class)

		if err != nil {
			return 0, fmt.Errorf("ticket class %s not found", passenger.Class)
		}

		totalPrice += stationsBetween * classPrice * passenger.NoOfPassengers

	}

	return totalPrice, nil
}
