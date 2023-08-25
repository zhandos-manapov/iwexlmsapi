package course

import (
	"context"
	"iwexlmsapi/database"
	"iwexlmsapi/models"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func FindOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := `SELECT course.name, course.course_id, course.level, course.description, course.agenda, level.level_name
  FROM course
  INNER JOIN level ON course.level = level.id
  WHERE course_id = $1`
	var course models.Course

	if err := database.Pool.QueryRow(context.Background(), query, id).Scan(&course.CourseId, &course.Name, &course.Level, &course.Description, &course.Agenda, &course.LevelName); err != nil {
		if err == pgx.ErrNoRows {
			return fiber.NewError(fiber.StatusBadRequest, "Курс не существует")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error: " + err.Error())
	}

	return c.JSON(course)
}

func FindMany(c *fiber.Ctx) error {
	query := `SELECT course_id, course.name, course.level, course.description, course.agenda, level.level_name 
	FROM course
	INNER JOIN level ON course.level = level.id`
	rows, err := database.Pool.Query(context.Background(), query)
	if err != nil {
		return err
	}

	courses := make([]models.Course, 0)
	for rows.Next() {
		var course models.Course
		if err := rows.Scan(&course.CourseId, &course.Name, &course.Level, &course.Description, &course.Agenda, &course.LevelName); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error: " + err.Error())
		}
		courses = append(courses, course)
	}

	if len(courses) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Что-то пошло не так")
	}

	return c.JSON(courses)
}

func CreateOne(c *fiber.Ctx) error {
	body := c.Locals("body")
	course := body.(*models.CourseSend)
	
	query := `SELECT course_id FROM course WHERE name = $1`
	
	var CourseID int64
	if err := database.Pool.QueryRow(context.Background(), query, course.Name).Scan(&CourseID); err != nil {
			if err != pgx.ErrNoRows {
					return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error: " + err.Error())
			}
	} else {
			return fiber.NewError(fiber.StatusBadRequest, "Курс с таким именем уже существует")
	}
	
	query = `
			INSERT INTO course (name, level, description, agenda)
			VALUES ($1, $2, $3, $4)
	`
	
	_, err := database.Pool.Exec(context.Background(), query, course.Agenda, course.Level, course.Description,  course.Name)
	if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error: " + err.Error())
	}
	
	return c.JSON(models.ServerError{Message: "Курс успешно создан"})
}

func UpdateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	body := c.Locals("body")
	course := body.(*models.CourseSend)
	
	existingCourseQuery := `SELECT course_id FROM course WHERE course_id = $1`
	existingCourseRow := database.Pool.QueryRow(context.Background(), existingCourseQuery, id)
	
	var existingCourseID int64
	if err := existingCourseRow.Scan(&existingCourseID); err != nil {
			if err == pgx.ErrNoRows {
					return fiber.NewError(fiber.StatusBadRequest, "Курс не найден")
			}
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error: " + err.Error())
	}
	
	updateQuery := `
			UPDATE course
			SET name = $1, level = $2, description = $3, agenda = $4
			WHERE course_id = $5
	`
	
	_, err := database.Pool.Exec(context.Background(), updateQuery, course.Name, course.Level, course.Description, course.Agenda, id)
	if err != nil {
			return err
	}
	
	return c.JSON(models.ServerError{Message: "Курс успешно обновлен"})
}


func DeleteOne(c *fiber.Ctx) error {
    id := c.Params("id")
    
    query := `SELECT course_id FROM course WHERE course_id = $1`
    
    var CourseID int64
    if err := database.Pool.QueryRow(context.Background(), query, id).Scan(&CourseID); err != nil {
        if err == pgx.ErrNoRows {
            return fiber.NewError(fiber.StatusBadRequest, "Курс не найден")
        }
        return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error: " + err.Error())
    }
    
    query = `DELETE FROM course WHERE course_id = $1`
    
    _, err := database.Pool.Exec(context.Background(), query, id)
    if err != nil {
        return err
    }
    
    return c.JSON(models.ServerError{Message: "Курс успешно удален"})
}
