-- password = $2a$10$M31dwIPy1MXx9w3Kn9ac3.5B8uW53xrwhnUrRO5.M7ZnKmrenq3zu
INSERT INTO users (id, email, password, created_at, updated_at, is_active) 
VALUES ('967d5bb5-3a7a-4d5e-8a6c-febc8c5b3f13', 't@t.io', '$2a$10$M31dwIPy1MXx9w3Kn9ac3.5B8uW53xrwhnUrRO5.M7ZnKmrenq3zu', '2019-10-01 15:36:38'::timestamp, '2019-10-01 15:36:38'::timestamp, true);
INSERT INTO users (id, email, password, created_at, updated_at, is_active) 
VALUES ('f1bfae38-e925-4556-8b05-a1d29f6e85fc', 'admin@t.io', '$2a$10$M31dwIPy1MXx9w3Kn9ac3.5B8uW53xrwhnUrRO5.M7ZnKmrenq3zu', '2019-10-01 15:36:38'::timestamp, '2019-10-01 15:36:38'::timestamp, true);

INSERT INTO shorts (code, original_url, visits, user_id, creator_ip, created_at, updated_at, deleted_at)
VALUES ('x3S', 'https://www.google.com', 0, '967d5bb5-3a7a-4d5e-8a6c-febc8c5b3f13', '192.168.0.1', '2019-10-01 15:36:38'::timestamp, '2019-10-01 15:36:38'::timestamp, NULL);

INSERT INTO keys (key, user_id, created_at, updated_at)
VALUES ('1OFGwtAJoSJdHzro14F4VftLO0nl4nWln8NUs7kYg3qZa76tW0D8F7EfGNNJaPOa', 'f1bfae38-e925-4556-8b05-a1d29f6e85fc', '2021-10-01 15:36:38'::timestamp, '2021-10-01 15:36:38'::timestamp);

INSERT INTO configs (name, value)
VALUES ('default_api_keys_quantity', '2');

INSERT INTO roles (id, name)
VALUES ('b8292ab3-eab3-4c16-9ee9-dfc7bde90e48', 'admin');

INSERT INTO role_users (role_id, user_id)
VALUES ('b8292ab3-eab3-4c16-9ee9-dfc7bde90e48', 'f1bfae38-e925-4556-8b05-a1d29f6e85fc');