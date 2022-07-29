ALTER TABLE blog_from_rss_items Modify column posted_at DATE;  
ALTER TABLE blog_from_manual_items Modify column posted_at DATE;  
ALTER TABLE rss_feeds DROP COLUMN created_at;

