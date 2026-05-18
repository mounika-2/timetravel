# TimeTravel Underwriting Platform

A temporal business underwriting platform built with:

* Go backend
* React frontend
* SQLite persistence
* Historical version reconstruction
* Immutable audit snapshots
* Dynamic underwriting records
* Underwriting change analysis

The system allows users to:

* Load business underwriting records
* View historical snapshots over time
* Reconstruct prior business states
* Edit historical versions
* Create new immutable versions
* Analyze operational and underwriting changes between versions

---

# Why This Project Exists

Insurance underwriting systems often need:

* historical auditing
* temporal reconstruction
* compliance tracking
* operational change analysis
* immutable record history

Traditional CRUD systems overwrite prior state.

This project demonstrates a temporal architecture where:

* every update creates a historical snapshot
* prior versions remain immutable
* business state can be reconstructed at any point in time
* underwriting changes can be analyzed over time

---

# Architecture

## Backend

* Language: Go
* Router: Gorilla Mux
* Database: SQLite
* API Style: REST

## Frontend

* React
* Dynamic underwriting form rendering
* Timeline-based UI
* Snapshot viewer
* Historical analysis tools

## Persistence Layer

SQLite stores:

* current business records
* immutable historical snapshots

---

# Core Features

## 1. Temporal Versioning

Every record update creates a new immutable snapshot.

Snapshots preserve:

* business attributes
* underwriting values
* operational changes
* timestamps

Example:

| Version | Employee Count | Payroll   |
| ------- | -------------- | --------- |
| v1      | 10             | 500,000   |
| v2      | 15             | 750,000   |
| v3      | 25             | 1,500,000 |

Users can reconstruct historical business state at any time.

---

## 2. Historical Timeline

The UI displays all historical versions sorted by date.

Users can:

* inspect historical snapshots
* compare versions
* edit historical state
* create new versions from prior snapshots

---

## 3. Immutable Audit History

Historical versions are never overwritten.

Every update creates:

* a new version
* a new timestamped snapshot
* preserved audit history

This models real enterprise underwriting systems.

---

## 4. Dynamic Schema Support

The frontend dynamically renders underwriting fields.

This allows:

* flexible underwriting data
* evolving schemas
* additional risk attributes without frontend rewrites

Examples of supported fields:

* employee_count
* annual_payroll
* hazardous_materials
* delivery_services
* overnight_baking
* business_hours
* general_liability_limit

The system is no longer tied to a fixed schema.

---

## 5. Underwriting Change Analysis

Users can analyze differences between two historical versions.

The system detects:

* changed fields
* added fields
* removed fields
* operational risk changes

Example analysis:

* employee_count changed from 10 to 25
* payroll exposure increased
* hazardous materials added
* delivery operations introduced

The analyzer also produces underwriting-oriented commentary.

---

# API Endpoints

## Load Current Record

```http
GET /api/v2/records/{id}
```

---

## Create or Update Record

```http
POST /api/v2/records/{id}
```

Example payload:

```json
{
  "employee_count": "25",
  "annual_payroll": "1500000",
  "delivery_services": "yes"
}
```

---

## List Historical Versions

```http
GET /api/v2/records/{id}/versions
```

---

## Load Specific Historical Snapshot

```http
GET /api/v2/records/{id}/versions/{version}
```

---

## Analyze Changes Between Versions

```http
POST /api/v2/records/{id}/analyze
```

Request body:

```json
{
  "from_version": 1,
  "to_version": 5
}
```

---

# Database Design

## records

Stores latest/current business state.

| Column     | Type     |
| ---------- | -------- |
| id         | INTEGER  |
| data       | TEXT     |
| created_at | DATETIME |

---

## record_versions

Stores immutable historical snapshots.

| Column     | Type     |
| ---------- | -------- |
| record_id  | INTEGER  |
| version    | INTEGER  |
| data       | TEXT     |
| created_at | DATETIME |

---

# Frontend Panels

The UI is organized into four main panels:

## 1. Historical Timeline

Displays chronological underwriting snapshots.

---

## 2. Snapshot Viewer

Displays full historical underwriting record data.

---

## 3. Update Existing Record

Allows editing a selected historical snapshot and creating a new immutable version.

---

## 4. Underwriting Analysis

Analyzes operational and underwriting changes between versions.

---

# Backward Compatibility

The system evolved from:

* fixed underwriting schema
* basic CRUD operations

into:

* dynamic underwriting records
* temporal versioning
* historical reconstruction

Backward compatibility was maintained through:

* API versioning
* non-breaking route evolution
* preserved REST structure

---

# Example Business Evolution

## Acme Plumbing LLC

Over time:

* employee count increased
* payroll exposure increased
* business hours expanded
* delivery operations introduced
* hazardous materials added
* liability limits increased

The system captures each stage historically.

---

# Local Development

## Backend

```bash
go run .
```

Server runs on:

```text
http://localhost:8000
```

---

## Frontend

```bash
npm install
npm run dev
```

Frontend runs on:

```text
http://localhost:5173
```

---

# Database Reset

Delete local SQLite database:

```bash
rm records.db
```

Restart backend:

```bash
go run .
```

Reseed data:

```bash
./seed.sh
```

---

# Demo Workflow

1. Load Record 1
2. View historical timeline
3. Select a historical snapshot
4. Edit underwriting fields
5. Create new immutable version
6. Compare two versions
7. Run underwriting analysis

---

# Technical Concepts Demonstrated

* Temporal data modeling
* Immutable audit systems
* Historical reconstruction
* REST API design
* Dynamic schema rendering
* Versioned persistence
* SQLite integration
* React state management
* Backend/frontend integration
* Underwriting risk analysis

---

# Future Enhancements

Potential future improvements:

* Authentication
* Multi-user access
* Real AI/LLM integration
* Diff visualizations
* Advanced underwriting scoring
* Role-based permissions
* Search/filtering
* Production database support
* Cloud deployment
* Event sourcing

---

# Project Status

Current status:

* Temporal versioning implemented
* Historical reconstruction implemented
* Immutable snapshots implemented
* Dynamic underwriting schema implemented
* Underwriting analysis implemented
* React UI implemented
* SQLite persistence implemented
* Seed data implemented
* API versioning implemented

The platform now functions as a complete temporal underwriting prototype.
