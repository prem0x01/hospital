##Hospital Management System API

This is a RESTful API for managing hospital users (Doctors and Receptionists) with JWT-based authentication, built using **Go** and **Gin** framework.

---

## Base URL

```

[http://localhost:PORT/api/v1](http://localhost:PORT/api/v1)

````

---

## Authentication Endpoints

### `POST /auth/register`

Register a new user account.

#### Request Body (JSON)

```json
{
  "name": "Prem Mankar",
  "email": "prem@example.com",
  "password": "strongpassword123",
  "role": "doctor"  // or "receptionist"
}
````

#### Success Response

* **Code:** `200 OK`

```json
{
  "message": "User registered successfully"
}
```

#### Error Responses

* **400 Bad Request**

```json
{
  "error": "Email already exists"
}
```

* **422 Unprocessable Entity**

```json
{
  "error": "Missing required fields"
}
```

---

### `POST /auth/login`

Login with existing credentials and receive a JWT token.

#### Request Body (JSON)

```json
{
  "email": "prem@example.com",
  "password": "strongpassword123"
}
```

#### Success Response

* **Code:** `200 OK`

```json
{
  "token": "JWT_TOKEN_HERE"
}
```

#### Error Responses

* **401 Unauthorized**

```json
{
  "error": "Invalid email or password"
}
```

* **422 Unprocessable Entity**

```json
{
  "error": "Missing email or password"
}
```

---

## Securing Routes

To access protected routes, include the token in the **Authorization** header:

```
Authorization: Bearer JWT_TOKEN_HERE
```

---

## Testing with Postman

1. Open Postman and create a `POST` request to:

   ```
   POST http://localhost:5433/api/v1/auth/register
   ```

   * Select **Body > raw > JSON**, and use the registration request body above.

2. Then, make a `POST` request to:

   ```
   POST http://localhost:5433/api/v1/auth/login
   ```

   * Copy the JWT token from the response.

3. For all **protected routes**:

   * Go to the **Authorization tab**
   * Set type: **Bearer Token**
   * Paste the copied token

---

## Tech Stack

* Go
* Gin
* JWT
* PostgreSQL

---

## Created by

**Prem Mankar**
GitHub: [prem0x01](https://github.com/prem0x01)

---

