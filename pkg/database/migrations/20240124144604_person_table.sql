-- +goose Up
-- +goose StatementBegin
CREATE TABLE People (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        first_name TEXT NOT NULL,
        last_name TEXT NOT NULL,
        patronymic TEXT,
        age INTEGER,
        gender TEXT,
        nationality TEXT
);

-- Создаем индексы
CREATE INDEX idx_first_name ON People (first_name);
CREATE INDEX idx_last_name ON People (last_name);
CREATE INDEX idx_patronymic ON People (patronymic);
CREATE INDEX idx_age ON People (age);
CREATE INDEX idx_gender ON People (gender);
CREATE INDEX idx_nationality ON People (nationality);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE People;
-- +goose StatementEnd
