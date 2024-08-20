package models

import (
	"time"
)

// Enums
type UserRoles string

const (
	HeadAdmin UserRoles = "HeadAdmin"
	Worker    UserRoles = "Worker"
	None      UserRoles = "None"
)

type OrderStatus string

const (
	Ordered   OrderStatus = "Ordered"
	Completed OrderStatus = "Completed"
	Pending   OrderStatus = "Pending"
	Canceled  OrderStatus = "Canceled"
)

type CakeSize string

const (
	Big   CakeSize = "BIG"
	Small CakeSize = "SMALL"
)

type CookieSize string

const (
	OneKG   CookieSize = "ONE_KG"
	TwoKG   CookieSize = "TWO_KG"
	ThreeKG CookieSize = "THREE_KG"
)

// Models
type User struct {
	ID                 int                 `gorm:"primaryKey"`
	Email              string              `gorm:"unique;not null"`
	Password           string              `gorm:"not null"`
	Username           string              `gorm:"unique;not null"`
	FirstName          string              `gorm:"not null"`
	LastName           string              `gorm:"not null"`
	Role               UserRoles           `gorm:"default:None"`
	IsActive           bool                `gorm:"default:false"`
	CreatedAt          time.Time           `gorm:"autoCreateTime"`
	UpdatedAt          time.Time           `gorm:"autoUpdateTime"`
	VerificationTokens []VerificationToken `gorm:"foreignKey:UserID"`
	Orders             []Order             `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "User"
}

type VerificationToken struct {
	ID        int       `gorm:"primaryKey"`
	Token     string    `gorm:"unique;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UserID    int
	User      User `gorm:"foreignKey:UserID"`
}

func (VerificationToken) TableName() string {
	return "VerificationToken"
}

type Product struct {
	ID              int    `gorm:"primaryKey"`
	Name            string `gorm:"not null"`
	Description     string
	CategoryID      int
	SubCategoryID   *int
	Price           *int
	Images          []Image             `gorm:"foreignKey:ProductID"`
	Category        Category            `gorm:"foreignKey:CategoryID"`
	Ingredients     []ProductIngredient `gorm:"foreignKey:ProductID"`
	SubCategory     *SubCategory        `gorm:"foreignKey:SubCategoryID"`
	OrderedProducts []OrderedProduct    `gorm:"foreignKey:ProductID"`
}

func (Product) TableName() string {
	return "Product"
}

type Order struct {
	ID                         int              `gorm:"primaryKey"`
	OrderedBy                  string           `gorm:"not null"`
	OrderedProducts            []OrderedProduct `gorm:"foreignKey:OrderID"`
	OrderDeliveryInformationID *int
	OrderDeliveryInformation   *OrderDeliveryInformation `gorm:"foreignKey:OrderDeliveryInformationID"`
	IsOrderVerified            bool                      `gorm:"not null"`
	Status                     OrderStatus               `gorm:"not null"`
	CompletedBy                *User                     `gorm:"foreignKey:UserID"`
	UserID                     *int
	OrderDateTime              *time.Time
	CreatedAt                  time.Time `gorm:"autoCreateTime"`
	VerificationToken          *string   `gorm:"unique"`
}

func (Order) TableName() string {
	return "Order"
}

type OrderDeliveryInformation struct {
	ID       int `gorm:"primaryKey"`
	City     *string
	Adresa   *string
	Zip      *string
	Email    *string
	PhoneNum *string
	Orders   []Order `gorm:"foreignKey:OrderDeliveryInformationID"`
}

func (OrderDeliveryInformation) TableName() string {
	return "OrderDeliveryInformation"
}

type OrderedProduct struct {
	ID          int `gorm:"primaryKey"`
	ProductID   int
	OrderID     int
	Description *string
	OptionID    *int    `gorm:"unique"`
	Quantity    int     `gorm:"default:1"`
	Product     Product `gorm:"foreignKey:ProductID"`
	Option      *Option `gorm:"foreignKey:OptionID"`
}

func (OrderedProduct) TableName() string {
	return "OrderedProduct"
}

type Option struct {
	ID               int `gorm:"primaryKey"`
	CakeSize         *CakeSize
	CookieSize       *CookieSize
	OrderedProductID *int            `gorm:"unique"`
	OrderedProduct   *OrderedProduct `gorm:"foreignKey:OrderedProductID"`
}

func (Option) TableName() string {
	return "Option"
}

type Category struct {
	ID            int           `gorm:"primaryKey"`
	Name          string        `gorm:"unique;not null"`
	SubCategories []SubCategory `gorm:"foreignKey:CategoryID"`
	Products      []Product     `gorm:"foreignKey:CategoryID"`
}

func (Category) TableName() string {
	return "Category"
}

type SubCategory struct {
	ID         int    `gorm:"primaryKey"`
	Name       string `gorm:"not null"`
	CategoryID int
	Category   Category  `gorm:"foreignKey:CategoryID"`
	Products   []Product `gorm:"foreignKey:SubCategoryID"`
}

func (SubCategory) TableName() string {
	return "SubCategory"
}

type Ingredient struct {
	ID        int                 `gorm:"primaryKey"`
	Name      string              `gorm:"not null"`
	IsAlergen bool                `gorm:"not null"`
	Products  []ProductIngredient `gorm:"foreignKey:IngredientID"`
}

func (Ingredient) TableName() string {
	return "Ingredient"
}

type ProductIngredient struct {
	ProductID    int
	IngredientID int
	Product      Product    `gorm:"foreignKey:ProductID"`
	Ingredient   Ingredient `gorm:"foreignKey:IngredientID"`

	// Composite primary key (custom)
	PrimaryKey [2]int `gorm:"primaryKey"`
}

func (ProductIngredient) TableName() string {
	return "ProductIngredient"
}

type Image struct {
	ID        int    `gorm:"primaryKey"`
	ImageUrl  string `gorm:"not null"`
	ProductID int
	Product   Product `gorm:"foreignKey:ProductID"`
}

func (Image) TableName() string {
	return "Image"
}
