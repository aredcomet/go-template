CREATE TABLE IF NOT EXISTS auth_user
(
    id           BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    first_name   VARCHAR(255),
    last_name    VARCHAR(255),
    username     VARCHAR(150)          NOT NULL,
    email        VARCHAR(254)          NOT NULL,
    phone        VARCHAR(25),
    password     VARCHAR(128)          NOT NULL,
    is_superuser BOOLEAN DEFAULT FALSE NOT NULL,
    is_active    BOOLEAN DEFAULT TRUE  NOT NULL,
    is_staff     BOOLEAN DEFAULT FALSE NOT NULL,
    is_bot       BOOLEAN DEFAULT FALSE NOT NULL,
    last_login   TIMESTAMP WITH TIME ZONE,
    created_at   TIMESTAMP WITH TIME ZONE,
    updated_at   TIMESTAMP WITH TIME ZONE,

    CONSTRAINT "unq-auth_user-username" UNIQUE (username),
    CONSTRAINT "unq-auth_user-email" UNIQUE (email)
);


CREATE TABLE IF NOT EXISTS auth_user_token
(
    id         BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id    BIGINT                   NOT NULL,
    token      VARCHAR(255)             NOT NULL,
    token_type VARCHAR(255)             NOT NULL,
    is_active  BOOLEAN                  NOT NULL DEFAULT TRUE,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,

    CONSTRAINT "fk-auth_user_token-user_id" FOREIGN KEY (user_id) REFERENCES auth_user (id)
);

-- trigger to disable old tokens on user password change or user is made inactive
CREATE OR REPLACE FUNCTION fn_auth_user_token_disable_old_tokens()
    RETURNS TRIGGER AS
$$
BEGIN
    UPDATE auth_user_token
    SET is_active = FALSE
    WHERE user_id = NEW.id
      AND is_active = TRUE;
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;

CREATE TRIGGER "trg-auth_user_token-disable_old_tokens"
    AFTER UPDATE OF password, is_active
    ON auth_user
    FOR EACH ROW
EXECUTE FUNCTION fn_auth_user_token_disable_old_tokens();
