// Code generated by SQLBoiler 4.10.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// BlogFromRSSItem is an object representing the database table.
type BlogFromRSSItem struct {
	ID           string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title        string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	PostedAt     time.Time `boil:"posted_at" json:"posted_at" toml:"posted_at" yaml:"posted_at"`
	SiteURL      string    `boil:"site_url" json:"site_url" toml:"site_url" yaml:"site_url"`
	ThumbnailURL string    `boil:"thumbnail_url" json:"thumbnail_url" toml:"thumbnail_url" yaml:"thumbnail_url"`
	ServiceName  string    `boil:"service_name" json:"service_name" toml:"service_name" yaml:"service_name"`

	R *blogFromRSSItemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L blogFromRSSItemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BlogFromRSSItemColumns = struct {
	ID           string
	Title        string
	PostedAt     string
	SiteURL      string
	ThumbnailURL string
	ServiceName  string
}{
	ID:           "id",
	Title:        "title",
	PostedAt:     "posted_at",
	SiteURL:      "site_url",
	ThumbnailURL: "thumbnail_url",
	ServiceName:  "service_name",
}

var BlogFromRSSItemTableColumns = struct {
	ID           string
	Title        string
	PostedAt     string
	SiteURL      string
	ThumbnailURL string
	ServiceName  string
}{
	ID:           "blog_from_rss_items.id",
	Title:        "blog_from_rss_items.title",
	PostedAt:     "blog_from_rss_items.posted_at",
	SiteURL:      "blog_from_rss_items.site_url",
	ThumbnailURL: "blog_from_rss_items.thumbnail_url",
	ServiceName:  "blog_from_rss_items.service_name",
}

// Generated where

var BlogFromRSSItemWhere = struct {
	ID           whereHelperstring
	Title        whereHelperstring
	PostedAt     whereHelpertime_Time
	SiteURL      whereHelperstring
	ThumbnailURL whereHelperstring
	ServiceName  whereHelperstring
}{
	ID:           whereHelperstring{field: "`blog_from_rss_items`.`id`"},
	Title:        whereHelperstring{field: "`blog_from_rss_items`.`title`"},
	PostedAt:     whereHelpertime_Time{field: "`blog_from_rss_items`.`posted_at`"},
	SiteURL:      whereHelperstring{field: "`blog_from_rss_items`.`site_url`"},
	ThumbnailURL: whereHelperstring{field: "`blog_from_rss_items`.`thumbnail_url`"},
	ServiceName:  whereHelperstring{field: "`blog_from_rss_items`.`service_name`"},
}

// BlogFromRSSItemRels is where relationship names are stored.
var BlogFromRSSItemRels = struct {
}{}

// blogFromRSSItemR is where relationships are stored.
type blogFromRSSItemR struct {
}

// NewStruct creates a new relationship struct
func (*blogFromRSSItemR) NewStruct() *blogFromRSSItemR {
	return &blogFromRSSItemR{}
}

// blogFromRSSItemL is where Load methods for each relationship are stored.
type blogFromRSSItemL struct{}

var (
	blogFromRSSItemAllColumns            = []string{"id", "title", "posted_at", "site_url", "thumbnail_url", "service_name"}
	blogFromRSSItemColumnsWithoutDefault = []string{"id", "title", "posted_at", "site_url", "thumbnail_url", "service_name"}
	blogFromRSSItemColumnsWithDefault    = []string{}
	blogFromRSSItemPrimaryKeyColumns     = []string{"id"}
	blogFromRSSItemGeneratedColumns      = []string{}
)

type (
	// BlogFromRSSItemSlice is an alias for a slice of pointers to BlogFromRSSItem.
	// This should almost always be used instead of []BlogFromRSSItem.
	BlogFromRSSItemSlice []*BlogFromRSSItem
	// BlogFromRSSItemHook is the signature for custom BlogFromRSSItem hook methods
	BlogFromRSSItemHook func(context.Context, boil.ContextExecutor, *BlogFromRSSItem) error

	blogFromRSSItemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	blogFromRSSItemType                 = reflect.TypeOf(&BlogFromRSSItem{})
	blogFromRSSItemMapping              = queries.MakeStructMapping(blogFromRSSItemType)
	blogFromRSSItemPrimaryKeyMapping, _ = queries.BindMapping(blogFromRSSItemType, blogFromRSSItemMapping, blogFromRSSItemPrimaryKeyColumns)
	blogFromRSSItemInsertCacheMut       sync.RWMutex
	blogFromRSSItemInsertCache          = make(map[string]insertCache)
	blogFromRSSItemUpdateCacheMut       sync.RWMutex
	blogFromRSSItemUpdateCache          = make(map[string]updateCache)
	blogFromRSSItemUpsertCacheMut       sync.RWMutex
	blogFromRSSItemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var blogFromRSSItemAfterSelectHooks []BlogFromRSSItemHook

var blogFromRSSItemBeforeInsertHooks []BlogFromRSSItemHook
var blogFromRSSItemAfterInsertHooks []BlogFromRSSItemHook

var blogFromRSSItemBeforeUpdateHooks []BlogFromRSSItemHook
var blogFromRSSItemAfterUpdateHooks []BlogFromRSSItemHook

var blogFromRSSItemBeforeDeleteHooks []BlogFromRSSItemHook
var blogFromRSSItemAfterDeleteHooks []BlogFromRSSItemHook

var blogFromRSSItemBeforeUpsertHooks []BlogFromRSSItemHook
var blogFromRSSItemAfterUpsertHooks []BlogFromRSSItemHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BlogFromRSSItem) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range blogFromRSSItemAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BlogFromRSSItem) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range blogFromRSSItemBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BlogFromRSSItem) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range blogFromRSSItemAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BlogFromRSSItem) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range blogFromRSSItemBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BlogFromRSSItem) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range blogFromRSSItemAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BlogFromRSSItem) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range blogFromRSSItemBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BlogFromRSSItem) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range blogFromRSSItemAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BlogFromRSSItem) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range blogFromRSSItemBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BlogFromRSSItem) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range blogFromRSSItemAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBlogFromRSSItemHook registers your hook function for all future operations.
func AddBlogFromRSSItemHook(hookPoint boil.HookPoint, blogFromRSSItemHook BlogFromRSSItemHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		blogFromRSSItemAfterSelectHooks = append(blogFromRSSItemAfterSelectHooks, blogFromRSSItemHook)
	case boil.BeforeInsertHook:
		blogFromRSSItemBeforeInsertHooks = append(blogFromRSSItemBeforeInsertHooks, blogFromRSSItemHook)
	case boil.AfterInsertHook:
		blogFromRSSItemAfterInsertHooks = append(blogFromRSSItemAfterInsertHooks, blogFromRSSItemHook)
	case boil.BeforeUpdateHook:
		blogFromRSSItemBeforeUpdateHooks = append(blogFromRSSItemBeforeUpdateHooks, blogFromRSSItemHook)
	case boil.AfterUpdateHook:
		blogFromRSSItemAfterUpdateHooks = append(blogFromRSSItemAfterUpdateHooks, blogFromRSSItemHook)
	case boil.BeforeDeleteHook:
		blogFromRSSItemBeforeDeleteHooks = append(blogFromRSSItemBeforeDeleteHooks, blogFromRSSItemHook)
	case boil.AfterDeleteHook:
		blogFromRSSItemAfterDeleteHooks = append(blogFromRSSItemAfterDeleteHooks, blogFromRSSItemHook)
	case boil.BeforeUpsertHook:
		blogFromRSSItemBeforeUpsertHooks = append(blogFromRSSItemBeforeUpsertHooks, blogFromRSSItemHook)
	case boil.AfterUpsertHook:
		blogFromRSSItemAfterUpsertHooks = append(blogFromRSSItemAfterUpsertHooks, blogFromRSSItemHook)
	}
}

