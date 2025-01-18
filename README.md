# Scalable E-commerce Microservice API

This platform handles various aspects of an online store, such as product catalog management, user authentication, shopping cart, payment processing, and order management. Each of these features will be implemented as separate microservices, allowing for independent development, deployment, and scaling.

## Table Of Contents

- [Microservices](#microservices)
- [Technologies used](#technologies-used)
- [Getting Started](#getting-started)

## Microservices

This application contains core services which include;

- **User service**

The user service handles authentication, resetting passwords and user profile.

- **Product service**

The product service takes care of products created by merchants as well as their shops.

- **Cart service**

The cart service allows the buyers or customers on the paltform to add items or products to cart, update cart and even remove products from cart.

- **Payment service**

The payment service handles transactions and payments by buyers for items in a shop. It enables buyers to be able to pay and order items from a shop.

- **Notification service**

The notification service is responsible for sending verification emails to the user. It also handles sending emails to the user updating them of an order.

- **API gateway**

The gateway which is built in NestJS is the entrypoint for the application.

### Technologies Used

- Go
- NestJs for building the API gateway and the product-service
- PostgresSQL as database
- Docker for containerization
- GRPC for building the services
- SQLC for generating Go code from sql queries

## Getting Started

1. Clone this repository.
2. Install dependencies for each service.
3. Configure environment variables (e.g., database connections, RabbitMQ settings).
4. Run the services:

- To start services written in go, run `make server`
- To start services written in NestJS, run `npm run service-name:dev`

[Roadmap URL](https://roadmap.sh/projects/scalable-ecommerce-platform)
