# Design Document: Talent Profile Page

## Overview

The Talent Profile Page feature extends the JobMatcher platform with comprehensive profile management capabilities for job seekers. The design leverages the existing Go/GraphQL backend and Nuxt 3 frontend architecture, adding new data models, GraphQL operations, and Vue components to support profile creation, editing, viewing, and discovery.

The system follows a layered architecture with clear separation between data persistence (GORM models), business logic (service layer), API interface (GraphQL resolvers), and presentation (Vue components). Profile data is stored in PostgreSQL with images hosted on Cloudinary, maintaining consistency with the existing platform infrastructure.

## Architecture

### System Components

```
┌─────────────────────────────────────────────────────────────┐
│                     Frontend (Nuxt 3)                        │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐      │
│  │ Profile View │  │ Profile Edit │  │ Search Page  │      │
│  │  Component   │  │  Component   │  │  Component   │      │
│  └──────┬───────┘  └──────┬───────┘  └──────┬───────┘      │
│         │                  │                  │              │
│         └──────────────────┼──────────────────┘              │
│                            │                                 │
│                    ┌───────▼────────┐                        │
│                    │  GraphQL Client │                       │
│                    │     (urql)      │                       │
│                    └───────┬────────┘                        │
└────────────────────────────┼──────────────────────────────────┘
                             │
                    ┌────────▼────────┐
                    │   GraphQL API   │
                    └────────┬────────┘
┌────────────────────────────┼──────────────────────────────────┐
│                  Backend (Go)                                  │
│                    ┌───────▼────────┐                         │
│                    │   Resolvers    │                         │
│                    └───────┬────────┘                         │
│                            │                                  │
│                    ┌───────▼────────┐                         │
│                    │ Profile Service│                         │
│                    └───────┬────────┘                         │
│                            │                                  │
│         ┌──────────────────┼──────────────────┐              │
│         │                  │                  │              │
│  ┌──────▼───────┐  ┌──────▼───────┐  ┌──────▼───────┐      │
│  │   Profile    │  │  Cloudinary  │  │ Validation   │      │
│  │  Repository  │  │    Utility   │  │   Service    │      │
│  └──────┬───────┘  └──────────────┘  └──────────────┘      │
│         │                                                    │
│  ┌──────▼───────┐                                           │
│  │  PostgreSQL  │                                           │
│  │   Database   │                                           │
│  └──────────────┘                                           │
└────────────────────────────────────────────────────────────┘
```

### Data Flow

1. **Profile Creation/Update**: User submits form → GraphQL mutation → Service validation → Repository persistence → Database storage
2. **Profile Viewing**: Client requests profile → GraphQL query → Repository fetch → Service enrichment → Response to client
3. **Image Upload**: User uploads file → GraphQL mutation with Upload scalar → Cloudinary upload → URL stored in database
4. **Profile Search**: Client searches → GraphQL query with filters → Repository filtered query → Paginated results

## Components and Interfaces

### Backend Components

#### 1. Data Models (GORM)

**TalentProfile Model** (extends User):
```go
type TalentProfile struct {
    gorm.Model
    UserID              uint      `json:"user_id" gorm:"unique;not null"`
    User                User      `json:"user" gorm:"foreignKey:UserID"`
    ProfessionalTitle   string    `json:"professional_title"`
    Bio                 string    `json:"bio" gorm:"type:text"`
    Location            string    `json:"location"`
    ProfileImageURL     string    `json:"profile_image_url"`
    ViewCount           int       `json:"view_count" gorm:"default:0"`
    CompletionPercent   int       `json:"completion_percent" gorm:"default:0"`
}
```

**Skill Model**:
```go
type Skill struct {
    gorm.Model
    ProfileID       uint      `json:"profile_id" gorm:"not null"`
    SkillName       string    `json:"skill_name" gorm:"not null"`
    ProficiencyLevel string   `json:"proficiency_level"` // Beginner, Intermediate, Advanced, Expert
    Category        string    `json:"category"` // Construction, House Keeping, Repairing, Management, Other
}
```

