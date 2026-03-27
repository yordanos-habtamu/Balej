# JobMatcher ATS

A fullstack Applicant Tracking System (ATS) with a Go backend and Nuxt frontend.

## Overview

JobMatcher is a modern, general-purpose Applicant Tracking System designed to help companies and recruiters manage their hiring process efficiently. The platform connects employers with job seekers across various industries including technology, healthcare, finance, education, marketing, and more.

**Key Features:**
- **General-Purpose Categories**: Supports 25+ industry categories for diverse hiring needs
- **Modern Tech Stack**: GraphQL API with real-time capabilities
- **User-Friendly Interface**: Beautiful, responsive design with Three.js animations
- **Role-Based Access**: Admin, Client (Employer), and Talent (Job Seeker) roles
- **Real-Time Communication**: Chat system with WebSocket subscriptions
- **File Upload Support**: CV/resume upload with Cloudinary integration

## Structure

- `backend/`: Go application (Gorm, GraphQL)
- `frontend/`: Nuxt.js application

## Prerequisites

- Go 1.23+
- Node.js 18+
- PostgreSQL

## Setup

### Backend

1. Navigate to `backend`:
   ```bash
   cd backend
   ```
2. Copy `.env.example` to `.env` and configure your database:
   ```bash
   cp .env.example .env
   ```
3. Run the server:
   ```bash
   go run cmd/server/main.go
   ```
   The server will start at `http://localhost:7000`.

### Frontend

1. Navigate to `frontend`:
   ```bash
   cd frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Copy `.env.example` to `.env`:
   ```bash
   cp .env.example .env
   ```
4. Run the development server:
   ```bash
   npm run dev
   ```
   The app will be available at `http://localhost:3000`.

## API Documentation

GraphQL Playground is available at `http://localhost:7000/` when the backend is running.
