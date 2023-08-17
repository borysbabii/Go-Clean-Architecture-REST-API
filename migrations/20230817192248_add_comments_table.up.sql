DROP TABLE IF EXISTS comments CASCADE;

CREATE TABLE comments
(
    comment_id UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    author_id  UUID                                               NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
    news_id    UUID                                               NOT NULL REFERENCES news (news_id) ON DELETE CASCADE,
    message    VARCHAR(1024)                                      NOT NULL CHECK ( message <> '' ),
    likes      BIGINT                   DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);