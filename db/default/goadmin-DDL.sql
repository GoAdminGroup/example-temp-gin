create table if not exists goadmin_menu
(
	id integer
		primary key autoincrement,
	parent_id INT default '0' not null,
	"order" INT default '0' not null,
	type INT default '0' not null,
	title CHAR(50) collate NOCASE not null,
	icon CHAR(50) collate NOCASE not null,
	uri CHAR(50) collate NOCASE default NULL,
	created_at TIMESTAMP default CURRENT_TIMESTAMP,
	updated_at TIMESTAMP default CURRENT_TIMESTAMP,
	header CHAR(150) default NULL
);

create table if not exists goadmin_operation_log
(
	id integer
		primary key autoincrement,
	user_id INT not null,
	path CHAR(255) collate NOCASE not null,
	method CHAR(10) collate NOCASE not null,
	ip CHAR(15) collate NOCASE not null,
	input text collate NOCASE not null,
	created_at TIMESTAMP default CURRENT_TIMESTAMP,
	updated_at TIMESTAMP default CURRENT_TIMESTAMP
);

create table if not exists goadmin_permissions
(
	id integer
		primary key autoincrement,
	name CHAR(50) collate NOCASE not null,
	slug CHAR(50) collate NOCASE not null,
	http_method CHAR(255) collate NOCASE default NULL,
	http_path text collate NOCASE,
	created_at TIMESTAMP default CURRENT_TIMESTAMP,
	updated_at TIMESTAMP default CURRENT_TIMESTAMP
);

create table if not exists goadmin_role_menu
(
	role_id INT not null,
	menu_id INT not null,
	created_at TIMESTAMP default CURRENT_TIMESTAMP,
	updated_at TIMESTAMP default CURRENT_TIMESTAMP
);

create table if not exists goadmin_role_permissions
(
	role_id INT not null,
	permission_id INT not null,
	created_at TIMESTAMP default CURRENT_TIMESTAMP,
	updated_at TIMESTAMP default CURRENT_TIMESTAMP
);

create table if not exists goadmin_role_users
(
	role_id INT not null,
	user_id INT not null,
	created_at TIMESTAMP default CURRENT_TIMESTAMP,
	updated_at TIMESTAMP default CURRENT_TIMESTAMP
);

create table if not exists goadmin_roles
(
	id integer
		primary key autoincrement,
	name CHAR(50) collate NOCASE not null,
	slug CHAR(50) collate NOCASE not null,
	created_at TIMESTAMP default CURRENT_TIMESTAMP,
	updated_at TIMESTAMP default CURRENT_TIMESTAMP
);

create table if not exists goadmin_session
(
	id integer
		primary key autoincrement,
	sid CHAR(50) default NULL,
	"values" CHAR(3000) default NULL,
	created_at TIMESTAMP default CURRENT_TIMESTAMP,
	updated_at TIMESTAMP default CURRENT_TIMESTAMP
);

create table if not exists goadmin_user_permissions
(
	user_id INT not null,
	permission_id INT not null,
	created_at TIMESTAMP default CURRENT_TIMESTAMP,
	updated_at TIMESTAMP default CURRENT_TIMESTAMP
);

