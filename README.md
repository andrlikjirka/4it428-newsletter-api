# 4it428-newsletter-api
API for a Newsletter platform (microservices architecture) 📰 

> Seminar work for the 4IT428 course

## Project description

This monorepo contains all microservices powering the Newsletter API. 

## How to run the project

### Development environment setup:
1. run `git clone <repository>`
2. run `cd 4it428-newsletter-api`
3. copy environment config `cp .env.sample .env`
4. edit `.env`
   - Obtain your Firebase Admin SDK private key (JSON file) from the Firebase console 
     - (Project settings -> Service accounts -> Generate new private key).
   - Move the downloaded JSON file into `secrets` directory.
   - Set `FIREBASE_CREDENTIALS=./secrets/your-key.json` in `.env` (adjust filename).
   - Set remaining env variables.
5. run: `docker compose -f docker-compose.dev.yml up --build`

### Production environment setup:
1. run `git clone <repository>`
2. run `cd 4it428-newsletter-api`
3. copy environment config `cp .env.sample .env`
4. edit `.env`
    - Obtain your Firebase Admin SDK private key (JSON file) from the Firebase console
      - (Project settings -> Service accounts -> Generate new private key).
    - Move the downloaded JSON file into `secrets` directory.
    - Set `FIREBASE_CREDENTIALS=./secrets/your-key.json` in `.env` (adjust filename).
    - Set remaining env variables.
5. run: `docker compose up`

## Architecture

### Component Diagram
<img src="docs/component-diagram.png" width="589" alt="component diagram">

### Deployment Diagram
<img src="docs/deployment-diagram.png" width="589" alt="deployment diagram">

## Migrations
### Mac
`migrate create -ext sql -dir db/migrations -seq <name_of_migration>`

`migrate -path db/migrations -database "postgres://postgres:admin@localhost:5432/newsletter-db?sslmode=disable" up `
### Windows
`.\go-migrate\migrate.exe create -ext sql -dir db/migrations -seq <name_of_migration>`

`.\go-migrate\migrate.exe -path db/migrations -database "postgres://postgres:admin@localhost:5432/newsletter-db?sslmode=disable" up`
