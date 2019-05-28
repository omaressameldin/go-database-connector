package firebase

import (
	"io"
	"os"
	"path/filepath"

	"golang.org/x/net/context"
)

func (f *Firebase) UploadFile(pathToFile string) error {
	file, err := os.Open(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	b, err := f.storage.DefaultBucket()
	if err != nil {
		return err
	}

	w := b.Object(filepath.Base(file.Name())).NewWriter(context.Background())
	w.CacheControl = "public, max-age=86400"
	w.ContentType = "image/jpg"

	if _, err = io.Copy(w, file); err != nil {
		return err
	}

	return w.Close()
}

func (f *Firebase) DeleteFile(url string) error {
	b, err := f.storage.DefaultBucket()
	if err != nil {
		return err
	}

	return b.Object(url).Delete(context.Background())
}
