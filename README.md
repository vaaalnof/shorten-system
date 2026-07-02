# shorten-url

## Table of Contents

- [About](#about)
- [Architecture](#architecture)
- [Services](#services)
- [Features](#features)
- [Repository Structure](#repository-structure)
- [Technology Stack](#technology-stack)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [Docker Setup](#docker-setup)
- [API Documentation](#api-documentation)

---

## About

### 🇺🇸 English

Shorten System is a URL shortening platform built with Go and a microservices architecture.

The project currently consists of two independent services:

- auth-service — Authentication and authorization service.
- shortener-service — URL shortening, redirection, analytics, reporting, and QR code generation service.

The platform is designed to provide a scalable and extensible foundation for building a modern URL shortener similar to Bitly, TinyURL, or Rebrandly.

### 🇮🇩 Bahasa Indonesia

Shorten System adalah platform pemendek URL yang dibangun menggunakan Go dan arsitektur microservices.

Saat ini proyek terdiri dari dua layanan utama yang berjalan secara terpisah:

- auth-service — Layanan autentikasi dan otorisasi pengguna.
- shortener-service — Layanan pemendek URL, redirect, analytics, reporting, dan pembuatan QR Code.

Platform ini dirancang sebagai fondasi yang scalable dan extensible untuk membangun layanan URL shortener modern seperti Bitly, TinyURL, atau Rebrandly.

---

## Architecture

```text
┌─────────────────┐
│     Client      │
│  Web / Mobile   │
└────────┬────────┘
         │
         ▼
┌──────────────────────┐
│  shortener-service   │
│                      │
│ • URL Shortener      │
│ • URL Redirect       │
│ • QR Code            │
│ • Analytics          │
│ • Reports            │
└──────────┬───────────┘
           │
           │ Validate Token
           ▼
┌──────────────────────┐
│    auth-service      │
│                      │
│ • Authentication     │
│ • Authorization      │
│ • JWT Validation     │
│ • User Management    │
└──────────────────────┘
```

---

## Services

### auth-service

🇺🇸 Authentication and authorization service responsible for managing user identities and securing access to protected resources.

🇮🇩 Layanan autentikasi dan otorisasi yang bertanggung jawab untuk mengelola identitas pengguna serta mengamankan akses ke resource yang dilindungi.

#### Responsibilities

- User registration
- User login
- Access token generation
- Access token validation
- User profile management
- Authorization checks

#### Main Features

- JWT-based authentication
- Secure password hashing
- Protected API access
- User management

---

### shortener-service

🇺🇸 URL shortening service responsible for link management, URL redirection, analytics collection, reporting, and QR code generation.

🇮🇩 Layanan pemendek URL yang bertanggung jawab untuk pengelolaan link, redirect URL, pengumpulan analytics, reporting dashboard, dan pembuatan QR Code.

#### Responsibilities

- Short URL creation
- URL redirection
- Password-protected URLs
- QR Code generation
- Analytics tracking
- Dashboard reporting

#### Main Features

- Custom short aliases
- URL expiration
- URL activation and deactivation
- Password-protected links
- QR Code generation
- Click analytics
- Unique visitor tracking
- Referrer reporting
- Country reporting
- Browser reporting
- Daily click statistics
- Top links dashboard

#### Current Reports

- Summary Report
- Daily Analytics Chart
- Referrer Report
- Country Report
- Browser Report
- Top Links Dashboard

---

## Features

### URL Management

Manage shortened URLs with flexible configuration and security options.

#### Main Features | Fitur Utama

- Create short URLs
- Update existing URLs
- Delete URLs
- Enable or disable URLs
- Custom short aliases
- URL expiration support
- Password-protected URLs

---

### URL Redirection

Fast and reliable URL redirection with built-in caching and security validation.

#### Main Features

- Public URL redirection
- Password verification flow
- Cache-first lookup strategy
- Redirect validation
- Analytics event publishing

---

### Analytics

Collect and process visitor analytics using an event-driven architecture.

#### Main Features

- Click tracking
- Unique visitor tracking
- Referrer tracking
- Country tracking
- Browser tracking
- Device tracking
- Daily analytics aggregation
- Asynchronous analytics processing
- NATS JetStream integration

---

### Reporting Dashboard

Built-in reporting endpoints for monitoring URL performance and audience insights.

#### Main Features

- Summary Report
- Daily Analytics Chart
- Referrer Report
- Country Report
- Browser Report
- Device Report
- Top Links Dashboard

---

### QR Code

Generate branded QR Codes directly from shortened URLs.

#### Main Features

- QR Code generation
- Custom QR styling
- Center logo support
- PNG output

---

### Security

Security features designed to protect users and resources.

#### Main Features

- JWT authentication
- Secure password hashing
- Authorization middleware
- Protected endpoints
- Resource ownership validation

---

### Performance

Optimized for scalability and high-traffic workloads.

#### Main Features

- Redis caching
- Event-driven architecture
- Asynchronous background processing
- Stateless API design
- Database read/write separation ready

---

### Developer Experience

Designed with maintainability and extensibility in mind.

#### Main Features

- Clean Architecture
- Repository Pattern
- Dependency Injection
- Structured Logging
- Configuration Management
- Environment-based setup

---

## Repository Structure

The platform is built using a microservices architecture where each service is maintained independently.

```text
shorten-url/
│
├── auth-service/
│   │
│   ├── cmd/
│   ├── internal/
│   │   ├── config/
│   │   ├── delivery/
│   │   ├── entity/
│   │   ├── exception/
│   │   ├── model/
│   │   ├── repository/
│   │   ├── security/
│   │   ├── usecase/
│   │   └── utils/
│   │
│   └── main.go
│
├── shortener-service/
│   │
│   ├── assets/
│   ├── cmd/
│   ├── internal/
│   │   ├── config/
│   │   ├── delivery/
│   │   ├── entity/
│   │   ├── exception/
│   │   ├── infra/
│   │   ├── model/
│   │   ├── repository/
│   │   ├── security/
│   │   ├── usecase/
│   │   ├── validation/
│   │   └── utils/
│   │
│   └── main.go
│
└── README.md
```

### Repositories

| Repository | Description |
|------------|-------------|
| `auth-service` | Authentication and authorization service responsible for user management, JWT generation, token validation, and access control. |
| `shortener-service` | URL shortening service responsible for URL management, redirection, analytics, reporting, and QR Code generation. |

### Common Directory Overview

| Directory | Description |
|------------|-------------|
| `cmd` | Application entry points and executable commands. |
| `config` | Application configuration, dependency initialization, and bootstrap logic. |
| `delivery` | HTTP controllers, middleware, and route definitions. |
| `entity` | Core business entities used throughout the application. |
| `exception` | Custom application error definitions and error handling utilities. |
| `infra` | Infrastructure integrations such as Redis, NATS, GeoIP, and external services. |
| `model` | Request and response models used by APIs. |
| `repository` | Database access layer and data persistence logic. |
| `security` | Security-related components including password hashing and token handling. |
| `usecase` | Business logic and application workflows. |
| `validation` | Custom validation rules and validators. |
| `utils` | Shared helper functions and reusable utilities. |
| `assets` | Static assets such as GeoIP databases, QR logos, and other resources. |

### Architecture Principles

| Principle | Description |
|------------|-------------|
| Clean Architecture | Business logic remains independent from frameworks and infrastructure. |
| Dependency Injection | Dependencies are initialized externally and injected where needed. |
| Repository Pattern | Data access is abstracted behind repository interfaces. |
| Separation of Concerns | Each layer has a clear and focused responsibility. |
| Event-Driven Analytics | Analytics processing is handled asynchronously through events. |
| Stateless API Design | Services do not store session state, making horizontal scaling easier. |

---

## Technology Stack

The platform is built using modern technologies focused on performance, scalability, and maintainability.

### Backend

| Technology | Purpose |
|------------|---------|
| Go | Primary programming language |
| Fiber | High-performance HTTP web framework |
| Viper | Configuration management |
| Validator | Request validation |
| Logrus | Structured logging |

### Database

| Technology | Purpose |
|------------|---------|
| PostgreSQL | Primary relational database |
| pgx | PostgreSQL driver and connection management |

### Cache

| Technology | Purpose |
|------------|---------|
| Redis | URL caching and performance optimization |

### Messaging & Event Streaming

| Technology | Purpose |
|------------|---------|
| NATS JetStream | Event-driven analytics processing |
| NATS Consumer | Background analytics aggregation |
| NATS Producer | Analytics event publishing |

### Security

| Technology | Purpose |
|------------|---------|
| JWT | Authentication and authorization |
| bcrypt | Password hashing |
| Custom Middleware | Route protection and authorization |

### Analytics

| Technology | Purpose |
|------------|---------|
| GeoLite2 | Country detection based on IP address |
| Event-Driven Architecture | Asynchronous analytics processing |
| Visitor Fingerprinting | Unique visitor identification |

### QR Code

| Technology | Purpose |
|------------|---------|
| go-qrcode | QR Code generation |
| nfnt/resize | Logo resizing and image processing |
| image/png | PNG rendering and export |


### Infrastructure

| Technology | Purpose |
|------------|---------|
| Docker | Containerization |
| Docker Compose | Local development environment |
| GitHub | Source code hosting |
| GitHub Releases | Version management and distribution |


### Architecture & Patterns

| Pattern | Purpose |
|----------|---------|
| Clean Architecture | Separation of business logic and infrastructure |
| Repository Pattern | Data access abstraction |
| Dependency Injection | Loose coupling between components |
| Event-Driven Architecture | Asynchronous processing |
| Stateless API Design | Horizontal scalability |
| Microservices Architecture | Independent service deployment |

---

## Getting Started

...

## Configuration

...

## Docker Setup

...

## API Documentation

### Base URLs

| Service | Base URL |
|----------|----------|
| auth-service | `http://localhost:3000` |
| shortener-service | `http://localhost:3001` |

---

# Auth-service

Authentication and authorization service.

## Authentication

| Method | Endpoint | Description |
|----------|----------|-------------|
| POST | `/api/v1/register` | Register a new user |
| POST | `/api/v1/login` | Authenticate user |
| GET | `/api/v1/validate-token` | Validate access token |
| GET | `/api/v1/me` | Get current authenticated user |
| POST | `/api/v1/refresh-token` | Refresh access token |
| POST | `/api/v1/logout` | Logout current session |

---

### Register

**Endpoint**

```http
POST /api/v1/register
```

**Request**

```http
POST /api/v1/register
Host: localhost:3000
X-Device-ID: android-device-001
Content-Type: application/json

{
  "email": "john.doe@example.com",
  "password": "secret123",
  "confirm_password": "secret123",
  "first_name": "John",
  "last_name": "Doe"
}
```

**Success Response (201 Created)**

```json
{
  "message": "registration success",
  "data": {
    "email": "john.doe@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "created_at": "2026-06-26 14:09:08"
  }
}
```

**Error Response (400 Bad Request)**

```json
{
  "message": "validation failed",
  "errors": {
    "confirm_password": "invalid value"
  }
}
```

**Error Response (409 Conflict)**

```json
{
  "message": "email already registered"
}
```

---

### Login

**Endpoint**

```http
POST /api/v1/login
```

**Request**

```http
POST /api/v1/login
Host: localhost:3000
X-Device-ID: iphone-device-001
X-Device-Type: iphone
User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 18_0 like Mac OS X)
Content-Type: application/json

{
  "email": "john.doe@example.com",
  "password": "secret123"
}
```

**Success Response (200 OK)**

```json
{
  "message": "login success",
  "data": {
    "access_token": "<access_token>",
    "refresh_token": "<refresh_token>"
  }
}
```

**Error Response (400 Bad Request)**

```json
{
  "message": "validation failed",
  "errors": {
    "email": "invalid email format"
  }
}
```

**Error Response (401 Unauthorized)**

```json
{
  "message": "invalid email or password"
}
```

---

### Validate Token

**Endpoint**

```http
GET /api/v1/validate-token
```

**Request**

```http
GET /api/v1/validate-token
Host: localhost:3000
Authorization: <access_token>
```

**Success Response (200 OK)**

```json
{
  "message": "token valid",
  "data": {
    "user_id": "<user_id>",
    "session_id": "<session_id>"
  }
}
```

**Error Response (401 Unauthorized)**

```json
{
  "message": "invalid token"
}
```

---

### Me

**Endpoint**

```http
GET /api/v1/me
```

**Request**

```http
GET /api/v1/me
Host: localhost:3000
Authorization: <access_token>
```

**Success Response (200 OK)**

```json
{
  "message": "profile fetched successfully",
  "data": {
    "id": "b11509f5-f143-4537-8fcc-92c7a63ea480",
    "email": "john.doe@example.com",
    "first_name": "john",
    "last_name": "doe",
    "avatar_url": null,
    "is_active": true,
    "email_verified": false,
    "created_at": "2026-06-12 22:14:55"
  }
}
```

**Error Response (401 Unauthorized)**

```json
{
  "message": "missing authorization token"
}
```

---

### Refresh Token

**Endpoint**

```http
POST /api/v1/refresh-token
```

**Request**

```http
POST /api/v1/refresh-token
Host: localhost:3000
X-Refresh-Token: <refresh_token>
```

**Success Response (200 OK)**

```json
{
  "message": "token refreshed successfully",
  "data": {
    "access_token": "<access_token>",
    "refresh_token": "<refresh_token>"
  }
}
```

**Error Response (401 Unauthorized)**

```json
{
  "message": "missing refresh token"
}
```

---

### Logout

**Endpoint**

```http
POST /api/v1/logout
```

**Request**

```http
POST /api/v1/logout
Host: localhost:3000
Authorization: <access_token>
```

**Success Response (200 OK)**

```json
{
  "message": "logout success"
}
```

**Error Response (401 Unauthorized)**

```json
{
  "message": "session not found"
}
```

---

# Shortener-service

## Public Endpoints

| Method | Endpoint | Description |
|----------|----------|-------------|
| GET | `/:short_code` | Redirect short URL |
| POST | `/:short_code/verify` | Verify password protected URL |

---

### Redirect URL

**Endpoint**

```http
GET /:short_code
```

**Request**

```http
GET /<short_code>
Host: localhost:3001
Referer: https://l.instagram.com/
User-Agent: Mozilla/5.0 (Linux; Android 15; SM-S928B) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/137.0.0.0 Mobile Safari/537.36 Instagram 389.0.0.0.77 Android
X-Forwarded-For: 36.68.0.10
X-Real-IP: 36.68.0.10
X-Forwarded-Proto: https
Accept-Language: id-ID,id;q=0.9,en;q=0.8
```

**Success Response (302 Found)**

```http
HTTP/1.1 302 Found
Location: https://www.youtube.com/
```

**Error Response (403 Forbidden)**

```json
{
  "message": "password required"
}
```

**Error Response (404 Not Found)**

```json
{
  "message": "shorturl not found"
}
```

---

### Verify Password Protected URL

**Endpoint**

```http
POST /:short_code/verify
```

**Request**

```http
POST /<short_code>/verify
Host: localhost:3001
Content-Type: application/json

{
  "password": "oke12345"
}
```

**Success Response (200 OK)**

```json
{
  "message": "password verified successfully",
  "data": "https://www.youtube.com/"
}
```

**Error Response (400 Bad Request)**

```json
{
  "message": "password is required"
}
```

**Error Response (400 Bad Request)**

```json
{
  "message": "invalid password"
}
```

**Error Response (400 Bad Request)**

```json
{
  "message": "shorturl does not have a password"
}
```

**Error Response (404 Not Found)**

```json
{
  "message": "shorturl not found"
}
```

---

## URL Management

| Method | Endpoint | Description |
|----------|----------|-------------|
| POST | `/api/v1/shorten/urls` | Create short URL |
| GET | `/api/v1/shorten/urls` | List URLs |
| GET | `/api/v1/shorten/urls/:id` | Get URL detail |
| PATCH | `/api/v1/shorten/urls/:id` | Update URL |
| PATCH | `/api/v1/shorten/urls/:id/password` | Update URL password |
| DELETE | `/api/v1/shorten/urls/:id/password` | Remove URL password |
| DELETE | `/api/v1/shorten/urls/:id` | Delete URL |
| GET | `/api/v1/shorten/urls/:id/qrcode` | Generate QR Code |

## Reports

| Method | Endpoint | Description |
|----------|----------|-------------|
| GET | `/api/v1/shorten/reports/:id/summary` | Summary report |
| GET | `/api/v1/shorten/reports/:id/chart` | Daily analytics chart |
| GET | `/api/v1/shorten/reports/:id/referrers` | Referrer report |
| GET | `/api/v1/shorten/reports/:id/countries` | Country report |
| GET | `/api/v1/shorten/reports/:id/devices` | Device report |
| GET | `/api/v1/shorten/reports/:id/browsers` | Browser report |
| GET | `/api/v1/shorten/reports/top-links` | Top links dashboard |

---

### Summary Report

**Endpoint**

```http
GET /api/v1/shorten/reports/:id/summary
```

**Request**

```http
GET /api/v1/shorten/reports/<short_id>/summary
Host: localhost:3001
Authorization: <access_token>
```

**Success Response (200 OK)**

```json
{
  "message": "report summary retrieved successfully",
  "data": {
    "total_clicks": 24,
    "unique_visitors": 5,
    "today_clicks": 8,
    "today_unique_visitors": 2
  }
}
```

**Error Response (401 Unauthorized)**

```json
{
  "message": "unauthorized"
}
```

**Error Response (404 Not Found)**

```json
{
  "message": "shorturl not found"
}
```

---

### Daily Analytics Chart

**Endpoint**

```http
GET /api/v1/shorten/reports/:id/chart
```

**Request**

```http
GET /api/v1/shorten/reports/<short_id>/chart
Host: localhost:3001
Authorization: <access_token>
```

**Success Response (200 OK)**

```json
{
  "message": "report chart retrieved successfully",
  "data": {
    "items": [
      {
        "date": "2026-06-20",
        "clicks": 5,
        "unique_visitors": 2
      },
      {
        "date": "2026-06-21",
        "clicks": 8,
        "unique_visitors": 3
      },
      {
        "date": "2026-06-22",
        "clicks": 11,
        "unique_visitors": 4
      }
    ]
  }
}
```

**Error Response (401 Unauthorized)**

```json
{
  "message": "unauthorized"
}
```

**Error Response (404 Not Found)**

```json
{
  "message": "shorturl not found"
}
```

---

### Referrer Report

**Endpoint**

```http
GET /api/v1/shorten/reports/:id/referrers
```

**Request**

```http
GET /api/v1/shorten/reports/<short_id>/referrers
Host: localhost:3001
Authorization: <access_token>
```

**Success Response (200 OK)**

```json
{
  "message": "report referrers retrieved successfully",
  "data": {
    "items": [
      {
        "referrer": "instagram",
        "clicks": 14
      },
      {
        "referrer": "direct",
        "clicks": 10
      }
    ]
  }
}
```

**Error Response (401 Unauthorized)**

```json
{
  "message": "unauthorized"
}
```

**Error Response (404 Not Found)**

```json
{
  "message": "shorturl not found"
}
```

---

### Country Report

**Endpoint**

```http
GET /api/v1/shorten/reports/:id/countries
```

**Request**

```http
GET /api/v1/shorten/reports/<short_id>/countries
Host: localhost:3001
Authorization: <access_token>
```

**Success Response (200 OK)**

```json
{
  "message": "report countries retrieved successfully",
  "data": {
    "items": [
      {
        "country": "Indonesia",
        "clicks": 18
      },
      {
        "country": "Singapore",
        "clicks": 4
      },
      {
        "country": "Other",
        "clicks": 2
      }
    ]
  }
}
```

**Error Response (401 Unauthorized)**

```json
{
  "message": "unauthorized"
}
```

**Error Response (404 Not Found)**

```json
{
  "message": "shorturl not found"
}
```

---

### Device Report

**Endpoint**

```http
GET /api/v1/shorten/reports/:id/devices
```

**Request**

```http
GET /api/v1/shorten/reports/<short_id>/devices
Host: localhost:3001
Authorization: <access_token>
```

**Success Response (200 OK)**

```json
{
  "message": "report devices retrieved successfully",
  "data": {
    "items": [
      {
        "device": "mobile",
        "clicks": 18
      },
      {
        "device": "desktop",
        "clicks": 5
      },
      {
        "device": "tablet",
        "clicks": 1
      }
    ]
  }
}
```

**Error Response (401 Unauthorized)**

```json
{
  "message": "unauthorized"
}
```

**Error Response (404 Not Found)**

```json
{
  "message": "shorturl not found"
}
```

---

### Browser Report

**Endpoint**

```http
GET /api/v1/shorten/reports/:id/browsers
```

**Request**

```http
GET /api/v1/shorten/reports/<short_id>/browsers
Host: localhost:3001
Authorization: <access_token>
```

**Success Response (200 OK)**

```json
{
  "message": "report browsers retrieved successfully",
  "data": {
    "items": [
      {
        "browser": "Chrome",
        "clicks": 15
      },
      {
        "browser": "Safari",
        "clicks": 6
      },
      {
        "browser": "Firefox",
        "clicks": 2
      },
      {
        "browser": "Other",
        "clicks": 1
      }
    ]
  }
}
```

**Error Response (401 Unauthorized)**

```json
{
  "message": "unauthorized"
}
```

**Error Response (404 Not Found)**

```json
{
  "message": "shorturl not found"
}
```

---

### Top Links Dashboard

**Endpoint**

```http
GET /api/v1/shorten/reports/top-links
```

**Request**

```http
GET /api/v1/shorten/reports/top-links
Host: localhost:3001
Authorization: <access_token>
```

**Success Response (200 OK)**

```json
{
  "message": "top links retrieved successfully",
  "data": {
    "items": [
      {
        "id": "232d2a8b-498e-429b-8dd6-ed002bcbaed6",
        "short_code": "test-youtube",
        "original_url": "https://www.youtube.com/",
        "clicks": 24,
        "created_at": "2026-06-25"
      },
      {
        "id": "7d6d3a1b-8b54-4f3e-9f5f-a123456789ab",
        "short_code": "portfolio",
        "original_url": "https://portfolio.example.com",
        "clicks": 18,
        "created_at": "2026-06-24"
      },
      {
        "id": "9f123456-1234-5678-9012-abcdef123456",
        "short_code": "github",
        "original_url": "https://github.com",
        "clicks": 12,
        "created_at": "2026-06-23"
      }
    ]
  }
}
```

**Error Response (401 Unauthorized)**

```json
{
  "message": "unauthorized"
}
```