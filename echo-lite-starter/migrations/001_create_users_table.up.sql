CREATE TABLE IF NOT EXISTS public.users (
    id UUID DEFAULT uuid_generate_v4(),
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'user',
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_email_key UNIQUE (email)
);

-- insert 1 dummy user
-- password= password
INSERT INTO public.users (email, password, role, created_at, updated_at)
VALUES ('sigit.priadi@vokal.ai', '$2a$10$kPYbvrcymOCcfpr727sljuRF6e.z6K9P.eN246b5BX2yGh8rocmUW', 'user', now(), now());