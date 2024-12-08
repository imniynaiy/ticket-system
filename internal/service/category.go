package service

func GetCategoryList() (categories []string, err error) {
	err = getPostRepo().Distinct("category").Find(&categories).Error

	return categories, err
}
