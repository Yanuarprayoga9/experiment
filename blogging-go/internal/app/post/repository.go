package repository

import (
	"database/sql"
	"fmt"
	models "myblog/internal/model"
	"time"

	"github.com/google/uuid"
)

// PostinganBlogRepository handles database operations for blog posts
type PostinganBlogRepository struct {
	db *sql.DB
}

// NewPostinganBlogRepository creates a new repository instance
func NewPostinganBlogRepository(db *sql.DB) *PostinganBlogRepository {
	return &PostinganBlogRepository{db: db}
}

// Create creates a new blog post
func (r *PostinganBlogRepository) Create(post *models.PostinganBlog) error {
	query := `
		INSERT INTO dbo.PostinganBlog (
			Id, Judul, Slug, Isi, PenulisId, StatusReview, 
			Visibilitas, JumlahTayang, UrlThumbnail, CreatedAt, UpdatedAt
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	post.Id = uuid.New()
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	_, err := r.db.Exec(query,
		post.Id, post.Judul, post.Slug, post.Isi, post.PenulisId,
		post.StatusReview, post.Visibilitas, post.JumlahTayang,
		post.UrlThumbnail, post.CreatedAt, post.UpdatedAt,
	)
	return err
}

// GetByID retrieves a blog post by ID
func (r *PostinganBlogRepository) GetByID(id uuid.UUID) (*models.PostinganBlog, error) {
	query := `
		SELECT Id, Judul, Slug, Isi, PenulisId, StatusReview, 
			   Visibilitas, JumlahTayang, UrlThumbnail, CreatedAt, UpdatedAt
		FROM dbo.PostinganBlog 
		WHERE Id = ? AND DeletedAt IS NULL`

	post := &models.PostinganBlog{}
	err := r.db.QueryRow(query, id).Scan(
		&post.Id, &post.Judul, &post.Slug, &post.Isi, &post.PenulisId,
		&post.StatusReview, &post.Visibilitas, &post.JumlahTayang,
		&post.UrlThumbnail, &post.CreatedAt, &post.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// GetBySlug retrieves a blog post by slug
func (r *PostinganBlogRepository) GetBySlug(slug string) (*models.PostinganBlog, error) {
	query := `
		SELECT Id, Judul, Slug, Isi, PenulisId, StatusReview, 
			   Visibilitas, JumlahTayang, UrlThumbnail, CreatedAt, UpdatedAt
		FROM dbo.PostinganBlog 
		WHERE Slug = ? AND DeletedAt IS NULL`

	post := &models.PostinganBlog{}
	err := r.db.QueryRow(query, slug).Scan(
		&post.Id, &post.Judul, &post.Slug, &post.Isi, &post.PenulisId,
		&post.StatusReview, &post.Visibilitas, &post.JumlahTayang,
		&post.UrlThumbnail, &post.CreatedAt, &post.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// GetAll retrieves blog posts with filtering, searching, and pagination
func (r *PostinganBlogRepository) GetAll(queries map[string]string) ([]*models.PostinganBlog, *models.PaginationInfo, error) {
	qb := services.NewQueryBuilder("dbo.PostinganBlog").
		AllowSort("CreatedAt", "UpdatedAt", "Judul", "JumlahTayang").
		SetSearch("", "Judul", "Isi").
		ApplyFromQuery(queries)

	// Build and execute count query
	countQuery, countArgs := qb.BuildCountQuery()
	var totalCount int
	err := r.db.QueryRow(countQuery, countArgs...).Scan(&totalCount)
	if err != nil {
		return nil, nil, err
	}

	// Build and execute select query
	selectQuery, selectArgs := qb.BuildSelectQuery()
	rows, err := r.db.Query(selectQuery, selectArgs...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var posts []*models.PostinganBlog
	for rows.Next() {
		post := &models.PostinganBlog{}
		err := rows.Scan(
			&post.Id, &post.Judul, &post.Slug, &post.Isi, &post.PenulisId,
			&post.StatusReview, &post.Visibilitas, &post.JumlahTayang,
			&post.UrlThumbnail, &post.CreatedAt, &post.UpdatedAt,
		)
		if err != nil {
			return nil, nil, err
		}
		posts = append(posts, post)
	}

	pagination := &models.PaginationInfo{
		CurrentPage: qb.Page,
		PageSize:    qb.Limit,
		TotalItems:  totalCount,
		TotalPages:  (totalCount + qb.Limit - 1) / qb.Limit,
	}

	return posts, pagination, nil
}

// Update updates a blog post
func (r *PostinganBlogRepository) Update(post *models.PostinganBlog) error {
	query := `
		UPDATE dbo.PostinganBlog 
		SET Judul = ?, Slug = ?, Isi = ?, StatusReview = ?, 
			Visibilitas = ?, JumlahTayang = ?, UrlThumbnail = ?, UpdatedAt = ?
		WHERE Id = ? AND DeletedAt IS NULL`

	post.UpdatedAt = time.Now()
	_, err := r.db.Exec(query,
		post.Judul, post.Slug, post.Isi, post.StatusReview,
		post.Visibilitas, post.JumlahTayang, post.UrlThumbnail,
		post.UpdatedAt, post.Id,
	)
	return err
}

// Delete soft deletes a blog post
func (r *PostinganBlogRepository) Delete(id uuid.UUID) error {
	query := `UPDATE dbo.PostinganBlog SET DeletedAt = ? WHERE Id = ?`
	_, err := r.db.Exec(query, time.Now(), id)
	return err
}

// IncrementViewCount increments the view count of a blog post
func (r *PostinganBlogRepository) IncrementViewCount(id uuid.UUID) error {
	query := `
		UPDATE dbo.PostinganBlog 
		SET JumlahTayang = JumlahTayang + 1, UpdatedAt = ?
		WHERE Id = ? AND DeletedAt IS NULL`
	_, err := r.db.Exec(query, time.Now(), id)
	return err
}

// GetByAuthor retrieves blog posts by author ID
func (r *PostinganBlogRepository) GetByAuthor(authorID uuid.UUID, queries map[string]string) ([]*models.PostinganBlog, *models.PaginationInfo, error) {
	qb := services.NewQueryBuilder("dbo.PostinganBlog").
		AllowSort("CreatedAt", "UpdatedAt", "Judul", "JumlahTayang").
		SetSearch("", "Judul", "Isi").
		AddFilter("PenulisId", authorID).
		ApplyFromQuery(queries)

	// Build and execute count query
	countQuery, countArgs := qb.BuildCountQuery()
	var totalCount int
	err := r.db.QueryRow(countQuery, countArgs...).Scan(&totalCount)
	if err != nil {
		return nil, nil, err
	}

	// Build and execute select query
	selectQuery, selectArgs := qb.BuildSelectQuery()
	rows, err := r.db.Query(selectQuery, selectArgs...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var posts []*models.PostinganBlog
	for rows.Next() {
		post := &models.PostinganBlog{}
		err := rows.Scan(
			&post.Id, &post.Judul, &post.Slug, &post.Isi, &post.PenulisId,
			&post.StatusReview, &post.Visibilitas, &post.JumlahTayang,
			&post.UrlThumbnail, &post.CreatedAt, &post.UpdatedAt,
		)
		if err != nil {
			return nil, nil, err
		}
		posts = append(posts, post)
	}

	pagination := &models.PaginationInfo{
		CurrentPage: qb.Page,
		PageSize:    qb.Limit,
		TotalItems:  totalCount,
		TotalPages:  (totalCount + qb.Limit - 1) / qb.Limit,
	}

	return posts, pagination, nil
}

// TagRepository handles database operations for tags
type TagRepository struct {
	db *sql.DB
}

// NewTagRepository creates a new tag repository instance
func NewTagRepository(db *sql.DB) *TagRepository {
	return &TagRepository{db: db}
}

// Create creates a new tag
func (r *TagRepository) Create(tag *models.Tag) error {
	query := `INSERT INTO dbo.Tag (Id, NamaTag) VALUES (?, ?)`
	tag.Id = uuid.New()
	_, err := r.db.Exec(query, tag.Id, tag.NamaTag)
	return err
}

// GetAll retrieves all tags with pagination
func (r *TagRepository) GetAll(queries map[string]string) ([]*models.Tag, *models.PaginationInfo, error) {
	qb := services.NewQueryBuilder("dbo.Tag").
		AllowSort("NamaTag").
		SetSearch("", "NamaTag").
		ApplyFromQuery(queries)

	// Build and execute count query
	countQuery, countArgs := qb.BuildCountQuery()
	var totalCount int
	err := r.db.QueryRow(countQuery, countArgs...).Scan(&totalCount)
	if err != nil {
		return nil, nil, err
	}

	// Build and execute select query
	selectQuery, selectArgs := qb.BuildSelectQuery()
	rows, err := r.db.Query(selectQuery, selectArgs...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var tags []*models.Tag
	for rows.Next() {
		tag := &models.Tag{}
		err := rows.Scan(&tag.Id, &tag.NamaTag)
		if err != nil {
			return nil, nil, err
		}
		tags = append(tags, tag)
	}

	pagination := &models.PaginationInfo{
		CurrentPage: qb.Page,
		PageSize:    qb.Limit,
		TotalItems:  totalCount,
		TotalPages:  (totalCount + qb.Limit - 1) / qb.Limit,
	}

	return tags, pagination, nil
}

// GetByID retrieves a tag by ID
func (r *TagRepository) GetByID(id uuid.UUID) (*models.Tag, error) {
	query := `SELECT Id, NamaTag FROM dbo.Tag WHERE Id = ?`
	tag := &models.Tag{}
	err := r.db.QueryRow(query, id).Scan(&tag.Id, &tag.NamaTag)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

// PostinganTagRepository handles many-to-many relationship between posts and tags
type PostinganTagRepository struct {
	db *sql.DB
}

// NewPostinganTagRepository creates a new repository instance
func NewPostinganTagRepository(db *sql.DB) *PostinganTagRepository {
	return &PostinganTagRepository{db: db}
}

// AddTagToPost adds a tag to a post
func (r *PostinganTagRepository) AddTagToPost(postID, tagID uuid.UUID) error {
	query := `INSERT INTO dbo.PostinganTag (PostinganId, TagId) VALUES (?, ?)`
	_, err := r.db.Exec(query, postID, tagID)
	return err
}

// RemoveTagFromPost removes a tag from a post
func (r *PostinganTagRepository) RemoveTagFromPost(postID, tagID uuid.UUID) error {
	query := `DELETE FROM dbo.PostinganTag WHERE PostinganId = ? AND TagId = ?`
	_, err := r.db.Exec(query, postID, tagID)
	return err
}

// GetTagsByPostID retrieves all tags for a specific post
func (r *PostinganTagRepository) GetTagsByPostID(postID uuid.UUID) ([]*models.Tag, error) {
	query := `
		SELECT t.Id, t.NamaTag 
		FROM dbo.Tag t
		INNER JOIN dbo.PostinganTag pt ON t.Id = pt.TagId
		WHERE pt.PostinganId = ?`

	rows, err := r.db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []*models.Tag
	for rows.Next() {
		tag := &models.Tag{}
		err := rows.Scan(&tag.Id, &tag.NamaTag)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

// PostinganLikeRepository handles blog post likes
type PostinganLikeRepository struct {
	db *sql.DB
}

// NewPostinganLikeRepository creates a new repository instance
func NewPostinganLikeRepository(db *sql.DB) *PostinganLikeRepository {
	return &PostinganLikeRepository{db: db}
}

// LikePost adds a like to a post
func (r *PostinganLikeRepository) LikePost(postID, userID uuid.UUID) error {
	query := `INSERT INTO dbo.PostinganLike (Id, PostinganId, PenggunaId) VALUES (?, ?, ?)`
	_, err := r.db.Exec(query, uuid.New(), postID, userID)
	return err
}

// UnlikePost removes a like from a post
func (r *PostinganLikeRepository) UnlikePost(postID, userID uuid.UUID) error {
	query := `DELETE FROM dbo.PostinganLike WHERE PostinganId = ? AND PenggunaId = ?`
	_, err := r.db.Exec(query, postID, userID)
	return err
}

// IsLikedByUser checks if a user has liked a specific post
func (r *PostinganLikeRepository) IsLikedByUser(postID, userID uuid.UUID) (bool, error) {
	query := `SELECT COUNT(*) FROM dbo.PostinganLike WHERE PostinganId = ? AND PenggunaId = ?`
	var count int
	err := r.db.QueryRow(query, postID, userID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetLikeCount returns the total number of likes for a post
func (r *PostinganLikeRepository) GetLikeCount(postID uuid.UUID) (int, error) {
	query := `SELECT COUNT(*) FROM dbo.PostinganLike WHERE PostinganId = ?`
	var count int
	err := r.db.QueryRow(query, postID).Scan(&count)
	return count, err
}

// KomentarBlogRepository handles blog comments
type KomentarBlogRepository struct {
	db *sql.DB
}

// NewKomentarBlogRepository creates a new repository instance
func NewKomentarBlogRepository(db *sql.DB) *KomentarBlogRepository {
	return &KomentarBlogRepository{db: db}
}

// Create creates a new comment
func (r *KomentarBlogRepository) Create(comment *models.KomentarBlog) error {
	query := `
		INSERT INTO dbo.KomentarBlog (
			Id, PostinganId, PenggunaId, KomentarIndukId, Konten, CreatedAt
		) VALUES (?, ?, ?, ?, ?, ?)`

	comment.Id = uuid.New()
	comment.CreatedAt = time.Now()

	_, err := r.db.Exec(query,
		comment.Id, comment.PostinganId, comment.PenggunaId,
		comment.KomentarIndukId, comment.Konten, comment.CreatedAt,
	)
	return err
}

// GetByPostID retrieves comments for a specific post
func (r *KomentarBlogRepository) GetByPostID(postID uuid.UUID, queries map[string]string) ([]*models.KomentarBlog, *models.PaginationInfo, error) {
	qb := services.NewQueryBuilder("dbo.KomentarBlog").
		AllowSort("CreatedAt").
		AddFilter("PostinganId", postID).
		ApplyFromQuery(queries)

	// Build and execute count query
	countQuery, countArgs := qb.BuildCountQuery()
	var totalCount int
	err := r.db.QueryRow(countQuery, countArgs...).Scan(&totalCount)
	if err != nil {
		return nil, nil, err
	}

	// Build and execute select query
	selectQuery, selectArgs := qb.BuildSelectQuery()
	rows, err := r.db.Query(selectQuery, selectArgs...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var comments []*models.KomentarBlog
	for rows.Next() {
		comment := &models.KomentarBlog{}
		err := rows.Scan(
			&comment.Id, &comment.PostinganId, &comment.PenggunaId,
			&comment.KomentarIndukId, &comment.Konten, &comment.CreatedAt,
		)
		if err != nil {
			return nil, nil, err
		}
		comments = append(comments, comment)
	}

	pagination := &models.PaginationInfo{
		CurrentPage: qb.Page,
		PageSize:    qb.Limit,
		TotalItems:  totalCount,
		TotalPages:  (totalCount + qb.Limit - 1) / qb.Limit,
	}

	return comments, pagination, nil
}

// Delete soft deletes a comment
func (r *KomentarBlogRepository) Delete(id uuid.UUID) error {
	query := `UPDATE dbo.KomentarBlog SET DeletedAt = ? WHERE Id = ?`
	_, err := r.db.Exec(query, time.Now(), id)
	return err
}

// KolaboratorBlogRepository handles blog collaborators
type KolaboratorBlogRepository struct {
	db *sql.DB
}

// NewKolaboratorBlogRepository creates a new repository instance
func NewKolaboratorBlogRepository(db *sql.DB) *KolaboratorBlogRepository {
	return &KolaboratorBlogRepository{db: db}
}

// AddCollaborator adds a collaborator to a blog post
func (r *KolaboratorBlogRepository) AddCollaborator(collab *models.KolaboratorBlog) error {
	query := `
		INSERT INTO dbo.KolaboratorBlog (
			Id, PostinganId, PenggunaId, PeranDalamPostingan, CreatedAt
		) VALUES (?, ?, ?, ?, ?)`

	collab.Id = uuid.New()
	collab.CreatedAt = time.Now()

	_, err := r.db.Exec(query,
		collab.Id, collab.PostinganId, collab.PenggunaId,
		collab.PeranDalamPostingan, collab.CreatedAt,
	)
	return err
}

// GetCollaboratorsByPostID retrieves collaborators for a specific post
func (r *KolaboratorBlogRepository) GetCollaboratorsByPostID(postID uuid.UUID) ([]*models.KolaboratorBlog, error) {
	query := `
		SELECT Id, PostinganId, PenggunaId, PeranDalamPostingan, CreatedAt
		FROM dbo.KolaboratorBlog 
		WHERE PostinganId = ? AND DeletedAt IS NULL`

	rows, err := r.db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var collaborators []*models.KolaboratorBlog
	for rows.Next() {
		collab := &models.KolaboratorBlog{}
		err := rows.Scan(
			&collab.Id, &collab.PostinganId, &collab.PenggunaId,
			&collab.PeranDalamPostingan, &collab.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		collaborators = append(collaborators, collab)
	}
	return collaborators, nil
}

// RemoveCollaborator removes a collaborator from a blog post
func (r *KolaboratorBlogRepository) RemoveCollaborator(id uuid.UUID) error {
	query := `UPDATE dbo.KolaboratorBlog SET DeletedAt = ? WHERE Id = ?`
	_, err := r.db.Exec(query, time.Now(), id)
	return err
}

// StatistikBlogRepository handles blog statistics
type StatistikBlogRepository struct {
	db *sql.DB
}

// NewStatistikBlogRepository creates a new repository instance
func NewStatistikBlogRepository(db *sql.DB) *StatistikBlogRepository {
	return &StatistikBlogRepository{db: db}
}

// CreateOrUpdate creates or updates daily statistics for a blog post
func (r *StatistikBlogRepository) CreateOrUpdate(stat *models.StatistikBlog) error {
	query := `
		MERGE dbo.StatistikBlog AS target
		USING (SELECT ? AS PostinganId, ? AS TanggalStat, ? AS JumlahTayang, ? AS JumlahKomentar, ? AS JumlahSuka) AS source
		ON target.PostinganId = source.PostinganId AND target.TanggalStat = source.TanggalStat
		WHEN MATCHED THEN
			UPDATE SET JumlahTayang = source.JumlahTayang, JumlahKomentar = source.JumlahKomentar, JumlahSuka = source.JumlahSuka
		WHEN NOT MATCHED THEN
			INSERT (Id, PostinganId, TanggalStat, JumlahTayang, JumlahKomentar, JumlahSuka)
			VALUES (?, source.PostinganId, source.TanggalStat, source.JumlahTayang, source.JumlahKomentar, source.JumlahSuka);`

	if stat.Id == uuid.Nil {
		stat.Id = uuid.New()
	}

	_, err := r.db.Exec(query,
		stat.PostinganId, stat.TanggalStat, stat.JumlahTayang, stat.JumlahKomentar, stat.JumlahSuka,
		stat.Id,
	)
	return err
}

// GetStatsByPostID retrieves statistics for a specific post within date range
func (r *StatistikBlogRepository) GetStatsByPostID(postID uuid.UUID, startDate, endDate time.Time) ([]*models.StatistikBlog, error) {
	query := `
		SELECT Id, PostinganId, TanggalStat, JumlahTayang, JumlahKomentar, JumlahSuka
		FROM dbo.StatistikBlog 
		WHERE PostinganId = ? AND TanggalStat BETWEEN ? AND ?
		ORDER BY TanggalStat DESC`

	rows, err := r.db.Query(query, postID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []*models.StatistikBlog
	for rows.Next() {
		stat := &models.StatistikBlog{}
		err := rows.Scan(
			&stat.Id, &stat.PostinganId, &stat.TanggalStat,
			&stat.JumlahTayang, &stat.JumlahKomentar, &stat.JumlahSuka,
		)
		if err != nil {
			return nil, err
		}
		stats = append(stats, stat)
	}
	return stats, nil
}

// GetTotalStatsByPostID retrieves aggregated statistics for a specific post
func (r *StatistikBlogRepository) GetTotalStatsByPostID(postID uuid.UUID) (*models.StatistikBlog, error) {
	query := `
		SELECT PostinganId, SUM(JumlahTayang) as TotalTayang, 
			   SUM(JumlahKomentar) as TotalKomentar, SUM(JumlahSuka) as TotalSuka
		FROM dbo.StatistikBlog 
		WHERE PostinganId = ?
		GROUP BY PostinganId`

	stat := &models.StatistikBlog{PostinganId: postID}
	err := r.db.QueryRow(query, postID).Scan(
		&stat.PostinganId, &stat.JumlahTayang, &stat.JumlahKomentar, &stat.JumlahSuka,
	)
	if err != nil {
		return nil, err
	}
	return stat, nil
}

// KomentarLikeRepository handles comment likes
type KomentarLikeRepository struct {
	db *sql.DB
}

// NewKomentarLikeRepository creates a new repository instance
func NewKomentarLikeRepository(db *sql.DB) *KomentarLikeRepository {
	return &KomentarLikeRepository{db: db}
}

// LikeComment adds a like to a comment
func (r *KomentarLikeRepository) LikeComment(commentID, userID uuid.UUID) error {
	query := `INSERT INTO dbo.KomentarLike (Id, KomentarId, PenggunaId, CreatedAt) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, uuid.New(), commentID, userID, time.Now())
	return err
}

