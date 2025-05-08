SELECT id, firebase_uid, email, firstname, lastname
FROM user_service.users
WHERE email = $1;
