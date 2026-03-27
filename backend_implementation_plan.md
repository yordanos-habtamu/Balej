# Balej_Backend // Implementation_Roadmap
> **Mission**: Transition from the high-fidelity mock frontend to a data-driven, Go-powered GraphQL intelligence engine.

## 01. Strategic Priorities // Phase Matrix
We will implement the backend using a modular Go architecture, prioritizing core connectivity first.

| Priority | Objective | Functional Node |
| :--- | :--- | :--- |
| **P0** | **Identity_Relay** | Multi-role Auth (JWT), Session Management |
| **P0** | **Job_Inventory** | CRUD for Jobs, Sector Status management |
| **P1** | **Talent_Sync** | Rich Profile management, Portfolio/CV storage |
| **P1** | **Candidacy_Flow** | Applications, Status updates (Interview/Offers) |
| **P2** | **Admin_Intelligence**| Moderation queue, Audit logs, System stats |
| **P2** | **Secure_Uplink** | Real token generation for secure link relay |

---

## 02. Feature_Decomposed // Backend Mapping

### 1. Identity & Role-Based Access (JWT)
*Frontend /admin, /client, and /talent routes are separated but rely on mock auth.*
- **Action**: Implement JWT middleware with role-claims (`ADMIN`, `CLIENT`, `TALENT`).
- **Storage**: User model with `role` and `is_verified` flags.
- **Goal**: Protect GraphQL resolvers based on the user's neural tier.

### 2. Job_Intelligence Core (CRUD)
*Frontend has Ultra-High Fidelity Job Management (`manage.vue`, `index.vue`).*
- **Action**: Implement `CreateJob`, `UpdateJob`, and `JobSearch` resolvers.
- **Metadata**: Support for `salary_range`, `location_protocol`, and `skill_tags`.
- **Search**: Advanced filtering by department and status (LIVE/ARCHIVE).

### 3. Application_Dynamics (Candidate Flow)
*Frontend features candidate management and invitation workflows.*
- **Action**: Build relational mapping between `Jobs` and `TalentNodes`.
- **States**: Implement status change triggers (Applied → Interviewing → Discovered).
- **Uplink**: Backend support for "Withdraw_Candidacy" logic with impact logs.

### 4. Talent_Portfolio_Sync
*Frontend has complex Profile Edit with Portfolio showcase.*
- **Action**: S3/Storage integration for Resume/CV binary storage.
- **Relational**: Link `Projects` and `Skills` to the Talent profile.
- **Intelligence**: Mock or real "Matching_Index" calculation for Job matches.

### 5. Secure_Relay_Service (Link Security)
*Frontend now generates a mock SecureToken for promoting jobs.*
- **Action**: Backend service to generate cryptographically signed tokens (`HS256`).
- **Relay**: A resolver to "De-authorize" or "Validate" a secure share link upon entry.

---

## 03. Technical_Architecture // Balej_Stack
- **Server**: Go (Golang) v1.21+
- **API Protocol**: GraphQL (Gqlgen) 
- **Database**: PostgreSQL (Relational integrity)
- **Caching**: Redis (For session state and frequent matches)
- **Tokenization**: JWT with custom role-based claims.

---

## 04. Immediate_Next_Steps (Execute_Order_66)
1. **Model_Audit**: Standardize `User`, `Job`, and `Application` schemas in `backend/graph/schema.graphqls`.
2. **Auth_Middleware**: Port the logic from `backend/auth` to the GraphQL context.
3. **Repository_Wireup**: Connect the `repository` layer to the existing PostgreSQL instance.
4. **Talent_Uplink**: Create the first "Real_Data" bridge for the `/talent/profile` page.

---
> SYSTEM_STATUS: // READY_FOR_IMPLEMENTATION
> HASH: {{ Math.random().toString(36).substr(2, 8).toUpperCase() }}
