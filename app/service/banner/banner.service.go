package banner_service

import (
	"errors"
	banner_dto "tasteplorer-internal-api/app/dto/banner"
	banner_model "tasteplorer-internal-api/app/model/banner"
	banner_repository "tasteplorer-internal-api/app/repository/banner"
)

func CreateBannerService(bannerRequestDto *banner_dto.BannerRequestDto) (*banner_dto.BannerDto, error) {
	banner := &banner_model.Banner{
		Title: bannerRequestDto.Title,
		Image: bannerRequestDto.Image,
	}

	err := banner_repository.CreateBanner(banner)

	if err != nil {
		return nil, errors.New("Something went wrong.")
	}

	bannerDto := &banner_dto.BannerDto{
		ID:        banner.ID,
		Title:     banner.Title,
		Image:     banner.Image,
		CreatedAt: banner.CreatedAt,
		UpdatedAt: banner.UpdatedAt,
		DeletedAt: banner.DeletedAt,
	}

	return bannerDto, nil
}

func FindAllBannerService(page uint, pageSize uint, search string) ([]banner_dto.BannerDto, int, error) {
	offset := pageSize * page

	banners, total, err := banner_repository.GetAllBanner(offset, pageSize, search)

	if err != nil {
		return nil, total, err
	}

	var bannerDtos []banner_dto.BannerDto

	if len(banners) == 0 {
		return []banner_dto.BannerDto{}, total, nil
	}

	for _, banner := range banners {
		bannerDtos = append(bannerDtos, banner_dto.BannerDto{
			ID:        banner.ID,
			Title:     banner.Title,
			Image:     banner.Image,
			CreatedAt: banner.CreatedAt,
			UpdatedAt: banner.UpdatedAt,
			DeletedAt: banner.DeletedAt,
		})
	}

	return bannerDtos, total, nil
}

func BannerDetailService(id uint) (*banner_dto.BannerDto, error) {
	banner, err := banner_repository.GetBannerById(id)

	if err != nil {
		return nil, err
	}

	bannerDto := &banner_dto.BannerDto{
		ID:        banner.ID,
		Title:     banner.Title,
		Image:     banner.Image,
		CreatedAt: banner.CreatedAt,
		UpdatedAt: banner.UpdatedAt,
		DeletedAt: banner.DeletedAt,
	}

	return bannerDto, nil
}

func UpdateBannerService(id uint, bannerRequestDto *banner_dto.BannerRequestDto) (*banner_dto.BannerDto, error) {
	banner := banner_model.Banner{
		Title: bannerRequestDto.Title,
		Image: bannerRequestDto.Image,
	}

	result, err := banner_repository.UpdateBanner(id, &banner)
	if err != nil {
		return nil, err
	}

	bannerDto := &banner_dto.BannerDto{
		ID:        result.ID,
		Title:     result.Title,
		Image:     result.Image,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		DeletedAt: result.DeletedAt,
	}

	return bannerDto, nil
}

func DeleteBannerService(id uint) error {
	err := banner_repository.DeleteBanner(id)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
