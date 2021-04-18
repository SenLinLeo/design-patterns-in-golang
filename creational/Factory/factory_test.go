package Factory

import "testing"

func TestNewRestaurant(t *testing.T) {
	NewRestaurant("q").GetFood()
	NewRestaurant("q").GetFood()
}
