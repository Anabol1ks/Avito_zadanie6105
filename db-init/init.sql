CREATE TABLE employee (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TYPE organization_type AS ENUM (
    'IE',
    'LLC',
    'JSC'
);

CREATE TABLE organization (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    type organization_type,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE organization_responsible (
    id SERIAL PRIMARY KEY,
    organization_id INT REFERENCES organization(id) ON DELETE CASCADE,
    user_id INT REFERENCES employee(id) ON DELETE CASCADE
);


INSERT INTO employee (username, first_name, last_name) 
VALUES ('user1', 'Имя', 'Фамилия');

INSERT INTO organization (name, description, type) 
VALUES ('Организация 1', 'Описание', 'LLC');

CREATE TABLE tender (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),  -- Используем UUID вместо SERIAL
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(50) NOT NULL,
    service_type VARCHAR(50),
    organization_id UUID NOT NULL,  -- Тип UUID
    creator_username VARCHAR(50) NOT NULL,  -- Тип VARCHAR(50)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT tender_organization_id_fkey FOREIGN KEY (organization_id)
        REFERENCES organization(id) ON DELETE CASCADE,
    CONSTRAINT tender_creator_username_fkey FOREIGN KEY (creator_username)
        REFERENCES employee(username) ON DELETE CASCADE
);

CREATE TABLE bid (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),  -- Используем UUID вместо SERIAL
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(50) NOT NULL,
    tender_id UUID REFERENCES tender(id) ON DELETE CASCADE,  -- Тип UUID
    organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,  -- Тип UUID
    creator_username VARCHAR(50) REFERENCES employee(username) ON DELETE CASCADE,  -- Тип VARCHAR(50)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
