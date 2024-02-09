# Write-And-Share

## Project Overview

Building a secure and scalable RESTful API to manage notes. 
Users should be able to perform CRUD operations on notes, share them, and search for notes based on keywords.

### Technical Requirements

- Implement RESTful APIs.
- Use a database for data storage (preferably MongoDB or PostgreSQL).
- Implement any authentication protocol.

### Additional Requirements
- Implement rate limiting and request throttling.
- Implement efficient search functionality, preferably using text indexing for high performance.

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


## Solution

To design the database for the given project, we need to consider the entities involved, their relationships, and the requirements of the API endpoints.


### Database Schema:

1. **Users Table**:
   - **Fields**: user_id (Primary Key), username, email, password_hash
   - This table stores information about registered users.

2. **Notes Table**:
   - **Fields**: note_id (Primary Key), user_id (Foreign Key), title, content, created_at, updated_at
   - Each note is associated with a user through the user_id field.
   - Users can have multiple notes, and each note belongs to a single user.

3. **Notes_To_User_Mapping Table**:
   - **Fields**: id (Primary Key), note_id (Foreign Key), user_id (Foreign Key)
   - This table stores information about notes shared by users with other users.
   - It establishes a many-to-many relationship between notes and users.
     

### Database Indexes:

1. Index on user_id field in the Notes table for efficient retrieval of notes by user.
2. Index on note_id field in the Notes_To_User_Mapping table for quick lookup of shared notes.
3. Index on user_id field in the Notes_To_User_Mapping table for efficient retrieval of notes shared with a particular user.


### Authentication and Authorization:

- Implement authentication using JWT (JSON Web Tokens) for user login and token-based authentication.
- Use rate limiting and request throttling middleware to prevent abuse and protect against denial-of-service attacks.


### Text Indexing and Search:

- Implement full-text indexing on the title and content fields of the Notes table for efficient keyword-based search.
- Use PostgreSQL's built-in full-text search capabilities or MongoDB's text search feature for performing fast and accurate searches.
