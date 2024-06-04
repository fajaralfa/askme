# ASKME

DESC:

YOU KNOW WHAT IT IS

## API

masih bisa berubah.

- Global request / response

header
```json
{
    "Content-Type": "application/json"
}
```

### POST /api/v1/register

Register

- REQUEST

body
```json
{
    "email": "example@email.com",
    "password": "rahasia"
}
```

- 200 OK

```json
{
    "status": "success",
    "data": null
}
```

- 401 UNAUTHORIZED

```json
{
    "status": "fail",
    "message": "register gagal",
    "data": {
        "email": "email sudah digunakan"
    }
}
```

### POST /api/v1/login

Login

- REQUEST

```json
{
    "email": "example@email.com",
    "password": "rahasia"
}
```

- 200 OK

```json
{
  "status": "success",
  "data": {
    "accessToken": "base64.encoded.jwt",
    "user": {
      "id": 1,
      "email": "example@email.com",
      "photo": ""
    }
  }
}
```

- 401 UNAUTHORIZED

```json
// body
{
    "status": "fail",
    "message": "email atau password salah",
    "data": null
}
```

### GET /api/v1/me

Get current logged in user info

- REQUEST

header

```json
{
    "Authorization": "Bearer {accessToken}"
}
```

- 200 OK

```json
{
    "status": "success",
    "data": {
        "user": {
            "id": 1,
            "email": "example@email.com",
            "photo": "hash.jpg"
        }
    }
}
```

- 401 UNAUTHORIZED

```json
{
    "status": "fail",
    "message": "token invalid" | "anda belum login"
}
```

### GET /api/v1/users/{email}

Get user data with email {email}

- 200 OK

```json
{
    "status": "success",
    "data": {
        "user": {
            "id": 1,
            "email": "example@email.com",
            "photo": "imagename.jpeg",
        }
    },
}
```

- 404 NOT FOUND

```json
{
    "status": "fail",
    "message": "pengguna tidak ditemukan!",
    "data": null
}
```

### POST /api/v1/questions

Send a question to target user

- REQUEST

```json
{
    "question": "pertanyaan?",
    "targetEmail": "example@email.com",
}
```

- 200 OK

```json
{
    "status": "success",
    "data": {
        "question": {
            "id": 1,
            "question": "pertanyaan",
            "created_at": "datetime",
            "user_id": 1
        }
    }
}
```

- 404 NOT FOUND

```json
{
    "status": "fail",
    "message": "pengguna dengan email ini tidak ditemukan"
}
```

### GET /api/v1/questions

Get all questions asked to current authenticated user

- REQUEST

header
```json
{
    "Authorization": "Bearer {accessToken}"
}
```

- 200 OK

```json
{
    "status": "success",
    "data": {
        "questions": [
            {
                "id": 1,
                "question": "question?",
                "created_at": "datetime",
                "user_id": 1
            }
        ],
    }
}
```

- 401 UNAUTHORIZED

```json
{
    "status": "fail",
    "message": "anda belum masuk!",
    "data": null
}
```

### DELETE /api/v1/questions

- REQUEST

header
```json
{
    "Authorization": "Bearer {accessToken}"
}
```

- 200 OK

```json
{
    "status": "success",
    "data": null
}
```

- 401 UNAUTHORIZED
```json
{
  "status": "fail",
  "message": "pertanyaan ini tidak ditanyakan ke kamu, maaf aja ya",
  "data": null
}
```