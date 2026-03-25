-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE items ADD COLUMN remarks TEXT;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE items DROP COLUMN remarks;