**Experience Model**:
```go
type Experience struct {
    gorm.Model
    ProfileID       uint      `json:"profile_id" gorm:"not null"`
    JobTitle        string    `json:"job_title" gorm:"not null"`
    CompanyName     string    `json:"company_name" gorm:"not null"`
    StartDate       time.Time `json:"start_date" gorm:"not null"`
    EndDate         *time.Time `json:"end_date"` // Nullable for current positions
    Description     string    `json:"description" gorm:"type:text"`
    IsCurrent       bool      `json:"is_current" gorm:"default:false"`
}
```

**WorkSample Model**:
```go
type WorkSample struct {
    gorm.Model
    ProfileID       uint      `json:"profile_id" gorm:"not null"`
    Title           string    `json:"title"`
    Description     string    `json:"description" gorm:"type:text"`
    ImageURL        string    `json:"image_url" gorm:"not null"`
    ExternalLink    string    `json:"external_link"`
    DisplayOrder    int       `json:"display_order" gorm:"default:0"`
}
```

#### 2. GraphQL Schema Extensions

```graphql
# Input Types
input CreateProfileInput {
  professional_title: String!
  bio: String!
  location: String!
  profile_image: Upload
}

input UpdateProfileInput {
  professional_title: String
  bio: String
  location: String
  profile_image: Upload
}

input AddSkillInput {
  skill_name: String!
  proficiency_level: String!
  category: String!
}

input AddExperienceInput {
  job_title: String!
  company_name: String!
  start_date: DateTime!
  end_date: DateTime
  description: String!
  is_current: Boolean!
}

input AddWorkSampleInput {
  title: String
  description: String
  image: Upload!
  external_link: String
}

input ProfileSearchFilter {
  skills: [String!]
  categories: [String!]
  location: String
  page: Int
  page_size: Int
}

# Output Types
type TalentProfile {
  id: ID!
  user: User!
  professional_title: String!
  bio: String!
  location: String!
  profile_image_url: String
  view_count: Int!
  completion_percent: Int!
  skills: [Skill!]!
  experiences: [Experience!]!
  work_samples: [WorkSample!]!
  created_at: DateTime!
  updated_at: DateTime!
}

type Skill {
  id: ID!
  skill_name: String!
  proficiency_level: String!
  category: String!
}

type Experience {
  id: ID!
  job_title: String!
  company_name: String!
  start_date: DateTime!
  end_date: DateTime
  description: String!
  is_current: Boolean!
}

type WorkSample {
  id: ID!
  title: String
  description: String
  image_url: String!
  external_link: String
  display_order: Int!
}

type ProfileSearchResult {
  profiles: [TalentProfile!]!
  total_count: Int!
  page: Int!
  page_size: Int!
}

# Queries
extend type Query {
  talentProfile(id: ID!): TalentProfile
  myProfile: TalentProfile
  searchProfiles(filter: ProfileSearchFilter!): ProfileSearchResult!
}

# Mutations
extend type Mutation {
  createProfile(input: CreateProfileInput!): TalentProfile!
  updateProfile(input: UpdateProfileInput!): TalentProfile!
  addSkill(input: AddSkillInput!): Skill!
  removeSkill(skill_id: ID!): Boolean!
  addExperience(input: AddExperienceInput!): Experience!
  updateExperience(experience_id: ID!, input: AddExperienceInput!): Experience!
  removeExperience(experience_id: ID!): Boolean!
  addWorkSample(input: AddWorkSampleInput!): WorkSample!
  removeWorkSample(work_sample_id: ID!): Boolean!
}
```

#### 3. Service Layer

