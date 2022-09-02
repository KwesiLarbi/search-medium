CREATE EXTENSION IF NOT EXISTS 'uuid-ossp';

DROP TABLE IF EXISTS article_schema.articles;

DROP SCHEMA IF EXISTS article_schema;

CREATE SCHEMA IF NOT EXISTS article_schema;

CREATE TABLE IF NOT EXISTS article_schema.articles (
    article_id uuid DEFAULT uuid_generate_v4(), 
    title VARCHAR, 
    username VARCHAR, 
    date_posted VARCHAR, 
    tag VARCHAR,
    article_image VARCHAR, 
    article_preview_text VARCHAR,
    user_profile_url VARCHAR,
    article_url VARCHAR,
    time_to_read VARCHAR,
    PRIMARY KEY (article_id)
);

INSERT INTO article_schema.articles (title, username, date_posted, tag, article_image, article_preview_text, user_profile_url, article_url, time_to_read)
  VALUES ('Uploading files directly to MongoDB through a Simple VueJS App', 'Kevin Takano', 'Mar 2', 'Vuejs', '../some/img/here.png', 'The problem of file upload directly to mongoDB wasn’t that much tricky but rather putting it all together. I found a lot of resources for individual parts of the problem', 'https://google.com/', 'https://google.com', '9 min read');