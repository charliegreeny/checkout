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


### Gotchas 
- Restart app if does not connect to DB 
    - Due to the spinning up of the app needing to connect to database error, occasionally will error out due to database image starting but not ready for connects 
    - **FIX:** 
        - wait ~1 minute for database to fully spin up and run `make up` again,
        - This will only spin up the app image and connect to the DB properly

#### Endpoints 

##### POST /cart

- creates a new cart 
- curl: 
    ``` 
    curl --location 'http://localhost:8080/cart/' \
        --header 'Content-Type: application/json' \
        --data '{
            "customerId": "customer_1",
            "lineItems": [{
                "itemSKU": "A",
                "quantity": 1
                },
                {
                "itemSKU": "B",
                "quantity": 3
                },
                {
                "itemSKU": "C",
                "quantity": 3
                }]
             }'
    ```

##### Get /cart/{id} (WIP)

- gets carts by id 
- curl: 
    ```
    curl --location 'http://localhost:8080/cart/6b9446ae-7001-4f50-b5bd-3bde7e5d2a2e
    ```