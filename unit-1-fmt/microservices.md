# Microservices

Microservices is a software architecture style where an application is built as a collection of small, independent services, each focusing on doing one specific thing well.

Instead of one big "monolithic" app, microservices break it down into multiple smaller apps (services) that work together.

## Key Ideas

### Small, independent units

Each service handles a single business capability (e.g., "User Service," "Payment Service," "Product Service").

You can update or replace a service without touching the rest.

### Independent deployment

Each service can be deployed, scaled, or restarted without affecting others.

### Communication via APIs

Services talk to each other using HTTP/REST, gRPC (google Remote Procedure Call), message queues, etc.

### Separate databases

Each service usually has its own database, avoiding direct sharing to reduce dependency.

### Polyglot freedom

Different services can be written in different programming languages and use different tech stacks if needed.

## Example

Imagine an e-commerce app:

- **User Service** → Manages login, profile, authentication.
- **Product Service** → Handles product catalog.
- **Order Service** → Manages orders.
- **Payment Service** → Processes payments.

Each one runs as its own small app, and they communicate via APIs.

## Benefits

- **Scalability**: Scale just the parts that need it.
- **Faster development**: Teams can work on different services in parallel.
- **Resilience**: If one service fails, the whole system might still function.

## Challenges

- More complex infrastructure (service discovery, monitoring, etc.).
- Communication overhead between services.
- Data consistency issues.
