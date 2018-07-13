package minio

import (
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	minio "github.com/minio/minio-go"
)

// SetupHandler -
func SetupHandler(e *echo.Group) {
	e.GET("", getBucketList)
	e.POST("", uploadFileEP)
	e.GET("/list", getAllObject)
}

func minioClient() (*minio.Client, error) {
	accessKey := os.Getenv("MINIO_ACCESSID")
	endPoint := os.Getenv("MINIO_ENDPOINT")
	secretKey := os.Getenv("MINIO_SECRETKEY")
	useSSL := os.Getenv("MINIO_USESSL")
	SSL, _ := strconv.ParseBool(useSSL)

	// Initiamainlize minio client object
	minioClient, err := minio.New(endPoint, accessKey, secretKey, SSL)
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}

func getBucketList(c echo.Context) error {
	mc, err := minioClient()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	result, err := getBucketListController(mc)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func uploadFileEP(c echo.Context) error {
	mc, err := minioClient()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusExpectationFailed, err.Error())
	}

	upload, err := UploadFile(mc, file, "test")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, upload)
}

func getAllObject(c echo.Context) error {
	mc, err := minioClient()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	bucket := c.QueryParam("bucket")

	result, err := GetAllFiles(mc, bucket)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
