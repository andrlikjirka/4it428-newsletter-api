SELECT id, title, description, created_at, updated_at
FROM newsletter_service.newsletters
WHERE id = $1;