package users

import (
	"project-e-commerces/configs"
	"project-e-commerces/entities"
	"project-e-commerces/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	userRepo := NewUsersRepo(db)

	t.Run("Insert User into Database", func(t *testing.T) {
		var mockInserUser entities.User

		mockInserUser.Name = "TestNameDB1"
		mockInserUser.Email = "TestDB@email.com"
		mockInserUser.Password = "TestPasswordDB1"

		res, err := userRepo.Create(mockInserUser)
		assert.Nil(t, err)
		assert.Equal(t, mockInserUser.Name, res.Name)
		assert.Equal(t, 1, int(res.ID))

	})
	t.Run("Insert User into Database", func(t *testing.T) {
		var mockInserUser entities.User

		mockInserUser.Name = "TestNameDB1"
		mockInserUser.Email = "TestDB@email.com"
		mockInserUser.Password = "TestPasswordDB1"

		res, err := userRepo.Create(mockInserUser)
		assert.Nil(t, err)
		assert.Equal(t, mockInserUser.Name, res.Name)
		assert.Equal(t, 0, int(res.ID))

	})

	t.Run("Login", func(t *testing.T) {
		res, err := userRepo.Login("TestDB@email.com", "TestPasswordDB1")
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

	t.Run("Select Users from Database", func(t *testing.T) {
		res, err := userRepo.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

	t.Run("Select User from Database", func(t *testing.T) {
		res, err := userRepo.Get(1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

	t.Run("Update User ", func(t *testing.T) {
		var mockUpdateUser entities.User
		mockUpdateUser.Name = "UPDATE TestNameDB1"
		mockUpdateUser.Password = "UPDATETestPasswordDB1"

		res, err := userRepo.Update(mockUpdateUser, 1)
		assert.Nil(t, err)
		assert.Equal(t, mockUpdateUser.Name, res.Name)
	})

	t.Run("Delete User", func(t *testing.T) {
		res, err := userRepo.Delete(1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

}
