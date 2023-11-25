DROP TABLE IF EXISTS survey;

DROP TABLE IF EXISTS question_type;

DROP TABLE IF EXISTS question;

DROP TABLE IF EXISTS answer;

DROP TABLE IF EXISTS results;



CREATE TABLE IF NOT EXISTS survey
(
    id uuid NOT NULL PRIMARY KEY,
    name text UNIQUE
);

INSERT INTO survey (id, name)
VALUES
    ('1e461708-6b04-45b9-a4fa-77c32c14d982', 'Опрос про товары'),
    ('1e461708-6b04-45b9-a4fa-77c32c14d487', 'Опрос про вид'),
    ('1e461708-6b04-45b9-a4fa-77c32c14d318', 'Опрос про что-то еще');

CREATE TABLE IF NOT EXISTS question_type
(
    id serial NOT NULL PRIMARY KEY,
    type text NOT NULL
);

INSERT INTO question_type (id, type)
VALUES
    (1, 'CSAT'),
    (2, 'NPS'),
    (3, 'CSI');

CREATE TABLE IF NOT EXISTS question
(
    id uuid NOT NULL PRIMARY KEY,
    type int NOT NULL,
    FOREIGN KEY (type) REFERENCES question_type(id) ON DELETE RESTRICT,
    id_survey uuid,
    FOREIGN KEY (id_survey) REFERENCES survey(id) ON DELETE RESTRICT,
    name text UNIQUE
);

INSERT INTO question (id, type, id_survey, name)
VALUES
    ('1e461708-6b04-45b9-a4fa-77c32c14d382', 1, '1e461708-6b04-45b9-a4fa-77c32c14d982', 'Крутой сайт?'),
    ('1e461708-6b04-45b9-a4fa-77c32c14d387', 1,'1e461708-6b04-45b9-a4fa-77c32c14d982', 'Круто дизайн?'),
    ('1e461708-6b04-45b9-a4fa-77c32c14d388', 2,'1e461708-6b04-45b9-a4fa-77c32c14d982', 'Ваше мнение?');

CREATE TABLE IF NOT EXISTS results
(
    id uuid NOT NULL PRIMARY KEY,
    user_id uuid,
    survey_id uuid,
    FOREIGN KEY (survey_id) REFERENCES survey(id) ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS answer
(
    id uuid NOT NULL PRIMARY KEY,
    question uuid,
    FOREIGN KEY (question) REFERENCES question(id) ON DELETE RESTRICT,
    result_id uuid,
    FOREIGN KEY (result_id) REFERENCES results(id) ON DELETE RESTRICT,
    answer int
);
