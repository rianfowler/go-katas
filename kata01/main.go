package main

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
