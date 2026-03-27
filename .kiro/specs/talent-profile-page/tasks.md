# Implementation Plan: Talent Profile Page

## Overview

This implementation plan breaks down the talent profile page feature into incremental coding tasks. The implementation follows a bottom-up approach: database models → repository layer → service layer → GraphQL API → frontend components. Each task builds on previous work, with testing integrated throughout to catch errors early.

## Tasks

- [-] 1. Set up database models and migrations
  - Create GORM models for TalentProfile, Skill, Experience, and WorkSample in `backend/models/profileModels.go`
  - Add foreign key relationships and constraints (CASCADE delete, date validation)
  - Create database migration file for new tables with indexes
  - Run migration to create tables in PostgreSQL
  - _Requirements: 1.1, 2.1, 3.1, 4.1, 8.1, 8.2_

- [ ]* 1.1 Write property test for profile model
  - **Property 1: Profile Data Round-Trip Consistency**
  - **Validates: Requirements 1.3**

- [ ] 2. Implement profile repository layer
  - [~] 2.1 Create ProfileRepository interface and implementation in `backend/repository/profileRepository.go`
    - Implement Create, Update, FindByID, FindByUserID, Search methods
    - Add IncrementViewCount method
    - _Requirements: 1.3, 1.4, 5.1, 5.4_
  
  - [~] 2.2 Create SkillRepository with CRUD operations
    - Implement Add, Remove, FindByProfile methods
    - Add duplicate detection logic
    - _Requirements: 2.1, 2.3, 2.4_
  
  - [~] 2.3 Create ExperienceRepository with CRUD operations
    - Implement Add, Update, Remove, FindByProfile methods
    - Add date validation at database level
    - _Requirements: 3.1, 3.3, 3.4_
  
  - [~] 2.4 Create WorkSampleRepository with CRUD operations
    - Implement Add, Remove, FindByProfile methods
    - _Requirements: 4.3, 4.5, 4.6_

- [ ]* 2.5 Write property tests for repository layer
  - **Property 8: Experience Date Invariant**
  - **Property 12: Resource Deletion Completeness**
  - **Validates: Requirements 3.3, 2.4, 4.6**

- [ ] 3. Implement service layer with business logic
  - [~] 3.1 Create ProfileService in `backend/service/profile/profileService.go`
    - Implement CreateProfile, UpdateProfile, GetProfile methods
    - Add profile completion percentage calculation
    - Implement input validation (required fields, length limits)
    - Add XSS sanitization for text inputs
    - _Requirements: 1.2, 1.3, 1.4, 1.5, 1.6, 5.6, 9.1, 10.2_
  
  - [~] 3.2 Create SkillService with validation
    - Implement AddSkill, RemoveSkill, GetSkillsByProfile methods
    - Add category validation against predefined enum
    - Implement skill limit enforcement (max 50 skills)
    - Add skill organization logic (group by category, sort by proficiency)
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  
  - [~] 3.3 Create ExperienceService with validation
    - Implement AddExperience, UpdateExperience, RemoveExperience methods
    - Add date validation (end_date >= start_date)
    - Implement chronological ordering (descending by start_date)
    - Add description length validation (max 1000 chars)
    - _Requirements: 3.1, 3.2, 3.3, 3.5, 3.6_
  
  - [~] 3.4 Create WorkSampleService with Cloudinary integration
    - Implement AddWorkSample, RemoveWorkSample methods
    - Add file type validation (JPEG, PNG, WebP)
    - Add file size validation (max 5MB)
    - Integrate with existing Cloudinary utility for uploads
    - Implement work sample limit enforcement (max 20 samples)
    - Add cleanup logic to remove Cloudinary images on deletion
    - _Requirements: 4.1, 4.2, 4.3, 4.4, 4.5, 4.6, 8.5_

- [ ]* 3.5 Write property tests for service layer validation
  - **Property 2: Required Field Validation**
  - **Property 3: Input Sanitization**
  - **Property 6: Skill Category Validation**
  - **Property 10: Work Sample File Type Validation**
  - **Validates: Requirements 1.2, 1.6, 2.5, 4.1, 9.1**

- [ ]* 3.6 Write property tests for service layer business logic
  - **Property 5: Skill Deduplication (Idempotence)**
  - **Property 7: Skills Organization by Category**
  - **Property 9: Experience Chronological Ordering**
  - **Property 16: Profile Completion Calculation**
  - **Validates: Requirements 2.3, 2.6, 3.5, 10.2**

