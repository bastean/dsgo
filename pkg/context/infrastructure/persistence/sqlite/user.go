package sqlite

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/model"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"index:idx_name,unique"`
	Role string
}

type UserTable struct {
	table *gorm.DB
}

func (db *UserTable) Save(user *user.User) error {
	newUser := &User{
		Name: user.Name.Value,
		Role: user.Role.Value,
	}

	result := db.table.Create(newUser)

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

func (db *UserTable) Update(user *user.User) error {
	result := db.table.Where(&User{Name: user.Name.Value}).Updates(user.ToPrimitives())

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

func (db *UserTable) Delete(name *user.Name) error {
	result := db.table.Where(&User{Name: name.Value}).Delete(&User{})

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

func (db *UserTable) Search(name *user.Name) (*user.User, error) {
	primitive := new(user.Primitive)

	result := db.table.Where(&User{Name: name.Value}).Scan(&primitive)

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

func NewUserTable(db *SQLite) (model.UserRepository, error) {
	err := db.Client.AutoMigrate(&User{})

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "NewUserTable",
			What:  "failure to run auto migration for user model",
			Who:   err,
		})
	}

	return &UserTable{
		table: db.Client.Model(&User{}),
	}, nil
}
