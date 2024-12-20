-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- Создание типа ENUM для ролей пользователей
CREATE TYPE user_role AS ENUM ('admin', 'organizer', 'participant');

-- Создание таблицы пользователей
CREATE TABLE users (
                       user_id SERIAL PRIMARY KEY,
                       username VARCHAR(50) NOT NULL UNIQUE,
                       email VARCHAR(100) NOT NULL UNIQUE,
                       password_hash VARCHAR(255) NOT NULL,
                       role user_role NOT NULL DEFAULT 'participant',
                       date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы виртуальных комнат
CREATE TABLE virtual_rooms (
                               room_id SERIAL PRIMARY KEY,
                               room_name VARCHAR(100) NOT NULL,
                               capacity INT NOT NULL
);

-- Создание таблицы мероприятий
CREATE TABLE events (
                        event_id SERIAL PRIMARY KEY,
                        event_name VARCHAR(100) NOT NULL,
                        description TEXT,
                        organizer_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
                        start_time TIMESTAMP NOT NULL,
                        end_time TIMESTAMP NOT NULL,
                        virtual_room_id INT REFERENCES virtual_rooms(room_id) ON DELETE SET NULL
);

-- Создание таблицы оборудования
CREATE TABLE equipment (
                           equipment_id SERIAL PRIMARY KEY,
                           equipment_name VARCHAR(100) NOT NULL,
                           user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE
);

-- Создание типа ENUM для ролей участников мероприятий
CREATE TYPE event_role AS ENUM ('participant', 'speaker', 'moderator');

-- Создание таблицы участников мероприятий
CREATE TABLE event_participants (
                                    event_id INT NOT NULL REFERENCES events(event_id) ON DELETE CASCADE,
                                    user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
                                    role_in_event event_role NOT NULL DEFAULT 'participant',
                                    PRIMARY KEY (event_id, user_id)
);

-- Создание типа ENUM для типов билетов
CREATE TYPE ticket_type AS ENUM ('standard', 'VIP');

-- Создание таблицы билетов
CREATE TABLE tickets (
                         ticket_id SERIAL PRIMARY KEY,
                         event_id INT NOT NULL REFERENCES events(event_id) ON DELETE CASCADE,
                         user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
                         ticket_type ticket_type NOT NULL DEFAULT 'standard',
                         price NUMERIC(10, 2) NOT NULL,
                         purchase_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы спонсоров
CREATE TABLE sponsors (
                          sponsor_id SERIAL PRIMARY KEY,
                          sponsor_name VARCHAR(100) NOT NULL,
                          contact_info VARCHAR(255),
                          event_id INT NOT NULL REFERENCES events(event_id) ON DELETE CASCADE,
                          contribution_amount NUMERIC(10, 2) NOT NULL
);

-- Создание таблицы логов
CREATE TABLE logs (
                      log_id SERIAL PRIMARY KEY,
                      table_name VARCHAR(50) NOT NULL,
                      operation_type VARCHAR(50) NOT NULL,
                      record_id INT NOT NULL,
                      changed_data JSONB,
                      timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание функции для записи логов
CREATE OR REPLACE FUNCTION log_changes()
    RETURNS TRIGGER AS $$
DECLARE
    record_id INT;
BEGIN
    IF TG_TABLE_NAME = 'users' THEN
        record_id := NEW.user_id;
    ELSIF TG_TABLE_NAME = 'events' THEN
        record_id := NEW.event_id;
    ELSE
        RAISE EXCEPTION 'Unsupported table: %', TG_TABLE_NAME;
    END IF;

    INSERT INTO logs (table_name, operation_type, record_id, changed_data, timestamp)
    VALUES (
               TG_TABLE_NAME,
               TG_OP,
               record_id,
               row_to_json(NEW),
               CURRENT_TIMESTAMP
           );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Триггер для вставки и обновления в таблице users
CREATE TRIGGER log_users_changes
    AFTER INSERT OR UPDATE ON users
                        FOR EACH ROW
                        EXECUTE FUNCTION log_changes();

-- Триггер для вставки и обновления в таблице events
CREATE TRIGGER log_events_changes
    AFTER INSERT OR UPDATE ON events
                        FOR EACH ROW
                        EXECUTE FUNCTION log_changes();

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS log_users_changes ON users;
DROP TRIGGER IF EXISTS log_events_changes ON events;
DROP FUNCTION IF EXISTS log_changes;
DROP TABLE IF EXISTS logs;
DROP TABLE IF EXISTS sponsors;
DROP TABLE IF EXISTS tickets;
DROP TABLE IF EXISTS event_participants;
DROP TABLE IF EXISTS equipment;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS virtual_rooms;
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS ticket_type;
DROP TYPE IF EXISTS event_role;
DROP TYPE IF EXISTS user_role;
-- +goose StatementEnd