**ProfileService Interface**:
```go
type ProfileService interface {
    CreateProfile(userID uint, input CreateProfileInput) (*TalentProfile, error)
    UpdateProfile(userID uint, input UpdateProfileInput) (*TalentProfile, error)
    GetProfile(profileID uint, viewerID *uint) (*TalentProfile, error)
    GetProfileByUserID(userID uint) (*TalentProfile, error)
    SearchProfiles(filter ProfileSearchFilter) (*ProfileSearchResult, error)
    CalculateCompletionPercent(profileID uint) (int, error)
    IncrementViewCount(profileID uint) error
}

type SkillService interface {
    AddSkill(profileID uint, input AddSkillInput) (*Skill, error)
    RemoveSkill(skillID uint, profileID uint) error
    GetSkillsByProfile(profileID uint) ([]*Skill, error)
}

type ExperienceService interface {
    AddExperience(profileID uint, input AddExperienceInput) (*Experience, error)
    UpdateExperience(experienceID uint, profileID uint, input AddExperienceInput) (*Experience, error)
    RemoveExperience(experienceID uint, profileID uint) error
    GetExperiencesByProfile(profileID uint) ([]*Experience, error)
}

type WorkSampleService interface {
    AddWorkSample(profileID uint, input AddWorkSampleInput) (*WorkSample, error)
    RemoveWorkSample(workSampleID uint, profileID uint) error
    GetWorkSamplesByProfile(profileID uint) ([]*WorkSample, error)
}
```

#### 4. Repository Layer

**ProfileRepository Interface**:
```go
type ProfileRepository interface {
    Create(profile *TalentProfile) error
    Update(profile *TalentProfile) error
    FindByID(id uint) (*TalentProfile, error)
    FindByUserID(userID uint) (*TalentProfile, error)
    Search(filter ProfileSearchFilter) ([]*TalentProfile, int, error)
    IncrementViewCount(id uint) error
}
```

### Frontend Components

#### 1. Pages

**`/pages/talent/profile/index.vue`** - Redirects to authenticated user's profile or login
**`/pages/talent/profile/[id].vue`** - Displays talent profile (view mode)
**`/pages/talent/profile/edit.vue`** - Profile editing interface
**`/pages/talents/search.vue`** - Client-facing talent search page

#### 2. Components

**`ProfileHeader.vue`** - Displays profile image, name, title, location
**`ProfileBio.vue`** - Displays bio section with edit capability
**`SkillsList.vue`** - Displays skills organized by category
**`SkillsEditor.vue`** - Interface for adding/removing skills
**`ExperienceTimeline.vue`** - Displays work history in timeline format
**`ExperienceEditor.vue`** - Form for adding/editing experience entries
**`WorkSamplesGallery.vue`** - Grid display of portfolio items
**`WorkSampleUploader.vue`** - Upload interface for work samples
**`ProfileCompletionIndicator.vue`** - Progress bar showing profile completion
**`ProfileSearchFilters.vue`** - Filter controls for talent search
**`ProfileCard.vue`** - Summary card for search results

#### 3. Composables

**`useProfile.ts`** - Profile data fetching and management
```typescript
export const useProfile = (profileId?: string) => {
  const client = useUrqlClient()
  const profile = ref<TalentProfile | null>(null)
  const loading = ref(false)
  const error = ref<Error | null>(null)

  const fetchProfile = async () => { /* ... */ }
  const updateProfile = async (input: UpdateProfileInput) => { /* ... */ }
  
  return { profile, loading, error, fetchProfile, updateProfile }
}
```

**`useSkills.ts`** - Skills management
**`useExperience.ts`** - Experience management
**`useWorkSamples.ts`** - Work samples management

## Data Models

### Database Schema

