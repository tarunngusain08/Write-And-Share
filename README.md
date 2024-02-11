# Write-And-Share

## Project Overview

Building a secure and scalable RESTful API to manage notes. 
Users should be able to perform CRUD operations on notes, share them, and search for notes based on keywords.

### Technical Requirements

- Implement RESTful APIs.
- Use a database for data storage (preferably MongoDB or PostgreSQL).
- Implement any authentication protocol with rate limiting and request throttling.
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


### Creating Tables

*Users Table*
The `users` table stores information about users including their unique identifier, username, and password.

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);
```

*Notes Table*
The `notes` table stores information about notes including their unique identifier, title, content, creator user ID, and timestamp.

```sql
CREATE TABLE notes (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    created_by INT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

*Notes to User Mapping Table*
The `notes_to_user_mapping` table establishes a many-to-many relationship between notes and users.

```sql
CREATE TABLE notes_to_user_mapping (
    id SERIAL PRIMARY KEY,
    note_id INT NOT NULL,
    username VARCHAR(255) NOT NULL
);
```

### Insert Dummy Users
```sql
INSERT INTO users (username, password) VALUES
    ('user1', 'password1'),
    ('user2', 'password2'),
    ('user3', 'password3');
```

### Insert Dummy Notes
```sql
INSERT INTO notes (title, content, created_by) VALUES
    ('Note 1', 'Content of note 1', 1),
    ('Note 2', 'Content of note 2', 2),
    ('Note 3', 'Content of note 3', 3);
```

### Insert Dummy Notes to User Mapping
```sql
INSERT INTO notes_to_user_mapping (note_id, username) VALUES
    (1, 'user1'),
    (2, 'user2'),
    (3, 'user3');
```

### Creating a GIN Index
To improve search performance on the `title` column of the `notes` table, you can create a GIN (Generalized Inverted Index) index using the `to_tsvector` function.
```sql
CREATE INDEX title_idx ON notes USING GIN(to_tsvector('english', title));
```
This index will tokenize the `title` column using the English text search configuration and build an inverted index for efficient full-text search operations.

## Explaining a Full-Text Search Query
To understand how PostgreSQL plans to execute a full-text search query on the `notes` table, you can use the `EXPLAIN` command.
```sql
EXPLAIN SELECT id, title FROM notes WHERE to_tsvector('english', title) @@ to_tsquery('english', 'keyword');
```
This command will provide insights into the query execution plan, including how PostgreSQL utilizes the GIN index created on the `title` column for matching the search term 'keyword' using full-text search.
