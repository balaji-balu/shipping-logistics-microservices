# Shipping Logistics Microservices 

Tech stack
- golang
- ent : ORM
- atlas : database migration
- postgres: db

##### Generate Schema
services\order-service\order>go run -mod=mod entgo.io/ent/cmd/ent new Order Product User
go run -mod=mod entgo.io/ent/cmd/ent new User 

##### Generate ent code 
Run the ent code generation tool to generate the Go code for the schemas.
services\order-service\order>go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/sc
hema