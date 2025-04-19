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
5. run: `docker compose -f docker-compose.dev.yml up --build`

### Production environment setup:
1. run `git clone <repository>`
2. run `cd 4it428-newsletter-api`
3. copy environment config `cp .env.sample .env`
4. edit `.env`
5. run: `docker compose up`

## Architecture

### Component Diagram
<img src="docs/component-diagram.png" width="589" alt="component diagram">

### Deployment Diagram
<img src="docs/deployment-diagram.png" width="589" alt="deployment diagram">
