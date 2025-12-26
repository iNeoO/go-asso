# API Quickstart

- Base URL: `http://localhost:5000/api/v1`
- Auth: `Authorization: Bearer <access_token>` or `refresh_token` cookie set by the login endpoint.
- Seed users/emails are defined in `platform/seed/seed.sql` (hashes are for the sample password `password`); run `make reset.db` to start fresh with seeded data.

## Auth

### Login
`POST /auth/login`

```bash
curl -i -c /tmp/planigramme_cookies.txt \
  -H "Content-Type: application/json" \
  -d '{"email":"alice@example.com","password":"password"}' \
  http://localhost:5000/api/v1/auth/login
```

Grab the access token from the JSON (`data.token`) for authenticated calls:

```bash
ACCESS_TOKEN=$(curl -s -c /tmp/planigramme_cookies.txt \
  -H "Content-Type: application/json" \
  -d '{"email":"alice@example.com","password":"password"}' \
  http://localhost:5000/api/v1/auth/login | jq -r '.data.token')
```

## Users (public)

- `GET /users` — list users

```bash
curl http://localhost:5000/api/v1/users
```

- `GET /users/{id}` — fetch one user

```bash
curl http://localhost:5000/api/v1/users/<user_id>
```

- `POST /users` — create a user

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"first_name":"New","last_name":"User","email":"new.user@example.com","password":"MyStrongPass123"}' \
  http://localhost:5000/api/v1/users
```

## Organizations (auth required)

Include either `-H "Authorization: Bearer $ACCESS_TOKEN"` or reuse the cookie jar `-b /tmp/planigramme_cookies.txt`.

- `GET /organizations` — list all organizations

```bash
curl -H "Authorization: Bearer $ACCESS_TOKEN" \
  http://localhost:5000/api/v1/organizations
```

- `GET /organizations/user` — list organizations for the authenticated user

```bash
curl -H "Authorization: Bearer $ACCESS_TOKEN" \
  http://localhost:5000/api/v1/organizations/user
```

- `GET /organizations/{id}` — fetch one organization

```bash
curl -H "Authorization: Bearer $ACCESS_TOKEN" \
  http://localhost:5000/api/v1/organizations/<org_id>
```

- `POST /organizations` — create an organization

```bash
curl -X POST \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"My New Org"}' \
  http://localhost:5000/api/v1/organizations
```

## Notes

- Use `make reset.db` to drop/recreate the database, run migrations, and reseed test data.
- If you prefer cookies over bearer tokens, pass `-c /tmp/planigramme_cookies.txt -b /tmp/planigramme_cookies.txt` on your `curl` commands after logging in.
