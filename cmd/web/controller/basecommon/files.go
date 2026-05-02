package basecommon

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/vo"
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
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

// 断点续传-
func UpHashFile(c *gin.Context) {
	var jsonp vo.FileHashVO
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	upPath := global.CONFIG.General.UploadFilePath
	path := upPath + "/tmp/" + jsonp.FileHash + "/"
	err = tool.CreatePath(path, 0666)
	if err != nil {
		response.FailErrDataJSON(c, response.ERROR_DEFAULT, "文件夹不存在", "")
		return
	}
	sFile := make([]string, 0)
	sFile, err = tool.GetAllFileName(path, sFile)
	if err != nil {
		response.FailErrDataJSON(c, response.ERROR_DEFAULT, "读取文件列表错误", "")
		return
	}
	c.JSON(200, struct {
		ShouldUploadFile bool     `json:"shouldUploadFile"`
		UploadedChunks   []string `json:"uploadedChunks"`
	}{
		ShouldUploadFile: true,
		UploadedChunks:   sFile,
	})
}

// 断点续传-保存快
func UpHashFileSave(c *gin.Context) {
	var jsonp vo.FileChunkHashVO
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrHttpCodeJSON(c, 500, err.Error(), "")
		return
	}
	upPath := global.CONFIG.General.UploadFilePath
	path := upPath + "/tmp/" + jsonp.FileHash + "/" + jsonp.Hash
	if tool.IsFile(path) {
		_ = os.Remove(path)
	}
	file, err := c.FormFile("chunk")
	if err != nil || file.Size < 1 {
		response.FailErrHttpCodeJSON(c, 500, "数据未存在", "")
		return
	}
	err = c.SaveUploadedFile(file, path) //0666
	if err != nil {
		if tool.IsFile(path) {
			_ = os.Remove(path)
		}
		response.FailErrHttpCodeJSON(c, 500, "写入文件失败", "")
		return
	}
	// Destination
	response.SuccessJSON(c, "")
}

// 断点续传-合并文件
func MergeHashFile(c *gin.Context) {
	var jsonp vo.FileChunkMergeVO
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrHttpCodeJSON(c, 500, "数据验证不通过", "")
		return
	}
	upPath := global.CONFIG.General.UploadFilePath
	path := upPath + "/tmp/" + jsonp.FileHash
	outFileName := upPath + "/" + tool.String(time.Now().Unix()) + "." + jsonp.FileName
	err = tool.FileMerge(path, jsonp.FileHash, outFileName, 0666) //读写权限
	if err != nil {
		response.FailErrHttpCodeJSON(c, 500, err.Error(), "")
		return
	}
	if tool.IsExists(path) && !global.CONFIG.General.UploadSecondPass {
		_ = os.RemoveAll(path)
	}
	response.SuccessJSON(c, outFileName)
}

// 小文件上传
func UpFileSingle(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailErrHttpCodeJSON(c, 500, err.Error(), "")
		return
	}
	if err = validateImageUpload(file); err != nil {
		response.FailErrHttpCodeJSON(c, 400, err.Error(), "")
		return
	}
	dst := global.CONFIG.General.UploadFilePath
	if !sys.IsExists(dst) {
		err = os.Mkdir(dst, 0755)
	}
	if err != nil {
		response.FailErrHttpCodeJSON(c, 500, err.Error(), "")
		return
	}

	fileName := uuid.NewString() + path.Ext(file.Filename)
	filePath := filepath.Join(dst, fileName)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		response.FailErrHttpCodeJSON(c, 500, err.Error(), "")
		return
	}

	filePath = "/api/upload/" + fileName
	if len(global.CONFIG.General.ImageDomain) > 0 {
		filePath = global.CONFIG.General.ImageDomain + filePath
	}
	c.JSON(200, struct {
		Name     string `json:"name"`
		Status   string `json:"status"`
		ThumbUrl string `json:"thumbUrl"`
		Url      string `json:"url"`
	}{
		Name:     file.Filename,
		Status:   "done",
		ThumbUrl: filePath,
		Url:      filePath,
	})
}
