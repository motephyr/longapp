package modelhelper

import (
	"log"

	"github.com/motephyr/longcare/app"
	"github.com/motephyr/longcare/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func SourceAll() models.SourceSlice {
	result, err := models.Sources().All(app.Http.Database)
	if err != nil {
		log.Println(err)
	}
	return result
}

func SourceCreate(source *models.Source) *models.Source {
	err := source.Insert(app.Http.Database, boil.Infer())
	if err != nil {
		log.Println(err)
	}
	return source
}

func SourceFind(id any) *models.Source {
	result, err := models.Sources(Where("id = ?", id)).One(app.Http.Database)
	if err != nil {
		log.Println(err)
	}
	return result
}

func SourceUpdate(id int, source *models.Source) *models.Source {

	source.ID = id
	rowsAff, _ := source.Update(app.Http.Database, boil.Infer())

	if rowsAff == 0 {
		return nil
	} else {
		return source
	}
}

func SourceDelete(id int, source *models.Source) bool {

	source.ID = id

	// Delete the pilot from the database
	rowsAff, _ := source.Delete(app.Http.Database)

	if rowsAff == 0 {
		return false
	} else {
		return true
	}

}
