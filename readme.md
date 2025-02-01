# Chacha App

## Overview

Chacha is a web platform designed for business registration. It allows users to:
- **Register and log in** to access a personal dashboard.
- **Submit business registration applications** (with a main name and three alternative names).
- **Track application statuses** (Pending, Approved, Rejected).
- **Allow multiple directors** per business.
- **Enable partners (authorities)** to review, approve, or reject applications with a reason.
- **Provide an admin dashboard** for managing users, applications, and overall statistics.

## Objectives

- **User Registration & Authentication:**  
  Allow users to create an account and log in to access the dashboard.
- **Business Registration:**  
  Enable users to submit business applications, with dynamic input for multiple directors and alternative names.
- **Partner & Admin Dashboards:**  
  - Partners can review and process business applications.
  - Admins can manage users, view all applications, and monitor platform statistics.

## Requirements

- **Backend:**  
  - Golang (1.16+ recommended)
  - PostgreSQL 15 (or PostgreSQL 17 if available)
- **Frontend:**  
  - Node.js (v14+ recommended)
  - Next.js, React, and Tailwind CSS
- **Docker & Docker Compose:** (optional for containerized development)
  - Docker Engine installed on Ubuntu

## Setup Instructions

### Local Setup Without Docker

1. **Clone the Repository:**
   ```bash
   git clone <repository_url>
   cd chacha


#### on first run uncomment this line in main.go
//  seeds.Seed()
