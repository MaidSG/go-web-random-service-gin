# go-web-random-service-gin
This is a simple Go application that serves as a random arithmetic expression generator. The application is built with the Gin framework and can be easily deployed with Docker.

# Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

**Prerequisites**
---
- Go (version 1.16 or later)
- Docker (version 20.10.5 or later)

**Installing**
  ---
1. Clone the repositor
 ```git clone https://github.com/yourusername/your-repo-name.git```
2. Navigate to the project directory
```cd your-repo-name```
3. Build the Docker image
```docker build -t your-repo-name .```
4. Run the Docker container
```docker run -p 8089:8089 your-repo-name```
The application will be available at http://localhost:8089.

**Built With**
---
Go - The programming language used
Gin - The web framework used
Docker - Used for containerization

**License**
---
This project is licensed under the MIT License - see the LICENSE.md file for details
