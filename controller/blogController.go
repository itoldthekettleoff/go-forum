package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/itoldthekettleoff/go-forum/database"
	"github.com/itoldthekettleoff/go-forum/model"
)

func BlogList(c *fiber.Ctx) error {
	var records []model.Blog
	context := fiber.Map{
		"status":  200,
		"message": "Blog List",
	}

	db := database.DBConn
	db.Find(&records)
	context["blogs"] = records

	c.SendStatus(200)
	return c.JSON(context)
}

func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"status":  200,
		"message": "Blog Create",
	}

	record := new(model.Blog)
	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")
		context["status"] = ""
		context["message"] = "Something went wrong :("
	}

	result := database.DBConn.Create(record)
	if result.Error != nil {
		log.Println("Error in writing record.")
	}

	context["message"] = "Record is saved successfully."
	context["data"] = record
	c.SendStatus(200)
	return c.JSON(context)
}

func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"status":  200,
		"message": "Blog Update",
	}

	id := c.Params("id")
	var record model.Blog
	database.DBConn.First(&record, id)
	if record.ID == 0 {
		log.Println("Record not found.")
		c.Status(400)
		context["status"] = 400
		return c.JSON(context)
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")
	}

	result := database.DBConn.Save(record)
	if result.Error != nil {
		log.Println("Error in updating record.")
	}

	context["message"] = "Record has been updated successfully."
	context["data"] = record
	c.SendStatus(200)
	return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"status":  200,
		"message": "Blog Delete",
	}

	id := c.Params("id")
	var record model.Blog
	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not found.")
		c.Status(400)
		context["status"] = 400
		context["message"] = "Record is not present."
		return c.JSON(context)
	}

	result := database.DBConn.Delete(record)
	if result.Error != nil {
		log.Println("Error in deleting record.")
		return c.JSON(context)
	}

	context["message"] = "Record has been deleted."
	c.SendStatus(200)
	return c.JSON(context)
}
