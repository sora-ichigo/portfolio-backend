package main

// func TestHandler(t *testing.T) {
// 	db, err := repository.NewDB()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	deleteAllRssFeeds(t, db)
//
// 	r := models.RSSFeed{
// 		ID:  uuid.New().String(),
// 		URL: "https://zenn.dev/ichigo_dev/feed",
// 	}
// 	err = r.Insert(context.Background(), db, boil.Infer())
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	// r = models.RSSFeed{
// 	// 	ID:  uuid.New().String(),
// 	// 	URL: "https://qiita.com/igsr5/feed",
// 	// }
// 	// err = r.Insert(context.Background(), db, boil.Infer())
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }
//
// 	r = models.RSSFeed{
// 		ID:  uuid.New().String(),
// 		URL: "https://note.com/ichigo341/rss",
// 	}
// 	err = r.Insert(context.Background(), db, boil.Infer())
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	err = handler(events.CloudWatchEvent{})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
//
// func deleteAllRssFeeds(t *testing.T, db *sql.DB) {
// 	t.Helper()
//
// 	if _, err := db.Exec("DELETE FROM rss_feeds"); err != nil {
// 		t.Fatal(err)
// 	}
// }
//
// // func TestGetBlogDataFromRSS(t *testing.T) {
// // 	tests := []struct {
// // 		name    string
// // 		url     string
// // 		isEmpty bool
// // 	}{
// // 		{
// // 			name:    "My Zenn Account",
// // 			url:     "https://zenn.dev/ichigo_dev/feed",
// // 			isEmpty: false,
// // 		},
// // 		// {
// // 		// 	name:    "My Qiita Account",
// // 		// 	url:     "https://qiita.com/igsr5/feed",
// // 		// 	isEmpty: false,
// // 		// },
// // 		{
// // 			name:    "My note Account",
// // 			url:     "https://note.com/ichigo341/rss",
// // 			isEmpty: false,
// // 		},
// // 	}
// //
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			i, err := getBlodDataFromRSSFeed(context.Background(), tt.url)
// // 			if err != nil {
// // 				t.Fatal(err)
// // 			}
// // 			if tt.isEmpty || len(i) <= 0 {
// // 				t.Fatalf("items is empty. url: %s", tt.url)
// // 			}
// // 		})
// // 	}
// // }
// //
// // func TestQiitaCrawler(t *testing.T) {
// // 	items, err := NewQiitaCrawler("https://qiita.com/igsr5/feed").Crawl() // 自分のQiitaアカウント
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// //
// // 	if len(items) <= 0 {
// // 		t.Fatal("items is empty")
// // 	}
// // }
//
// func TestConvertRSSToApiURL(t *testing.T) {
// 	// e.g. rssUrl → https://qiita.com/igsr5/feed
// 	// e.g. igsr5
// 	// e.g. https://qiita.com/api/v2/users/igsr5/items"
// 	rssUrl := "https://qiita.com/igsr5/feed"
// 	apiUrl := convertRSSToApiURL(rssUrl)
//
// 	if apiUrl != "https://qiita.com/api/v2/users/igsr5/items" {
// 		t.Fatal("failed to convert RSS URL to API URL.")
// 	}
// }
//
// func TestGetOgpImageUrl(t *testing.T) {
// 	ogp, err := opengraph.Fetch("https://github.com/")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	fmt.Println(ogp.Image[0].URL, err)
// }
