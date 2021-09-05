package models

//TestData .. is the object used to test data objects CRUD
type TestData struct {
	ID            string  `json:"id,omitempty" bson:"_id"`
	Name          string  `json:"name,omitempty"`
	Price         float32 `json:"price,omitempty"`
	PostalAddress string  `json:"postalAddress,omitempty" bson:"postalAddress"`
}
