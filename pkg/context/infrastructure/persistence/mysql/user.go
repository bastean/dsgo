package mysql

import (
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

func (mySQL *User) Save(user *user.User) error {
	new := &UserModel{
		Name: user.Name.Value,
		Role: user.Role.Value,
	}

	result := mySQL.table.Create(new)

	switch {
	case errors.Is(result.Error, gorm.ErrDuplicatedKey):
		return errors.NewAlreadyExist(&errors.Bubble{
			Where: "Save",
			What:  "already registered",
			Who:   result.Error,
		})
	case result.Error != nil:
		return errors.NewInternal(&errors.Bubble{
			Where: "Save",
			What:  "failure to save a user",
			Why: errors.Meta{
				"Name": user.Name.Value,
			},
			Who: result.Error,
		})
	}

	return nil
}

func (mySQL *User) Update(user *user.User) error {
	result := mySQL.table.Where(&UserModel{Name: user.Name.Value}).Updates(user.ToPrimitives())

	if result.Error != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Update",
			What:  "failure to update a user",
			Why: errors.Meta{
				"Name": user.Name.Value,
			},
			Who: result.Error,
		})
	}

	return nil
}

func (mySQL *User) Delete(name *user.Name) error {
	result := mySQL.table.Where(&UserModel{Name: name.Value}).Unscoped().Delete(&UserModel{})

	if result.Error != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Delete",
			What:  "failure to delete a user",
			Why: errors.Meta{
				"Name": name.Value,
			},
			Who: result.Error,
		})
	}

	return nil
}

func (mySQL *User) Search(name *user.Name) (*user.User, error) {
	primitive := new(user.Primitive)

	result := mySQL.table.Where(&UserModel{Name: name.Value}).First(&primitive)

	switch {
	case errors.Is(result.Error, gorm.ErrRecordNotFound):
		return nil, errors.NewNotExist(&errors.Bubble{
			Where: "Search",
			What:  "not found",
			Why: errors.Meta{
				"Index": name.Value,
			},
			Who: result.Error,
		})
	case result.Error != nil:
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Search",
			What:  "failure to search a user",
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
			What:  "failure to create an aggregate from a primitive",
			Why: errors.Meta{
				"Primitive": primitive,
				"Index":     name.Value,
			},
			Who: err,
		})
	}

	return user, nil
}

func UserTable(mySQL *MySQL) (repository.User, error) {
	err := mySQL.Session.AutoMigrate(&UserModel{})

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "UserTable",
			What:  "failure to run auto migration for user model",
			Who:   err,
		})
	}

	return &User{
		table: mySQL.Session.Model(&UserModel{}).Session(&gorm.Session{}),
	}, nil
}
