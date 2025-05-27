UPDATE newsletter_service.posts
SET title = $1, content = $2, html_content = $3, published = $4, updated_at = NOW()
WHERE id = $5;
