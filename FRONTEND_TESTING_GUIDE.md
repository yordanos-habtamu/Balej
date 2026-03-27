# Frontend Testing Guide

## Overview

All frontend pages have been updated to use mock data instead of backend connections for testing purposes. This allows you to test the UI without needing the backend server running.

## Pages Updated with Mock Data

### ✅ **Connected Pages (Now Using Mock Data)**

1. **Admin Clients** (`/admin/clients/index.vue`)
   - **Mock Data**: 6 client users with names, emails, phone numbers, and locations
   - **Features**: Search functionality, user cards with profile images
   - **Test**: Search for clients by name, email, or ID

2. **Talent Dashboard** (`/talent/dashboard.vue`)
   - **Mock Data**: 6 job listings across different categories
   - **Features**: Job search, category filtering, job type filtering
   - **Test**: Browse jobs, filter by category/type, search functionality

3. **Client Dashboard** (`/client/dashboard.vue`)
   - **Mock Data**: 6 talent profiles with contact information
   - **Features**: Talent search, profile viewing, action buttons
   - **Test**: Search for talents, view profiles, test action buttons

4. **Admin Dashboard** (`/admin/dashboard.vue`)
   - **Mock Data**: System statistics and job listings
   - **Features**: Charts, statistics, job listings with applicant counts
   - **Test**: View dashboard stats, check chart functionality

5. **Old Client Jobs Index** (`/admin/_old/client_jobs_index.vue`)
   - **Mock Data**: 6 job listings (duplicate of talent dashboard data)
   - **Features**: Job browsing and filtering
   - **Test**: Browse jobs, test filtering functionality

6. **Client Jobs Create** (`/client/jobs/create.vue`)
   - **Mock Data**: Complete job posting form with validation
   - **Features**: Job creation form, preview functionality, requirement management
   - **Test**: Create job postings, add/remove requirements, preview before submission

## Pages Still Using Mock Data (No Backend Connection)

### **Admin Pages**
- `/admin/jobs/index.vue` - Uses hardcoded mock data
- `/admin/talents/[name].vue` - UI only
- `/admin/users/index.vue` - UI only
- All compliance, configuration, moderation, roles, and settings pages

### **Client Pages**
- `/client/jobs/[_id]/` - Job detail pages (UI only)
- `/client/organization/*` - Organization management (UI only)
- `/client/projects/*` - Project management (UI only)
- `/client/settings/*` - Client settings (UI only)
- `/client/talents/*` - Talent browsing (UI only)
- `/client/communication/*` - Communication features (UI only)

### **Talent Pages**
- `/talent/jobs/[_id]/` - Job application pages (UI only)
- `/talent/profile/*` - Profile management (UI only)
- `/talent/projects/*` - Project pages (UI only)
- `/talent/settings/*` - Talent settings (UI only)
- `/talent/communication/*` - Communication features (UI only)

## Testing Instructions

### **Navigation from Landing Page**

The main landing page (`/`) now includes direct navigation links to all dashboards:

- **Admin Dashboard**: `/admin/dashboard`
- **Client Dashboard**: `/client/dashboard` 
- **Talent Dashboard**: `/talent/dashboard`

### **1. Admin Pages Testing**
```bash
# Navigate to admin pages
/admin/clients/index.vue
/admin/dashboard.vue
/admin/jobs/index.vue
```

**Test Actions:**
- Search for clients by name/email
- View client profiles
- Check dashboard statistics
- Browse job listings
- Test category filtering

### **2. Client Pages Testing**
```bash
# Navigate to client pages
/client/dashboard.vue
/client/jobs/create.vue
/client/jobs/index.vue (if exists)
```

**Test Actions:**
- Browse talent profiles
- Search for talents
- Create new job postings
- Test job creation form validation
- Preview job postings before submission

### **3. Talent Pages Testing**
```bash
# Navigate to talent pages
/talent/dashboard.vue
/talent/jobs/index.vue (if exists)
```

**Test Actions:**
- Browse job listings
- Filter jobs by category/type
- Search for specific jobs
- View job details

### **4. System Pages Testing**
```bash
# Navigate to system pages
/system/access-denied.vue
/system/invite-expired.vue
/system/job-archived.vue
/system/network-failure.vue
/system/project-closed.vue
/system/suspended.vue
```

**Test Actions:**
- Test error page navigation
- Check system status pages

## Mock Data Structure

### **Users (Clients/Talents)**
```javascript
{
  first_name: 'Sarah',
  last_name: 'Johnson',
  email: 'sarah.johnson@techcorp.com',
  phone_number: '+1 (555) 123-4567',
  address: 'San Francisco, CA',
  role: 'client'
}
```

### **Jobs**
```javascript
{
  job_type: 'Full-time',
  title: 'Senior Software Engineer',
  applicationCount: 24,
  category: 'Technology'
}
```

### **Job Creation Form**
```javascript
{
  title: 'Job Title',
  jobType: 'full-time',
  categories: ['technology', 'engineering'],
  location: 'Remote',
  minSalary: '70000',
  maxSalary: '120000',
  description: 'Job description...',
  requirements: ['Requirement 1', 'Requirement 2']
}
```

## Backend Connection Status

### **Commented Out Imports**
All GraphQL-related imports have been commented out:
```javascript
// import { gql } from 'graphql-tag'
// import { CombinedError } from '@urql/core'
// import { useUrqlClient } from '@/composables/useUrqlClient'
```

### **Commented Out Queries**
All GraphQL queries have been commented out:
```javascript
// const USER_QUERY = gql`...`
// const JOB_QUERY = gql`...`
```

### **Commented Out Functions**
All backend fetch functions have been commented out:
```javascript
// async function loadUsers() { ... }
// async function loadJobs() { ... }
```

### **Mock Data Usage**
All pages now use reactive mock data objects instead of API calls.

## Next Steps

Once you're satisfied with the frontend testing, we can:

1. **Reconnect to Backend**: Uncomment the GraphQL imports and queries
2. **Implement Missing Pages**: Connect the remaining UI-only pages
3. **Add Real Data**: Replace mock data with actual backend responses
4. **Add Error Handling**: Implement proper error states and loading indicators
5. **Add Authentication**: Connect login/register flows to backend
6. **Add Real-time Features**: Implement WebSocket subscriptions for chat

## Testing Checklist

- [ ] Test admin clients page with search functionality
- [ ] Test talent dashboard with job filtering
- [ ] Test client dashboard with talent browsing
- [ ] Test admin dashboard charts and statistics
- [ ] Test job creation form with validation
- [ ] Test job preview functionality
- [ ] Test category and job type filters
- [ ] Test responsive design on mobile devices
- [ ] Test form validation and error states
- [ ] Test navigation between pages

## Notes

- All pages use the same mock data structure for consistency
- Form validation is implemented using HTML5 validation
- Charts use Chart.js with mock data
- Search functionality works with mock data arrays
- No backend server required for frontend testing
- All routes are accessible without authentication for testing