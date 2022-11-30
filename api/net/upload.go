package net

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"path/filepath"

	"github.com/alfg/openencoder/api/data"
	"github.com/alfg/openencoder/api/config"
	"github.com/alfg/openencoder/api/types"
	"github.com/alfg/openencoder/api/helpers"
)

func MoveFile(sourcePath, destPath string) error {
    inputFile, err := os.Open(sourcePath)
    if err != nil {
        return fmt.Errorf("Couldn't open source file: %s", err)
    }
    outputFile, err := os.Create(destPath)
    if err != nil {
        inputFile.Close()
        return fmt.Errorf("Couldn't open dest file: %s", err)
    }
    defer outputFile.Close()
    _, err = io.Copy(outputFile, inputFile)
    inputFile.Close()
    if err != nil {
        return fmt.Errorf("Writing to output file failed: %s", err)
    }
    // The copy was successful, so now delete the original file
    err = os.Remove(sourcePath)
    if err != nil {
        return fmt.Errorf("Failed removing original file: %s", err)
    }
    return nil
}

// Upload uploads a job based on the driver setting.
func Upload(job types.Job) error {
	db := data.New()
	driver, err := db.Settings.GetSetting(types.StorageDriver)
	if err != nil {
		return errors.New("no driver set")
	}

	if driver.Value == "s3" {
		if err := s3Upload(job); err != nil {
			return err
		}
		return nil
	} else if driver.Value == "ftp" {
		if err := ftpUpload(job); err != nil {
			return err
		}
		return nil
	} else if driver.Value == "local" {
		if err := localUpload(job); err != nil {
			return err
		}
		return nil
	}
	return errors.New("no driver set")
}

// GetUploader gets the upload function.
func s3Upload(job types.Job) error {
	// Get credentials from settings.
	db := data.New()
	settings := db.Settings.GetSettings()

	config := S3Config{
		AccessKey:      types.GetSetting(types.S3AccessKey, settings),
		SecretKey:      types.GetSetting(types.S3SecretKey, settings),
		Provider:       types.GetSetting(types.S3Provider, settings),
		Region:         types.GetSetting(types.S3OutboundBucketRegion, settings),
		InboundBucket:  types.GetSetting(types.S3InboundBucket, settings),
		OutboundBucket: types.GetSetting(types.S3OutboundBucket, settings),
	}

	// Get job data.
	j, err := db.Jobs.GetJobByGUID(job.GUID)
	if err != nil {
		log.Error(err)
		return err
	}
	encodeID := j.EncodeID

	s3 := NewS3(config)
	go trackTransferProgress(encodeID, s3)
	err = s3.Upload(job)
	close(progressCh)

	return err
}

// GetFTPUploader sets the FTP upload function.
func ftpUpload(job types.Job) error {
	db := data.New()
	settings := db.Settings.GetSettings()

	addr := types.GetSetting(types.FTPAddr, settings)
	user := types.GetSetting(types.FTPUsername, settings)
	pass := types.GetSetting(types.FTPPassword, settings)

	f := NewFTP(addr, user, pass)
	err := f.Upload(job)
	return err
}

// GetLocalUploader sets the Local upload function.
func localUpload(job types.Job) error {
	db := data.New()
	settings := db.Settings.GetSettings()

	configPath := types.GetSetting(types.LocalPath, settings)
	tmpPath := helpers.GetTmpPath(config.Get().WorkDirectory, job.GUID)
	
	filelist := []string{}
	filepath.Walk(tmpPath+"dst", func(path string, f os.FileInfo, err error) error {
		if isDirectory(path) {
			return nil
		}
		filelist = append(filelist, path)
		return nil
	})
	
	for _, file := range filelist {
		os.MkdirAll(configPath+"/"+job.GUID, 0644)
		MoveFile(file, configPath+"/"+job.GUID+"/")
	}
	return nil
}
