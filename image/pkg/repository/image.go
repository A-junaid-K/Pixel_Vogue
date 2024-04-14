package repository


type ImageRepository struct{
	DB *Gorm.DB
}

func NewImageRepository(DB *Gorm.DB) ImageRepository{
	return &ImageRepository{
		DB : DB,
	}
}

func (ir *ImageRepository)UploadImage(){
	
}