package class

import (
	"context"
	"fmt"
	"iwexlmsapi/database"
	"iwexlmsapi/models"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func getEnrolledStudents(c *fiber.Ctx) error {
	id := c.Params("id")
	query := `
	SELECT enrolled_student.*,
  	role.role_name
	FROM (
    SELECT e.cycle_id,
      e.student_id,
      e.enrollment_date,
      e.cancelled,
      COALESCE(e.cancellation_reason, '') AS cancellation_reason,
      u.first_name,
      u.last_name,
      u.email,
      u.contact_number,
      u.date_of_birth,
      u.role,
      u.is_active
    FROM enrollment AS e
      INNER JOIN users AS u ON e.student_id = u.id
    WHERE e.cycle_id = $1
  ) AS enrolled_student
  INNER JOIN role ON enrolled_student.role = role.id`
	rows, err := database.Pool.Query(context.Background(), query, id)
	defer rows.Close()
	if err != nil {
		return err
	}
	enrollments, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.EnrolledStudent])
	if err != nil {
		return err
	}
	return c.JSON(enrollments)
}

func enrollStudents(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return err
	}
	students := c.Locals("body").(*models.EnrollStudentsDTO)

	query := strings.Builder{}
	query.WriteString("INSERT INTO enrollment (cycle_id, student_id) VALUES ")

	for i := range students.Students {
		query.WriteString(fmt.Sprintf("($1, $%d),", i+2))
	}
	queryString := query.String()
	queryString = queryString[:len(queryString)-1]
	queryString += ";"

	queryParams := []any{idInt}
	for _, student_id := range students.Students {
		queryParams = append(queryParams, student_id)
	}

	if tag, err := database.Pool.Exec(context.Background(), queryString, queryParams...); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.ErrInternalServerError
	}
	return c.JSON(models.RespMsg{Message: "Студенты успешно зачислены"})
}

func findMany(ctx *fiber.Ctx) error {
	query := `
	SELECT course_cycle.id,
		course_cycle.course_id,
		course_cycle.branch_id,
		course_cycle.description,
		course_cycle.start_date,
		course_cycle.end_date,
		course_cycle.open_for_enrollment,
		course_cycle.course_code,
		branch_office.name as branch_name,
		course.name as course_name
	FROM course_cycle
		INNER JOIN branch_office ON course_cycle.branch_id = branch_office.id
		INNER JOIN course ON course_cycle.course_id = course.course_id`
	rows, err := database.Pool.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		return err
	}
	classes, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.ClassDB])
	if err != nil {
		return err
	}
	if len(classes) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Классы не найдены")
	}
	return ctx.JSON(classes)
}

func findOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := `
	SELECT course_cycle.id,
		course_cycle.course_id,
		course_cycle.branch_id,
		course_cycle.description,
		course_cycle.start_date,
		course_cycle.end_date,
		course_cycle.open_for_enrollment,
		course_cycle.course_code,
		branch_office.name as branch_name,
		course.name as course_name
	FROM course_cycle
		INNER JOIN branch_office ON course_cycle.branch_id = branch_office.id
		INNER JOIN course ON course_cycle.course_id = course.course_id
	WHERE course_cycle.id = $1`
	class := models.ClassDB{}
	if err := database.Pool.QueryRow(context.Background(), query, id).Scan(
		&class.ID,
		&class.CourseID,
		&class.BranchID,
		&class.Description,
		&class.StartDate,
		&class.EndDate,
		&class.OpenForEnrollment,
		&class.CourseCode,
		&class.BranchName,
		&class.CourseName,
	); err != nil {
		return err
	}
	return c.JSON(class)
}

func createOne(c *fiber.Ctx) error {
	class := c.Locals("body").(*models.CreateClassDTO)
	query := `
	INSERT INTO course_cycle (
    description,
    start_date,
    end_date,
    open_for_enrollment,
    course_code,
    branch_id,
    course_id
  )
	VALUES($1, $2, $3, $4, $5, $6, $7)
	RETURNING id`
	if err := database.Pool.QueryRow(
		context.Background(),
		query,
		class.Description,
		class.StartDate,
		class.EndDate,
		class.BranchID,
		class.CourseID,
	).Scan(&class.ID); err != nil {
		return err
	}
	return c.JSON(class)
}

func updateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	// TODO open_for_enrollment
	class := c.Locals("body").(*models.UpdateClassDTO)
	if class.Description == "" &&
		class.StartDate == "" &&
		class.EndDate == "" &&
		class.BranchID == 0 &&
		class.CourseID == 0 &&
		class.CourseCode == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Не указаны данные для обновления")
	}
	query := strings.Builder{}
	query.WriteString("UPDATE course_cycle SET")
	queryParams := []any{id}

	if class.Description != "" {
		query.WriteString(fmt.Sprintf(" description=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, class.Description)
	}

	if class.StartDate != "" {
		query.WriteString(fmt.Sprintf(" start_date=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, class.StartDate)
	}

	if class.EndDate != "" {
		query.WriteString(fmt.Sprintf(" end_date=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, class.EndDate)
	}

	if class.BranchID != 0 {
		query.WriteString(fmt.Sprintf(" branch_id=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, class.BranchID)
	}

	if class.CourseID != 0 {
		query.WriteString(fmt.Sprintf(" course_id=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, class.CourseID)
	}

	if class.CourseCode != "" {
		query.WriteString(fmt.Sprintf(" course_code=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, class.CourseCode)
	}
	queryString := query.String()
	queryString = queryString[:len(queryString)-1]
	queryString += " WHERE id=$1"

	if tag, err := database.Pool.Exec(context.Background(), queryString, queryParams...); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.ErrInternalServerError
	}
	return c.JSON(models.RespMsg{Message: "Класс успешно обновлен"})
}

func deleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "DELETE FROM course_cycle WHERE id = $1"
	if tag, err := database.Pool.Exec(context.Background(), query, id); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.NewError(fiber.StatusNotFound, "Класс не найден")
	}
	return c.JSON(models.RespMsg{Message: "Класс успешно удален"})
}
