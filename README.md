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
