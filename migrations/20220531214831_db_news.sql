-- +goose Up
-- +goose StatementBegin
create table news
(
    slug              varchar(255) not null
        primary key,
    title             varchar(255),
    author            text,
    preview_text      text,
    content           text,
    reading_time      text,
    published_at      timestamp with time zone,
    image             text,
    photo_report      text[],
    similar_news_slug text[],
    tags_title        text[]
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE news;
-- +goose StatementEnd