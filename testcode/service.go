package main

import (
    "github.com/pkg/errors"
    "github.com/umi0410/how-to-backend-in-go/testcode/ent"
)

var (
	ErrInvalidValue = errors.New("올바르지 않은 값입니다")
	ErrInvalidUserID = errors.Wrap(ErrInvalidValue, "유저 ID가 조건을 만족하지 않습니다")
	ErrInvalidPassword = errors.Wrap(ErrInvalidValue, "비밀번호는 6자리 이상이어야합니다")
	ErrInvalidUserName = errors.Wrap(ErrInvalidValue, "유저 이름이 조건을 만족하지 않습니다")
)

// 편의상 서비스는 interface를 이용하지 않고
// 단순히 struct를 이용
type UserService struct{
    // UserRepository는 interface이기 때문에 mocking 될 수 있는 등 좀 더 유연하다.
    userRepository UserRepository
}

// 의존성 주입 형태 - Good
func NewUserService(userRepository UserRepository) *UserService{
    return &UserService{
        userRepository: userRepository,
    }
}

// 의존성을 주입하지 않는 형태 - Bad
func NewUserServiceWithoutInjection(client *ent.UserClient) *UserService{
    return &UserService{
        userRepository: NewUserRepository(client),
    }
}

func (s *UserService) Create(input *UserCreateInput) (*ent.User, error){
    // 도메인/비즈니스 룰, 로직 구현, 조합

    // 유저 ID에 대한 룰
    if len(input.ID) < 6 || 20 < len(input.ID) {
        return nil, ErrInvalidUserID
    }

    // 유저 이름에 대한 룰
    if len(input.Name) < 2 || 20 < len(input.Name) {
        return nil, ErrInvalidUserName
    }

    // 유저 패스워드에 대한 룰
    if len(input.Password) < 6 {
        return nil, ErrInvalidPassword
    }

    // Data Access Layer를 호출
    // 이 부분은 추상화된 Interface에 의존하기 때문에 Mocking될 수 있다.
    user, err := s.userRepository.Create(input)
    if err != nil {
        return nil, err
    }

    return user, nil
}