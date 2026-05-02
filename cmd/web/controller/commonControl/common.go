package commonControl

import (
	"bootpkg/common/global"
	"bootpkg/common/middleware"
	"bootpkg/common/response"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kirinlabs/utils/sys"
)

const maxImageUploadSize = 5 << 20 // 5MB

var allowedImageExt = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".webp": true,
}

var allowedImageMIME = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/gif":  true,
	"image/webp": true,
}

func validateImageUpload(fileHeader *multipart.FileHeader) error {
	if fileHeader == nil {
		return fmt.Errorf("file is required")
	}
	if fileHeader.Size <= 0 {
		return fmt.Errorf("file is empty")
	}
	if fileHeader.Size > maxImageUploadSize {
		return fmt.Errorf("file too large, max 5MB")
	}

	ext := strings.ToLower(path.Ext(fileHeader.Filename))
	if !allowedImageExt[ext] {
		return fmt.Errorf("unsupported file extension: %s", ext)
	}

	f, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("open file failed: %w", err)
	}
	defer f.Close()

	buf := make([]byte, 512)
	n, err := f.Read(buf)
	if err != nil {
		return fmt.Errorf("read file failed: %w", err)
	}
	mimeType := http.DetectContentType(buf[:n])
	if !allowedImageMIME[mimeType] {
		return fmt.Errorf("unsupported content type: %s", mimeType)
	}

	return nil
}

// GetCSRFToken 返回 CSRF token 给前端
func GetCSRFToken(c *gin.Context) {
	token := c.Writer.Header().Get(middleware.CSRFTokenHeaderKey)
	if token == "" {
		if err := middleware.SetCSRFTokenCookie(c); err != nil {
			c.JSON(500, gin.H{"message": "failed to generate csrf token"})
			return
		}
		token = c.Writer.Header().Get(middleware.CSRFTokenHeaderKey)
	}
	c.JSON(200, gin.H{
		"csrf_token": token,
	})
}

func UploadFileSingle(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = validateImageUpload(file); err != nil {
		response.FailErrHttpCodeJSON(c, 400, err.Error(), "")
		return
	}
	dst := "./upload"
	if !sys.IsExists(dst) {
		os.Mkdir(dst, 0755)
	}
	//更改文件名称
	fileName := uuid.NewString() + path.Ext(file.Filename)
	filePath := filepath.Join(dst, fileName)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		response.FailErrHttpCodeJSON(c, 500, err.Error(), "")
		return
	}
	newPath := strings.Replace(dst, "./", "/api/", 1) + "/" + fileName
	if len(global.CONFIG.General.ImageDomain) > 0 {
		newPath = global.CONFIG.General.ImageDomain + newPath
	}
	c.JSON(200, struct {
		Name     string `json:"name"`
		Status   string `json:"status"`
		ThumbUrl string `json:"thumbUrl"`
		Url      string `json:"url"`
	}{
		Name:     file.Filename,
		Status:   "done",
		ThumbUrl: newPath,
		Url:      newPath,
	})
}

// 多文件上传
func UploadFileSingleMax(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	for _, file := range files {
		log.Println(file.Filename)
	}
}
