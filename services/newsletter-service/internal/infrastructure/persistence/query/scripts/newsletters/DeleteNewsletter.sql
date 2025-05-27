DELETE FROM newsletter_service.newsletters
WHERE id = $1 AND user_id = $2;
