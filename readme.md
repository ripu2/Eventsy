# Eventsy

## About
Planning an event but tired of the chaos? Say hello to **Eventsy**—your go-to event management solution! Whether it's a small meetup or a large gathering, handling registrations, tracking attendees, and managing event details can be a headache. **Eventsy makes it effortless.**

With just a few clicks, you can:
✅ Create and customize events
✅ Manage attendees with ease
✅ Allow users to register and join seamlessly

Built with Go, Eventsy provides a robust and scalable backend to ensure your event runs smoothly. No more juggling spreadsheets—focus on making your event a success!

## Tech Stack
- **Backend:** Go (Golang)
- **Framework:** net/http
- **Database:** PostgreSQL / MySQL (Specify accordingly)
- **Authentication:** JWT / OAuth (If applicable)

## Running Locally

### Prerequisites
- **Go** (Ensure you have Go installed) - [Download Go](https://golang.org/dl/)
- **Database** (Ensure your chosen database is installed and running)

### Steps to Run

1. **Clone the repository:**
   ```sh
   git clone https://github.com/ripu2/Eventsy.git
   cd Eventsy
   ```

2. **Set up environment variables:**
   - Create a `.env` file in the root directory and add the necessary environment variables.
   - Refer to `.env.example` for guidance.

3. **Install dependencies:**
   ```sh
   go mod tidy
   ```

4. **Run the server:**
   ```sh
   go run cmd/server/main.go
   ```
   The server should now be running at `http://localhost:8080` (or your specified port).

## Environment Variables
Ensure you set the following in your `.env` file:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=eventsy_db
SERVER_PORT=8080
```
Modify these according to your local setup.

## Postman Collection
To test API endpoints, import the provided Postman collection (`Eventsy.postman_collection.json`) from the repository.

**⚠ Important:** Before using, replace the base URL in Postman requests with your local server:
```
http://localhost:<your-port>
```

## Contributing
Contributions are welcome! Please fork the repository and create a pull request with your changes.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
