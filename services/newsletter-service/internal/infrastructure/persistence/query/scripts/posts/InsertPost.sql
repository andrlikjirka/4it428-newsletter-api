INSERT INTO newsletter_service.posts (id, newsletter_id, title, content, html_content, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5,  NOW(), null);
