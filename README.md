# Initial Full Stack App Template

This project is a work in progress with goal of creating a basic full stack application that is under $50/month to run on AWS and can be used as base template for creating full stack applications.

## Tech Stack

### Database

- PostgreSQL
- goose for database migrations

### Backend

- Golang
- chi for routing/middleware

### Frontend

- Golang standard library templates
- HTMX
- Tailwind CSS
- AlpineJS

### CI/CD

- Github Actions
- Docker
- Terraform for setting up AWS resources

## TODO

1. Setup infrastructure enough to deploy single endpoint app
2. Add basic frontend (home page, htmx, static files)
3. Add oauth login using AWS Cognito
4. Add database
5. Add basic page that fetches from database
