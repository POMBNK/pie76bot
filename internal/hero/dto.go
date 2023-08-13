package hero

type ToCreateDTO struct {
	Name string
	Luck int
}

func CreateHeroDTO(dto ToCreateDTO) Hero {
	return Hero{
		Name: dto.Name,
		Luck: dto.Luck,
	}
}
