# ASKME

## API

REST API Design (Kemungkinan berubah - ubah).

- Global request / response

```json
// header
{
    "Content-Type": "application/json"
}
```

### POST /api/v1/login

Login

- REQUEST

```json
// body
{
    "email": string,
    "password": string,
}
```

- 200 OK

```json
// body
{
    "accessToken": string, // encoded jwt
    "expireIn": int // time in second
}
```

- 401 UNAUTHORIZED

```json
// body
{
    "error": {
        "messages": [
            "email atau password salah"
        ]
    }
}
```

### POST /api/v1/register

Register

- REQUEST

```json
// body
{
    "email": string,
    "password": string,
    "photo": file|null, // optional
}
```

- 200 OK

```json
{}
```

- 401 UNAUTHORIZED

```json
// body
{
    "error": {
        "messages": [
            "password kurang dari 8 karakter",
            "email telah digunakan",
        ]
    }
}
```

### GET /api/v1/questions

Get all questions asked to current authenticated user

- REQUEST

```json
// header
{
    "Authorization": "Bearer {accessToken}"
}
```

- 200 OK

```json
// body
{
    "questions": [
        {
            "id": int,
            "date": date,
            "text": string
        }
    ],
}
```

- 401 UNAUTHORIZED

```json
{
    "error": {
        "messages": [
            "anda belum masuk!"
        ]
    }
}
```

### GET /api/v1/questions/{id}

Get one question asked to current user

- REQUEST

```json
// header
{
    "Authorization": "Bearer {accessToken}"
}
```

- 200 OK

```json
{
    "question": {
        "id": int,
        "date": date,
        "text": string,
    }
}
```

- 401 UNAUTHORIZED

```json
{
    "error": {
        "messages": [
            "anda belum masuk!",
            "pertanyaan ini ditanyakan ke orang lain!",
        ]
    }
}
```

### GET /api/v1/user/{email}

Get user data with email {email}

- 200 OK

```json
// body
{
    "user": {
        "email": string,
        "photo": string,
    }
}
```

- 404 NOT FOUND

```json
// body
{
    "error": {
        "messages": [
            "pengguna tidak ditemukan!"
        ]
    }
}
```

### POST /api/v1/questions

Send a question to target user

- REQUEST

```json
// body
{
    "question": string,
    "userEmail": string,
}
```

- 200 OK

```json
// body
{
    "question": {
        "id": int,
        "date": date,
        "text": string,
    },
    "user": {
        "email": string,
        "photo": string,
    },
}
```

- 404 NOT FOUND

```json
// body
{
    "error": {
        "messages": [
            "pengguna tidak ditemukan!"
        ]
    }
}
```
