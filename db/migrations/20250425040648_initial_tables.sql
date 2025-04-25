CREATE TABLE employees (
    id BIGSERIAL PRIMARY KEY,
    name TEXT,
    email TEXT UNIQUE NOT NULL,
    is_admin BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT TRUE,
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
    description TEXT,
    status TEXT,
    needs_maintenance BOOLEAN DEFAULT FALSE,
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
    needs_maintenance BOOLEAN DEFAULT FALSE,
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
