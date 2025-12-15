# Database Schema

This document outlines the database schema for the donation campaign project.

## Tables

### `users`

| Column | Data Type | Constraints | Description |
|---|---|---|---|
| `id` | SERIAL | PRIMARY KEY | Unique identifier for the user. |
| `email` | VARCHAR(255) | UNIQUE, NOT NULL | User's email address. |
| `name` | VARCHAR(255) | NOT NULL | User's full name. |
| `phone` | VARCHAR(255) | | User's phone number. |
| `created_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Timestamp of when the user was created. |
| `updated_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Timestamp of when the user was last updated. |

### `roles`

| Column | Data Type | Constraints | Description |
|---|---|---|---|
| `id` | SERIAL | PRIMARY KEY | Unique identifier for the role. |
| `name` | VARCHAR(255) | UNIQUE, NOT NULL | Name of the role (e.g., 'admin', 'editor'). |
| `location` | JSONB | | JSON object describing location-based permissions. |
| `permissions` | JSONB | | JSON object describing the permissions for the role. |

### `user_roles`

This is a join table to manage the many-to-many relationship between users and roles.

| Column | Data Type | Constraints | Description |
|---|---|---|---|
| `user_id` | INTEGER | NOT NULL, PRIMARY KEY, FOREIGN KEY | Foreign key referencing the `users` table. |
| `role_id` | INTEGER | NOT NULL, PRIMARY KEY, FOREIGN KEY | Foreign key referencing the `roles` table. |

### `campaigns`

| Column | Data Type | Constraints | Description |
|---|---|---|---|
| `id` | SERIAL | PRIMARY KEY | Unique identifier for the campaign. |
| `name` | JSONB | NOT NULL | Name of the campaign, can be multilingual. |
| `description` | JSONB | | Description of the campaign, can be multilingual. |
| `location` | JSONB | | Location of the campaign. |
| `created_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Timestamp of when the campaign was created. |
| `updated_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Timestamp of when the campaign was last updated. |

### `donation_boxes`

| Column | Data Type | Constraints | Description |
|---|---|---|---|
| `id` | SERIAL | PRIMARY KEY | Unique identifier for the donation box. |
| `campaign_id` | INTEGER | NOT NULL, FOREIGN KEY | Foreign key referencing the `campaigns` table. |
| `name` | JSONB | NOT NULL | Name of the donation box, can be multilingual. |
| `expected_donations` | REAL | | The expected amount of donations for this box. |
| `donation_size` | REAL | | The size or unit of a single donation. |
| `created_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Timestamp of when the donation box was created. |
| `updated_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Timestamp of when the donation box was last updated. |

### `donations`

| Column | Data Type | Constraints | Description |
|---|---|---|---|
| `id` | SERIAL | PRIMARY KEY | Unique identifier for the donation. |
| `donation_box_id` | INTEGER | NOT NULL, FOREIGN KEY | Foreign key referencing the `donation_boxes` table. |
| `promised_amount` | REAL | | The amount the donator promised to donate. |
| `received_amount` | REAL | | The amount actually received from the donator. |
| `donator` | JSONB | | Information about the donator. |
| `created_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Timestamp of when the donation was made. |
| `updated_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Timestamp of when the donation was last updated. |

### `projects`

| Column | Data Type | Constraints | Description |
|---|---|---|---|
| `id` | SERIAL | PRIMARY KEY | Unique identifier for the project. |
| `donation_box_id` | INTEGER | NOT NULL, FOREIGN KEY | Foreign key referencing the `donation_boxes` table. |
| `name` | JSONB | NOT NULL | Name of the project, can be multilingual. |
| `description` | JSONB | | Description of the project, can be multilingual. |
| `location` | JSONB | | Location of the project. |
| `amount` | REAL | | The budget or amount allocated to the project. |
| `technical_progress` | INTEGER | | Technical progress of the project (e.g., percentage). |
| `financial_progress` | INTEGER | | Financial progress of the project (e.g., percentage). |
| `status` | VARCHAR(255) | | Current status of the project. |
| `implementation__date` | TIMESTAMPTZ | | The start date of the project implementation. |
| `expected_end_date` | TIMESTAMPTZ | | The expected end date of the project. |
| `created_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Timestamp of when the project was created. |
| `updated_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Timestamp of when the project was last updated. |

### `contents`

| Column | Data Type | Constraints | Description |
|---|---|---|---|
| `id` | SERIAL | PRIMARY KEY | Unique identifier for the content. |
| `campaign_id` | INTEGER | FOREIGN KEY | Foreign key referencing the `campaigns` table. |
| `donation_box_id` | INTEGER | FOREIGN KEY | Foreign key referencing the `donation_boxes` table. |
| `project_id` | INTEGER | FOREIGN KEY | Foreign key referencing the `projects` table. |
| `title` | JSONB | | Title of the content, can be multilingual. |
| `description` | JSONB | | Description of the content, can be multilingual. |
| `created_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Timestamp of when the content was created. |
| `updated_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Timestamp of when the content was last updated. |

**Constraint:** `check_one_parent` - Ensures that each content is associated with exactly one parent (either a campaign, a donation box, or a project).

### `attachments`

| Column | Data Type | Constraints | Description |
|---|---|---|---|
| `id` | SERIAL | PRIMARY KEY | Unique identifier for the attachment. |
| `content_id` | INTEGER | NOT NULL, FOREIGN KEY | Foreign key referencing the `contents` table. |
| `name` | VARCHAR(255) | UNIQUE, NOT NULL | Name of the attachment. |
| `file_link` | VARCHAR(255) | | Link to the file. |
| `file_blob` | BYTEA | | The file stored as a binary large object. |
| `created_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Timestamp of when the attachment was created. |
| `updated_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | Timestamp of when the attachment was last updated. |

**Constraint:** `check_one_attachment_type` - Ensures that each attachment has either a `file_link` or a `file_blob`, but not both.

## Entity-Relationship Diagram

For a visual representation of the database schema, please refer to the following files:

- [Mermaid Chart](./database_schema_chart.md)
- [PNG Image](./db.png)
