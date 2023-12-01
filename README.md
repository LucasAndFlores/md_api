# MD API

This API allows you to perform the following actions:

- Create a new user
- Authenticate a user
- Utilize the `/audio` routes to store, list all stored audios, and stream a specific audio.

### How to Run

1. Create a `.env` file using the variables from `.env.example`.
2. Run the following command:

```bash
docker-compose up
```

#### Next steps for the project:

[] Implement encrypt and decrypt functions for sensitive data (e.g., email, password, and cellphone).
[] Add unit and end-to-end (E2E) tests for user-related functionalities.
[] Implement the `/auth/` route to generate JWT tokens.
