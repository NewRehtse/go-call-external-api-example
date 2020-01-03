package csv

import (
	"encoding/csv"
	"fmt"
	"os"

	githubcli "newrehtse.info/go-training/github/internal"
)

type saveCsvRepository struct {
}

func NewSaveCsvRepository() githubcli.SaveRepo {
	return &saveCsvRepository{}
}

func (s saveCsvRepository) StoreRepositoriesResponse(list githubcli.RepositoriesResponse, dest string) error {
	if dest == "" {
		return nil
	}

	file, err := os.Create(dest)
	if nil != err {
		return err
	}

	defer func() {
		e := file.Close()

		if nil != e {
			err = e
		}
	}()

	w := csv.NewWriter(file)

	headers := list.GetHeaders()
	if err := w.Write(headers); err != nil {
		fmt.Println("Error", err)
	}

	items := list.Repositories
	for _, item := range items {
		w.Write(item.ToSlice())
	}

	w.Flush()
	return nil
}
