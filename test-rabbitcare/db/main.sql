-- 1. List Customer and Product Without Sale
--  Using the UNION operator, in one list return all customers who do not
--  have an invoice and all products that were not sold.
--  For each customer without an invoice, return:
--    the string customer
--    the customer id
--    the customer_name
--  For each product without an invoice, return:
--    the string product
--    the product id
--    the product_name
--  The columns must be in the order shown, but row order does not matter.
-- Table: customer
--    column name
--                      column type
--                                     key / NULL
--   id
--                      int
--                                     PK
--   customer_name
--                      varchar(255)
--   city_id
--                      int
--                                     FK
--   customer_address
--                      varchar(255)
--                      varchar(255)
--   contact_person
--                                     N
--   email
--                      varchar(128)
--   phone
--                      varchar(128)

-- Table: product
--     column name
--                         column type
--                                        key / NULL
--   id
--                         int
--                                        РК
--   sku
--                        varchar(32)
--   product_name
--                        varchar(128)
--   product_description
--                         text
--                         decimal(8,2)
--   current_price
--   quantity_in_stock
--                         int

-- Table: invoice
--    column name
--                      column type
--                                      key / NULL
--   id
--                      int
--                                      PK
--   invoice_number
--                      varchar(255)
--   customer_id
--                      int
--                                      FK
--   user_account_id
--                      int
--   total_price
--                      decimal(8,2)
--   time_issued
--                      varchar
--                                      N
--   time_due
--                      varchar
--                                      N
--   time_paid
--                      varchar
--                                      N
--   time_canceled
--                      varchar
--                                      N
--   time_refunded
--                      varchar
--                                       N
-- invoice.customer_id references customer.id



-- To retrieve the required information, we can use two separate SELECT statements for customers and products, and combine the results using the UNION operator. The first SELECT statement retrieves customers who do not have an invoice, and the second SELECT statement retrieves products that were not sold. Here's the SQL query:

sql
Copy code
SELECT 'customer' AS type, id, customer_name
FROM customer
WHERE id NOT IN (SELECT DISTINCT customer_id FROM invoice)
UNION
SELECT 'product' AS type, id, product_name
FROM product
WHERE id NOT IN (SELECT DISTINCT product_id FROM invoice_item)
In the first SELECT statement, we use a subquery to find all customer IDs that appear in the invoice table, and then retrieve customers whose IDs do not appear in that subquery. In the second SELECT statement, we use a similar subquery to find all product IDs that appear in the invoice_item table, and then retrieve products whose IDs do not appear in that subquery.

Note that we use the keyword AS to rename the type column in each SELECT statement, so that the resulting UNION combines two columns with the same name. Also note that we assume there is an invoice_item table that links invoices to products, based on the information provided.