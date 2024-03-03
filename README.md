
# payments-gateway

The "payments-gateway" project is a payment gateway system that facilitates and secures online financial transactions between customers, merchants, and payment service providers. Its main objective is to provide a secure and unified platform for processing payments through various methods and managing cash flow among the involved parties.

## Architecture:

payments-gateway is implemented using the hexagonal architecture pattern, separating the core business logic from the external dependencies. It consists of multiple layers, including the domain layer, application layer, and infrastructure layer. Each layer is designed to be independent and interchangeable, allowing for easy testing and maintenance.

### domain driven design
The project implements Domain-Driven Design (DDD), an approach to software development that emphasizes understanding and modeling the core business domain of an application. In DDD, the focus is on creating a rich and expressive domain model that accurately reflects the business requirements and rules.

### IoC
Inversion of Control (IoC) is a software design principle where control over the execution flow is delegated to a container or framework. Dependency Injection (DI) is a common implementation of IoC, where dependencies are provided externally to a component rather than created internally. This allows for better flexibility, testability, and decoupling of components in the codebase. So, by applying Dependency Injection, you're allowing components to receive their dependencies externally, promoting flexibility and modular design.


## Tech Stack
Go
Fiber
Swagger
Docker




## Run Locally

Clone the project

```bash
  git clone https://github.com/jcmexdev/payments-gateway
```

Go to the project directory

```bash
  cd payments-gateway
```


Prepare Environments File (run make command on terminal)

```bash
make
```
Or copy env files manually

```bash
cp ./payments/.env.example ./payments/.env
```


Install dependencies

```bash
 docker-compose up

```

## Running Tests

To run tests, run the following command

```bash
go test -v ./bank
```
## Documentation

Once you execute compose command you can found swagger docs on

http://localhost:5000/docs

Postman collections can be found in the `payments/docs` folder

## API Reference

### Important Note
The bank API only manages two accounts, each identified by their card number and you should use that numbers to make request:

```
Account with card number "1111-2222-3333-4444":
Customer Name: John Client
Balance: $1000

Account with card number "1234-5678-1234-5678":
Customer Name: John Merchant
Balance: $500
```

### My Bank Api
/health:

Description: Endpoint used to check the system's health status.
Functionality: Allows clients to verify if the system is up and running, available to process requests.

/balance:

Description: Endpoint to retrieve the balance of a bank account.
Functionality: Enables users to retrieve information about the available balance in a specific bank account.

/deposit:

Description: Endpoint to make a deposit into a bank account.
Functionality: Allows users to make deposits into a bank account by specifying the account number and the amount to deposit.

/withdraw:

Description: Endpoint to make a withdrawal from a bank account.
Functionality: Allows users to make withdrawals from a bank account by specifying the account number and the amount to withdraw.

/transfer:

Description: Endpoint to transfer funds between two bank accounts.
Functionality: Allows users to transfer funds from one bank account to another by specifying the source account number, the amount to transfer, and the destination account number.

### Payments Api
/transaction:

Description: Endpoint to create a payment transaction.
Functionality: Allows users to initiate a payment transaction between two parties by specifying the source account, the destination account, and the amount to transfer.

/transaction/{transactionId}:

Description: Endpoint to retrieve details of a specific transaction.
Functionality: Allows users to query detailed information about a specific payment transaction using its unique identifier.

/transaction/{transactionId}/refund:

Description: Endpoint to process a refund for a transaction.
Functionality: Allows users to request a refund for a specific transaction using its unique identifier.


## Feedback
I acknowledge that there are areas where I can improve, particularly in terms of test coverage for my payment API. Unfortunately, due to time constraints, I was only able to focus on implementing tests for the banking functionalities. However, I developed the system with the flexibility to easily add more tests, especially for the payment services.

I structured the architecture of the system with dependency injection in mind. This means that I designed the services responsible for payment functionalities in such a way that allows me to test them in isolation by injecting mock dependencies. With this approach, I can thoroughly test each implementation of the payment services, ensuring their correctness and robustness.

Moving forward, I plan to allocate more time and resources to expand our test coverage to include the payment functionalities. By leveraging dependency injection and building upon the existing testing framework, I aim to enhance the reliability and stability of our payment API.

## Striving for Excellence: Embracing a Comprehensive Approach

In my approach to the project, I intentionally exerted some extra effort with the aim of showcasing my proficiency in the programming language. Typically, for an MVP (Minimum Viable Product), I would have opted for a project structure that isn't as comprehensive as an hexagonal architecture. However, I chose to implement a hexagonal architecture to demonstrate my skills and capabilities in designing robust and scalable software systems. While this decision required more time and effort upfront, I believe it will pay off in the long run by providing a solid foundation for future development and allowing for easier integration of new features and functionalities. Ultimately, my goal was to create a high-quality solution that not only meets the immediate needs of the project but also sets the stage for future growth and expansion.
