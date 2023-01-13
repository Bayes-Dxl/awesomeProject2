package main

// Service /*
type Service struct {
	description    string
	durationMonths int
	monthlyFee     float64
	features       []string
}

/*  1
type Service struct {
	description    string
	durationMonths int
	monthlyFee     float64
}
*/

func (s Service) getName() string {
	return s.description
}

func (s Service) getCost(recur bool) float64 {
	if recur {
		return s.monthlyFee * float64(s.durationMonths)
	}
	return s.monthlyFee
}
