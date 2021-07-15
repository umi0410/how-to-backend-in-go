package main

import (
    "github.com/golang/mock/gomock"
    log "github.com/sirupsen/logrus"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/umi0410/how-to-backend-in-go/testcode/ent"
    "testing"
)

type ManualMockUserRepository struct{}

func (m *ManualMockUserRepository) Create(input *UserCreateInput) (*ent.User, error) {
    log.Info("직접 Mocking. Args: ", input)
    return &ent.User{}, nil
}

func TestUserService_CreateWithManualMock(t *testing.T){
    s := NewUserService(new(ManualMockUserRepository))

    _, err := s.Create(&UserCreateInput{ID: "jinsu_umi", Name: "Jinsu Park", Password: "123123123"})
    // 이런 식으로 로그를 통해 하나 하나 테스트 결과를 확인하는 것이 아니라
    //t.Log("에러가 존재하지 않아야합니다.")
    //t.Log("err == nil 인지 확인하십시오. err == nil: " , err == nil)
    // assert를 이용해 자동으로 성공/실패를 판단하십시오.
    assert.NoError(t, err)
}

func TestUserService_CreateWithMockery(t *testing.T){
    mockUserRepository := &MockUserRepository{}
    // method를 문자열 자체로 설정해야해서 safe하지 않음.
    mockUserRepository.On("Create", mock.Anything).Run(func(args mock.Arguments) {
        t.Log("testify/mock과 mockery를 이용한 Mocking. Args: ", args.Get(0))
    }).Return(&ent.User{}, nil)
    // 해당 이름의 유저가 있는지 확인
    s := NewUserService(mockUserRepository)

    _, err := s.Create(&UserCreateInput{ID: "jinsu_umi", Name: "Jinsu Park", Password: "123123123"})
    assert.NoError(t, err)
}

func TestUserService_CreateWithMockgen(t *testing.T){
    // gomock controller을 만들고 Finish 시켜주는 등의 불편함 존재.
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

	mockUserRepository := NewGomockRepository(ctrl)
	// 원래의 type을 기반으로 method가 safe하게 제공됨.
	mockUserRepository.EXPECT().Create(gomock.Any()).DoAndReturn(
		func(input *UserCreateInput) (*ent.User, error) {
			t.Log("Gomock을 이용한 Mocking. Args: ", input)
			return &ent.User{}, nil
		})
    // 해당 이름의 유저가 있는지 확인
    s := NewUserService(mockUserRepository)

    _, err := s.Create(&UserCreateInput{ID: "jinsu_umi", Name: "Jinsu Park", Password: "123123123"})
    assert.NoError(t, err)
}

func TestUserService_유저_생성(t *testing.T){
    // repository mocking
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

	mockUserRepository := NewGomockRepository(ctrl)
	// 원래의 type을 기반으로 method가 safe하게 제공됨.
	mockUserRepository.EXPECT().Create(gomock.Any()).DoAndReturn(
		func(input *UserCreateInput) (*ent.User, error) {
			t.Log("Gomock을 이용한 Mocking. Args: ", input)
			return &ent.User{}, nil
		})
    // 해당 이름의 유저가 있는지 확인
    s := NewUserService(mockUserRepository)

    t.Run("성공", func(t *testing.T) {
        _, err := s.Create(&UserCreateInput{ID: "jinsu_umi", Name: "Jinsu Park", Password: "123123123"})
        assert.NoError(t, err)
    })

    t.Run("에러) 너무 짧은 아이디", func(t *testing.T) {
        _, err := s.Create(&UserCreateInput{ID: "short", Name: "Jinsu Park", Password: "123123123"})
        assert.ErrorIs(t, err, ErrInvalidUserID)
        assert.ErrorIs(t, err, ErrInvalidValue)
    })

    t.Run("에러) 너무 긴 이름", func(t *testing.T) {
        _, err := s.Create(&UserCreateInput{ID: "jinsu_umi", Name: "니노막시무스 카이저 쏘제 쏘냐도르앤 스파르타 죽지 않아 나는 죽지않아 오오오오 나는 카이저 쏘제", Password: "123123123"})
        assert.ErrorIs(t, err, ErrInvalidUserName)
        assert.ErrorIs(t, err, ErrInvalidValue)
    })

    t.Run("에러) 너무 짧은 비밀번호들", func(t *testing.T) {
        errorPasswords := []string{
            "123",
            "abc",
            "a1b2",
            "asd",
        }
        for _, errorPassword := range errorPasswords {
            t.Run("테스트 케이스) " + errorPassword, func(t *testing.T) {
                _, err := s.Create(&UserCreateInput{ID: "jinsu_umi", Name: "Jinsu Park", Password: errorPassword})
                assert.ErrorIs(t, err, ErrInvalidPassword)
                assert.ErrorIs(t, err, ErrInvalidValue)
            })
        }
    })
}