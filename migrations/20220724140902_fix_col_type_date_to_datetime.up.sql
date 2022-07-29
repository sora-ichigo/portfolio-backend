ALTER TABLE blog_from_rss_items Modify column posted_at DATETIME;  
ALTER TABLE blog_from_manual_items Modify column posted_at DATETIME;  
ALTER TABLE rss_feeds ADD COLUMN created_at DATETIME;