- [ ] 4. Extend GraphQL schema and create resolvers
  - [~] 4.1 Update `backend/graph/schema.graphqls` with profile types
    - Add TalentProfile, Skill, Experience, WorkSample types
    - Add input types for create/update operations
    - Add ProfileSearchFilter and ProfileSearchResult types
    - Add Upload scalar for file uploads (if not already present)
    - _Requirements: 1.1, 2.1, 3.1, 4.1, 6.1_
  
  - [~] 4.2 Implement profile query resolvers in `backend/graph/schema.resolvers.go`
    - Implement talentProfile(id: ID!) query
    - Implement myProfile query (uses authenticated user context)
    - Implement searchProfiles(filter: ProfileSearchFilter!) query
    - Add view count increment on profile access
    - Implement access control (hide contact info for unauthenticated users)
    - _Requirements: 5.1, 5.2, 5.4, 6.1, 6.2, 6.3, 6.4, 7.2, 7.4_
  
  - [~] 4.3 Implement profile mutation resolvers
    - Implement createProfile mutation
    - Implement updateProfile mutation with authorization check
    - Add rate limiting middleware (10 updates per hour)
    - _Requirements: 1.3, 1.4, 9.3, 9.4_
  
  - [~] 4.4 Implement skill mutation resolvers
    - Implement addSkill mutation
    - Implement removeSkill mutation with authorization check
    - _Requirements: 2.1, 2.4, 9.4_
  
  - [~] 4.5 Implement experience mutation resolvers
    - Implement addExperience mutation
    - Implement updateExperience mutation with authorization check
    - Implement removeExperience mutation with authorization check
    - _Requirements: 3.1, 3.2, 9.4_
  
  - [~] 4.6 Implement work sample mutation resolvers
    - Implement addWorkSample mutation with file upload
    - Implement removeWorkSample mutation with authorization check
    - _Requirements: 4.3, 4.5, 4.6, 9.4_

- [ ]* 4.7 Write property tests for GraphQL API
  - **Property 13: Profile Retrieval Completeness**
  - **Property 14: Unauthenticated Access Data Filtering**
  - **Property 15: View Count Increment**
  - **Property 23: Authorization Enforcement**
  - **Property 24: Rate Limiting Enforcement**
  - **Validates: Requirements 5.1, 5.2, 5.4, 9.3, 9.4**

- [ ]* 4.8 Write property tests for search functionality
  - **Property 17: Search Result Correctness**
  - **Property 18: Multi-Filter Conjunction**
  - **Property 19: Pagination Consistency**
  - **Property 20: Search Result Ordering**
  - **Validates: Requirements 6.1, 6.2, 6.3, 6.4, 6.6**

- [~] 5. Checkpoint - Backend complete, ensure all tests pass
  - Run all backend unit tests and property tests
  - Verify GraphQL schema compiles and generates types
  - Test API endpoints manually with GraphQL playground
  - Ensure all tests pass, ask the user if questions arise

- [ ] 6. Implement frontend TypeScript types and composables
  - [~] 6.1 Create TypeScript interfaces in `frontend/types/profile.ts`
    - Define TalentProfile, Skill, Experience, WorkSample interfaces
    - Define input types matching GraphQL schema
    - _Requirements: 1.1, 2.1, 3.1, 4.1_
  
  - [~] 6.2 Create GraphQL queries and mutations in `frontend/graphql/profile.ts`
    - Write TALENT_PROFILE_QUERY with all fields
    - Write MY_PROFILE_QUERY
    - Write SEARCH_PROFILES_QUERY with filters
    - Write CREATE_PROFILE_MUTATION
    - Write UPDATE_PROFILE_MUTATION
    - Write skill, experience, and work sample mutations
    - _Requirements: 1.3, 5.1, 6.1_
  
  - [~] 6.3 Create useProfile composable in `frontend/composables/useProfile.ts`
    - Implement fetchProfile, updateProfile, createProfile methods
    - Add loading and error state management
    - Add profile completion calculation helper
    - _Requirements: 1.3, 1.4, 5.1, 5.6_
  
  - [~] 6.4 Create useSkills composable in `frontend/composables/useSkills.ts`
    - Implement addSkill, removeSkill, fetchSkills methods
    - Add skill organization logic (group by category)
    - _Requirements: 2.1, 2.4, 2.6_
  
  - [~] 6.5 Create useExperience composable in `frontend/composables/useExperience.ts`
    - Implement addExperience, updateExperience, removeExperience methods
    - Add date validation helpers
    - _Requirements: 3.1, 3.2, 3.3_
  
  - [~] 6.6 Create useWorkSamples composable in `frontend/composables/useWorkSamples.ts`
    - Implement addWorkSample, removeWorkSample methods
    - Add file validation helpers (type, size)
    - _Requirements: 4.1, 4.2, 4.3, 4.6_

