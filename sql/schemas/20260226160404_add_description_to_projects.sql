-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE projects.projects 
ADD COLUMN description TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE projects.projects
DROP COLUMN description;
-- +goose StatementEnd
