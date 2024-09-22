# Switch development
- Example: https://gofr.dev/docs/quick-start/configuration
``` APP_ENV=dev go run main.go ```


# Redis configuration
 - Optional: Remove existing redis (if we are have already)
``` sudo docker rm gofr-redis```
- Create new redis container in docker 
    ``` sudo docker run --name gofr-redis -p 6379:6379 -d redis ```

# Database configuration

1 - create new postgres container with docker
    ``` sudo docker run --name gofr-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=root123 -e POSTGRES_DB=test_db -p 5433:5432 -d postgres:15 ```

- 1.1: create new customers table
    ``` docker exec -it gofr-postgres psql -U postgres -d test_db -c "CREATE TABLE customers (id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL);" ```
- 1.2: Insert one record to customers table
    ``` docker exec -it gofr-postgres psql -U postgres -d test_db -c "INSERT INTO customers (name) VALUES ('John Doe');" ```
- 1.3 : Optional: Select a customer
    ``` docker exec -it gofr-postgres psql -U postgres -d test_db -c "SELECT * FROM customers;" ```


# Connection
- 1: Check running containers
    ``` sudo docker ps ```, or ``` sudo docker ps -a ```


# Example CRUD 
- https://github.com/gofr-dev/gofr/tree/main/examples/using-add-rest-handlers

# Articles
- 1: https://levelup.gitconnected.com/why-use-gofr-for-golang-backend-c6489640ee91
- 2: https://nyadgar.com/posts/go-profiling-like-a-pro/?ref=dailydev#summary
- 3: Channel: https://www.dolthub.com/blog/2024-06-21-channel-three-ways/?ref=dailydev
- 4: Patterns: https://blog.algomaster.io/p/20-patterns-to-master-dynamic-programming?ref=dailydev