UPDATE user_service.users
SET firstname = $1, lastname = $2
WHERE email = $3