-- migrate:up

-- 1. 大学マスタ
CREATE TABLE universities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    max_period INT DEFAULT 5 NOT NULL
);

-- 2. 学部マスタ（追加）
CREATE TABLE faculties (
    id SERIAL PRIMARY KEY,
    university_id INT NOT NULL REFERENCES universities(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL
);

-- 3. ユーザーテーブル
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    share_code VARCHAR(10) UNIQUE NOT NULL,
    university_id INT REFERENCES universities(id) ON DELETE SET NULL,
    faculty_id INT REFERENCES faculties(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 4. 授業マスタ（シラバス）
CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    university_id INT NOT NULL REFERENCES universities(id) ON DELETE CASCADE,
    faculty_id INT REFERENCES faculties(id) ON DELETE SET NULL,
    subject_name VARCHAR(100) NOT NULL,
    room VARCHAR(50)
);

-- 5. 履修登録テーブル（ユーザと授業の中間テーブル）
CREATE TABLE registrations (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    course_id INT NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    day VARCHAR(10) NOT NULL,
    period INT NOT NULL,
    UNIQUE (user_id, day, period)
);

-- 6. 友達関係テーブル（ユーザとユーザの中間テーブル）
CREATE TABLE friendships (
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    friend_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, friend_id)
);

-- メタデータコメント
COMMENT ON TABLE universities IS '大学マスタ情報';
COMMENT ON TABLE faculties IS '学部マスタ情報';
COMMENT ON TABLE users IS 'アプリを利用するユーザー情報';
COMMENT ON COLUMN users.faculty_id IS '所属学部（任意）';
COMMENT ON TABLE courses IS '大学・学部ごとの授業（シラバス）情報';
COMMENT ON COLUMN courses.faculty_id IS '対象学部（NULLの場合は全学共通科目）';
COMMENT ON TABLE registrations IS 'ユーザーごとの時間割（コマ）登録情報';
COMMENT ON TABLE friendships IS 'ユーザー同士の友達関係（相互登録パターン）';

-- migrate:down
DROP TABLE IF EXISTS friendships;
DROP TABLE IF EXISTS registrations;
DROP TABLE IF EXISTS courses;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS faculties;
DROP TABLE IF EXISTS universities;