```sql
-- Talent Profiles Table
CREATE TABLE talent_profiles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    professional_title VARCHAR(255),
    bio TEXT,
    location VARCHAR(255),
    profile_image_url VARCHAR(500),
    view_count INTEGER DEFAULT 0,
    completion_percent INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Skills Table
CREATE TABLE skills (
    id SERIAL PRIMARY KEY,
    profile_id INTEGER NOT NULL REFERENCES talent_profiles(id) ON DELETE CASCADE,
    skill_name VARCHAR(255) NOT NULL,
    proficiency_level VARCHAR(50),
    category VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    UNIQUE(profile_id, skill_name)
);

-- Experiences Table
CREATE TABLE experiences (
    id SERIAL PRIMARY KEY,
    profile_id INTEGER NOT NULL REFERENCES talent_profiles(id) ON DELETE CASCADE,
    job_title VARCHAR(255) NOT NULL,
    company_name VARCHAR(255) NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP,
    description TEXT,
    is_current BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CHECK (end_date IS NULL OR end_date >= start_date)
);

-- Work Samples Table
CREATE TABLE work_samples (
    id SERIAL PRIMARY KEY,
    profile_id INTEGER NOT NULL REFERENCES talent_profiles(id) ON DELETE CASCADE,
    title VARCHAR(255),
    description TEXT,
    image_url VARCHAR(500) NOT NULL,
    external_link VARCHAR(500),
    display_order INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Indexes for performance
CREATE INDEX idx_skills_profile_id ON skills(profile_id);
CREATE INDEX idx_skills_category ON skills(category);
CREATE INDEX idx_experiences_profile_id ON experiences(profile_id);
CREATE INDEX idx_work_samples_profile_id ON work_samples(profile_id);
CREATE INDEX idx_talent_profiles_user_id ON talent_profiles(user_id);
```

### TypeScript Interfaces (Frontend)

```typescript
interface TalentProfile {
  id: string
  user: User
  professionalTitle: string
  bio: string
  location: string
  profileImageUrl?: string
  viewCount: number
  completionPercent: number
  skills: Skill[]
  experiences: Experience[]
  workSamples: WorkSample[]
  createdAt: string
  updatedAt: string
}

interface Skill {
  id: string
  skillName: string
  proficiencyLevel: 'Beginner' | 'Intermediate' | 'Advanced' | 'Expert'
  category: 'Construction' | 'House Keeping' | 'Repairing' | 'Management' | 'Other'
}

interface Experience {
  id: string
  jobTitle: string
  companyName: string
  startDate: string
  endDate?: string
  description: string
  isCurrent: boolean
}

interface WorkSample {
  id: string
  title?: string
  description?: string
  imageUrl: string
  externalLink?: string
  displayOrder: number
}
```


## Correctness Properties

A property is a characteristic or behavior that should hold true across all valid executions of a system—essentially, a formal statement about what the system should do. Properties serve as the bridge between human-readable specifications and machine-verifiable correctness guarantees.

### Property 1: Profile Data Round-Trip Consistency

*For any* valid profile data (including professional title, bio, location), when submitted to the Profile_System and then retrieved, the returned data should match the submitted data exactly.

**Validates: Requirements 1.3, 2.1, 3.1, 4.5**

### Property 2: Required Field Validation

*For any* profile submission where one or more required fields (professional_title, bio, location) are empty or missing, the Profile_System should reject the submission and return a validation error.

**Validates: Requirements 1.2**

### Property 3: Input Sanitization

*For any* text input containing potentially malicious content (XSS payloads, SQL injection attempts, script tags), the Profile_System should sanitize the input before storage and the retrieved value should not contain executable code.

**Validates: Requirements 1.6, 9.1**

### Property 4: Profile Update Preservation

*For any* existing profile, when updated with new data, the updated profile should contain the new data while preserving the original creation timestamp and user ID.

**Validates: Requirements 1.4**

### Property 5: Skill Deduplication (Idempotence)

*For any* skill, adding it to a profile twice should result in the same state as adding it once—the profile should contain exactly one instance of that skill.

**Validates: Requirements 2.3**

### Property 6: Skill Category Validation

*For any* skill submission with a category not in the predefined set (Construction, House Keeping, Repairing, Management, Other), the Profile_System should reject the submission.

**Validates: Requirements 2.5, 9.5**

### Property 7: Skills Organization by Category

*For any* profile with multiple skills, when retrieved and displayed, skills should be grouped by category and within each category ordered by proficiency level (Expert > Advanced > Intermediate > Beginner).

**Validates: Requirements 2.6**

### Property 8: Experience Date Invariant

*For any* experience entry, if an end_date is provided, it must be greater than or equal to the start_date. Submissions violating this invariant should be rejected.

**Validates: Requirements 3.3**

### Property 9: Experience Chronological Ordering

*For any* profile with multiple experience entries, when retrieved, the experiences should be ordered by start_date in descending order (most recent first).

**Validates: Requirements 3.5**

