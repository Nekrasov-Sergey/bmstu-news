-- +goose Up
-- +goose StatementBegin
create table db_news
(
    slug               text not null
        constraint db_news_pkey
            primary key,
    title              text,
    author             text,
    preview_text       text,
    content            text,
    reading_time       text,
    published_at       date,
    image              text
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE db_news CASCADE;
-- +goose StatementEnd
