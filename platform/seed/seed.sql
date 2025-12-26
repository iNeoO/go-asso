BEGIN;

INSERT INTO roles_enum (id, name) VALUES
    ('CREATOR', 'creator'),
    ('ADMINISTRATOR', 'admin'),
    {'TEAM_MEMBER', 'teamMember'},
    ('VALIDATED', 'validated'),
    ('NOT_VALIDATED', 'notValidated')
ON CONFLICT (name) DO NOTHING;

INSERT INTO organizations (id, name) VALUES
    ('aaaaaaa1-aaaa-aaaa-aaaa-aaaaaaaaaaa1', 'Handball Club'),
    ('aaaaaaa2-aaaa-aaaa-aaaa-aaaaaaaaaaa2', 'Football Club'),
    ('aaaaaaa3-aaaa-aaaa-aaaa-aaaaaaaaaaa3', 'Tennis Club'),
    ('aaaaaaa4-aaaa-aaaa-aaaa-aaaaaaaaaaa4', 'Rugby Club')
ON CONFLICT (id) DO NOTHING;

INSERT INTO users (id, email, first_name, last_name, password_hash, is_email_verified) VALUES
    ('bbbbbbb1-bbbb-bbbb-bbbb-bbbbbbbbbbb1', 'alice@example.com', 'Alice', 'Alice', '$2a$12$8fUdRrsL9Z9qzsRJLykU2.QdMoSXatsUaHgqk8P5ueS1gIcY//mkS', TRUE),
    ('bbbbbbb2-bbbb-bbbb-bbbb-bbbbbbbbbbb2', 'bob@example.com', 'Bob', 'Bob', '$2a$12$8fUdRrsL9Z9qzsRJLykU2.QdMoSXatsUaHgqk8P5ueS1gIcY//mkS', TRUE),
    ('bbbbbbb3-bbbb-bbbb-bbbb-bbbbbbbbbbb3', 'carol@example.com', 'Carol', 'Carol', '$2a$12$8fUdRrsL9Z9qzsRJLykU2.QdMoSXatsUaHgqk8P5ueS1gIcY//mkS', TRUE),
    ('bbbbbbb4-bbbb-bbbb-bbbb-bbbbbbbbbbb4', 'dave@example.com', 'Dave', 'Dave', '$2a$12$8fUdRrsL9Z9qzsRJLykU2.QdMoSXatsUaHgqk8P5ueS1gIcY//mkS', TRUE),
    ('bbbbbbb5-bbbb-bbbb-bbbb-bbbbbbbbbbb5', 'eve@example.com', 'Eve', 'Eve', '$2a$12$8fUdRrsL9Z9qzsRJLykU2.QdMoSXatsUaHgqk8P5ueS1gIcY//mkS', TRUE),
    ('bbbbbbb6-bbbb-bbbb-bbbb-bbbbbbbbbbb6', 'frank@example.com', 'Frank', 'Frank', '$2a$12$8fUdRrsL9Z9qzsRJLykU2.QdMoSXatsUaHgqk8P5ueS1gIcY//mkS', TRUE),
    ('bbbbbbb7-bbbb-bbbb-bbbb-bbbbbbbbbbb7', 'grace@example.com', 'Grace', 'Grace', '$2a$12$8fUdRrsL9Z9qzsRJLykU2.QdMoSXatsUaHgqk8P5ueS1gIcY//mkS', TRUE),
    ('bbbbbbb8-bbbb-bbbb-bbbb-bbbbbbbbbbb8', 'heidi@example.com', 'Heidi', 'Heidi', '$2a$12$8fUdRrsL9Z9qzsRJLykU2.QdMoSXatsUaHgqk8P5ueS1gIcY//mkS', TRUE),
    ('bbbbbbb9-bbbb-bbbb-bbbb-bbbbbbbbbbb9', 'ivan@example.com', 'Ivan', 'Ivan', '$2a$12$8fUdRrsL9Z9qzsRJLykU2.QdMoSXatsUaHgqk8P5ueS1gIcY//mkS', TRUE),
    ('bbbbbb10-bbbb-bbbb-bbbb-bbbbbbbbbbb0', 'judy@example.com', 'Judy', 'Judy', '$2a$12$8fUdRrsL9Z9qzsRJLykU2.QdMoSXatsUaHgqk8P5ueS1gIcY//mkS', TRUE)
