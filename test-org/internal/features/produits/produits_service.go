package modules

type ProduitsService interface {
	DoBusinessLogic() string
}

type produitsService struct {
	// Add repository here if needed
}

func NewProduitsService() ProduitsService {
	return &produitsService{}
}

func (s *produitsService) DoBusinessLogic() string {
	return "Hello from Produits Service!"
}