### Property 10: Work Sample File Type Validation

*For any* file upload, if the file type is not JPEG, PNG, or WebP, the Profile_System should reject the upload and return a validation error.

**Validates: Requirements 4.1, 9.2**

### Property 11: Work Sample Storage Round-Trip

*For any* valid work sample upload, the Profile_System should store the image to Cloudinary and save the returned URL, and subsequent retrieval should return a valid accessible image URL.

**Validates: Requirements 4.3**

### Property 12: Resource Deletion Completeness

*For any* work sample or skill, when deleted from a profile, subsequent queries for that profile should not include the deleted item, and for work samples, the associated storage resource should be removed.

**Validates: Requirements 2.4, 4.6**

### Property 13: Profile Retrieval Completeness

*For any* existing profile ID, when requested by an authenticated client, the Profile_System should return all profile sections (bio, skills, experiences, work_samples) populated with their respective data.

**Validates: Requirements 5.1**

### Property 14: Unauthenticated Access Data Filtering

*For any* profile request from a non-authenticated user, the returned profile data should exclude contact information (phone_number, email) while including all other public profile data.

**Validates: Requirements 5.2**

### Property 15: View Count Increment

*For any* profile, each view request should increment the view_count by exactly 1, and the view_count should never decrease.

**Validates: Requirements 5.4**

### Property 16: Profile Completion Calculation

*For any* profile, the completion_percent should equal the sum of: bio present (20%) + skills present (20%) + experience present (20%) + work_samples present (20%) + contact_info present (20%).

**Validates: Requirements 5.6, 10.2**

### Property 17: Search Result Correctness

*For any* search query with skill filters, all returned profiles should contain at least one skill matching the search criteria, and no profiles without matching skills should be returned.

**Validates: Requirements 6.1, 6.2**

### Property 18: Multi-Filter Conjunction

*For any* search with multiple filters (skills AND category AND location), all returned profiles should satisfy ALL filter criteria simultaneously.

**Validates: Requirements 6.3**

### Property 19: Pagination Consistency

*For any* search result set, when paginated with page_size N, page 1 and page 2 should contain no overlapping profiles, and the total count should equal the sum of all profiles across all pages.

**Validates: Requirements 6.4**

### Property 20: Search Result Ordering

*For any* search result set, profiles should be ordered by completion_percent in descending order, with ties broken by updated_at (most recent first).

**Validates: Requirements 6.6**

### Property 21: Profile URL Uniqueness

*For any* two different users, their generated profile URLs should be unique and map to their respective profiles when accessed.

**Validates: Requirements 7.1, 7.3**

### Property 22: Automatic Profile Creation

*For any* new user registration with role "talent", the Profile_System should automatically create an associated empty profile record with completion_percent = 0.

**Validates: Requirements 8.2**

### Property 23: Authorization Enforcement

*For any* user attempting to edit a profile that does not belong to them, the Profile_System should deny access and return an authorization error without modifying any data.

**Validates: Requirements 9.4**

### Property 24: Rate Limiting Enforcement

*For any* user making more than 10 profile update requests within a 1-hour window, the 11th and subsequent requests should be rejected with a rate limit error until the window resets.

**Validates: Requirements 9.3**

### Property 25: Audit Logging Completeness

*For any* profile modification operation (create, update, delete), the Profile_System should create an audit log entry containing the user ID, operation type, timestamp, and affected resource ID.

**Validates: Requirements 9.6**

## Error Handling

### Error Categories

**Validation Errors** (HTTP 400):
- Missing required fields
- Invalid data formats (dates, URLs, email)
- Field length violations (bio > 2000 chars, description > 1000 chars)
- Invalid enum values (category, proficiency_level)
- File type or size violations
- Date constraint violations (end_date < start_date)

**Authentication Errors** (HTTP 401):
- Missing or invalid JWT token
- Expired authentication token
- Unauthenticated access to protected mutations

**Authorization Errors** (HTTP 403):
- Attempting to edit another user's profile
- Attempting to delete another user's resources
- Rate limit exceeded

