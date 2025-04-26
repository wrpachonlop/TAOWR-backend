-- +goose Up
CREATE TABLE employees (
    id BIGSERIAL PRIMARY KEY,
    full_name TEXT,
    email TEXT UNIQUE NOT NULL,
    phone TEXT,
    address TEXT,
    postal_code TEXT,
    drivers_license TEXT UNIQUE NOT NULL,
    sin TEXT UNIQUE NOT NULL,
    birth_date TIMESTAMPTZ,
    is_admin BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT TRUE,
    start_date TIMESTAMPTZ,
    job_title TEXT,
    type_contract TEXT,
    salary FLOAT,
    institution_no TEXT,
    account_no TEXT,
    transit_no TEXT,
    bank_account_name TEXT,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE emergency_contacts (
    id BIGSERIAL PRIMARY KEY,
    employee_id BIGINT REFERENCES employees(id),
    name TEXT,
    phone TEXT,
    address TEXT,
    relationship TEXT,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE suppliers (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    contact TEXT,
    email TEXT UNIQUE,
    phone TEXT,
    address TEXT,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE tools (
    id BIGSERIAL PRIMARY KEY,
    name TEXT,
    brand TEXT,
    description TEXT,
    status TEXT,
    needs_maintenance BOOLEAN,
    serial_number TEXT,
    maintenance_period INT,
    employee_id BIGINT REFERENCES employees(id),
    assigned_at TIMESTAMPTZ,
    supplier_id BIGINT REFERENCES suppliers(id),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE trucks (
    id BIGSERIAL PRIMARY KEY,
    license_plate TEXT UNIQUE,
    truck_model TEXT,
    status TEXT,
    needs_maintenance BOOLEAN,
    employee_id BIGINT REFERENCES employees(id),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE maintenance_logs (
    id BIGSERIAL PRIMARY KEY,
    item_id BIGINT NOT NULL,
    item_type TEXT NOT NULL,
    description TEXT,
    date TIMESTAMPTZ,
    performed_by_id BIGINT REFERENCES employees(id),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- +goose Down
DROP TABLE IF EXISTS maintenance_logs;
DROP TABLE IF EXISTS trucks;
DROP TABLE IF EXISTS tools;
DROP TABLE IF EXISTS suppliers;
DROP TABLE IF EXISTS emergency_contacts;
DROP TABLE IF EXISTS employees;
