package mediaService

const (
	ACCESS_LEVEL_PUBLIC  string = "public-read"
	ACCESS_LEVEL_PRIVATE        = "private"
)

type UploadedURL struct {
	Bucket   string
	Region   string
	Host     string
	Prefix   string
	FileName string
}
