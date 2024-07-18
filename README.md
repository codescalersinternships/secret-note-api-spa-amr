# Secret Notes Web Application


This project is a web application for securely sharing self-destructing secret notes. It is built using Go for the backend and Vue for the frontend. The backend provides a RESTful API with Swagger documentation, and the frontend offers a simple, responsive UI. This README overviews the project setup, dependencies, and how to run it using Docker.

## Features
- **User Authentication:** Sign up and sign in functionality.
- **Note Creation:** Create notes with text content, maximum view limits, and expiration time.
- **Note Viewing:** Retrieve and view notes securely using unique URLs.
- **Note Self-Destruction:** Notes self-destruct after the specified number of views or expiration time.
- **Security:** Secure URL generation for every note.
- **Rate Limiting:** Implements rate limiting to prevent abuse.

## Requirements
- Go (1.16 or later)
- Node.js (14.x or later) with npm
- Docker and Docker Compose (for containerized setup)

## Installation
1. **Clone the repository:**
``` bash
git clone github.com/codescalersinternships/secret-note-api-spa-amr
```
2. **Build and run the Docker containers:**
``` bash
docker-compose up --build
```
This command will build the Docker images and start the containers for the application and necessary services.

3. **Access the application:**

Once the containers are up and running. 
- access the web application at [http://localhost:8090/](http://localhost:8090/)
- Access the backend server at
[http://localhost:8000/](http://localhost:8000/)


## Usage

### Signing Up

1. **Navigate to the Sign-Up Page**:
    - Visit `http://localhost:3000/signup` in your web browser.
2. **Fill in the Sign-Up Form**:
    - Enter your desired username and password.
3. **Submit the Form**:
    - Click the "Sign Up" button to create your account.
4. **Redirection to Sign-In Page**:
    - After a successful sign-up, you will be redirected to the sign-in page.

### Signing In

1. **Navigate to the Sign-In Page**:
    - Visit `http://localhost:3000/signin` in your web browser.
2. **Fill in the Sign-In Form**:
    - Enter your username and password.
3. **Submit the Form**:
    - Click the "Sign In" button to log into your account.
4. **Redirection to Note Creation Page**:
    - After a successful sign-in, you will be redirected to the note creation page.

### Creating a Note

1. **Navigate to the Note Creation Page**:
    - If you are not already there, visit `http://localhost:3000/create`.
2. **Fill in the Note Form**:
    - Enter your note content, maximum views, and expiration time.
3. **Submit the Form**:
    - Click the "Create Note" button to generate a unique URL for your note.
4. **View the Note URL**:
    - A unique URL for your note will be displayed. You can share this URL with others.

### Viewing a Note

1. **Visit the Note URL**:
    - Open the unique URL generated during note creation in your web browser.
2. **View the Note**:
    - The content of the note will be displayed. If the note has expired or reached its view limit, you will see an appropriate message.



## API Documentation
Swagger documentation is available at http://localhost:8000/swagger/index.html when the backend server is running.

## Tests
### Running Backend Tests:
Navigate to backend directory and run the tests:
```
cd backend
go test ./...
```
### Running Cypress End-to-End Tests

1. **Install Cypress**:
    - Ensure you have Cypress installed as a development dependency:
    ```bash
    npm install cypress --save-dev
    ```

2. **Open Cypress Test Runner**:
    - Run the following command to open the Cypress Test Runner:
    ```bash
    npx cypress open
    ```

3. **Run the Tests**:
    - In the Cypress Test Runner, click on the test spec file to run the tests. Cypress will launch a browser and execute the tests.

## Contribution
Feel free to open issues or submit pull requests with improvements.

