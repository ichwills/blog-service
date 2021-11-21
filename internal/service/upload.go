package service

import (
	"Pro/blog-service/global"
	"Pro/blog-service/pkg/upload"
	"errors"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported.")
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit.")
	}
	UploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(UploadSavePath) {
		if err := upload.CreateSavePath(UploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}
	if upload.CheckPermission(UploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}
	dst := UploadSavePath + "/" + fileName
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}
	accesssUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accesssUrl}, nil
}
