package faker

import (
	"time"

	"github.com/fajarcandraaa/pizza_hub/helpers"
	"github.com/fajarcandraaa/pizza_hub/internal/entity"
	"github.com/google/uuid"
)

const (
	FakeInitialsChef01 = "Chef01"
	FakeNameChef01     = "So Klienz"
	FakeUsernameChef01 = "chef01"
	FakeInitialsChef02 = "Chef02"
	FakeNameChef02     = "Bai Gon"
	FakeUsernameChef02 = "chef02"
	FakeInitialsChef03 = "Chef03"
	FakeNameChef03     = "Sun Light"
	FakeUsernameChef03 = "chef03"

	FakePassword = "123456"

	FakeNameMenu01 = "Pizza Cheese"
	FakeCodeMenu01 = "pzzachse"
	FakeNameMenu02 = "Pizza BBQ"
	FakeCodeMenu02 = "pzzabq"
)

func FakeChef() []*entity.Chef {
	var (
		fakeusers []*entity.Chef
	)
	psswd, err := helpers.HashPassword(FakePassword)
	if err != nil {
		return nil
	}

	fake01 := &entity.Chef{
		ID:            uuid.NewString(),
		CheffInitials: FakeInitialsChef01,
		CheffName:     FakeNameChef01,
		Username:      FakeUsernameChef01,
		Password:      psswd,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}
	fakeusers = append(fakeusers, fake01)

	fake02 := &entity.Chef{
		ID:            uuid.NewString(),
		CheffInitials: FakeInitialsChef02,
		CheffName:     FakeNameChef02,
		Username:      FakeUsernameChef02,
		Password:      psswd,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}
	fakeusers = append(fakeusers, fake02)

	fake03 := &entity.Chef{
		ID:            uuid.NewString(),
		CheffInitials: FakeInitialsChef03,
		CheffName:     FakeNameChef03,
		Username:      FakeUsernameChef03,
		Password:      psswd,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}
	fakeusers = append(fakeusers, fake03)

	return fakeusers
}

func FakeMenu() []*entity.Menu {
	var (
		menus []*entity.Menu
	)

	menu01 := &entity.Menu{
		ID:        uuid.NewString(),
		MenuName:  FakeNameMenu01,
		MenuCode:  FakeCodeMenu01,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	menus = append(menus, menu01)

	menu02 := &entity.Menu{
		ID:        uuid.NewString(),
		MenuName:  FakeNameMenu02,
		MenuCode:  FakeCodeMenu02,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	menus = append(menus, menu02)

	return menus
}
