// Create User
curl -i -X POST -H "Content-Type: application/json" -d '{
"username":"Budi",
"email":"budikurniawan238@gmail.com",
"phone":"081290858473",
"password":"payphone16"
}' http://localhost:8089/api/v1/users

// Login
curl -i -X POST -H "Content-Type: application/json" -d '{"email":"budikurniawan238@gmail.com","password":"payphone16"}' http://localhost:8080/login

//
https://git.heroku.com/stormy-tundra-37100.git
https://stormy-tundra-37100.herokuapp.com/

CREATE TABLE saved_orders1 (
id integer,
create_dtm timestamp without time zone,
user_id character varying(50),
outlet_id character varying(50),
saved_orders_id character varying(50),
name character varying(50),
phone character varying(20),
orders json,
table_id character varying(20)
);

CREATE TABLE sales1 (
id integer,
create_dtm timestamp without time zone,
sales_id character varying(50),
user_id character varying(50),
outlet_id character varying(50),
sales_type character varying(50),
customer_id character varying(50),
products json,
subtotal integer,
total_diskon integer,
total_bill integer,
payment_method character varying(50),
payment_due_date character varying(50),
total_payment integer,
exchange integer,
notes character varying(100),
total_buy_cost integer,
payment_date character varying(20),
total_tax json,
reward_id character varying(50),
points_redeem integer
);

CREATE TABLE onlinesales1 (
id integer,
create_dtm timestamp without time zone,
sales_id character varying(50),
user_id character varying(50),
outlet_id character varying(50),
customer_id character varying(50),
customer json,
products json,
subtotal integer,
total_diskon integer,
total_tax json,
total_bill integer,
payment_method character varying(50),
payment_account character varying(50),
payment_due_date character varying(50),
total_payment integer,
expedition character varying(50),
service character varying(50),
weight integer,
delivery_cost integer,
notes character varying(100),
total_buy_cost integer,
payment_date character varying(50),
reward_id character varying(50),
points_redeem integer,
order_status character varying(50),
shipment_number character varying(50)
);

type Subscribers1 struct {
ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
UserId string `gorm:"size:50;" json:"user_id"`
OwnerName string `gorm:"size:100;" json:"owner_name"`
FcmToken string `gorm:"size:200;" json:"fcm_token"`
IdcardName string `gorm:"size:50;" json:"idcard_name"`
IdcardNumber string `gorm:"size:50;" json:"idcard_number"`
BankHolderName string `gorm:"size:256;" json:"bank_holder_name"`
BankName string `gorm:"size:256;" json:"bank_name"`
BankAccount string `gorm:"size:256;" json:"bank_account"`
ReferralCode string `gorm:"size:256;" json:"referral_code"`
Email string `gorm:"size:100;" json:"email"`
SecretPassword string `gorm:"size:100;" json:"secret_password,omitempty"`
IdcardImage string `gorm: json:"idcard_image"`
Created_Date time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create_dtm"`
}

type Onlinesales1 struct {
ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create_dtm"`
SalesId string `gorm:"size:50;" json:"sales_id"`
UserId string `gorm:"size:50;" json:"user_id"`
OutletId string `gorm:"size:50;" json:"outlet_id"`
CustomerId string `gorm:"size:50;" json:"customer_id"`
Customer string `gorm:"size:50;" json:"customer"`
Products string `gorm:"not null;" json:"products,"`
Subtotal uint64 `gorm:json:"subtotal,"`
TotalDiskon uint64 `gorm:json:"total_diskon,"`
TotalTax uint64 `gorm:json:"total_tax,"`
TotalBill uint64 `gorm:json:"total_bill,"`
PaymentMethod string `gorm:"size:50;not null;" json:"payment_method"`
PaymentAccount string `gorm:"size:50;not null;" json:"payment_account"`
PaymentDueDate string `gorm:"size:50;not null;" json:"payment_due_date"`
TotalPayment uint64 `gorm:json:"total_payment,"`
Expedition string `gorm:"size:50;not null;" json:"expedition"`
Service string `gorm:"size:50;not null;" json:"service"`
Weight uint64 `gorm:json:"weight,"`
DeliveryCost uint64 `gorm:json:"delivery_cost,"`
Notes string `gorm:"size:100;" json:"notes"`
TotalBuyCost uint64 `gorm:json:"total_buy_cost,"`
PaymentDate string `gorm:"size:100;" json:"payment_date"`
RewardId string `gorm:"size:50;not null;" json:"reward_id"`
PointsRedeem uint64 `gorm:json:"points_redeem,"`
OrderStatus string `gorm:"size:50;not null;" json:"order_status"`
ShipmentNumber string `gorm:"size:50;not null;" json:"shipment_number"`
}
type Sales1 struct {
ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create_dtm"`
SalesId string `gorm:"size:50;" json:"sales_id"`
UserId string `gorm:"size:50;" json:"user_id"`
OutletId string `gorm:"size:50;" json:"outlet_id"`
SalesType string `gorm:"size:50;" json:"sales_type"`
CustomerId string `gorm:"size:50;" json:"customer_id"`
Products string `gorm:"not null;" json:"products,"`
Subtotal uint64 `gorm:json:"subtotal,"`
TotalDiskon uint64 `gorm:json:"total_diskon,"`
TotalBill uint64 `gorm:json:"total_bill,"`
PaymentMethod string `gorm:"size:50;not null;" json:"payment_method"`
PaymentDueDate string `gorm:"size:50;not null;" json:"payment_due_date"`
TotalPayment uint64 `gorm:json:"total_payment,"`
Exchange uint64 `gorm:json:"exchange,"`
Notes string `gorm:"size:100;" json:"notes"`
TotalBuyCost uint64 `gorm:json:"total_buy_cost,"`
PaymentDate string `gorm:"size:100;" json:"payment_date"`
TotalTax uint64 `gorm:json:"total_tax,"`
RewardId string `gorm:"size:50;not null;" json:"reward_id"`
PointsRedeem uint64 `gorm:json:"points_redeem,"`
}
type SavedOrder struct {
ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
UserId string `gorm:"size:50;not null;" json:"user_id"`
OutletId string `gorm:"size:50;not null;" json:"outlet_id"`
Saved_orders_id string `gorm:"size:50;" json:"saved_orders_id"`
Name string `gorm:"size:100;" json:"name"`
Phone string `gorm:"size:20;null;" json:"phone"`
TableId string `gorm:"size:20;null;" json:"table_id"`
CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create_dtm"`
}
