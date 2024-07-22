package sql

import (
	"fmt"

	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/repository"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Name string `gorm:"index:idx_name,unique"`
	Role string
}

type User struct {
	table *gorm.DB
}

func (database *User) Save(user *user.User) error {
	new := &UserModel{
		Name: user.Name.Value,
		Role: user.Role.Value,
	}

	result := database.table.Create(new)

	switch {
	case errors.Is(result.Error, gorm.ErrDuplicatedKey):
		return errors.NewAlreadyExist(&errors.Bubble{
			Where: "Save",
			What:  fmt.Sprintf("%s already registered", user.Name.Value),
			Who:   result.Error,
		})
	case result.Error != nil:
		return errors.NewInternal(&errors.Bubble{
			Where: "Save",
			What:  "Failure to save a user",
			Why: errors.Meta{
				"Name": user.Name.Value,
			},
			Who: result.Error,
		})
	}

	return nil
}

func (database *User) Update(user *user.User) error {
	result := database.table.Where(&UserModel{Name: user.Name.Value}).Updates(user.ToPrimitives())

	if result.Error != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Update",
			What:  "Failure to update a user",
			Why: errors.Meta{
				"Name": user.Name.Value,
			},
			Who: result.Error,
		})
	}

	return nil
}

func (database *User) Delete(name *user.Name) error {
	result := database.table.Where(&UserModel{Name: name.Value}).Unscoped().Delete(&UserModel{})

	if result.Error != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Delete",
			What:  "Failure to delete a user",
			Why: errors.Meta{
				"Name": name.Value,
			},
			Who: result.Error,
		})
	}

	return nil
}

func (database *User) Search(name *user.Name) (*user.User, error) {
	primitive := new(user.Primitive)

	result := database.table.Where(&UserModel{Name: name.Value}).First(&primitive)

	switch {
	case errors.Is(result.Error, gorm.ErrRecordNotFound):
		return nil, errors.NewNotExist(&errors.Bubble{
			Where: "Search",
			What:  fmt.Sprintf("%s not found", name.Value),
			Why: errors.Meta{
				"Index": name.Value,
			},
			Who: result.Error,
		})
	case result.Error != nil:
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Search",
			What:  "Failure to search a user",
			Why: errors.Meta{
				"Name": name.Value,
			},
			Who: result.Error,
		})
	}

	user, err := user.FromPrimitives(primitive)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Search",
			What:  "Failure to create an aggregate from a primitive",
			Why: errors.Meta{
				"Primitive": primitive,
				"Index":     name.Value,
			},
			Who: err,
		})
	}

	return user, nil
}

func UserTable(database *Database) (repository.User, error) {
	err := database.Session.AutoMigrate(&UserModel{})

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "UserTable",
			What:  "Failure to run auto migration for user model",
			Who:   err,
		})
	}

	return &User{
		table: database.Session.Model(&UserModel{}).Session(&gorm.Session{}),
	}, nil
}
