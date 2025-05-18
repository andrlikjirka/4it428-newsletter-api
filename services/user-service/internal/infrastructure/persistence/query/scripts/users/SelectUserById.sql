SELECT id, email, firstname, lastname, firebase_uid
FROM user_service.users
WHERE id = $1