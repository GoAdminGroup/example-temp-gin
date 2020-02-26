create table if not exists demo_test
(
    id         integer primary key autoincrement,
    grade_name varchar(30) not null,
    grade_desc varchar(120)         default null,
    created_at datetime    not null default (STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')),
    updated_at datetime             default null
);

create table if not exists demo_grade
(
    id         integer primary key autoincrement,
    grade_name varchar(30) not null,
    grade_desc varchar(120)         default null,
    created_at datetime    not null default (STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')),
    updated_at datetime             default null
);

create table if not exists demo_class
(
    id               integer primary key autoincrement,
    class_name       varchar(30) not null,
    class_desc       varchar(120)         default null,
    grade_id         integer     not null,
    class_time_start varchar(16) not null,
    class_time_end   varchar(16) not null,
    created_at       datetime    not null default current_timestamp,
    updated_at       datetime             default null
);

create table if not exists demo_student
(
    id         integer primary key autoincrement,
    stu_name   varchar(30) not null,
    stu_age    integer     not null,
    stu_sex    integer              default 0 check (stu_sex in (0, 1, 2)), -- 0:unknown 1:Male 2:Female
    created_at datetime    not null default current_timestamp,
    updated_at datetime             default null
);

create table if not exists demo_student_class
(
    id         integer primary key autoincrement,
    class_id   integer  not null,
    stu_id     integer  not null,
    created_at datetime not null default current_timestamp,
    updated_at datetime          default null
);

create table if not exists demo_student_score
(
    id         integer primary key autoincrement,
    stu_id     integer  not null,
    class_id   integer  not null,
    stu_score  int      not null,
    created_at datetime not null default current_timestamp,
    updated_at datetime          default null
);