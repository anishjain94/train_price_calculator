package price_calculator

import "testing"

func TestAddTrainConfig(t *testing.T) {
	trainConfig := NewTrainConfig()

	trainId := "101"
	stations := []string{"StationA", "StationB", "StationC"}

	// Step 3: Add train configuration
	trainConfig.AddTrainConfig(trainId, stations)

	// Step 4: Verify the train configuration was added correctly
	if len(trainConfig.config) != 1 {
		t.Errorf("Expected TrainConfig to have 1 entry, got %d", len(trainConfig.config))
	}

	if _, exists := trainConfig.config[trainId]; !exists {
		t.Errorf("Expected TrainConfig to contain train ID %s", trainId)
	}

	if !equalSlices(trainConfig.config[trainId].Stations, stations) {
		t.Errorf("Expected stations %v for train ID %s, got %v", stations, trainId, trainConfig.config[trainId])
	}
}

func TestCalculatePrice(t *testing.T) {
	tests := []struct {
		name           string
		trainId        string
		stations       []string
		source         string
		destination    string
		noOfPassengers int
		expectedPrice  int
		expectedError  string
	}{
		{
			name:           "Valid route with 2 stations in between",
			trainId:        "101",
			stations:       []string{"StationA", "StationB", "StationC", "StationD"},
			source:         "StationA",
			destination:    "StationD",
			noOfPassengers: 2,
			expectedPrice:  60, // 3 stations * 10 * 2 passengers
			expectedError:  "",
		},
		{
			name:           "Source station not found",
			trainId:        "101",
			stations:       []string{"StationA", "StationB", "StationC"},
			source:         "StationX",
			destination:    "StationC",
			noOfPassengers: 1,
			expectedPrice:  0,
			expectedError:  "source station StationX not found",
		},
		{
			name:           "Destination station not found",
			trainId:        "101",
			stations:       []string{"StationA", "StationB", "StationC"},
			source:         "StationA",
			destination:    "StationY",
			noOfPassengers: 1,
			expectedPrice:  0,
			expectedError:  "destination station StationY not found",
		},
		{
			name:           "Source station is after destination station",
			trainId:        "101",
			stations:       []string{"StationA", "StationB", "StationC"},
			source:         "StationC",
			destination:    "StationA",
			noOfPassengers: 1,
			expectedPrice:  0,
			expectedError:  "source station StationC is after destination station StationA",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup train configuration
			trainConfig := NewTrainConfig()
			trainConfig.AddTrainConfig(tt.trainId, tt.stations)

			// Setup pricing strategy
			priceCalculator := &PriceCalculator{
				strategy: NewFixedPerStationStrategy(&trainConfig, &trainConfig),
			}

			price, err := priceCalculator.CalculatePrice(tt.trainId, tt.source, tt.destination, []PassengerType{
				{
					Class:          GeneralClass,
					NoOfPassengers: tt.noOfPassengers,
				},
			})

			// Check price
			if price != tt.expectedPrice {
				t.Errorf("Expected price %d, got %d", tt.expectedPrice, price)
			}

			// Check error
			if err != nil && err.Error() != tt.expectedError {
				t.Errorf("Expected error '%s', got '%s'", tt.expectedError, err.Error())
			} else if err == nil && tt.expectedError != "" {
				t.Errorf("Expected error '%s', got nil", tt.expectedError)
			}
		})
	}
}

// Helper function to compare two slices
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
