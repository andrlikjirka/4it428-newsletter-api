INSERT INTO newsletter_service.newsletters (id, title, description, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW());
