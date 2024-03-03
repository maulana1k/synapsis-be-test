# Synapsis Backend API Test

This is a REST API built using [**Golang Fiber**](https://gofiber.io/) with **SQLite** database and [**GORM**](https://gorm.io/index.html) as relational mapping. It provides endpoints for various functionalities including 
- Customer can view product list by product category
- Customer can add product to shopping cart
- Customers can see a list of products that have been added to the shopping cart
- Customer can delete product list in shopping cart
- Customers can checkout and make payment transactions
- Login and register customers


## Setup and Running the Program

### Prerequisites

- Go installed on your system. You can download it from [here](https://golang.org/dl/).
- Docker installed on your system. You can download it from [here](https://www.docker.com/products/docker-desktop).

### How to use

1. Clone the repository:

    ```bash
    git clone github.com/maulana1k/synapsis-be-test
    cd synapsis-be-test
    ```
    ```bash
    go mod download
    go run .
    ```

### Run using docker
2. Build the Docker image:

    ```bash
    docker build -t synapsis-backend .
    ```

3. Run the Docker container:

    ```bash
    docker run -d -p 3000:3000 synapsis-backend
    ```

    This will start the API server on port 3000.

### API Spec

The API specification is documented using Swagger. 
https://app.swaggerhub.com/apis/maulana1k/synapsis-be-test/1.0.0