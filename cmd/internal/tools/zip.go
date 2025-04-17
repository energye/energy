//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package tools

import (
	"archive/tar"
	"archive/zip"
	"compress/bzip2"
	"compress/gzip"
	"errors"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cmd/internal/term"
	"github.com/pterm/pterm"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func ExtractUnZip(filePath, targetPath string, rmRootDir bool, files ...string) error {
	if rc, err := zip.OpenReader(filePath); err == nil {
		defer rc.Close()
		multi := pterm.DefaultMultiPrinter
		defer multi.Stop()
		fileTotalProcessBar := pterm.DefaultProgressbar.WithWriter(multi.NewWriter())
		writeFileProcessBar := pterm.DefaultProgressbar.WithWriter(multi.NewWriter())
		var createWriteFile = func(info fs.FileInfo, path string, file io.Reader) error {
			targetFileName := filepath.Join(targetPath, path)
			if info.IsDir() {
				os.MkdirAll(targetFileName, info.Mode())
				return nil
			}
			fDir, name := filepath.Split(targetFileName)
			if !IsExist(fDir) {
				os.MkdirAll(fDir, 0755)
			}
			if targetFile, err := os.Create(targetFileName); err == nil {
				defer targetFile.Close()
				var (
					total = 100
					c     int
					cn    int
				)
				wfpb, err := writeFileProcessBar.WithCurrent(0).WithTotal(total).Start("Write File " + name)
				if err != nil {
					return err
				}
				WriteFile(file, targetFile, info.Size(), func(totalLength, processLength int64) {
					process := int((float64(processLength) / float64(totalLength)) * 100)
					if process > c {
						c = process
						wfpb.Add(process)
						cn++
					}
				})
				if cn < total {
					wfpb.Add(total - cn)
				}
				wfpb.Stop()
				return nil
			} else {
				return err
			}
		}
		// 所有文件
		if len(files) == 0 {
			zipFiles := rc.File
			extractFilesProcessBar, err := fileTotalProcessBar.WithTotal(len(zipFiles)).Start("Extract File")
			if err != nil {
				return err
			}
			multi.Start()
			for _, f := range zipFiles {
				extractFilesProcessBar.Increment() // +1
				r, _ := f.Open()
				var name string
				if rmRootDir {
					// 移除压缩包内的根文件夹名
					name = filepath.Clean(f.Name[strings.Index(f.Name, "/")+1:])
				} else {
					name = filepath.Clean(f.Name)
				}
				if err := createWriteFile(f.FileInfo(), name, r.(io.Reader)); err != nil {
					return err
				}
				_ = r.Close()
			}
			return nil
		} else {
			extractFilesProcessBar, err := fileTotalProcessBar.WithTotal(len(files)).Start("Extract File")
			if err != nil {
				return err
			}
			multi.Start()
			// 指定名字的文件
			for i := 0; i < len(files); i++ {
				extractFilesProcessBar.Increment() // +1
				if f, err := rc.Open(files[i]); err == nil {
					info, _ := f.Stat()
					if err := createWriteFile(info, files[i], f); err != nil {
						return err
					}
					_ = f.Close()
				} else {
					return err
				}
			}
			return nil
		}
	} else {
		return err
	}
}

func WriteFile(r io.Reader, w *os.File, totalLength int64, callback func(totalLength, processLength int64)) {
	buf := make([]byte, 1024*10)
	var written int64
	for {
		nr, err := r.Read(buf)
		if nr > 0 {
			nw, err := w.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			callback(totalLength, written)
			if err != nil {
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if err != nil {
			break
		}
	}
}

func tarFileReader(filePath string) (*tar.Reader, func(), error) {
	reader, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error: cannot open file, error=[%v]\n", err)
		return nil, nil, err
	}
	if filepath.Ext(filePath) == ".gz" {
		gr, err := gzip.NewReader(reader)
		if err != nil {
			fmt.Printf("error: cannot open gzip file, error=[%v]\n", err)
			return nil, nil, err
		}
		return tar.NewReader(gr), func() {
			gr.Close()
			reader.Close()
		}, nil
	} else {
		return tar.NewReader(reader), func() {
			reader.Close()
		}, nil
	}
}

func tarFileCount(filePath string) (int, error) {
	tarReader, clos, err := tarFileReader(filePath)
	if err != nil {
		return 0, err
	}
	defer clos()
	var count int
	for {
		_, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return 0, err
		}
		count++
	}
	return count, nil
}

func filePathInclude(compressPath string, files ...string) (string, bool) {
	if len(files) == 0 {
		return compressPath, true
	} else {
		for _, file := range files {
			f := file
			tIdx := strings.LastIndex(f, "/*")
			if tIdx != -1 {
				f = f[:tIdx]
				if f[0] == '/' {
					if strings.Index(compressPath, f[1:]) == 0 {
						return compressPath, true
					}
				}
				if strings.Index(compressPath, f) == 0 {
					return strings.Replace(compressPath, f, "", 1), true
				}
			} else {
				if f[0] == '/' {
					if compressPath == f[1:] {
						return f, true
					}
				}
				if compressPath == f {
					f = f[strings.Index(f, "/")+1:]
					return f, true
				}
			}
		}
	}
	return "", false
}

func dir(path string) string {
	path = strings.ReplaceAll(path, "\\", string(filepath.Separator))
	lastSep := strings.LastIndex(path, string(filepath.Separator))
	return path[:lastSep]
}

func ExtractUnTar(filePath, targetPath string, files ...string) error {
	term.Logger.Info("Read Files Number")
	fileCount, err := tarFileCount(filePath)
	println(fileCount)
	if err != nil {
		return err
	}
	tarReader, clos, err := tarFileReader(filePath)
	if err != nil {
		return err
	}
	defer clos()
	multi := pterm.DefaultMultiPrinter
	defer multi.Stop()
	fileTotalProcessBar := pterm.DefaultProgressbar.WithWriter(multi.NewWriter())
	writeFileProcessBar := pterm.DefaultProgressbar.WithWriter(multi.NewWriter())
	extractFilesProcessBar, err := fileTotalProcessBar.WithTotal(fileCount).Start("Extract File")
	if err != nil {
		return err
	}
	multi.Start()
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		extractFilesProcessBar.Increment()
		// 去除压缩包内的一级目录
		compressPath := filepath.Clean(header.Name[strings.Index(header.Name, "/")+1:])
		includePath, isInclude := filePathInclude(compressPath, files...)
		if !isInclude {
			continue
		}
		info := header.FileInfo()
		targetFile := filepath.Join(targetPath, includePath)
		if info.IsDir() {
			if err = os.MkdirAll(targetFile, info.Mode()); err != nil {
				return err
			}
		} else {
			fDir := dir(targetFile)
			_, err = os.Stat(fDir)
			if os.IsNotExist(err) {
				if err = os.MkdirAll(fDir, info.Mode()); err != nil {
					return err
				}
			}
			file, err := os.OpenFile(targetFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
			if err != nil {
				return err
			}
			var (
				total = 100
				c     int
				cn    int
			)
			_, fileName := filepath.Split(targetFile)
			wfpb, err := writeFileProcessBar.WithCurrent(0).WithTotal(total).Start("Write File " + fileName)
			if err != nil {
				return err
			}
			WriteFile(tarReader, file, header.Size, func(totalLength, processLength int64) {
				process := int((float64(processLength) / float64(totalLength)) * 100)
				if process > c {
					c = process
					wfpb.Add(process)
					cn++
				}
			})
			if cn < total {
				wfpb.Add(total - cn)
			}
			wfpb.Stop()
			file.Sync()
			file.Close()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// UnBz2ToTar 释放bz2文件到tar
func UnBz2ToTar(bz2FilePath string, callback func(totalLength, processLength int64)) (string, error) {
	fileBz2, err := os.Open(bz2FilePath)
	if err != nil {
		return "", err
	}
	defer fileBz2.Close()
	dirName := fileBz2.Name()
	dirName = dirName[:strings.LastIndex(dirName, ".")]
	if !IsExist(dirName) {
		r := bzip2.NewReader(fileBz2)
		w, err := os.Create(dirName)
		if err != nil {
			return "", err
		}
		defer w.Close()
		WriteFile(r, w, 0, callback)
	} else {
		term.Section.Println("File already exists")
	}
	return dirName, nil
}

// 提取文件
func ExtractFiles(sourcePath, targetPath string, extractOSConfig []string) error {
	ext := filepath.Ext(sourcePath)
	switch ext {
	case ".tar":
		return ExtractUnTar(sourcePath, targetPath, extractOSConfig...)
	case ".zip":
		return ExtractUnZip(sourcePath, targetPath, false, extractOSConfig...)
	case ".7z":
		// 7z 直接解压目录内所有文件
	}
	return errors.New("not module")
}