// OneG returns a single blogFromRSSItem record from the query using the global executor.
func (q blogFromRSSItemQuery) OneG(ctx context.Context) (*BlogFromRSSItem, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single blogFromRSSItem record from the query.
func (q blogFromRSSItemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BlogFromRSSItem, error) {
	o := &BlogFromRSSItem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for blog_from_rss_items")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all BlogFromRSSItem records from the query using the global executor.
func (q blogFromRSSItemQuery) AllG(ctx context.Context) (BlogFromRSSItemSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all BlogFromRSSItem records from the query.
func (q blogFromRSSItemQuery) All(ctx context.Context, exec boil.ContextExecutor) (BlogFromRSSItemSlice, error) {
	var o []*BlogFromRSSItem

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BlogFromRSSItem slice")
	}

	if len(blogFromRSSItemAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all BlogFromRSSItem records in the query using the global executor
func (q blogFromRSSItemQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all BlogFromRSSItem records in the query.
func (q blogFromRSSItemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count blog_from_rss_items rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q blogFromRSSItemQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q blogFromRSSItemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if blog_from_rss_items exists")
	}

	return count > 0, nil
}

// BlogFromRSSItems retrieves all the records using an executor.
func BlogFromRSSItems(mods ...qm.QueryMod) blogFromRSSItemQuery {
	mods = append(mods, qm.From("`blog_from_rss_items`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`blog_from_rss_items`.*"})
	}

	return blogFromRSSItemQuery{NewQuery(q)}
}

// FindBlogFromRSSItemG retrieves a single record by ID.
func FindBlogFromRSSItemG(ctx context.Context, iD string, selectCols ...string) (*BlogFromRSSItem, error) {
	return FindBlogFromRSSItem(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindBlogFromRSSItem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBlogFromRSSItem(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*BlogFromRSSItem, error) {
	blogFromRSSItemObj := &BlogFromRSSItem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `blog_from_rss_items` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, blogFromRSSItemObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from blog_from_rss_items")
	}

	if err = blogFromRSSItemObj.doAfterSelectHooks(ctx, exec); err != nil {
		return blogFromRSSItemObj, err
	}

	return blogFromRSSItemObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *BlogFromRSSItem) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BlogFromRSSItem) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no blog_from_rss_items provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(blogFromRSSItemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	blogFromRSSItemInsertCacheMut.RLock()
	cache, cached := blogFromRSSItemInsertCache[key]
	blogFromRSSItemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			blogFromRSSItemAllColumns,
			blogFromRSSItemColumnsWithDefault,
			blogFromRSSItemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(blogFromRSSItemType, blogFromRSSItemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(blogFromRSSItemType, blogFromRSSItemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `blog_from_rss_items` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `blog_from_rss_items` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `blog_from_rss_items` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, blogFromRSSItemPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into blog_from_rss_items")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for blog_from_rss_items")
	}

CacheNoHooks:
	if !cached {
		blogFromRSSItemInsertCacheMut.Lock()
		blogFromRSSItemInsertCache[key] = cache
		blogFromRSSItemInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single BlogFromRSSItem record using the global executor.
// See Update for more documentation.
func (o *BlogFromRSSItem) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the BlogFromRSSItem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BlogFromRSSItem) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	blogFromRSSItemUpdateCacheMut.RLock()
	cache, cached := blogFromRSSItemUpdateCache[key]
	blogFromRSSItemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			blogFromRSSItemAllColumns,
			blogFromRSSItemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update blog_from_rss_items, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `blog_from_rss_items` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, blogFromRSSItemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(blogFromRSSItemType, blogFromRSSItemMapping, append(wl, blogFromRSSItemPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update blog_from_rss_items row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for blog_from_rss_items")
	}

	if !cached {
		blogFromRSSItemUpdateCacheMut.Lock()
		blogFromRSSItemUpdateCache[key] = cache
		blogFromRSSItemUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q blogFromRSSItemQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q blogFromRSSItemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for blog_from_rss_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for blog_from_rss_items")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o BlogFromRSSItemSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BlogFromRSSItemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), blogFromRSSItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `blog_from_rss_items` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, blogFromRSSItemPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in blogFromRSSItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all blogFromRSSItem")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *BlogFromRSSItem) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

var mySQLBlogFromRSSItemUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BlogFromRSSItem) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no blog_from_rss_items provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(blogFromRSSItemColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLBlogFromRSSItemUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	blogFromRSSItemUpsertCacheMut.RLock()
	cache, cached := blogFromRSSItemUpsertCache[key]
	blogFromRSSItemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			blogFromRSSItemAllColumns,
			blogFromRSSItemColumnsWithDefault,
			blogFromRSSItemColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			blogFromRSSItemAllColumns,
			blogFromRSSItemPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert blog_from_rss_items, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`blog_from_rss_items`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `blog_from_rss_items` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(blogFromRSSItemType, blogFromRSSItemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(blogFromRSSItemType, blogFromRSSItemMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for blog_from_rss_items")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(blogFromRSSItemType, blogFromRSSItemMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for blog_from_rss_items")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for blog_from_rss_items")
	}

CacheNoHooks:
	if !cached {
		blogFromRSSItemUpsertCacheMut.Lock()
		blogFromRSSItemUpsertCache[key] = cache
		blogFromRSSItemUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single BlogFromRSSItem record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *BlogFromRSSItem) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single BlogFromRSSItem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BlogFromRSSItem) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BlogFromRSSItem provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), blogFromRSSItemPrimaryKeyMapping)
	sql := "DELETE FROM `blog_from_rss_items` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from blog_from_rss_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for blog_from_rss_items")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q blogFromRSSItemQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q blogFromRSSItemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no blogFromRSSItemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from blog_from_rss_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for blog_from_rss_items")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o BlogFromRSSItemSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BlogFromRSSItemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(blogFromRSSItemBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), blogFromRSSItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `blog_from_rss_items` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, blogFromRSSItemPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from blogFromRSSItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for blog_from_rss_items")
	}

	if len(blogFromRSSItemAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *BlogFromRSSItem) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no BlogFromRSSItem provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BlogFromRSSItem) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBlogFromRSSItem(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BlogFromRSSItemSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty BlogFromRSSItemSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BlogFromRSSItemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BlogFromRSSItemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), blogFromRSSItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `blog_from_rss_items`.* FROM `blog_from_rss_items` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, blogFromRSSItemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BlogFromRSSItemSlice")
	}

	*o = slice

	return nil
}

// BlogFromRSSItemExistsG checks if the BlogFromRSSItem row exists.
func BlogFromRSSItemExistsG(ctx context.Context, iD string) (bool, error) {
	return BlogFromRSSItemExists(ctx, boil.GetContextDB(), iD)
}

// BlogFromRSSItemExists checks if the BlogFromRSSItem row exists.
func BlogFromRSSItemExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `blog_from_rss_items` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if blog_from_rss_items exists")
	}

	return exists, nil
}
