package bq

import (
	"encoding/json"
	"time"

	"cloud.google.com/go/bigquery"
)

type Tax struct {
	Name     string `json:"name"`
	Rate     int64  `json:"rate"`
	Cannabis bool   `json:"cannabis"`
	Strict   bool   `json:"strict"`
}

type Modifier struct {
	ID            string  `json:"id"`
	SourceID      string  `json:"source_id"`
	Type          string  `json:"type"`
	Name          string  `json:"name"`
	Amount        int64   `json:"amount"`
	AmountType    string  `json:"amount_type"`
	Coupon        bool    `json:"coupon"`
	Code          string  `json:"code"`
	UseLimit      int64   `json:"uselimit"`
	TotalUseLimit int64   `json:"totaluselimit"`
	Conditional   *string `json:"conditional"`
}

type LoyaltyModifier struct {
	*Modifier
	OrderID    string `json:"order_id"`
	SKU        string `json:"sku"`
	LocationID string `json:"location_id"`
}

type TestingResult struct {
	TestingResultTHC          string `json:"testing_result_thc"`
	TestingResultTHCA         string `json:"testing_result_thca"`
	TestingResultDTHC         string `json:"testing_result_dthc"`
	TestingResultCBD          string `json:"testing_result_cbd"`
	TestingResultCBDA         string `json:"testing_result_cbda"`
	TestingResultCBN          string `json:"testing_result_cbn"`
	TestingResultLabName      string `json:"testing_result_lab_name"`
	TestingResultTestingBatch string `json:"testing_result_testing_batch"`
	TestingResultTestedOn     Date   `json:"testing_result_tested_on" bq:"date"`
}

type Item struct {
	*TestingResult
	Modifiers              []*Modifier        `json:"modifiers"`
	LoyaltyModifiers       []*LoyaltyModifier `json:"loyalty_modifiers"`
	ProductTags            string             `json:"product_tags"`
	SKU                    string             `json:"sku"`
	Qty                    int64              `json:"qty"`
	TargetQty              int64              `json:"target_qty"`
	THC                    int64              `json:"thc"`
	Price                  int64              `json:"price"`
	Vendor                 string             `json:"vendor"`
	VendorLicense          string             `json:"vendor_license"`
	ProductType            string             `json:"product_type"`
	Name                   string             `json:"name"`
	Variant                string             `json:"variant"`
	Type                   string             `json:"type"`
	CPC                    string             `json:"cpc"`
	Cannabis               bool               `json:"cannabis"`
	IngredientsName        string             `json:"ingredients_name"`
	Created                DateTime           `json:"created" bq:"datetime"`
	PesticideChecked       bool               `json:"pesticide_checked"`
	Harvest                string             `json:"harvest"`
	HarvestDate            string             `json:"harvest_date"`
	FinalPrice             int64              `json:"final_price"`
	CostEach               int64              `json:"cost_each"`
	PackageTag             string             `json:"package_tag"`
	TrackingUOM            string             `json:"tracking_uom"`
	Category               string             `json:"category"`
	Subcategory            string             `json:"subcategory"`
	Tax                    int64              `json:"tax"`
	ItemDiscount           int64              `json:"item_discount"`
	CartDiscount           int64              `json:"cart_discount"`
	PackagingCount         int64              `json:"packaging_count"`
	PackagingTHCWeight     int64              `json:"packaging_thc_weight"`
	PackagingConvertWeight int64              `json:"packaging_convert_weight"`
	PackagingUOM           string             `json:"packaging_uom"`
	PricingGroupID         string             `json:"pricing_group_id"`
	PricingGroupName       string             `json:"pricing_group_name"`
	PricingGroupType       string             `json:"pricing_group_type"`
	PricingGroupUOM        string             `json:"pricing_group_uom"`
	LoyaltyPoints          int64              `json:"loyalty_points"`
	LoyaltyCash            int64              `json:"loyalty_cash"`
}

type Register struct {
	RegisterID   string `json:"register_id"`
	RegisterName string `json:"register_name"`
}

type Budtender struct {
	BudtenderID        string `json:"budtender_id"`
	BudtenderFirstName string `json:"budtender_first_name"`
	BudtenderLastName  string `json:"budtender_last_name"`
}