- [ ] 7. Create reusable Vue components
  - [~] 7.1 Create ProfileHeader component in `frontend/components/profile/ProfileHeader.vue`
    - Display profile image, name, professional title, location
    - Add edit button (conditional on ownership)
    - Style with Tailwind CSS matching existing design system
    - _Requirements: 1.1, 7.2_
  
  - [~] 7.2 Create ProfileBio component in `frontend/components/profile/ProfileBio.vue`
    - Display bio text with proper formatting
    - Add character count display in edit mode
    - _Requirements: 1.1, 1.5_
  
  - [~] 7.3 Create SkillsList component in `frontend/components/profile/SkillsList.vue`
    - Display skills grouped by category
    - Show proficiency level badges
    - Add remove button in edit mode
    - _Requirements: 2.1, 2.6_
  
  - [~] 7.4 Create SkillsEditor component in `frontend/components/profile/SkillsEditor.vue`
    - Form for adding new skills
    - Category and proficiency level dropdowns
    - Skill count display (X/50)
    - _Requirements: 2.1, 2.2, 2.5_
  
  - [~] 7.5 Create ExperienceTimeline component in `frontend/components/profile/ExperienceTimeline.vue`
    - Display experiences in timeline format
    - Show date ranges with "Present" for current positions
    - Add edit/delete buttons in edit mode
    - _Requirements: 3.1, 3.2, 3.5_
  
  - [~] 7.6 Create ExperienceEditor component in `frontend/components/profile/ExperienceEditor.vue`
    - Form for adding/editing experience entries
    - Date pickers with validation
    - "Current position" checkbox
    - Character count for description (X/1000)
    - _Requirements: 3.1, 3.2, 3.3, 3.6_
  
  - [~] 7.7 Create WorkSamplesGallery component in `frontend/components/profile/WorkSamplesGallery.vue`
    - Grid layout for work sample thumbnails
    - Lightbox/modal for full-size viewing
    - Add delete button in edit mode
    - _Requirements: 4.5, 4.7_
  
  - [~] 7.8 Create WorkSampleUploader component in `frontend/components/profile/WorkSampleUploader.vue`
    - File upload interface with drag-and-drop
    - Image preview before upload
    - File type and size validation with error messages
    - Upload progress indicator
    - Work sample count display (X/20)
    - _Requirements: 4.1, 4.2, 4.3, 4.4_
  
  - [~] 7.9 Create ProfileCompletionIndicator component in `frontend/components/profile/ProfileCompletionIndicator.vue`
    - Progress bar showing completion percentage
    - Breakdown of completed sections
    - Prompts for incomplete sections
    - Completion badge at 100%
    - _Requirements: 5.6, 10.1, 10.2, 10.3, 10.5_

- [ ] 8. Create profile view page
  - [~] 8.1 Create profile view page in `frontend/pages/talent/profile/[id].vue`
    - Fetch profile data using useProfile composable
    - Display ProfileHeader, ProfileBio, SkillsList, ExperienceTimeline, WorkSamplesGallery
    - Show ProfileCompletionIndicator
    - Conditional rendering of edit controls based on ownership
    - Handle loading and error states
    - Implement 404 handling for non-existent profiles
    - _Requirements: 5.1, 5.3, 5.5, 7.2, 7.3_

- [ ]* 8.2 Write unit tests for profile view page
  - Test profile data rendering
  - Test ownership-based edit controls visibility
  - Test error handling (404, network errors)
  - _Requirements: 5.3, 5.5, 7.2_

- [ ] 9. Create profile edit page
  - [~] 9.1 Create profile edit page in `frontend/pages/talent/profile/edit.vue`
    - Fetch user's profile using myProfile query
    - Display editable ProfileHeader, ProfileBio
    - Include SkillsEditor, ExperienceEditor, WorkSampleUploader
    - Add save button with loading state
    - Implement form validation before submission
    - Show success/error notifications
    - Redirect to view page after successful save
    - _Requirements: 1.1, 1.2, 1.3, 1.4, 2.1, 3.1, 4.3_

