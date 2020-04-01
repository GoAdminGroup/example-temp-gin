CREATE TABLE IF NOT EXISTS demo_grade
(
    `id`         int         NOT NULL AUTO_INCREMENT,
    `grade_name` varchar(30) not null,
    `grade_desc` varchar(120)     default null,
    `created_at` timestamp   NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp   NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` timestamp   NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

INSERT INTO demo_grade (id, grade_name, grade_desc, created_at, updated_at, deleted_at)
VALUES (1, 'First-year junior', 'First-year junior high school', '2020-03-26 10:03:17', '2020-03-26 10:03:09', null);
INSERT INTO demo_grade (id, grade_name, grade_desc, created_at, updated_at, deleted_at)
VALUES (2, 'Second-year junior', 'Second-year junior high school ', '2020-03-26 10:03:23', '2020-03-26 10:03:19', null);
INSERT INTO demo_grade (id, grade_name, grade_desc, created_at, updated_at, deleted_at)
VALUES (3, 'Third-year junior', 'Third-year junior high school', '2020-03-26 10:03:29', '2020-03-26 10:03:24', null);

CREATE TABLE IF NOT EXISTS demo_class
(
    `id`               int          NOT NULL AUTO_INCREMENT,
    `class_name`       varchar(30)  not null,
    `class_desc`       varchar(120)          default null,
    `grade_id`         integer      not null,
    `class_time_start` varchar(16)  not null,
    `class_time_end`   varchar(16)  not null,
    `created_at`       timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`       timestamp    NULL     DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`       timestamp(0) NULL     DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

INSERT INTO demo_class (id, class_name, class_desc, grade_id, class_time_start, class_time_end, created_at, updated_at,
                        deleted_at)
VALUES (1, 'English', 'English', 1, '9:00', '9:45', '2020-03-26 10:03:45', '2020-03-26 10:03:33', null);
INSERT INTO demo_class (id, class_name, class_desc, grade_id, class_time_start, class_time_end, created_at, updated_at,
                        deleted_at)
VALUES (2, 'Mathematics', 'Mathematics', 2, '10:00', '10:45', '2020-03-26 10:04:03', '2020-03-26 10:03:49', null);

CREATE TABLE IF NOT EXISTS demo_student
(
    `id`         int         NOT NULL AUTO_INCREMENT,
    `stu_name`   varchar(30) not null,
    `stu_age`    integer     not null,
    `stu_sex`    integer          default 0 check (stu_sex in (0, 1, 2)), -- 0:unknown 1:Male 2:Female
    `created_at` timestamp   NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp   NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` timestamp   NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

INSERT INTO demo_student (id, stu_name, stu_age, stu_sex, created_at, updated_at, deleted_at)
VALUES (1, 'Marry Kamilah', 10, 2, '2020-03-26 10:04:30', '2020-03-26 10:04:19', null);
INSERT INTO demo_student (id, stu_name, stu_age, stu_sex, created_at, updated_at, deleted_at)
VALUES (2, 'Li Reece', 10, 1, '2020-03-26 10:04:42', '2020-03-26 10:04:33', null);
INSERT INTO demo_student (id, stu_name, stu_age, stu_sex, created_at, updated_at, deleted_at)
VALUES (3, 'Lily Sandy', 10, 2, '2020-03-26 10:04:49', '2020-03-26 10:04:44', null);
INSERT INTO demo_student (id, stu_name, stu_age, stu_sex, created_at, updated_at, deleted_at)
VALUES (4, 'Yii Class', 10, 2, '2020-03-26 10:05:06', '2020-03-26 10:04:50', null);

CREATE TABLE IF NOT EXISTS demo_student_class
(
    `id`         int       NOT NULL AUTO_INCREMENT,
    `class_id`   integer   not null,
    `stu_id`     integer   not null,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS demo_student_score
(
    `id`         int       NOT NULL AUTO_INCREMENT,
    `stu_id`     integer   not null,
    `class_id`   integer   not null,
    `stu_score`  int       not null,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;