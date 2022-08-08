package automobileAdModel

import (
	"context"
	"matar/clients"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	Total_amount uint32 `json:"amount" validate:"required"`
	Unit         string `json:"unit" validate:"required,oneof=MAD USD EUR"`
}

type topSpeed struct {
	Amount uint16 `json:"amount" validate:"max=10000,min=1"`
	Unit   string `json:"unit" validate:"required,oneof=km mile"`
}

type generalSpecification struct {
	Doors                uint8    `json:"Doors" validate:"max=50,min=1"`
	TopSpeed             topSpeed `json:"top_speed" bson:"top_speed"`
	ZeroToHundredKmInSec uint16   `json:"zero_to_hundred_km_in_sec" bson:"zero_to_hundred_km_in_sec"  validate:"max=10000,min=1"`
	GearNumber           uint8    `json:"gear_number" bson:"gear_number" validate:"max=50,min=1"`
	FrontRimSizeInch     uint8    `json:"front_rim_size_inch" bson:"front_rim_size_inch" validate:"max=100,min=1"`
	RearRimSizeInch      uint8    `json:"rear_rim_size_inch" bson:"rear_rim_size_inch" validate:"max=100,min=1"`
	FrontTyre            string   `json:"front_tyre" bson:"front_tyre"`
	RearTyre             string   `json:"rear_tyre" bson:"rear_tyre"`
}

type engine struct {
	Engine                          string    `json:"engine" validate:"required"`
	Cc                              *uint16   `json:"cc,omitempty" validate:"max=60000,min=1"`
	TrunckVolumeInLitre             *uint32   `json:"trunck_volume_in_litre,omitempty" bson:"trunck_volume_in_litre" validate:"max=1000000,min=1"`
	Cylinders                       *uint8    `json:"cylinders,omitempty" validate:"max=200,min=1"`
	HorsePower                      *topSpeed `json:"top_speed,omitempty" bson:"top_speed"`
	Aspiration                      *string   `json:"aspiration,omitempty"`
	FuelConsumtionLitrePerHundredKm *float32  `json:"fuel_consumtion_litre_per_hundred_Km,omitempty" bson:"fuel_consumtion_litre_per_hundred_Km" validate:"max=1000,min=0"`
	Co2EmissionGramPerKm            *float32  `json:"co2_emission_gram_per_km,omitempty" bson:"co2_emission_gram_per_km" validate:"max=10000,min=0"`
}

type dimension struct {
	Length uint32 `json:"length"`
	Width  uint32 `json:"width"`
	Height uint32 `json:"height"`
	Unit   string `json:"unit" validate:"oneof=meter milimeter centimeter"`
}

type specification struct {
	General            *generalSpecification `json:"general,omitempty"`
	Engine             engine                `json:"engine" validate:"required"`
	Dimension          *dimension            `json:"dimension,omitempty"`
	SafetyAndSecurity  []string              `json:"safety_and_security" bson:"safety_and_security"`
	DrivingAndInterior []string              `json:"driving_and_interior" bson:"driving_and_interior"`
	Outside            []string              `json:"outside"`
}

type AutomobileAd struct {
	Id               primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Title            string             `json:"title" validate:"required"`
	UserId           string             `json:"user_id" bson:"user_id" validate:"required"`
	Brand            brand              `json:"brand" validate:"required"`
	BodyType         []string           `json:"body_type" bson:"body_type" validate:"required,min=1,max=2oneof=sedan convertible_roadstar sports_supercar cuv_crossover suv_muv hatchback wagon_stationwagon pickup van_minivan bus_minibus truck motorcycle atv"`
	Address          address            `json:"address" validate:"required"`
	Model            model              `json:"model" validate:"required"`
	Milage           milage             `json:"milage" validate:"required"`
	Price            price              `json:"price" validate:"required"`
	Images           []string           `json:"images" validate:"required,min=1,max=20"`
	ContactNo        []string           `json:"contact_no" bson:"contact_no" validate:"required,min=1,max=20"`
	FuelType         string             `json:"fuel_type" bson:"fuel_type" validate:"required,oneof=petrol diesel cng lpg hybrid electric"`
	Color            string             `json:"color" validate:"required"`
	Transmission     string             `json:"transmission" validate:"required,oneof=automatic manual"`
	WheelDrive       string             `json:"wheel_drive" bson:"wheel_drive" validate:"required,oneof=awd 4wd rwd fwd"`
	UsageCondition   string             `json:"usage_condition" bson:"usage_condition" validate:"required,oneof=new used"`
	RegistrationCard string             `json:"registration_card" bson:"registration_card"`
	SellerComments   string             `json:"seller_comments" bson:"seller_comments"`
	SeatCapacity     uint16             `json:"seat_capacity" validate:"max=1000,min=1"`
	Specification    *specification     `json:"specification,omitempty"`
	Active           bool               `json:"active"`
	CreatedAt        time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at" bson:"updated_at"`
}

func CreateAutomobileAdIndexes(client *mongo.Client) {
	col := clients.GetMongoCollection(client, AutomobileAdCollectionName)
	col.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.M{"address.city": 1, "body_type": 1, "brand.name": 1, "price.total_amount": 1},
	})
	col.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.M{"body_type": 1, "address.city": 1, "brand.name": 1, "price.total_amount": 1},
	})
	col.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.M{"brand.name": 1, "address.city": 1, "body_type": 1, "price.total_amount": 1},
	})
}