type Employee struct {
	EmployeeID        string `json:"employee_id"`
	EmployeeFirstName string `json:"employee_first_name"`
	EmployeeLastName  string `json:"employee_last_name"`
}

type Patient struct {
	PatientID                string  `json:"patient_id"`
	PatientFirstName         string  `json:"patient_first_name"`
	PatientLastName          string  `json:"patient_last_name"`
	PatientLicense           string  `json:"patient_license"`
	PatientLoyaltyID         *string `json:"patient_loyalty_id"`
	PatientLoyaltyPoints     int64   `json:"patient_loyalty_points"`
	PatientMembershipType    string  `json:"patient_membership_type"`
	PatientMembershipPrimary bool    `json:"patient_membership_primary"`
	PatientTaxExempt         bool    `json:"patient_tax_exempt"`
	PatientPlantCount        int64   `json:"patient_plant_count"`
}

type Caregiver struct {
	CaregiverID string `json:"caregiver_id"`
}

type Payment struct {
	PaymentType         string   `json:"payment_type"`
	PaymentTender       int64    `json:"payment_tender"`
	PaymentTenderCharge int64    `json:"payment_tender_charge"`
	PaymentChange       int64    `json:"payment_change"`
	PaymentCreated      DateTime `json:"payment_created"`
}

type Transaction struct {
	*Employee
	*Budtender
	*Register
	*Patient
	*Payment
	*Caregiver
	OrderID               string      `json:"order_id"`
	TransactionID         string      `json:"transaction_id"`
	Created               DateTime    `json:"created"`
	Completed             DateTime    `json:"completed"`
	CheckedIn             *DateTime   `json:"checked_in"`
	LocationID            string      `json:"location_id"`
	Status                string      `json:"status"`
	Resident              bool        `json:"resident"`
	DiscountCannabis      int64       `json:"discount_cannabis"`
	DiscountNonCannabis   int64       `json:"discount_non_cannabis"`
	Note                  string      `json:"note"`
	THC                   int64       `json:"thc"`
	Tax                   int64       `json:"tax"`
	Medical               bool        `json:"medical"`
	SubtotalCannabis      int64       `json:"subtotal_cannabis"`
	SubtotalNonCannabis   int64       `json:"subtotal_non_cannabis"`
	CustomerGroupID       string      `json:"customer_group_id"`
	CustomerGroupName     string      `json:"customer_group_name"`
	PointsEarned          int64       `json:"points_earned"`
	LoyaltyRedeemedPoints int64       `json:"loyalty_redeemed_points"`
	LoyaltyRedeemedRate   int64       `json:"loyalty_redeemed_rate"`
	ChargeID              *string     `json:"charge_id"`
	OnlineID              *string     `json:"online_id"`
	Modifiers             []*Modifier `json:"modifiers"`
	LoyaltyModifiers      []*Modifier `json:"loyalty_modifiers"`
	Items                 []*Item     `json:"items"`
	Taxes                 []*Tax      `json:"taxes"`
	CannabisTaxes         int64       `json:"cannabis_taxes"`
	NonCannabisTaxes      int64       `json:"non_cannabis_taxes"`
}

func (t *Transaction) Save() (map[string]bigquery.Value, string, error) {
	var d map[string]bigquery.Value
	b, err := json.Marshal(t)
	if err != nil {
		return nil, "", err
	}

	if err := json.Unmarshal(b, &d); err != nil {
		return nil, "", err
	}

	return d, "", nil
}

type Date struct {
	time.Time
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format("2006-01-02"))
}

// func (d *Date) UnmarshalJSON(data []byte) error {
// 	parsed, err := time.Parse("2006-01-02", string(data))
// 	if err != nil {
// 		return err
// 	}
// 	d.Time = parsed
// 	return nil
// }

type DateTime struct {
	time.Time
}

func (d *DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format("2006-01-02 15:04:05"))
}

// func (d *DateTime) UnmarshalJSON(data []byte) error {
// 	parsed, err := time.Parse("2006-01-02 15:04:05", string(data))
// 	if err != nil {
// 		return err
// 	}
// 	d.Time = parsed
// 	return nil
// }
