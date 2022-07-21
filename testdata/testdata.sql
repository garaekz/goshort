-- password = $2a$10$M31dwIPy1MXx9w3Kn9ac3.5B8uW53xrwhnUrRO5.M7ZnKmrenq3zu
INSERT INTO users (id, email, password, created_at, updated_at, is_active) 
VALUES ('967d5bb5-3a7a-4d5e-8a6c-febc8c5b3f13', 't@t.io', '$2a$10$M31dwIPy1MXx9w3Kn9ac3.5B8uW53xrwhnUrRO5.M7ZnKmrenq3zu', '2019-10-01 15:36:38'::timestamp, '2019-10-01 15:36:38'::timestamp, true);

INSERT INTO shorts (code, original_url, visits, user_id, creator_ip, created_at, updated_at, deleted_at)
VALUES ('x3S', 'https://www.google.com', 0, '967d5bb5-3a7a-4d5e-8a6c-febc8c5b3f13', '192.168.0.1', '2019-10-01 15:36:38'::timestamp, '2019-10-01 15:36:38'::timestamp, NULL);

INSERT INTO keys (id, user_id, key, created_at, updated_at, is_active)
VALUES ('e0bb80ec-75a6-4348-bfc3-6ac1e89b195e', '967d5bb5-3a7a-4d5e-8a6c-febc8c5b3f13', '1OFGwtAJoSJdHzro14F4VftLO0nl4nWln8NUs7kYg3qZa76tW0D8F7EfGNNJaPOa', '2019-10-01 15:36:38'::timestamp, '2019-10-01 15:36:38'::timestamp, true);