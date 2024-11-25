package utils

import (
	"os"
	"regexp"
)

type FileInfo struct {
	IsDir bool
	Name  string
}

type FileList []FileInfo

// keep or delete files what match a regex
func FileListGlob(fl *FileList, regex string, keep_match bool) FileList {
	var out FileList
	for _, f := range *fl {
		found, _ := regexp.MatchString(regex, f.Name)
		if found == keep_match {
			out = append(out, f)
		}

	}
	return out
}

func FilesInDir(dir string, show_dirs bool, recursive bool) FileList {
	var out FileList

	if dir[len(dir)-1] != '/' {
		dir = dir + "/"
	}

	items, _ := os.ReadDir(dir)

	for _, v := range items {

		// if not a directory or if we want directories
		if !v.IsDir() || (v.IsDir() && show_dirs) {
			out = append(out,
				FileInfo{
					IsDir: v.IsDir(),
					Name:  v.Name(),
				})
		}

		// if we are a directory and plan to recurse we want to add 2 lists
		if v.IsDir() && recursive {
			new_files := FilesInDir(dir+v.Name(), show_dirs, recursive)
			for _, new_v := range new_files {
				new_path := v.Name() + "/" + new_v.Name

				out = append(out,
					FileInfo{
						IsDir: new_v.IsDir,
						Name:  new_path,
					})

			}
		}
	}

	return out
}
