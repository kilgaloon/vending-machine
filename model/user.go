package model

import (
	"errors"
	"math"

	"github.com/kilgaloon/atm/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrUserAlreadyExist     = errors.New("User with same username already exist")
	ErrUserNotFound         = errors.New("User not found")
	ErrInvalidDepositAmount = errors.New("Invalid deposit amount")

	// roles
	Seller = "seller"
	Buyer  = "buyer"
)

type Role string

func (r Role) IsSeller() bool {
	return r.String() == Seller
}

func (r Role) IsBuyer() bool {
	return r.String() == Buyer
}

func (r Role) String() string {
	return string(r)
}

var (
	allowedDepositAmounts = []uint64{100, 50, 20, 10, 5}
)

type Deposit uint64

func (d Deposit) IsValid() bool {
	dep := uint64(d)
	for _, a := range allowedDepositAmounts {
		if a == dep {
			return true
		}
	}

	return false
}

type BuyRequest struct {
	ProductID uint   `json:"product_id"`
	Amount    uint64 `json:"amount"`
}

type BuyResponse struct {
	TotalSpent uint64   `json:"total_spent"`
	ProductID  uint     `json:"products"`
	Change     []uint64 `json:"change"`
}

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Deposit  uint64 `json:"deposit"`
	Role     Role   `json:"role"`
}

func (u *User) Find() (*User, error) {
	db := utils.DBConnect()

	// find user by provided username
	result := db.Where("username = ? OR id = ?", u.Username, u.ID).Find(u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}

	return u, nil
}

func (u *User) Create() (*User, error) {
	db := utils.DBConnect()

	// find user by provided username
	result := db.Where("username = ?", u.Username).First(&User{})

	if result.RowsAffected > 0 {
		return nil, ErrUserAlreadyExist
	}

	// hash password
	pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return nil, err
	}

	u.Password = string(pass)

	result = db.Create(u)
	if result.Error != nil {
		return nil, result.Error
	}

	return u, nil
}

func (u *User) Update() (*User, error) {
	db := utils.DBConnect()

	// find user by provided username
	result := db.Where("username = ?", u.Username).First(&User{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}

	// hash password
	if u.Password != "" {
		pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
		if err != nil {
			return nil, err
		}

		u.Password = string(pass)
	}

	result = db.Save(u)
	if result.Error != nil {
		return nil, result.Error
	}

	return u, nil
}

func (u *User) Delete() error {
	db := utils.DBConnect()

	// find user by provided username
	result, err := u.Find()
	if err != nil {
		return err
	}

	deleteResult := db.Delete(result)
	if deleteResult.Error != nil {
		return deleteResult.Error
	}

	return nil
}

func (u *User) DepositAmount(d Deposit) (uint64, error) {
	db := utils.DBConnect()

	if !d.IsValid() {
		return 0, ErrInvalidDepositAmount
	}

	user, err := u.Find()
	if err != nil {
		return 0, err
	}

	user.Deposit = user.Deposit + uint64(d)

	result := db.Save(user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.Deposit, nil
}

func (u *User) ResetDeposit() error {
	u, err := u.Find()
	if err != nil {
		return err
	}

	u.Deposit = 0
	_, err = u.Update()
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Buy(req BuyRequest) (*BuyResponse, error) {
	if u.ID == 0 {
		u, _ = u.Find()
	}

	db := utils.DBConnect()

	product := &Product{}
	// find product
	db.Where("id = ?", req.ProductID).Find(product)
	if product.AmountAvailable < req.Amount {
		return nil, errors.New("Not enough products available")
	}

	// check does user has that much money
	if (req.Amount * product.Cost) > u.Deposit {
		return nil, errors.New("Not enough money")
	}

	// this logic can go into product model also
	product.AmountAvailable = product.AmountAvailable - req.Amount
	_, err := product.Update()
	if err != nil {
		return nil, err
	}

	diff := u.Deposit - (req.Amount * product.Cost)

	c := []uint64{}
	change := makeChange(diff, c)

	u.ResetDeposit()

	resp := &BuyResponse{
		TotalSpent: req.Amount * product.Cost,
		ProductID:  product.ID,
		Change:     change,
	}

	return resp, nil

}

func makeChange(diff uint64, change []uint64) []uint64 {
	for _, a := range allowedDepositAmounts {
		d := diff % a
		if d == diff {
			continue
		}

		if d == 0 {
			n := int(diff / a)
			for i := 0; i < n; i++ {
				change = append(change, a)
			}

			return change
		}

		n := int(math.Floor(float64(diff / a)))
		for i := 0; i < n; i++ {
			change = append(change, a)
		}

		//change = append(change, a)
		return makeChange(d, change)
	}

	return change
}
