-- +goose Up
ALTER TABLE tools ADD COLUMN brand TEXT;

-- +goose Down
ALTER TABLE tools DROP COLUMN brand;