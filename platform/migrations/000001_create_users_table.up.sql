CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SET TIME ZONE 'Europe/Paris';

CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    is_email_verified BOOLEAN DEFAULT FALSE NOT NULL,
    password_hash TEXT NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE sessions (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_hash TEXT NOT NULL UNIQUE,
    revoked_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX sessions_user_id_idx ON sessions(user_id);
CREATE INDEX sessions_expires_at_idx ON sessions(expires_at);

CREATE TABLE roles_enum (
    id VARCHAR(50) DEFAULT NOT NULL PRIMARY KEY,
);


CREATE TABLE organizations (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE user_organizations (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    role_id VARCHAR(50) NOT NULL REFERENCES roles_enum(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY (user_id, organization_id)
);

CREATE INDEX user_organizations_organization_id_idx ON user_organizations(organization_id);
CREATE INDEX user_organizations_user_id_idx ON user_organizations(user_id);

CREATE TABLE activities (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    creator_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT NULL,
    starts_at TIMESTAMPTZ NOT NULL,
    duration_minutes INT  NOT NULL CHECK (duration_minutes > 0),
    capacity INT NOT NULL CHECK (capacity >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX activities_organization_id_starts_at_idx ON activities (organization_id, starts_at);

CREATE TABLE registrations_status_enum (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
);

CREATE TABLE registrations (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    activity_id UUID NOT NULL REFERENCES activities(id) ON DELETE CASCADE,
    status_id UUID NOT NULL REFERENCES registrations_status_enum(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY (user_id, activity_id)
);

CREATE INDEX registrations_user_id_idx ON registrations(user_id);
CREATE INDEX registrations_activity_id_idx ON registrations(activity_id);

INSERT INTO roles_enum (id, name) VALUES
    ('CREATOR'),
    ('ADMINISTRATOR'),
    ('TEAM_MEMBER'),
    ('VALIDATED'),
    ('NOT_VALIDATED')

INSERT INTO registrations_status_enum (id) VALUES
    ('PENDING'),
    ('CONFIRMED'),
    ('CANCELLED');