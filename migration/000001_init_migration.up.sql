CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    location JSONB,
    permissions JSONB
);

CREATE TABLE user_roles (
    user_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);

CREATE TABLE campaigns (
    id SERIAL PRIMARY KEY,
    name JSONB NOT NULL,
    description JSONB,
    location JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE donation_boxes (
    id SERIAL PRIMARY KEY,
    campaign_id INTEGER NOT NULL,
    name JSONB NOT NULL,
    expected_donations REAL,
    donation_size REAL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (campaign_id) REFERENCES campaigns(id) ON DELETE CASCADE
);

CREATE TABLE donations (
    id SERIAL PRIMARY KEY,
    donation_box_id INTEGER NOT NULL,
    promised_amount REAL,
    received_amount REAL,
    donator JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (donation_box_id) REFERENCES donation_boxes(id) ON DELETE CASCADE
);

CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    donation_box_id INTEGER NOT NULL,
    name JSONB NOT NULL,
    description JSONB,
    location JSONB,
    amount REAL,
    technical_progress INTEGER,
    financial_progress INTEGER,
    status VARCHAR(255),
    implementation_start_date TIMESTAMPTZ,
    expected_end_date TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (donation_box_id) REFERENCES donation_boxes(id) ON DELETE CASCADE
);

CREATE TABLE contents (
    id SERIAL PRIMARY KEY,
    campaign_id INTEGER,
    donation_box_id INTEGER,
    project_id INTEGER,
    title JSONB,
    description JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (campaign_id) REFERENCES campaigns(id) ON DELETE CASCADE,
    FOREIGN KEY (donation_box_id) REFERENCES donation_boxes(id) ON DELETE CASCADE,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
    CONSTRAINT check_one_parent CHECK (
        (CASE WHEN campaign_id IS NOT NULL THEN 1 ELSE 0 END) +
        (CASE WHEN donation_box_id IS NOT NULL THEN 1 ELSE 0 END) +
        (CASE WHEN project_id IS NOT NULL THEN 1 ELSE 0 END)
        = 1
    )
);

CREATE TABLE attachments (
    id SERIAL PRIMARY KEY,
    content_id INTEGER NOT NULL,
    name VARCHAR(255) UNIQUE NOT NULL,
    file_link VARCHAR(255),
    file_blob BYTEA,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (content_id) REFERENCES contents(id) ON DELETE CASCADE,
    CONSTRAINT check_one_attachment_type CHECK (
        (CASE WHEN file_link IS NOT NULL THEN 1 ELSE 0 END) +
        (CASE WHEN file_blob IS NOT NULL THEN 1 ELSE 0 END)
        = 1
    )
);

INSERT INTO roles (name, permissions, location) VALUES ('admin', '{"all": true}', '{"all": true}');
