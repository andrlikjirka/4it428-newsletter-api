SELECT p.id, p.newsletter_id, p.title, p.content, p.html_content, p.published, p.created_at, p.updated_at
FROM newsletter_service.posts p
         JOIN newsletter_service.newsletters n ON p.newsletter_id = n.id
WHERE p.id = $1 AND p.newsletter_id = $2 AND n.user_id = $3
