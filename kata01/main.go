package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type OfferType string

const (
	BuyOneGetOneFree OfferType = "BuyOneGetOneFree"
	MultiBuyDiscount OfferType = "MutliBuyDiscount"
	PercentagOff     OfferType = "PercentageOff"
)

type Item struct {
	ID            string `yaml:"id"`
	Name          string `yaml:"name"`
	PriceInCents  int    `yaml:"priceInCents"`
	OfferGroupsID string `yaml:"offerGroupId"`
}

type OfferGroup struct {
	GroupId string    `yaml:"groupId"`
	Type    OfferType `yaml:"type"`
}

type Cart struct {
	Items map[string]int `yaml:"items"`
}

type Config struct {
	Items       []Item       `yaml:"items"`
	OfferGroups []OfferGroup `yaml:"offerGroups"`
	Cart        Cart         `yaml:"cart"`
}

// LoadItems extracts items from the configuration.
func LoadItems(config *Config) []Item {
	return config.Items
}

// LoadOfferGroups extracts offer groups from the configuration.
func LoadOfferGroups(config *Config) []OfferGroup {
	return config.OfferGroups
}

// LoadCart extracts the cart from the configuration.
func LoadCart(config *Config) Cart {
	return config.Cart
}

// LoadConfig reads the YAML file and unmarshals it into a Config struct.
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func main() {
	// Load the configuration from a YAML file.
	config, err := LoadConfig("samples/cart01.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Log the loaded configuration pieces.
	log.Printf("Items: %+v", LoadItems(config))
	log.Printf("OfferGroups: %+v", LoadOfferGroups(config))
	log.Printf("Cart: %+v", LoadCart(config))
}
