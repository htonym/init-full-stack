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

- [x] Read configuration from AWS Parameter store
  - Update the config.NewAppConfig to read from parameter store if not running locally
- [ ] Add database
  - [x] Widget detail page
  - [ ] Create widget page
  - [ ] Delete Widget page
  - [ ] Edit Widget page
  - [ ] Write terraform module for db
- [ ] Add basic page that fetches from database
