package modelhelper

import (
	"database/sql"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/gofiber/fiber/v2"
	"github.com/motephyr/longapp/app"
	"github.com/motephyr/longapp/models"
	"github.com/motephyr/longapp/utils"
	"github.com/sujit-baniya/xid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Client struct { // Our example struct, you can use "-" to ignore a field
	Timestamp    null.String `csv:"timestamp"`
	MacAddress   null.String `csv:"mac_address"`
	UserID       int         `csv:"user_id"`
	Action       null.String `csv:"action"`
	Room         null.String `csv:"room"`
	VideoUrl     null.String `csv:"video_url"`
	CameraStatus null.String `csv:"camera_status"`
	NotUsed      null.String `csv:"-"`
}

func HandleUploadIndividualFile(c *fiber.Ctx, file *multipart.FileHeader) error {
	fileParts := strings.Split(file.Filename, ".")
	datestring := fileParts[0]

	id := xid.New().String()
	id = id + "." + fileParts[1]
	fileName := filepath.Join(app.Http.Server.UploadPath, id)
	// Check for errors
	if err := c.SaveFile(file, fileName); err != nil {
		return err
		// c.Status(500).Send("Something went wrong 😭")
	}
	clientsFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	clients := []*Client{}

	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
		panic(err)
	}

	err = utils.Tx(app.Http.Database.DB, func(tx *sql.Tx) error {

		if _, err := models.Sources(Where("datestring = ?", datestring)).DeleteAllG(); err != nil {
			return err // The helper will automatically rollback for you here
		}

		if _, err := models.Groups(Where("datestring = ?", datestring)).DeleteAllG(); err != nil {
			return err // The helper will automatically rollback for you here
		}

		for _, client := range clients {
			source := models.Source{}
			source.Idstring = client.MacAddress
			source.Timestring = client.Timestamp
			source.Datestring = null.StringFrom(datestring)
			if client.UserID != -1 {
				source.Action = client.Action

				source.Space = client.Room
				source.URL = client.VideoUrl

			} else {
				source.Action = client.CameraStatus
			}
			if err := source.InsertG(boil.Infer()); err != nil {
				return err // The helper will automatically rollback for you here
			}
		}

		return nil // The helper will automatically call commit since you returned nil
	})

	if err != nil { // Load clients from file
		c.JSON("source upload fail.")
	}

	// if _, err := clientsFile.Seek(0, 0); err != nil { // Go to the start of the file
	// 	panic(err)
	// }

	// csvContent, err := gocsv.MarshalString(&clients) // Get all clients as CSV string
	// //err = gocsv.MarshalFile(&clients, clientsFile) // Use this to save the CSV back to the file
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(csvContent) // Display all clients as CSV string

	return c.JSON("source upload successfully.")
}
