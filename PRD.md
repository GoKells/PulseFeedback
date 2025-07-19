---

**Project Name:** PulseFeedback
**Owner:** Kelly (SaaS Indie Hacker / Junior Engineer Portfolio)
**Timeline:** 4 Weeks
**Status:** In Progress

---

## 🧑‍💼 Problem Statement

Small businesses often lack tools to systematically collect, manage, and analyze customer feedback. Existing solutions are either too expensive or too complex. Most businesses still rely on ad-hoc Google Forms, spreadsheets, or anecdotal feedback from staff.

---

## 🎯 Objective

**Build a simple, clean, SaaS-style feedback platform** where customers can leave reviews, business owners can track feedback trends through a dashboard, and admins can moderate content.

The project’s purpose is to showcase fullstack development skills in Go, TypeScript, SQL, and Tailwind through a realistic business problem.

---

## 🏁 Success Criteria

| Metric                   | Goal                                       |
| ------------------------ | ------------------------------------------ |
| ✅ Feature Completeness   | Full CRUD on feedback, business, users     |
| ✅ Technical Completeness | Auth, API, dashboards, clean UI            |
| ✅ Deployment             | Public, usable demo                        |
| ✅ Presentation Ready     | Clear GitHub README, LinkedIn post, resume |

---

## 🔑 Key Features & Requirements

### 1️⃣ User Authentication

* Sign up / Login / Logout
* Password hashing & JWT sessions
* Role-based access (admin / business / user)

---

### 2️⃣ Business Profiles

* Business owners can create/manage their business profile
* Each business has its own feedback feed
* Admin can manage all businesses

---

### 3️⃣ Feedback Submission System

* Customers can leave feedback on a business profile:

  * Text
  * 1-5 star rating
  * Optional tags ("Service", "Cleanliness", etc.)
* Public-facing feedback page per business
* Admin can moderate (edit/delete)

---

### 4️⃣ Admin Dashboard (Analytics)

* Total feedback count
* Average rating (last 30 days / all time)
* Ratings over time (line graph)
* Common tags (bar chart)
* Search + filter feedback (by date, rating)

---

### 5️⃣ UI & UX

* Clean, modern UI with **TailwindCSS**
* Mobile-friendly
* Intuitive navigation: Dashboard / Feedback / Analytics / Settings

---

### 6️⃣ Optional (Stretch)

* CSV Export
* Business invites team members
* Email notifications on new feedback

---

## 🏗️ Technical Stack

### **Frontend**

| Tech         | Purpose                  |
| ------------ | ------------------------ |
| Vite + React | SPA Framework            |
| TypeScript   | Strong typing, modern DX |
| TailwindCSS  | Styling                  |
| React Query  | Data fetching / caching  |
| Chart.js     | Analytics graphs         |

### **Backend**

| Tech       | Purpose                         |
| ---------- | ------------------------------- |
| Go (Fiber) | Fast HTTP API server            |
| PostgreSQL | Relational data, business logic |
| GORM/sqlx  | ORM / DB Queries                |
| JWT        | Auth Tokens                     |

### **Infra / Tooling**

| Tool    | Purpose              |
| ------- | -------------------- |
| Docker  | Local Dev Containers |
| Railway | Backend Deployment   |
| Vercel  | Frontend Deployment  |

---

## 🔄 User Flows

### **Customer:**

1. Visit business profile → Submit feedback (rating, comments, tags)

### **Business Owner:**

1. Login → View dashboard → See analytics
2. Moderate feedback → Respond/delete

### **Admin:**

1. Login → View all businesses
2. Moderate feedback globally
3. Manage users and businesses

---

## 🔢 Database Schema (Simplified)

```plaintext
Users
- id
- email
- hashed_password
- role (user, business, admin)

Businesses
- id
- owner_id (FK: Users)
- name
- description

Feedback
- id
- business_id (FK: Businesses)
- rating (1-5)
- comment
- tags (array)
- created_at
```

---

## 🚢 Deliverables

| Deliverable    | Description                           |
| -------------- | ------------------------------------- |
| Live App Demo  | Fully deployed, public-facing app     |
| GitHub Repo    | Clean, documented codebase            |
| README.md      | Clear instructions, features list     |
| LinkedIn Post  | Short, focused showcase post          |
| Resume Bullets | Results-driven, tech-specific wording |

---

## 💬 How to Explain in Interview

> PulseFeedback is a SaaS-style feedback analytics platform I built to practice real-world software design and delivery. I handled the fullstack: Go APIs, relational data modeling with Postgres, JWT auth, and a modern React + TypeScript frontend with data dashboards. It demonstrates my ability to think through multi-user flows, data aggregation, and clean UI/UX — exactly what’s needed in product engineering roles.

---

## 📅 Timeline (Milestones)

| Week   | Focus                           |
| ------ | ------------------------------- |
| Week 1 | Auth, DB schema, UI scaffolding |
| Week 2 | CRUD for feedback/business      |
| Week 3 | Analytics dashboards            |
| Week 4 | Polish, deploy, write-ups       |

---