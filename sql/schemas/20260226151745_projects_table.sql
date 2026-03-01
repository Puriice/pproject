-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE projects.projects (
	id		UUID	UNIQUE NOT NULL DEFAULT gen_random_uuid(),
	name	TEXT	NOT NULL,
	picture	TEXT,

	PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE projects.projects;
-- +goose StatementEnd
