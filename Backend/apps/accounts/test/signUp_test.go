package accounts

import (
	"GoTransact/apps/accounts/models"

	// "GoTransact/apps/accounts/models"
	utils "GoTransact/apps/accounts"
	"GoTransact/pkg/db"

	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignupsuccess(t *testing.T) {
	SetupTestDb()
	ClearDatabase()
	utils.AccountLogInit()
	utils.InitValidation()
	input := utils.RegisterInput{
		FirstName:   "testfirstname",
		LastName:    "testlastname",
		Email:       "test@gmail.com",
		Companyname: "trellissoft",
		Password:    "Password@123",
	}

	status, message, data := utils.Signup(input)

	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, "User created successfully", message)
	assert.Equal(t, map[string]interface{}{}, data)

	var user models.User
	err := db.DB.Where("email = ?", input.Email).First(&user).Error
	assert.NoError(t, err)
	assert.Equal(t, input.FirstName, user.FirstName)
	assert.Equal(t, input.LastName, user.LastName)
	assert.Equal(t, input.Email, user.Email)

	var company models.Company
	err = db.DB.Where("name = ?", input.Companyname).First(&company).Error
	assert.NoError(t, err)
	assert.Equal(t, input.Companyname, company.Name)
	ClearDatabase()
	CloseTestDb()
}

func TestSignupEmailAreadyExist(t *testing.T) {
	SetupTestDb()
	ClearDatabase()
	utils.AccountLogInit()
	utils.InitValidation()
	existingUser := models.User{
		FirstName: "testfirstname",
		LastName:  "testlastname",
		Email:     "test@gmail.com",
		Company: models.Company{
			Name: "trellissoft",
		},
	}

	db.DB.Create(&existingUser)

	input := utils.RegisterInput{
		FirstName:   "otherfirstname",
		LastName:    "otherlastname",
		Email:       "test@gmail.com",
		Companyname: "Google",
		Password:    "Password@123",
	}

	status, message, data := utils.Signup(input)

	assert.Equal(t, http.StatusBadRequest, status)
	assert.Equal(t, "email already exists", message)
	assert.Equal(t, map[string]interface{}{}, data)
	ClearDatabase()
	CloseTestDb()
}

func TestSignupCompanyAlreadyExist(t *testing.T) {
	SetupTestDb()
	ClearDatabase()
	utils.AccountLogInit()
	utils.InitValidation()

	existingUser := models.User{
		FirstName: "testfirstname",
		LastName:  "testlastname",
		Email:     "test@gmail.com",
		Company: models.Company{
			Name: "trellissoft",
		},
	}

	db.DB.Create(&existingUser)

	input := utils.RegisterInput{
		FirstName:   "othername",
		LastName:    "otherlastname",
		Email:       "testother@gmail.com",
		Companyname: "trellissoft",
		Password:    "Password@123",
	}

	status, message, data := utils.Signup(input)

	assert.Equal(t, http.StatusBadRequest, status)
	assert.Equal(t, "company already exists", message)
	assert.Equal(t, map[string]interface{}{}, data)
	ClearDatabase()
	CloseTestDb()
}

func TestSignupInvaldPassword(t *testing.T) {
	SetupTestDb()
	ClearDatabase()
	utils.AccountLogInit()
	utils.InitValidation()

	input := utils.RegisterInput{
		FirstName:   "otherfirstname",
		LastName:    "otherlastname",
		Email:       "test@gmail.com",
		Companyname: "trellissoft",
		Password:    "password@123",
	}

	status, message, data := utils.Signup(input)

	assert.Equal(t, http.StatusBadRequest, status)
	assert.Equal(t, "Password should contain atleast one upper case character,one lower case character,one number and one special character", message)
	assert.Equal(t, map[string]interface{}{}, data)
	ClearDatabase()
	CloseTestDb()
}
