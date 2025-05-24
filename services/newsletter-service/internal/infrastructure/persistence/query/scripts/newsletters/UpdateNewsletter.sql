UPDATE newsletter_service.newsletters
SET title = $1, description = $2, updated_at = $3
WHERE id = $4;
