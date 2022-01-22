package users

import (
	"project-e-commerces/entities"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetAll() ([]entities.User, error) {
	users := []entities.User{}
	ur.db.Find(&users)
	return users, nil
}

func (ur *UserRepository) Get(userId int) (entities.User, error) {
	user := entities.User{}
	ur.db.Find(&user, userId)
	return user, nil
}

func (ur *UserRepository) Create(newUser entities.User) (entities.User, error) {

	if err := ur.db.Save(&newUser).Error; err != nil {
		return newUser, nil
	}
	return newUser, nil
}

func (ur *UserRepository) Login(email, password string) (entities.User, error) {
	var user entities.User
	getPass := entities.User{}
	ur.db.Select("password").Where("Email = ?", email).Find(&getPass)
	bcrypt.CompareHashAndPassword([]byte(getPass.Password), []byte(password))
	ur.db.Where("Email = ?", email).Find(&user)

	newCart := entities.Cart{
		User_id:       user.ID,
		Total_Product: 0,
		Total_price:   0,
	}
	ur.db.Save(&newCart)

	return user, nil
}
func (ur *UserRepository) Update(updateUser entities.User, userId int) (entities.User, error) {
	user := entities.User{}
	ur.db.Find(&user, "id=?", userId)

	user.Name = updateUser.Name
	user.Password = updateUser.Password

	ur.db.Save(&user)
	return updateUser, nil
}

func (ur *UserRepository) Delete(userId int) (entities.User, error) {
	user := entities.User{}
	ur.db.Find(&user, "id=?", userId)
	ur.db.Delete(&user)
	return user, nil
}
