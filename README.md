# music-library-app

Music-library

## Description

This project is a REST API for music storage. It is written in Go and uses PostgreSQL and Docker. The project includes creation, reading with pagination, updating, and deletion of music.


## How to Run


### Step 1 
Сreate a .env file in the root of the project with the following content:
``` .env
DB_PASSWORD=your_postgres_password

PGADMIN_DEFAULT_EMAIL=your_pgadmin_email
PGADMIN_DEFAULT_PASSWORD=your_pgadmin_password
```

### Step 2
Make sure you have Docker and Docker Compose installed. Then run the following command:
```
docker-compose up --build
```

### Step 3
In order to check the database operation go to http://localhost:5050
Swagger UI is on http://localhost:8080/swagger/index.html