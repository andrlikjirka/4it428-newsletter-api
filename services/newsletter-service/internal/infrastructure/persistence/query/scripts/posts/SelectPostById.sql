SELECT id, newsletter_id, title, content, html_content, published, created_at, updated_at
FROM newsletter_service.posts
WHERE id = $1 AND newsletter_id = $2