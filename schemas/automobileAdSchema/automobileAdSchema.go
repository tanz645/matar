package automobileAdSchema

import (
	"context"
	"fmt"
	"matar/clients"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var AutomobileAdCollectionName = "automobile_ad"

type address struct {
	Country     string `json:"country" validate:"required"`
	StateRegion string `json:"state_region" bson:"state_region" validate:"required"`
	City        string `json:"city" validate:"required"`
}

type brand struct {
	Name string `json:"name" validate:"required"`
	Logo string `json:"logo" validate:"required"`
}

type model struct {
	Name string `json:"name" validate:"required"`
	Year uint16 `json:"year" validate:"required"`
}

type milage struct {
	Amount uint32 `json:"amount" validate:"required"`
	Unit   string `json:"unit" validate:"required,oneof=km mile"`
}

type price struct {
	TotalAmount uint32 `json:"total_amount" bson:"total_amount" validate:"required"`
	Unit        string `json:"unit" validate:"required,oneof=MAD USD EUR"`
}

type topSpeed struct {
	Amount uint16 `json:"amount" validate:"max=10000,min=1"`
	Unit   string `json:"unit" validate:"required,oneof=km mile"`
}

type fuelTank struct {
	Amount float32 `json:"amount" validate:"max=10000,min=1"`
	Unit   string  `json:"unit" validate:"oneof=gal litre"`
}

type generalSpecification struct {
	Doors                uint8    `json:"doors" validate:"max=50,min=1"`
	TopSpeed             topSpeed `json:"top_speed" bson:"top_speed"`
	ZeroToHundredKmInSec float32  `json:"zero_to_hundred_km_in_sec" bson:"zero_to_hundred_km_in_sec"  validate:"max=10000,min=1"`
	TrunckVolumeInLitre  float32  `json:"trunck_volume_in_litre" bson:"trunck_volume_in_litre" validate:"max=1000000,min=1"`
	FuelTank             fuelTank `json:"fuel_tank" bson:"fuel_tank"`
	GearNumber           uint8    `json:"gear_number" bson:"gear_number" validate:"max=50,min=1"`
	FrontRimSizeInch     uint8    `json:"front_rim_size_inch" bson:"front_rim_size_inch" validate:"max=100,min=1"`
	RearRimSizeInch      uint8    `json:"rear_rim_size_inch" bson:"rear_rim_size_inch" validate:"max=100,min=1"`
	FrontTyre            string   `json:"front_tyre" bson:"front_tyre"`
	RearTyre             string   `json:"rear_tyre" bson:"rear_tyre"`
}

type engine struct {
	Engine         string   `json:"engine" validate:"required"`
	Cc             *float32 `json:"cc,omitempty" validate:"max=60000,min=1"`
	Cylinders      *uint8   `json:"cylinders,omitempty" validate:"max=200,min=1"`
	HorsePower     *float32 `json:"horse_power,omitempty" bson:"horse_power"`
	Torque         string   `json:"torque"`
	Aspiration     *string  `json:"aspiration,omitempty"`
	FuelConsumtion *string  `json:"fuel_consumtion,omitempty" bson:"fuel_consumtion"`
	Co2Emission    *string  `json:"co2_emission,omitempty" bson:"co2_emission"`
}

type dimension struct {
	Length         string `json:"length"`
	Width          string `json:"width"`
	Height         string `json:"height"`
	Weight         string `json:"weight"`
	GroundClerance string `json:"ground_clerance" bson:"ground_clerance"`
	TurningCircle  string `json:"turning_circle" bson:"turning_circle"`
}

type specification struct {
	General            *generalSpecification `json:"general,omitempty"`
	Engine             engine                `json:"engine" validate:"required"`
	Dimension          *dimension            `json:"dimension,omitempty"`
	Suspension         []string              `json:"suspension" validate:"max=10"`
	SafetyAndSecurity  []string              `json:"safety_and_security" bson:"safety_and_security" validate:"max=100"`
	ComfortAndInterior []string              `json:"comfort_and_interior" bson:"comfort_and_interior" validate:"max=100"`
	Driving            []string              `json:"driving" validate:"max=100"`
	OtherFeatures      []string              `json:"other_features" bson:"other_features" validate:"max=100"`
}

type AutomobileAd struct {
	Id               primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Title            string             `json:"title" validate:"required"`
	UserId           string             `json:"user_id" bson:"user_id"`
	Brand            brand              `json:"brand" validate:"required"`
	BodyType         []string           `json:"body_type" bson:"body_type" validate:"required,min=1,max=2,dive,oneof=sedan coupe convertible_roadstar sports_supercar cuv_crossover suv_muv hatchback wagon_stationwagon pickup van_minivan bus_minibus truck motorcycle atv"`
	Address          address            `json:"address" validate:"required"`
	Model            model              `json:"model" validate:"required"`
	Milage           milage             `json:"milage" validate:"required"`
	Price            price              `json:"price" validate:"required"`
	Images           []string           `json:"images" validate:"required,min=1,max=20"`
	ContactNo        []string           `json:"contact_no" bson:"contact_no" validate:"required,min=1,max=20,dive,e164"`
	FuelType         string             `json:"fuel_type" bson:"fuel_type" validate:"required,oneof=petrol_gasoline diesel cng lpg hybrid electric"`
	Color            string             `json:"color" validate:"required,hexcolor"`
	Transmission     string             `json:"transmission" validate:"required,oneof=automatic manual"`
	WheelDrive       string             `json:"wheel_drive" bson:"wheel_drive" validate:"required,oneof=awd 4wd rwd fwd"`
	UsageCondition   string             `json:"usage_condition" bson:"usage_condition" validate:"required,oneof=new used"`
	RegistrationCard string             `json:"registration_card" bson:"registration_card"`
	SellerComments   string             `json:"seller_comments" bson:"seller_comments"`
	SeatCapacity     uint16             `json:"seat_capacity" validate:"max=1000,min=1"`
	Specification    specification      `json:"specification" validate:"required"`
	Active           bool               `json:"active"`
	CreatedAt        time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at" bson:"updated_at"`
}

type AutomobileAdGeneral struct {
	Id               primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Title            string             `json:"title"`
	Brand            brand              `json:"brand"`
	BodyType         []string           `json:"body_type" bson:"body_type"`
	Address          address            `json:"address"`
	Model            model              `json:"model"`
	Milage           milage             `json:"milage"`
	Price            price              `json:"price"`
	Images           []string           `json:"images"`
	ContactNo        []string           `json:"contact_no" bson:"contact_no"`
	FuelType         string             `json:"fuel_type" bson:"fuel_type"`
	Color            string             `json:"color"`
	Transmission     string             `json:"transmission"`
	WheelDrive       string             `json:"wheel_drive" bson:"wheel_drive"`
	UsageCondition   string             `json:"usage_condition" bson:"usage_condition"`
	RegistrationCard string             `json:"registration_card" bson:"registration_card"`
	SellerComments   string             `json:"seller_comments" bson:"seller_comments"`
	SeatCapacity     uint16             `json:"seat_capacity"`
	Specification    specification      `json:"specification"`
	CreatedAt        time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at" bson:"updated_at"`
}

type AutomobileAdInListing struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	Title        string             `json:"title"`
	Brand        brand              `json:"brand"`
	BodyType     []string           `json:"body_type" bson:"body_type"`
	Address      address            `json:"address"`
	Model        model              `json:"model"`
	Milage       milage             `json:"milage"`
	Price        price              `json:"price"`
	Images       []string           `json:"images"`
	FuelType     string             `json:"fuel_type" bson:"fuel_type"`
	Color        string             `json:"color"`
	Transmission string             `json:"transmission"`
	WheelDrive   string             `json:"wheel_drive" bson:"wheel_drive"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
}

type SearchAutomobileAdGeneral struct {
	Limit        uint16 `form:"limit" validate:"required,max=50,min=1"`
	Page         uint16 `form:"page" validate:"required,min=1"`
	BrandName    string `form:"brand_name"`
	CityName     string `form:"city_name"`
	BodyType     string `form:"body_type"`
	FuelType     string `form:"fuel_type"`
	Transmission string `form:"transmission" validate:"omitempty,oneof=automatic manual"`
	WheelDrive   string `json:"wheel_drive" bson:"wheel_drive"  validate:"omitempty,oneof=awd 4wd rwd fwd"`
	SortBy       string `form:"sort_by" validate:"required,oneof=price.total_amount milage.amount created_at"`
	SortOrder    int8   `form:"sort_order" validate:"required,min=-1,max=1"`
}

type UpdateAutomobileAdActiveStatus struct {
	Active bool `json:"active"`
}
type DeleteAutomobileAdImage struct {
	Url string `json:"url" validate:"required"`
}

func CreateAutomobileAdIndexes(ctx context.Context, client *mongo.Client) {
	col := clients.GetMongoCollection(client, AutomobileAdCollectionName)
	models := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{Key: "address.city", Value: bsonx.Int32(1)},
				{Key: "body_type", Value: bsonx.Int32(1)},
				{Key: "brand.name", Value: bsonx.Int32(1)},
				{Key: "transmission", Value: bsonx.Int32(1)},
				{Key: "fuel_type", Value: bsonx.Int32(1)},
				{Key: "wheel_drive", Value: bsonx.Int32(1)},
				{Key: "price.total_amount", Value: bsonx.Int32(1)},
				{Key: "milage.amount", Value: bsonx.Int32(1)},
			},
		},
		{
			Keys: bsonx.Doc{
				{Key: "body_type", Value: bsonx.Int32(1)},
				{Key: "address.city", Value: bsonx.Int32(1)},
				{Key: "brand.name", Value: bsonx.Int32(1)},
				{Key: "transmission", Value: bsonx.Int32(1)},
				{Key: "fuel_type", Value: bsonx.Int32(1)},
				{Key: "wheel_drive", Value: bsonx.Int32(1)},
				{Key: "price.total_amount", Value: bsonx.Int32(1)},
				{Key: "milage.amount", Value: bsonx.Int32(1)},
			},
		},
		{
			Keys: bsonx.Doc{
				{Key: "brand.name", Value: bsonx.Int32(1)},
				{Key: "address.city", Value: bsonx.Int32(1)},
				{Key: "body_type", Value: bsonx.Int32(1)},
				{Key: "transmission", Value: bsonx.Int32(1)},
				{Key: "fuel_type", Value: bsonx.Int32(1)},
				{Key: "wheel_drive", Value: bsonx.Int32(1)},
				{Key: "price.total_amount", Value: bsonx.Int32(1)},
				{Key: "milage.amount", Value: bsonx.Int32(1)},
			},
		},
		{
			Keys: bsonx.Doc{
				{Key: "transmission", Value: bsonx.Int32(1)},
				{Key: "address.city", Value: bsonx.Int32(1)},
				{Key: "body_type", Value: bsonx.Int32(1)},
				{Key: "brand.name", Value: bsonx.Int32(1)},
				{Key: "fuel_type", Value: bsonx.Int32(1)},
				{Key: "wheel_drive", Value: bsonx.Int32(1)},
				{Key: "price.total_amount", Value: bsonx.Int32(1)},
				{Key: "milage.amount", Value: bsonx.Int32(1)},
			},
		},
		{
			Keys: bsonx.Doc{
				{Key: "user_id", Value: bsonx.Int32(1)},
				{Key: "created_at", Value: bsonx.Int32(1)},
			},
		},
	}
	_, err := col.Indexes().CreateMany(ctx, models)
	if err != nil {
		fmt.Println(err)
	}
}
