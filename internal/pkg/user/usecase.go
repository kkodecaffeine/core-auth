package user

// UseCase interface definition
type UseCase interface {
	GetUserBy(ID *int64) (User, error)
}

type usecase struct {
	repo Repository
}

func (u *usecase) GetUserBy(ID *int64) (User, error) {
	return User{}, nil
}

// NewUseCase returns new UseCase implementation
func NewUseCase(userRepo Repository) UseCase {
	return &usecase{repo: userRepo}
}

var _ UseCase = &usecase{}
