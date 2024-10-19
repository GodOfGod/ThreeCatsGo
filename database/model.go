package database

const CREATE_USER_TABLE = `
	CREATE TABLE users (
		id VARCHAR(50) PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		create_time DATETIME,
		avatar_path TEXT NOT NULL,
		pets_id TEXT,
		event_id TEXT
	);
`

const CREATE_EVENT_TABLE = `
	CREATE TABLE events (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title TEXT NOT NULL,
		image_path TEXT,
		file_path TEXT,
		record TEXT NOT NULL,
		user_id TEXT NOT NULL,
		pet_id TEXT NOT NULL
	);
`

const CREATE_PETS_TABLE = `
	CREATE TABLE pets (
		id VARCHAR(20) PRIMARY KEY,
		nick_name VARCHAR(100),
		birthday VARCHAR(20),
		gender INT,
		avatar_path VARCHAR(120)
	);
`

const INSERT_USER = `
	INSERT INTO users(id, username, create_time, avatar_path, pets_id, event_id) VALUES (?, ?, ?, ?, ?, ?)
`

const SELECT_USER_ID = `
	SELECT id, username, create_time, avatar_path, pets_id, event_id from users WHERE id = ?
`

const INSERT_PET = `
INSERT INTO pets (id, nick_name, gender, avatar_path, birthday) VALUES (?, ?, ?, ?, ?)`

const UPDATE_USER_PETS = `
	UPDATE users
SET pets_id = CASE 
                 WHEN pets_id IS NULL OR pets_id = '' THEN ?
                 ELSE CONCAT(pets_id, ',', ?)
              END
WHERE id = ?;
`

const SELECT_PETS_BY_ID = `
SELECT id, nick_name, birthday, gender, avatar_path from pets WHERE FIND_IN_SET(id, ?);
`