- [ ]* 9.2 Write unit tests for profile edit page
  - Test form validation
  - Test save functionality
  - Test error handling
  - _Requirements: 1.2, 1.6_

- [ ] 10. Create profile index page with redirect logic
  - [~] 10.1 Create profile index in `frontend/pages/talent/profile/index.vue`
    - Check authentication status
    - Redirect authenticated talents to their own profile (/talent/profile/[their-id])
    - Redirect unauthenticated users to login page
    - _Requirements: 7.4_

- [ ] 11. Create talent search page for clients
  - [~] 11.1 Create ProfileSearchFilters component in `frontend/components/profile/ProfileSearchFilters.vue`
    - Skill search input with autocomplete
    - Category checkboxes
    - Location filter
    - Apply/Clear filters buttons
    - _Requirements: 6.1, 6.2, 6.3_
  
  - [~] 11.2 Create ProfileCard component in `frontend/components/profile/ProfileCard.vue`
    - Display profile summary: name, title, primary skills, profile image
    - Show completion percentage badge
    - Link to full profile view
    - _Requirements: 6.5_
  
  - [~] 11.3 Create talent search page in `frontend/pages/talents/search.vue`
    - Display ProfileSearchFilters
    - Fetch and display search results using searchProfiles query
    - Render ProfileCard for each result
    - Implement pagination controls
    - Handle empty results state
    - _Requirements: 6.1, 6.2, 6.3, 6.4, 6.5, 6.6_

- [ ]* 11.4 Write unit tests for search functionality
  - Test filter application
  - Test pagination
  - Test empty results handling
  - _Requirements: 6.1, 6.3, 6.4_

- [ ] 12. Implement automatic profile creation on user registration
  - [~] 12.1 Update user registration mutation in `backend/graph/schema.resolvers.go`
    - After creating new user with role "talent", automatically create empty profile
    - Set completion_percent to 0
    - Handle errors gracefully (rollback user creation if profile creation fails)
    - _Requirements: 8.2_

- [ ]* 12.2 Write property test for automatic profile creation
  - **Property 22: Automatic Profile Creation**
  - **Validates: Requirements 8.2**

- [ ] 13. Add audit logging for profile modifications
  - [~] 13.1 Create audit log model in `backend/models/auditModels.go`
    - Define AuditLog model with user_id, operation_type, resource_type, resource_id, timestamp
    - _Requirements: 9.6_
  
  - [~] 13.2 Implement audit logging in service layer
    - Add logging calls to all Create, Update, Delete operations
    - Log user ID, operation type, and affected resource
    - _Requirements: 9.6_

- [ ]* 13.3 Write property test for audit logging
  - **Property 25: Audit Logging Completeness**
  - **Validates: Requirements 9.6**

- [ ] 14. Update existing talent dashboard with profile link
  - [~] 14.1 Update `frontend/pages/talent/dashboard.vue`
    - Ensure profile link in header navigates to user's profile
    - Update user icon link to point to /talent/profile
    - _Requirements: 7.1, 7.4_

- [ ] 15. Integration and end-to-end testing
  - [~] 15.1 Create integration test suite
    - Test complete profile creation flow (create → view → edit → update)
    - Test work sample upload and deletion flow
    - Test search and filter flow from client perspective
    - Test authorization (attempt to edit another user's profile)
    - Test rate limiting (make 11 update requests)
    - _Requirements: 1.3, 4.3, 6.1, 9.3, 9.4_

- [~] 16. Final checkpoint - Ensure all tests pass
  - Run all backend tests (unit + property + integration)
  - Run all frontend tests (unit + integration)
  - Manually test all user flows in browser
  - Verify profile URLs are shareable
  - Test on different screen sizes (responsive design)
  - Ensure all tests pass, ask the user if questions arise

## Notes

- Tasks marked with `*` are optional and can be skipped for faster MVP
- Each task references specific requirements for traceability
- Property tests validate universal correctness properties across all inputs
- Unit tests validate specific examples, edge cases, and error conditions
- Integration tests validate end-to-end user flows
- The implementation follows a layered architecture: models → repository → service → API → frontend
- All components should match the existing JobMatcher design system (Tailwind CSS with border-2 border-black style)
