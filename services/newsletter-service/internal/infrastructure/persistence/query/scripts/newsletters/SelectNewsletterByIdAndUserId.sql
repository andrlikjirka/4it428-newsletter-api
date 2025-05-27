SELECT id, title, description, created_at, updated_at, user_id
FROM newsletter_service.newsletters
WHERE id = $1 AND user_id = $2;