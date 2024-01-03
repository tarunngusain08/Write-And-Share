# Write-And-Share

## Project Overview

You have been tasked with building a secure and scalable RESTful API to manage notes. Users should be able to perform CRUD operations on notes, share them, and search for notes based on keywords.

### Technical Requirements

- Implement a RESTful API using a framework (e.g., Express, DRF, Spring).
- Use a database for data storage (preferably MongoDB or PostgreSQL).
- Implement any authentication protocol with rate limiting and request throttling.
- Implement efficient search functionality, preferably using text indexing for high performance.
- Write unit tests and integration tests for API endpoints using a testing framework.

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

### Deliverables

- A GitHub repository with your code.
- A README file with:
  - Details explaining the choice of framework, database, and any 3rd party tools.
  - Instructions on how to run your code and tests.
  - Any necessary setup files or scripts for local or test environment deployment.

### Evaluation Criteria

- Correctness: Does the code meet requirements and work as expected?
- Performance: Does the code use rate limiting and request throttling for high traffic?
- Security: Does the code implement secure authentication and authorization?
- Quality: Is the code well-organized, maintainable, and easy to understand?
- Completeness: Does the code include unit, integration, and end-to-end tests for all endpoints?
- Search Functionality: Does the code implement text indexing and search functionality for keyword-based notes search?
