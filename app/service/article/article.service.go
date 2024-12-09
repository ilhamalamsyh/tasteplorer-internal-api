package article_service

import (
	"errors"
	article_dto "tasteplorer-internal-api/app/dto/article"
	article_model "tasteplorer-internal-api/app/model/article"
	article_repository "tasteplorer-internal-api/app/repository/article"
)

func CreateArticleService(articleRequestDto *article_dto.ArticleRequestDto) (*article_dto.ArticleDto, error) {
	article := &article_model.Article{
		Title:       articleRequestDto.Title,
		Description: articleRequestDto.Description,
		ImageUrl:    articleRequestDto.ImageUrl,
	}

	err := article_repository.CreateArticle(article)

	if err != nil {
		return nil, errors.New("something went wrong")
	}

	articleDto := &article_dto.ArticleDto{
		ID:          article.ID,
		Title:       article.Title,
		Description: article.Description,
		ImageUrl:    article.ImageUrl,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
		DeletedAt:   article.DeletedAt,
	}

	return articleDto, nil
}

func FindAllArticleService(page uint, pageSize uint, search string) ([]article_dto.ArticleDto, int, error) {
	offset := pageSize * page

	articles, total, err := article_repository.GetAllArticle(offset, pageSize, search)

	if err != nil {
		return nil, total, err
	}

	var articleDtos []article_dto.ArticleDto

	if len(articles) == 0 {
		return []article_dto.ArticleDto{}, total, nil
	}

	for _, article := range articles {
		articleDtos = append(articleDtos, article_dto.ArticleDto{
			ID:          article.ID,
			Title:       article.Title,
			Description: article.Description,
			ImageUrl:    article.ImageUrl,
			CreatedAt:   article.CreatedAt,
			UpdatedAt:   article.UpdatedAt,
			DeletedAt:   article.DeletedAt,
		})
	}

	return articleDtos, total, nil
}

func ArticleDetailService(id uint) (*article_dto.ArticleDto, error) {
	article, err := article_repository.GetArticleById(id)

	if err != nil {
		return nil, err
	}

	articleDto := &article_dto.ArticleDto{
		ID:          article.ID,
		Title:       article.Title,
		Description: article.Description,
		ImageUrl:    article.ImageUrl,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
		DeletedAt:   article.DeletedAt,
	}

	return articleDto, nil
}

func UpdateArticleService(id uint, articleRequestDto *article_dto.ArticleRequestDto) (*article_dto.ArticleDto, error) {
	article := article_model.Article{
		Title:       articleRequestDto.Title,
		Description: articleRequestDto.Description,
		ImageUrl:    articleRequestDto.ImageUrl,
	}

	result, err := article_repository.UpdateArticle(id, &article)
	if err != nil {
		return nil, err
	}

	articleDto := &article_dto.ArticleDto{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		ImageUrl:    result.ImageUrl,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
		DeletedAt:   result.DeletedAt,
	}

	return articleDto, nil
}

func DeleteArticleService(id uint) error {
	err := article_repository.DeleteArticle(id)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
