-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE widgets (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    
    -- Audit Fields
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE components (
    id SERIAL PRIMARY KEY,
    widget_id INTEGER NOT NULL REFERENCES widgets(id),
    name VARCHAR(100) NOT NULL,
    complexity INTEGER DEFAULT 1,

    -- Audit Fields
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE widgets;
DROP TABLE components;

DROP SEQUENCE IF EXISTS widgets_id_seq;
DROP SEQUENCE IF EXISTS components_id_seq;
-- +goose StatementEnd
