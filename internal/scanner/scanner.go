package scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func Scan(dir string, debug, human bool, maxDepth int) error {
	dirSizes := make(map[string]int64)

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
			// Calculate actual disk usage (block-aligned size)
			const blockSize = 4096
			size := info.Size()
			// Round up to nearest block
			diskSize := ((size + blockSize - 1) / blockSize) * blockSize

			if debug {
				fmt.Printf("FILE: %s (%d bytes, %d disk)\n", path, size, diskSize)
			}
			// Add disk size to all parent directories
			currentDir := filepath.Dir(path)
			for {
				dirSizes[currentDir] += diskSize
				parent := filepath.Dir(currentDir)
				if parent == currentDir {
					break
				}
				currentDir = parent
			}
		} else {
			// Directories don't add their own overhead to the count
			// The space they use is already accounted for in the file system
			if debug {
				fmt.Printf("DIR: %s\n", path)
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	// Create list of directories to show based on depth
	var dirsToShow []string
	for dirPath := range dirSizes {
		relPath, _ := filepath.Rel(dir, dirPath)
		depth := 0
		if relPath != "." {
			depth = strings.Count(relPath, string(os.PathSeparator)) + 1
		}

		// Include this directory if within depth limit
		if maxDepth < 0 || depth <= maxDepth {
			dirsToShow = append(dirsToShow, dirPath)
		}
	}

	// Sort directories, but put the root directory last like du does
	sort.Slice(dirsToShow, func(i, j int) bool {
		// Root directory (.) should come last
		if dirsToShow[i] == dir {
			return false
		}
		if dirsToShow[j] == dir {
			return true
		}
		return dirsToShow[i] < dirsToShow[j]
	})

	// Print directory sizes
	for _, dirPath := range dirsToShow {
		size := dirSizes[dirPath]
		// Format path like du does (with ./ prefix for relative paths)
		displayPath := dirPath
		if relPath, err := filepath.Rel(dir, dirPath); err == nil {
			if relPath == "." {
				displayPath = "."
			} else {
				displayPath = "./" + relPath
			}
		}

		if human {
			fmt.Printf("%s\t%s\n", formatSize(size), displayPath)
		} else {
			fmt.Printf("%d\t%s\n", size/1024, displayPath)
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
