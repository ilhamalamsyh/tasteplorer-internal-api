package article_repository

import (
	"context"
	"fmt"
	article_model "tasteplorer-internal-api/app/model/article"
	"tasteplorer-internal-api/platform/database"
	"time"

	"github.com/jackc/pgx/v4"
)

func CreateArticle(article *article_model.Article) error {
	article.CreatedAt = time.Now()
	article.UpdatedAt = article.CreatedAt

	sql := `INSERT INTO articles (title, description, image_url, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, title, description, image_url, created_at, updated_at, deleted_at`

	err := database.DB.QueryRow(context.Background(), sql, article.Title, article.Description, article.ImageUrl, article.CreatedAt, article.UpdatedAt).Scan(&article.ID, &article.Title, &article.Description, &article.ImageUrl, &article.CreatedAt, &article.UpdatedAt, &article.DeletedAt)

	return err
}

func GetArticleById(id uint) (*article_model.Article, error) {
	var article article_model.Article
	err := database.DB.QueryRow(context.Background(), "SELECT id, title, description, image_url, created_at, updated_at, deleted_at FROM articles WHERE deleted_at IS NULL AND id = $1", id).Scan(&article.ID, &article.Title, &article.Description, &article.ImageUrl, &article.CreatedAt, &article.UpdatedAt, &article.DeletedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("article not found: %v", err)
		}

		return nil, err
	}

	return &article, nil
}

func GetAllArticle(page uint, pageSize uint, searchKeyword string) ([]article_model.Article, int, error) {
	sql := `
		SELECT id, title, description, image_url, created_at, updated_at, deleted_at
		FROM articles
		WHERE title ILIKE $1
		AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	countQuery := `SELECT COUNT(*)
		FROM articles
		WHERE title ILIKE $1
		AND deleted_at IS NULL`

	rows, err := database.DB.Query(context.Background(), sql, "%"+searchKeyword+"%", pageSize, page)

	if err != nil {
		return nil, 0, fmt.Errorf("error querying article data: %w", err)
	}

	defer rows.Close()

	var articles []article_model.Article

	for rows.Next() {
		var article article_model.Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Description, &article.ImageUrl, &article.CreatedAt, &article.UpdatedAt, &article.DeletedAt); err != nil {
			return nil, 0, fmt.Errorf("error scaning rows: %v", err)
		}

		articles = append(articles, article)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating rows: %w", err)
	}

	if rows == nil {
		return []article_model.Article{}, 0, nil
	}

	var totalCount int
	err = database.DB.QueryRow(context.Background(), countQuery, "%"+searchKeyword+"%").Scan(&totalCount)

	if err != nil {
		return nil, 0, err
	}

	// total := len(banners)

	return articles, totalCount, nil
}

func UpdateArticle(id uint, article *article_model.Article) (*article_model.Article, error) {
	article.UpdatedAt = time.Now()

	query := `
		UPDATE articles
		SET title = $1, description=$2, image_url = $3, updated_at = $4
		WHERE id=$5
		RETURNING id
	`

	var updatedId uint
	err := database.DB.QueryRow(context.Background(), query, article.Title, article.Description, article.ImageUrl, article.UpdatedAt, id).Scan(&updatedId)

	if err != nil {
		return nil, fmt.Errorf("error updating article: %v", err)
	}

	if updatedId == 0 {
		return nil, fmt.Errorf("article not found")
	}

	err = database.DB.QueryRow(context.Background(), "SELECT id, title, description, image_url, created_at, updated_at, deleted_at FROM articles WHERE deleted_at IS NULL AND id = $1", id).Scan(&article.ID, &article.Title, &article.Description, &article.ImageUrl, &article.CreatedAt, &article.UpdatedAt, &article.DeletedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("article not found: %v", err)
		}

		return nil, err
	}

	return article, nil
}

func DeleteArticle(id uint) error {
	var article article_model.Article
	err := database.DB.QueryRow(context.Background(), "SELECT id, title, description, image_url, created_at, updated_at, deleted_at FROM articles WHERE deleted_at IS NULL AND id = $1", id).Scan(&article.ID, &article.Title, &article.Description, &article.ImageUrl, &article.CreatedAt, &article.UpdatedAt, &article.DeletedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("article not found: %v", err)
		}

		return err
	}

	query := `
		UPDATE articles
		SET deleted_at = $1
		WHERE id = $2
	`

	currentTime := time.Now()
	result, err := database.DB.Exec(context.Background(), query, currentTime, id)

	if err != nil {
		return fmt.Errorf("error updating article: %v", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("error updating article: no rows in result set")
	}

	return nil
}
