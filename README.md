# ASKME

DESC:

YOU KNOW WHAT IT IS

## API

REST API Design (Kemungkinan berubah - ubah).

- Global request / response

```json
// header
{
    "Content-Type": "application/json"
}
```

### POST /api/v1/register

Register

- REQUEST

```json
// body
{
    "email": string,
    "password": string
}
```

- 200 OK

```json
{
    "status": "success",
    "data": null,
    "message": "register berhasil, silakan login"
}
```

- 401 UNAUTHORIZED

```json
// body
{
    "status": "fail",
    "data": {
        "email": "email sudah digunakan",
        "password": "password kurang dari 8 karakter"
    },
    "message": "register gagal",
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
    "status": "success",
    "data": {
        "accessToken": string, // encoded jwt
    }
}
```

- 401 UNAUTHORIZED

```json
// body
{
    "status": "fail",
    "message": "email atau password salah"
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
    "status": "success",
    "data": {
        "questions": [
            {
                "id": int,
                "date": date,
                "text": string
            }
        ],
    }
}
```

- 401 UNAUTHORIZED

```json
{
    "status": "fail",
    "message": "anda belum masuk!"
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
    "status": "success",
    "data": {
        "question": {
            "id": int,
            "date": date,
            "text": string,
        }
    }
}
```

- 401 UNAUTHORIZED

```json
{
    "status": "fail",
    "message": "anda belum masuk!"|"pertanyaan ini ditanyakan ke orang lain!"
}
```

### GET /api/v1/user/{email}

Get user data with email {email}

- 200 OK

```json
// body
{
    "status": "success",
    "data": {
        "user": {
            "email": string,
            "photo": string,
        }
    },
}
```

- 404 NOT FOUND

```json
// body
{
    "status": "fail",
    "message": "pengguna tidak ditemukan!"
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
    "status": "success",
    "data": {
        "question": {
            "id": int,
            "date": date,
            "text": string,
        },
        "user": {
            "email": string,
            "photo": string,
        }
    }
}
```

- 404 NOT FOUND

```json
// body
{
    "status": "fail",
    "message": "pengguna tidak ditemukan!"
}
```
