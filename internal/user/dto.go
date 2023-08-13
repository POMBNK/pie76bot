package user

type ToCreateDTO struct {
	TelegramID string
}

func CreateUserDTO(dto ToCreateDTO) User {
	return User{
		TelegramID: dto.TelegramID,
	}
}