ON CONFLICT (id) DO NOTHING;

INSERT INTO user_organizations (id, user_id, organization_id, role_id) VALUES
    ('c1c1c1c1-0000-0000-0000-000000000001', 'bbbbbbb1-bbbb-bbbb-bbbb-bbbbbbbbbbb1', 'aaaaaaa1-aaaa-aaaa-aaaa-aaaaaaaaaaa1', (SELECT id FROM roles_enum WHERE name = 'creator')),
    ('c1c1c1c1-0000-0000-0000-000000000002', 'bbbbbbb2-bbbb-bbbb-bbbb-bbbbbbbbbbb2', 'aaaaaaa1-aaaa-aaaa-aaaa-aaaaaaaaaaa1', (SELECT id FROM roles_enum WHERE name = 'admin')),
    ('c1c1c1c1-0000-0000-0000-000000000003', 'bbbbbbb3-bbbb-bbbb-bbbb-bbbbbbbbbbb3', 'aaaaaaa1-aaaa-aaaa-aaaa-aaaaaaaaaaa1', (SELECT id FROM roles_enum WHERE name = 'valided')),

    ('c1c1c1c1-0000-0000-0000-000000000004', 'bbbbbbb4-bbbb-bbbb-bbbb-bbbbbbbbbbb4', 'aaaaaaa2-aaaa-aaaa-aaaa-aaaaaaaaaaa2', (SELECT id FROM roles_enum WHERE name = 'creator')),
    ('c1c1c1c1-0000-0000-0000-000000000005', 'bbbbbbb5-bbbb-bbbb-bbbb-bbbbbbbbbbb5', 'aaaaaaa2-aaaa-aaaa-aaaa-aaaaaaaaaaa2', (SELECT id FROM roles_enum WHERE name = 'admin')),
    ('c1c1c1c1-0000-0000-0000-000000000006', 'bbbbbbb6-bbbb-bbbb-bbbb-bbbbbbbbbbb6', 'aaaaaaa2-aaaa-aaaa-aaaa-aaaaaaaaaaa2', (SELECT id FROM roles_enum WHERE name = 'valided')),

    ('c1c1c1c1-0000-0000-0000-000000000007', 'bbbbbbb7-bbbb-bbbb-bbbb-bbbbbbbbbbb7', 'aaaaaaa3-aaaa-aaaa-aaaa-aaaaaaaaaaa3', (SELECT id FROM roles_enum WHERE name = 'creator')),
    ('c1c1c1c1-0000-0000-0000-000000000008', 'bbbbbbb8-bbbb-bbbb-bbbb-bbbbbbbbbbb8', 'aaaaaaa3-aaaa-aaaa-aaaa-aaaaaaaaaaa3', (SELECT id FROM roles_enum WHERE name = 'admin')),
    ('c1c1c1c1-0000-0000-0000-000000000009', 'bbbbbbb9-bbbb-bbbb-bbbb-bbbbbbbbbbb9', 'aaaaaaa3-aaaa-aaaa-aaaa-aaaaaaaaaaa3', (SELECT id FROM roles_enum WHERE name = 'valided')),

    ('c1c1c1c1-0000-0000-0000-000000000010', 'bbbbbb10-bbbb-bbbb-bbbb-bbbbbbbbbbb0', 'aaaaaaa4-aaaa-aaaa-aaaa-aaaaaaaaaaa4', (SELECT id FROM roles_enum WHERE name = 'creator')),
    ('c1c1c1c1-0000-0000-0000-000000000011', 'bbbbbbb1-bbbb-bbbb-bbbb-bbbbbbbbbbb1', 'aaaaaaa4-aaaa-aaaa-aaaa-aaaaaaaaaaa4', (SELECT id FROM roles_enum WHERE name = 'admin')),
    ('c1c1c1c1-0000-0000-0000-000000000012', 'bbbbbbb2-bbbb-bbbb-bbbb-bbbbbbbbbbb2', 'aaaaaaa4-aaaa-aaaa-aaaa-aaaaaaaaaaa4', (SELECT id FROM roles_enum WHERE name = 'notValided'))
ON CONFLICT (user_id, organization_id) DO NOTHING;

COMMIT;
