-- DATABASE POSTGRESQL

-- USERS

CREATE TABLE IF NOT EXISTS USERS 
(US_ID_USER 			SERIAL PRIMARY KEY
,US_FIRST_NAME          VARCHAR(100)
,US_LAST_NAME          VARCHAR(100)
,US_EMAIL				VARCHAR(200) UNIQUE
,US_USERNAME			VARCHAR(100) UNIQUE
,US_HASH				VARCHAR(500)
,US_PERSONAL_ID			VARCHAR(50)	UNIQUE
,US_BIRTHDAY_DATE		DATE
,US_CREATE_DATETIME		TIMESTAMP
,US_CREATE_USER			VARCHAR(100)
,US_LAST_UPDATE			TIMESTAMP);

CREATE USER API_USER WITH PASSWORD 'API_USER_STORAGE';

GRANT ALL PRIVILEGES ON DATABASE users TO API_USER;

GRANT INSERT, UPDATE, DELETE, SELECT ON TABLE users TO API_USER;

GRANT USAGE, SELECT ON SEQUENCE users_us_id_user_seq TO API_USER;
