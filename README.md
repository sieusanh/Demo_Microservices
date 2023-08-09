api gateway [ip_address]:8000
account-service 9001
product-service 9002
order-service 9003

===================================================

Enpoints Overview:

+ Common Feature: Authentication, Authorization. 

+ Account Service:
Common: Greeting, Register, Login
User: GetAccountInfo
Admin: Count, Find, FindById, UpdateById, RemoveAll

+ Product Service:
Common: Greeting, Find
User: Create, UpdateByQuery, DeleteById
Admin: AggByCategory, AggByRating, AggByBrand, AggByTime

+ Order Service:
Common: Greeting
User: Create, UpdateById
Admin: Find, FindById, DeleteById
