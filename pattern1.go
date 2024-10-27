func getAnimal(animalType string) (Animal, error) {
	switch strings.ToLower(animalType) {
	case "monkey":
		return Monkey{}, nil
	case "shark":
		return Shark{}, nil
	case "eagle":
		return Eagle{}, nil
	case "bear":
		return Bear{}, nil
	case "whale":
		return Whale{}, nil
	default:
		return UnknownAnimal{}, errors.New("unknown animal type")
	}
}