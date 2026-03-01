-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE projects.projects 
ADD CONSTRAINT name_unique
UNIQUE (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE projects.projects
DROP CONSTRAINT name_unique;
-- +goose StatementEnd
