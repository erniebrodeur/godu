package analyzer

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type FileInfo struct {
	Path      string
	Size      int64
	DiskSize  int64
	Extension string
}

type TypeStats struct {
	Extension string
	Count     int
	TotalSize int64
}

func Run(dir string, debug, human bool) error {
	var files []FileInfo

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if debug {
				fmt.Printf("SKIP: %s (error: %v)\n", path, err)
			}
			return nil
		}

		if debug {
			fmt.Printf("VISIT: %s\n", path)
		}

		if !info.IsDir() {
			const blockSize = 4096
			size := info.Size()
			diskSize := ((size + blockSize - 1) / blockSize) * blockSize

			var ext string
			filename := filepath.Base(path)

			if strings.HasPrefix(filename, ".") && !strings.Contains(filename[1:], ".") {
				// Pure dotfiles without extensions (like .gitkeep, .bashrc)
				ext = "dotfiles"
			} else if strings.HasPrefix(filename, ".") {
				// Dotfiles with extensions (like .gitignore.sample)
				ext = strings.ToLower(filepath.Ext(filename))
			} else {
				// Regular files
				ext = strings.ToLower(filepath.Ext(path))
				if ext == "" {
					ext = "no-ext"
				}
			}

			files = append(files, FileInfo{
				Path:      path,
				Size:      size,
				DiskSize:  diskSize,
				Extension: ext,
			})

			if debug {
				fmt.Printf("FILE: %s (%d bytes, %d disk) ext=%s\n", path, size, diskSize, ext)
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	typeMap := make(map[string]*TypeStats)

	for _, file := range files {
		ext := file.Extension

		if typeMap[ext] == nil {
			typeMap[ext] = &TypeStats{
				Extension: ext,
				Count:     0,
				TotalSize: 0,
			}
		}

		typeMap[ext].Count++
		typeMap[ext].TotalSize += file.DiskSize
	}

	var types []*TypeStats
	for _, stats := range typeMap {
		types = append(types, stats)
	}

	sort.Slice(types, func(i, j int) bool {
		return types[i].TotalSize > types[j].TotalSize
	})

	var totalSize int64
	for _, t := range types {
		totalSize += t.TotalSize
	}

	for _, t := range types {
		percentage := float64(t.TotalSize) / float64(totalSize) * 100

		if human {
			fmt.Printf("%s\t%d\t%.1f%%\t%s\n",
				formatSize(t.TotalSize), t.Count, percentage, t.Extension)
		} else {
			fmt.Printf("%dK\t%d\t%.1f%%\t%s\n",
				t.TotalSize/1024, t.Count, percentage, t.Extension)
		}
	}

	return nil
}

func formatSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%dB", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%c", float64(size)/float64(div), "KMGTPE"[exp])
}
