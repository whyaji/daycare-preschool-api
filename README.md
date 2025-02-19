# Daycare Preschool API

This project is a Daycare Preschool API built using Go, Fiber, MySQL, and JWT for authentication.

## Features

- User authentication with JWT
- CRUD operations for managing daycare and preschool data
- Secure and efficient API endpoints

## Technologies Used

- [Go](https://golang.org/)
- [Fiber](https://gofiber.io/)
- [MySQL](https://www.mysql.com/)
- [JWT](https://jwt.io/)

## Getting Started

### Prerequisites

- Go 1.24.0 or higher
- MySQL database
- Air [Air](https://github.com/air-verse/air)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/whyaji/daycare-preschool-api
   ```
2. Navigate to the project directory:
   ```sh
   cd daycare-preschool-api
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```

### Configuration

1. Create a `.env` file in the root directory based on the `.env.example` file:
   ```sh
   cp .env.example .env
   ```
2. Run script initiate project database:
   ```sh
   go run .\scripts\initiate_project_database\initiate_project_database_main.go
   ```

### Running the Application

1. Run the application using [Air](https://github.com/air-verse/air):

   ```sh
   go run main.go
   ```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

## Contact

For any inquiries, please contact [whyaji](https://github.com/whyaji).
