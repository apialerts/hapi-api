# Hapi API

Hey look, an API written in GO on Google Cloud Run with a Postgres Database.

## Environment Variables

#### Local Development
Use the Google Cloud SQL Auth Proxy

```bash
DB_USER=hapi
DB_PASSWORD=xxxx
```

#### Production Cloud Run
Use the Google Cloud SQL Connection String via UNIX sockets

```bash
CLOUD_SQL_INSTANCE=xxx
DB_USER=hapi
DB_PASSWORD=xxxx
```

## Local Deploy

Use gcloud auth to authenticate with Google Cloud

```bash
gcloud run deploy api \
    --source . \
    --project hapi-9cd84 \
    --region us-central1 \
    --allow-unauthenticated
```

## Public API

// Todo - secure export request

### Export

POST http://localhost:8080/export

Request body:
```json
{
  "endpoints": [
    {
      "id": 1,
      "name": "API Alerts",
      "url": "https://api.apialerts.com/health/cron",
      "method": "GET",
      "isActive": true,
      "slowThresholdMs": 500,
      "createdAt": 0
    },
    {
      "id": 2,
      "name": "Stripe",
      "url": "https://api.stripe.com/healthcheck",
      "method": "GET",
      "isActive": true,
      "slowThresholdMs": 200,
      "createdAt": 0
    }
  ]
}
```

Response Body:
```json
{
    "code": "5OMRDTSE",
    "expiresAt": "2025-12-29T12:42:34.058193313Z"
}
```

### Import

GET http://localhost:8080/import/5OMRDTSE

Response:
```json
{
  "endpoints": [
    {
      "id": 1,
      "name": "API Alerts",
      "url": "https://api.apialerts.com/health/cron",
      "method": "GET",
      "isActive": true,
      "slowThresholdMs": 500,
      "createdAt": 0
    },
    {
      "id": 2,
      "name": "Stripe",
      "url": "https://api.stripe.com/healthcheck",
      "method": "GET",
      "isActive": true,
      "slowThresholdMs": 200,
      "createdAt": 0
    }
  ]
}
```