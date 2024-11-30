package banner_repository

import (
	"context"
	"fmt"
	banner_model "tasteplorer-internal-api/app/model/banner"
	"tasteplorer-internal-api/platform/database"
	"time"

	"github.com/jackc/pgx/v4"
)

func CreateBanner(banner *banner_model.Banner) error {
	banner.CreatedAt = time.Now()
	banner.UpdatedAt = banner.CreatedAt

	sql := `INSERT INTO banners (title, image, created_at, updated_at) 
	VALUES ($1, $2, $3, $4)
	RETURNING id, title, image, created_at, updated_at, deleted_at`

	err := database.DB.QueryRow(context.Background(), sql, banner.Title, banner.Image, banner.CreatedAt, banner.UpdatedAt).Scan(&banner.ID, &banner.Title, &banner.Image, &banner.CreatedAt, &banner.UpdatedAt, &banner.DeletedAt)

	return err
}

func GetBannerById(id uint) (*banner_model.Banner, error) {
	var banner banner_model.Banner
	err := database.DB.QueryRow(context.Background(), "SELECT id, title, image, created_at, updated_at, deleted_at FROM banners WHERE deleted_at IS NULL AND id = $1", id).Scan(&banner.ID, &banner.Title, &banner.Image, &banner.CreatedAt, &banner.UpdatedAt, &banner.DeletedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("banner not found: %v", err)
		}

		return nil, err
	}

	return &banner, nil
}

func GetAllBanner(page uint, pageSize uint, searchKeyword string) ([]banner_model.Banner, int, error) {
	sql := `
		SELECT id, title, image, created_at, updated_at, deleted_at
		FROM banners
		WHERE title ILIKE $1
		AND deleted_at IS NULL
		ORDER BY created_at
		LIMIT $2 OFFSET $3
	`

	countQuery := `SELECT COUNT(*)
		FROM banners
		WHERE title ILIKE $1
		AND deleted_at IS NULL`

	rows, err := database.DB.Query(context.Background(), sql, "%"+searchKeyword+"%", pageSize, page)

	if err != nil {
		return nil, 0, fmt.Errorf("error querying banner data: %w", err)
	}

	defer rows.Close()

	var banners []banner_model.Banner

	for rows.Next() {
		var banner banner_model.Banner
		if err := rows.Scan(&banner.ID, &banner.Title, &banner.Image, &banner.CreatedAt, &banner.UpdatedAt, &banner.DeletedAt); err != nil {
			return nil, 0, fmt.Errorf("error scaning rows: %v", err)
		}

		banners = append(banners, banner)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating rows: %w", err)
	}

	if rows == nil {
		return []banner_model.Banner{}, 0, nil
	}

	var totalCount int
	err = database.DB.QueryRow(context.Background(), countQuery, "%"+searchKeyword+"%").Scan(&totalCount)

	if err != nil {
		return nil, 0, err
	}

	// total := len(banners)

	return banners, totalCount, nil
}

func UpdateBanner(id uint, banner *banner_model.Banner) (*banner_model.Banner, error) {
	banner.UpdatedAt = time.Now()

	query := `
		UPDATE banners
		SET title = $1, image = $2, updated_at = $3
		WHERE id=$4
		RETURNING id
	`

	var updatedId uint
	err := database.DB.QueryRow(context.Background(), query, banner.Title, banner.Image, banner.UpdatedAt, id).Scan(&updatedId)

	if err != nil {
		return nil, fmt.Errorf("error updating banner: %v", err)
	}

	if updatedId == 0 {
		return nil, fmt.Errorf("banner not found")
	}

	err = database.DB.QueryRow(context.Background(), "SELECT id, title, image, created_at, updated_at, deleted_at FROM banners WHERE deleted_at IS NULL AND id = $1", id).Scan(&banner.ID, &banner.Title, &banner.Image, &banner.CreatedAt, &banner.UpdatedAt, &banner.DeletedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("banner not found: %v", err)
		}

		return nil, err
	}

	return banner, nil
}

func DeleteBanner(id uint) error {
	var banner banner_model.Banner
	err := database.DB.QueryRow(context.Background(), "SELECT id, title, image, created_at, updated_at, deleted_at FROM banners WHERE deleted_at IS NULL AND id = $1", id).Scan(&banner.ID, &banner.Title, &banner.Image, &banner.CreatedAt, &banner.UpdatedAt, &banner.DeletedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("banner not found: %v", err)
		}

		return err
	}

	query := `
		UPDATE banners
		SET deleted_at = $1
		WHERE id = $2
	`

	currentTime := time.Now()
	result, err := database.DB.Exec(context.Background(), query, currentTime, id)

	if err != nil {
		return fmt.Errorf("error updating banner: %v", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("error updating banner: no rows in result set")
	}

	return nil
}
