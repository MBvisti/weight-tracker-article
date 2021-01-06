package api

// WeightService contains the methods of the user service
type WeightService interface {
}

// WeightRepository is what lets our service do db operations without knowing anything about the implementation
type WeightRepository interface {
}

// weightService contains the actual implementations of the methods defined in WeightService
type weightService struct {
	storage WeightRepository
}

func NewWeightService(weightRepo WeightRepository) WeightService {
	return &weightService{storage: weightRepo}
}
