# Checkout
Checkout System Task

### Prerequisites
- Go (v1.22.1)
- Docker

### Run Locally

- On first time start up run ```make docker-network``` 

- Then run `make up` to start the app
    - Spins up app image and mySQL images with docker-compose
    - Build app docker image from `./DockerFile`
    - init db from `./scripts/init_db.sql`

- App will be running on `http://localhost:8080`


#### Endpoints 

##### POST /cart

- creates a new cart 
- curl: 
    ``` 
    curl --location 'http://localhost:8080/cart/' \
    --header 'Content-Type: application/json' \
    --data '{
        "customerId": "dffc100d-787b-4ca4-8d29-bf783c855535",
        "lineItems": [
        { 
            "itemSKU": "A",
            "quantity": 5
        },
        { 
            "itemSKU": "B",
            "quantity":  3
        }]}'
    ```

##### Get /cart/{id} (WIP)

- gets carts by id 
- curl: 
    ```
    curl --location 'http://localhost:8080/cart/6b9446ae-7001-4f50-b5bd-3bde7e5d2a2e
    ```