**Not Found Errors** (HTTP 404):
- Profile ID does not exist
- Skill/Experience/WorkSample ID does not exist
- User ID does not exist

**Conflict Errors** (HTTP 409):
- Duplicate skill addition
- Profile already exists for user

**Server Errors** (HTTP 500):
- Database connection failures
- Cloudinary upload failures
- Unexpected service errors

### Error Response Format

All errors should follow a consistent GraphQL error format:

```json
{
  "errors": [
    {
      "message": "Validation failed: bio exceeds maximum length of 2000 characters",
      "extensions": {
        "code": "VALIDATION_ERROR",
        "field": "bio",
        "constraint": "max_length",
        "limit": 2000
      }
    }
  ]
}
```

### Error Handling Strategy

1. **Input Validation**: Validate all inputs at the GraphQL resolver level before passing to service layer
2. **Service Layer Errors**: Service layer returns typed errors that resolvers translate to GraphQL errors
3. **Database Errors**: Catch and translate database constraint violations to user-friendly messages
4. **External Service Errors**: Implement retry logic for Cloudinary uploads with exponential backoff
5. **Graceful Degradation**: If profile images fail to load, display placeholder images
6. **User Feedback**: Provide specific, actionable error messages to guide users in correcting issues

## Testing Strategy

### Dual Testing Approach

The testing strategy employs both unit tests and property-based tests as complementary approaches:

- **Unit tests**: Verify specific examples, edge cases, and error conditions
- **Property tests**: Verify universal properties across all inputs
- Together they provide comprehensive coverage: unit tests catch concrete bugs, property tests verify general correctness

### Property-Based Testing

**Library Selection**: 
- Backend (Go): Use `gopter` for property-based testing
- Frontend (TypeScript): Use `fast-check` for property-based testing

**Configuration**:
- Minimum 100 iterations per property test
- Each property test must reference its design document property
- Tag format: `// Feature: talent-profile-page, Property {number}: {property_text}`

**Property Test Implementation**:
Each correctness property listed above must be implemented as a single property-based test. For example:

```go
// Feature: talent-profile-page, Property 1: Profile Data Round-Trip Consistency
func TestProfileDataRoundTrip(t *testing.T) {
    properties := gopter.NewProperties(nil)
    properties.Property("profile data round-trip", prop.ForAll(
        func(title, bio, location string) bool {
            // Generate valid profile data
            input := CreateProfileInput{
                ProfessionalTitle: title,
                Bio: bio,
                Location: location,
            }
            // Submit to system
            profile, err := profileService.CreateProfile(userID, input)
            if err != nil {
                return false
            }
            // Retrieve and verify
            retrieved, err := profileService.GetProfile(profile.ID, &userID)
            return err == nil &&
                retrieved.ProfessionalTitle == title &&
                retrieved.Bio == bio &&
                retrieved.Location == location
        },
        gen.AnyString(),
        gen.AnyString(),
        gen.AnyString(),
    ))
    properties.TestingRun(t, gopter.ConsoleReporter(false))
}
```

### Unit Testing

**Focus Areas**:
- Specific examples demonstrating correct behavior
- Edge cases: empty inputs, boundary values (50 skills, 20 work samples, 2000 char bio)
- Error conditions: invalid file types, unauthorized access, rate limiting
- Integration points: Cloudinary uploads, database transactions, GraphQL resolvers

**Unit Test Balance**:
Avoid writing too many unit tests for scenarios covered by property tests. Focus unit tests on:
- Specific error messages and error codes
- Integration with external services (Cloudinary, database)
- UI component rendering and user interactions
- Specific examples from requirements (e.g., current position with no end date)

### Test Coverage Goals

- **Backend**: 80% code coverage minimum
- **Frontend**: 70% code coverage minimum
- **Property Tests**: All 25 correctness properties implemented
- **Integration Tests**: Critical user flows (profile creation, editing, viewing, search)

### Testing Workflow

1. **Development**: Write property tests alongside implementation
2. **Pre-commit**: Run unit tests and property tests locally
3. **CI/CD**: Run full test suite including integration tests
4. **Regression**: Maintain property tests as regression suite for future changes
