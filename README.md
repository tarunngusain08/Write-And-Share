# Write-And-Share

## Project Overview

Building a secure and scalable RESTful API to manage notes. 
Users should be able to perform CRUD operations on notes, share them, and search for notes based on keywords.

### Technical Requirements

- Implement RESTful APIs.
- Use a database for data storage (preferably MongoDB or PostgreSQL).
- Implement any authentication protocol with rate limiting and request throttling.
- Implement efficient search functionality, preferably using text indexing for high performance.
- Write unit tests for API endpoints using a testing framework.

### API Endpoints

*Authentication Endpoints:*

- `POST /api/auth/signup`: Create a new user account.
- `POST /api/auth/login`: Log in to an existing user account and receive an access token.

*Note Endpoints:*

- `GET /api/notes`: Get a list of all notes for the authenticated user.
- `GET /api/notes/:id`: Get a note by ID for the authenticated user.
- `POST /api/notes`: Create a new note for the authenticated user.
- `PUT /api/notes/:id`: Update an existing note by ID for the authenticated user.
- `DELETE /api/notes/:id`: Delete a note by ID for the authenticated user.
- `POST /api/notes/:id/share`: Share a note with another user for the authenticated user.
- `GET /api/search?q=:query`: Search for notes based on keywords for the authenticated user.
