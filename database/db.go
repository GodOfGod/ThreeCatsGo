package database

const CREATE_DB = `
	create database THREE_CATS;
	use THREE_CATS;
`

const CREATE_DB_USER = `
	create user 'root'@'%' identified by '0604064Li;';
	grant all privileges on *.* to 'root'@'%';
	flush privileges;
`

const CREATE_USER = `
	create table users (
		id varchar(50) primary key,
		username varchar(50),
		create_time varchar(20),
		avatar_path text,
		pets_id text,
		event_id text
	);
`
const CREATE_EVENT = `
	create table events (
		id varchar(30) primary key,
		title text,
		image_path text,
		file_ids text,
		record text,
		user_id varchar(50),
		pet_id varchar(30),
		is_draft int,
		date varchar(20),
		finished int
	);
`

const CREATE_FILE = `
	create table file (
		id varchar(30) primary key,
		file_name text,
		file_path text,
		event_id varchar(30),
		pet_id varchar(30),
		user_id varchar(50)
	);
`

const CREATE_PETS = `
	create table pets (
		id varchar(30) primary key,
		nick_name varchar(100),
		birthday varchar(20),
		gender int,
		avatar_path text
	);
`

const CREATE_QUESTIONNAIRE = `
	create table questionnaire (id varchar(30) primary key, date varchar(50), config_id varchar(30), questionnaire text);
`

const CREATE_QUESTIONNAIRE_CONFIG = `create table questionnaire_config (id varchar(30) primary key, config_fields text, title text);`

const CREATE_QUESTION_ITEM = `create table question_item (id varchar(30) primary key, field varchar(50), title text, input_type varchar(20), options text);`
