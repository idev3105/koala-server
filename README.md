# Movie App Backend

Welcome to the backend of the Movie App! This project provides the server-side functionality for managing movies, users, and reviews.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Technologies](#technologies)
- [Architecture](#architecture)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Introduction

The Movie App Backend is a RESTful API built to handle the data and business logic for a movie application. It allows users to browse movies, read reviews, and manage their profiles.

## Features

- User authentication and authorization
- CRUD operations for movies
- User reviews and ratings
- Search and filter movies
- Secure and scalable architecture

## Technologies

- **Programming Language:** [Node.js](https://nodejs.org/)
- **Framework:** [Express.js](https://expressjs.com/)
- **Database:** [MongoDB](https://www.mongodb.com/)
- **Authentication:** [JWT](https://jwt.io/)
- **Environment Management:** [dotenv](https://www.npmjs.com/package/dotenv)

## Architecture

The project follows a clean architecture approach, with the following layers:

- **API Layer:** Handles HTTP requests and responses.
- **Service Layer:** Contains business logic.
- **Repository Layer:** Handles data access.
- **Model Layer:** Defines the data structure and rules.

## Installation

To get a local copy up and running, follow these steps:

1. **Clone the repository:**

   ```sh
   git clone https://github.com/your-username/movie-app-backend.git
   ```

2. **Navigate to the project directory:**

   ```sh
   cd movie-app-backend
   ```

3. **Install dependencies:**

   ```sh
   npm install
   ```

4. **Set up environment variables:**
   Create a `.env` file in the root directory and add the following:

   ```env
   PORT=5000
   MONGODB_URI=your_mongodb_uri
   JWT_SECRET=your_jwt_secret
   ```

5. **Start the server:**
   ```sh
   npm start
   ```

## Usage

Once the server is running, you can access the API at `http://localhost:5000`.

## API Endpoints

Here are some of the main endpoints available in the API:

- **Movies:**

  - `GET /api/movies` - Get all movies
  - `GET /api/movies/:id` - Get a single movie by ID
  - `POST /api/movies` - Create a new movie
  - `PUT /api/movies/:id` - Update a movie by ID
  - `DELETE /api/movies/:id` - Delete a movie by ID

- **Users:**

  - `POST /api/users/register` - Register a new user
  - `POST /api/users/login` - Login a user
  - `GET /api/users/profile` - Get user profile (requires authentication)

- **Reviews:**
  - `POST /api/movies/:id/reviews` - Add a review to a movie (requires authentication)
  - `GET /api/movies/:id/reviews` - Get all reviews for a movie

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Your Name - [your-email@example.com](mailto:your-email@example.com)

Project Link: [https://github.com/your-username/movie-app-backend](https://github.com/your-username/movie-app-backend)