// UnlikeComment removes a like from a comment
func (r *KomentarLikeRepository) UnlikeComment(commentID, userID uuid.UUID) error {
	query := `DELETE FROM dbo.KomentarLike WHERE KomentarId = ? AND PenggunaId = ?`
	_, err := r.db.Exec(query, commentID, userID)
	return err
}

// IsLikedByUser checks if a user has liked a specific comment
func (r *KomentarLikeRepository) IsLikedByUser(commentID, userID uuid.UUID) (bool, error) {
	query := `SELECT COUNT(*) FROM dbo.KomentarLike WHERE KomentarId = ? AND PenggunaId = ?`
	var count int
	err := r.db.QueryRow(query, commentID, userID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetLikeCount returns the total number of likes for a comment
func (r *KomentarLikeRepository) GetLikeCount(commentID uuid.UUID) (int, error) {
	query := `SELECT COUNT(*) FROM dbo.KomentarLike WHERE KomentarId = ?`
	var count int
	err := r.db.QueryRow(query, commentID).Scan(&count)
	return count, err
}

// BlogAggregateRepository provides complex queries and aggregations
type BlogAggregateRepository struct {
	db *sql.DB
}

// NewBlogAggregateRepository creates a new repository instance
func NewBlogAggregateRepository(db *sql.DB) *BlogAggregateRepository {
	return &BlogAggregateRepository{db: db}
}

// GetPostWithDetails retrieves a blog post with all related details
func (r *BlogAggregateRepository) GetPostWithDetails(postID uuid.UUID, userID *uuid.UUID) (*models.BlogPostWithDetails, error) {
	// Main post query
	postQuery := `
		SELECT p.Id, p.Judul, p.Slug, p.Isi, p.PenulisId, p.StatusReview, 
			   p.Visibilitas, p.JumlahTayang, p.UrlThumbnail, p.CreatedAt, p.UpdatedAt,
			   u.Nama as PenulisNama, u.Email as PenulisEmail
		FROM dbo.PostinganBlog p
		LEFT JOIN dbo.Pengguna u ON p.PenulisId = u.Id
		WHERE p.Id = ? AND p.DeletedAt IS NULL`

	post := &models.BlogPostWithDetails{}
	var penulisNama, penulisEmail sql.NullString

	err := r.db.QueryRow(postQuery, postID).Scan(
		&post.Id, &post.Judul, &post.Slug, &post.Isi, &post.PenulisId,
		&post.StatusReview, &post.Visibilitas, &post.JumlahTayang,
		&post.UrlThumbnail, &post.CreatedAt, &post.UpdatedAt,
		&penulisNama, &penulisEmail,
	)
	if err != nil {
		return nil, err
	}

	// Set author info if available
	if penulisNama.Valid {
		post.Penulis = &models.Pengguna{
			Id:    post.PenulisId,
			Nama:  penulisNama.String,
			Email: penulisEmail.String,
		}
	}

	// Get tags
	tagQuery := `
		SELECT t.Id, t.NamaTag 
		FROM dbo.Tag t
		INNER JOIN dbo.PostinganTag pt ON t.Id = pt.TagId
		WHERE pt.PostinganId = ?`

	tagRows, err := r.db.Query(tagQuery, postID)
	if err != nil {
		return nil, err
	}
	defer tagRows.Close()

	for tagRows.Next() {
		tag := &models.Tag{}
		err := tagRows.Scan(&tag.Id, &tag.NamaTag)
		if err != nil {
			return nil, err
		}
		post.Tags = append(post.Tags, tag)
	}

	// Get like count
	likeCountQuery := `SELECT COUNT(*) FROM dbo.PostinganLike WHERE PostinganId = ?`
	err = r.db.QueryRow(likeCountQuery, postID).Scan(&post.LikeCount)
	if err != nil {
		return nil, err
	}

	// Get comment count
	commentCountQuery := `SELECT COUNT(*) FROM dbo.KomentarBlog WHERE PostinganId = ? AND DeletedAt IS NULL`
	err = r.db.QueryRow(commentCountQuery, postID).Scan(&post.CommentCount)
	if err != nil {
		return nil, err
	}

	// Check if liked by current user (if user is provided)
	if userID != nil {
		isLikedQuery := `SELECT COUNT(*) FROM dbo.PostinganLike WHERE PostinganId = ? AND PenggunaId = ?`
		var likedCount int
		err = r.db.QueryRow(isLikedQuery, postID, *userID).Scan(&likedCount)
		if err != nil {
			return nil, err
		}
		post.IsLiked = likedCount > 0
	}

	return post, nil
}

// GetCommentsWithDetails retrieves comments for a post with user details and like info
func (r *BlogAggregateRepository) GetCommentsWithDetails(postID uuid.UUID, userID *uuid.UUID, queries map[string]string) ([]*models.CommentWithDetails, *models.PaginationInfo, error) {
	// Build base query with QueryBuilder
	qb := services.NewQueryBuilder("dbo.KomentarBlog k").
		AllowSort("CreatedAt").
		AddFilter("k.PostinganId", postID).
		ApplyFromQuery(queries)

	// Custom select fields for join
	selectFields := []string{
		"k.Id", "k.PostinganId", "k.PenggunaId", "k.KomentarIndukId",
		"k.Konten", "k.CreatedAt", "u.Nama", "u.Email",
	}

	// Modify the query to include joins
	baseQuery := fmt.Sprintf(`
		SELECT %s
		FROM dbo.KomentarBlog k
		LEFT JOIN dbo.Pengguna u ON k.PenggunaId = u.Id
		WHERE k.PostinganId = ? AND k.DeletedAt IS NULL`,
		strings.Join(selectFields, ", "))

	// Add search conditions if any
	var args []interface{}
	args = append(args, postID)
	argIndex := 2

	if qb.Search != "" && len(qb.SearchFields) > 0 {
		baseQuery += " AND k.Konten LIKE ?"
		args = append(args, "%"+qb.Search+"%")
		argIndex++
	}

	// Add sorting
	if qb.Sort != "" && qb.isAllowedSort(qb.Sort) {
		baseQuery += fmt.Sprintf(" ORDER BY k.%s %s", qb.Sort, qb.Order)
	}

	// Add pagination
	offset := (qb.Page - 1) * qb.Limit
	baseQuery += fmt.Sprintf(" OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", offset, qb.Limit)

	// Get total count
	countQuery := `
		SELECT COUNT(*) 
		FROM dbo.KomentarBlog k 
		WHERE k.PostinganId = ? AND k.DeletedAt IS NULL`
	
	var totalCount int
	err := r.db.QueryRow(countQuery, postID).Scan(&totalCount)
	if err != nil {
		return nil, nil, err
	}

	// Execute main query
	rows, err := r.db.Query(baseQuery, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var comments []*models.CommentWithDetails
	for rows.Next() {
		comment := &models.CommentWithDetails{}
		var userName, userEmail sql.NullString

		err := rows.Scan(
			&comment.Id, &comment.PostinganId, &comment.PenggunaId,
			&comment.KomentarIndukId, &comment.Konten, &comment.CreatedAt,
			&userName, &userEmail,
		)
		if err != nil {
			return nil, nil, err
		}

		// Set user info if available
		if userName.Valid {
			comment.Pengguna = &models.Pengguna{
				Id:    comment.PenggunaId,
				Nama:  userName.String,
				Email: userEmail.String,
			}
		}

		// Get like count for this comment
		likeCountQuery := `SELECT COUNT(*) FROM dbo.KomentarLike WHERE KomentarId = ?`
		err = r.db.QueryRow(likeCountQuery, comment.Id).Scan(&comment.LikeCount)
		if err != nil {
			return nil, nil, err
		}

		// Check if liked by current user
		if userID != nil {
			isLikedQuery := `SELECT COUNT(*) FROM dbo.KomentarLike WHERE KomentarId = ? AND PenggunaId = ?`
			var likedCount int
			err = r.db.QueryRow(isLikedQuery, comment.Id, *userID).Scan(&likedCount)
			if err != nil {
				return nil, nil, err
			}
			comment.IsLiked = likedCount > 0
		}

		comments = append(comments, comment)
	}

	pagination := &models.PaginationInfo{
		CurrentPage: qb.Page,
		PageSize:    qb.Limit,
		TotalItems:  totalCount,
		TotalPages:  (totalCount + qb.Limit - 1) / qb.Limit,
	}

	return comments, pagination, nil
}

// GetPopularPosts retrieves popular posts based on views, likes, and comments
func (r *BlogAggregateRepository) GetPopularPosts(limit int, days int) ([]*models.BlogPostWithDetails, error) {
	query := `
		SELECT TOP(?) p.Id, p.Judul, p.Slug, p.Isi, p.PenulisId, p.StatusReview, 
			   p.Visibilitas, p.JumlahTayang, p.UrlThumbnail, p.CreatedAt, p.UpdatedAt,
			   u.Nama as PenulisNama, u.Email as PenulisEmail,
			   COALESCE(likes.LikeCount, 0) as LikeCount,
			   COALESCE(comments.CommentCount, 0) as CommentCount
		FROM dbo.PostinganBlog p
		LEFT JOIN dbo.Pengguna u ON p.PenulisId = u.Id
		LEFT JOIN (
			SELECT PostinganId, COUNT(*) as LikeCount 
			FROM dbo.PostinganLike 
			GROUP BY PostinganId
		) likes ON p.Id = likes.PostinganId
		LEFT JOIN (
			SELECT PostinganId, COUNT(*) as CommentCount 
			FROM dbo.KomentarBlog 
			WHERE DeletedAt IS NULL 
			GROUP BY PostinganId
		) comments ON p.Id = comments.PostinganId
		WHERE p.StatusReview = 'diterbitkan' 
			AND p.DeletedAt IS NULL 
			AND p.CreatedAt >= DATEADD(day, ?, GETDATE())
		ORDER BY (p.JumlahTayang + COALESCE(likes.LikeCount, 0) * 2 + COALESCE(comments.CommentCount, 0) * 3) DESC`

	rows, err := r.db.Query(query, limit, -days)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.BlogPostWithDetails
	for rows.Next() {
		post := &models.BlogPostWithDetails{}
		var penulisNama, penulisEmail sql.NullString

		err := rows.Scan(
			&post.Id, &post.Judul, &post.Slug, &post.Isi, &post.PenulisId,
			&post.StatusReview, &post.Visibilitas, &post.JumlahTayang,
			&post.UrlThumbnail, &post.CreatedAt, &post.UpdatedAt,
			&penulisNama, &penulisEmail, &post.LikeCount, &post.CommentCount,
		)
		if err != nil {
			return nil, err
		}

		// Set author info if available
		if penulisNama.Valid {
			post.Penulis = &models.Pengguna{
				Id:    post.PenulisId,
				Nama:  penulisNama.String,
				Email: penulisEmail.String,
			}
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// SearchPosts performs full-text search across posts
func (r *BlogAggregateRepository) SearchPosts(searchTerm string, filters map[string]interface{}, queries map[string]string) ([]*models.BlogPostWithDetails, *models.PaginationInfo, error) {
	qb := services.NewQueryBuilder("dbo.PostinganBlog p").
		AllowSort("CreatedAt", "UpdatedAt", "Judul", "JumlahTayang").
		SetSearch(searchTerm, "p.Judul", "p.Isi").
		ApplyFromQuery(queries)

	// Add additional filters
	for key, value := range filters {
		qb.AddFilter("p."+key, value)
	}

	// Always filter for published posts in search
	qb.AddFilter("p.StatusReview", "diterbitkan")

	// Build custom query with joins
	baseSelectQuery := `
		SELECT p.Id, p.Judul, p.Slug, p.Isi, p.PenulisId, p.StatusReview, 
			   p.Visibilitas, p.JumlahTayang, p.UrlThumbnail, p.CreatedAt, p.UpdatedAt,
			   u.Nama as PenulisNama, u.Email as PenulisEmail,
			   COALESCE(likes.LikeCount, 0) as LikeCount,
			   COALESCE(comments.CommentCount, 0) as CommentCount
		FROM dbo.PostinganBlog p
		LEFT JOIN dbo.Pengguna u ON p.PenulisId = u.Id
		LEFT JOIN (
			SELECT PostinganId, COUNT(*) as LikeCount 
			FROM dbo.PostinganLike 
			GROUP BY PostinganId
		) likes ON p.Id = likes.PostinganId
		LEFT JOIN (
			SELECT PostinganId, COUNT(*) as CommentCount 
			FROM dbo.KomentarBlog 
			WHERE DeletedAt IS NULL 
			GROUP BY PostinganId
		) comments ON p.Id = comments.PostinganId`

	// Build WHERE conditions
	var whereConditions []string
	var args []interface{}
	argIndex := 1

	// Add filters
	for field, value := range qb.Filters {
		whereConditions = append(whereConditions, fmt.Sprintf("%s = ?", field))
		args = append(args, value)
		argIndex++
	}

	// Add search conditions
	if qb.Search != "" {
		searchConditions := []string{
			"p.Judul LIKE ?",
			"p.Isi LIKE ?",
		}
		searchTerm := "%" + qb.Search + "%"
		for range searchConditions {
			args = append(args, searchTerm)
			argIndex++
		}
		whereConditions = append(whereConditions, "("+strings.Join(searchConditions, " OR ")+")")
	}

	// Add soft delete filter
	whereConditions = append(whereConditions, "p.DeletedAt IS NULL")

	baseQuery := baseSelectQuery
	if len(whereConditions) > 0 {
		baseQuery += " WHERE " + strings.Join(whereConditions, " AND ")
	}

	// Get total count
	countQuery := `
		SELECT COUNT(*) 
		FROM dbo.PostinganBlog p 
		WHERE ` + strings.Join(whereConditions, " AND ")

	var totalCount int
	err := r.db.QueryRow(countQuery, args...).Scan(&totalCount)
	if err != nil {
		return nil, nil, err
	}

	// Add sorting and pagination
	if qb.Sort != "" && qb.isAllowedSort(qb.Sort) {
		baseQuery += fmt.Sprintf(" ORDER BY %s %s", qb.Sort, qb.Order)
	}

	offset := (qb.Page - 1) * qb.Limit
	baseQuery += fmt.Sprintf(" OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", offset, qb.Limit)

	// Execute query
	rows, err := r.db.Query(baseQuery, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var posts []*models.BlogPostWithDetails
	for rows.Next() {
		post := &models.BlogPostWithDetails{}
		var penulisNama, penulisEmail sql.NullString

		err := rows.Scan(
			&post.Id, &post.Judul, &post.Slug, &post.Isi, &post.PenulisId,
			&post.StatusReview, &post.Visibilitas, &post.JumlahTayang,
			&post.UrlThumbnail, &post.CreatedAt, &post.UpdatedAt,
			&penulisNama, &penulisEmail, &post.LikeCount, &post.CommentCount,
		)
		if err != nil {
			return nil, nil, err
		}

		// Set author info if available
		if penulisNama.Valid {
			post.Penulis = &models.Pengguna{
				Id:    post.PenulisId,
				Nama:  penulisNama.String,
				Email: penulisEmail.String,
			}
		}

		posts = append(posts, post)
	}

	pagination := &models.PaginationInfo{
		CurrentPage: qb.Page,
		PageSize:    qb.Limit,
		TotalItems:  totalCount,
		TotalPages:  (totalCount + qb.Limit - 1) / qb.Limit,
	}

	return posts, pagination, nil
}