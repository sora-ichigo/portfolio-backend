CREATE TABLE IF NOT EXISTS blog_from_manual_items
(
  id VARCHAR(36) NOT NULL PRIMARY KEY ,
  title VARCHAR(255) NOT NULL,
  posted_at DATE NOT NULL,
  site_url VARCHAR(255) NOT NULL,
  thumbnail_url VARCHAR(255) NOT NULL,
  service_name VARCHAR(36) NOT NULL
)
;
