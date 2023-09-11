package users

import (
	"context"
	"fmt"
	"iwexlmsapi/database"
	"iwexlmsapi/models"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func findOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := `
	SELECT users.id,
		users.first_name,
		users.last_name,
		users.email,
		users.contact_number,
		users.date_of_birth,
		users.is_active,
		users.role,
		role.role_name
	FROM users
		INNER JOIN role ON users.role = role.id
	WHERE (users.id = $1)`
	user := models.UserDB{}
	if err := database.Pool.QueryRow(context.Background(), query, id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.ContactNumber,
		&user.DateOfBirth,
		&user.IsActive,
		&user.Role,
		&user.RoleName,
	); err != nil {
		return err
	}
	return c.JSON(user)
}

func findMany(c *fiber.Ctx) error {
	query := `
	SELECT users.id,
		users.first_name,
		users.last_name,
		users.email,
		users.contact_number,
		users.date_of_birth,
		users.is_active,
		users.role,
		role.role_name
	FROM users
		INNER JOIN role ON users.role = role.id`
	rows, err := database.Pool.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		return err
	}
	users, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByNameLax[models.UserDB])
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Пользователи не найдены ")
	}
	return c.JSON(users)
}

func updateOne(c *fiber.Ctx) error {
	id := c.Params("id")
	user := c.Locals("body").(*models.UpdateUserDTO)

	if (*user == models.UpdateUserDTO{}) {
		return fiber.NewError(fiber.StatusBadRequest, "Не указаны данные для обновления")
	}

	isActive := false
	q := "SELECT is_active FROM users WHERE id = $1"
	if err := database.Pool.QueryRow(context.Background(), q, id).Scan(&isActive); err != nil {
		return err
	}

	query := strings.Builder{}
	query.WriteString("UPDATE users SET")
	queryParams := []any{id}

	if user.FirstName != "" {
		query.WriteString(fmt.Sprintf(" first_name=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, user.FirstName)
	}

	if user.LastName != "" {
		query.WriteString(fmt.Sprintf(" last_name=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, user.LastName)
	}

	if user.ContactNumber != "" {
		query.WriteString(fmt.Sprintf(" contact_number=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, user.ContactNumber)
	}

	if user.DateOfBirth != "" {
		query.WriteString(fmt.Sprintf(" date_of_birth=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, user.DateOfBirth)
	}

	if user.Email != "" {
		query.WriteString(fmt.Sprintf(" email=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, user.Email)
	}

	if user.Role != 0 {
		query.WriteString(fmt.Sprintf(" role=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, user.Role)
	}

	if user.IsActive != isActive {
		query.WriteString(fmt.Sprintf(" is_active=$%d,", len(queryParams)+1))
		queryParams = append(queryParams, user.IsActive)
	}

	queryString := query.String()
	queryString = queryString[:len(queryString)-1]
	queryString += " WHERE id=$1"

	fmt.Println(queryString)

	if tag, err := database.Pool.Exec(context.Background(), queryString, queryParams...); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.NewError(fiber.StatusNotFound, "Пользователь не найден")
	}
	return c.JSON(models.RespMsg{Message: "Данные пользователя успешно обновлены"})
}

func deleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	query := `DELETE FROM users WHERE id = $1`
	if tag, err := database.Pool.Exec(context.Background(), query, id); err != nil {
		return err
	} else if tag.RowsAffected() < 1 {
		return fiber.NewError(fiber.StatusNotFound, "Пользователь не найден")
	}
	return c.JSON(models.RespMsg{Message: "Пользователь успешно удален "})
}

func filterUsers(c *fiber.Ctx) error {
	var filter models.UsersFilter
	if err := c.QueryParser(&filter); err != nil {
		return err
	}

	query := `
        SELECT 
        u.id,
        u.first_name,
        u.last_name,
        u.email,
        u.contact_number,
        u.date_of_birth,
        u.is_active,
        u.role,
        r.role_name,
        e.cycle_id,
        e.student_id,
        cc.course_code,
        cc.course_id AS course_course_id,
        c.name AS course_name
        FROM users u
        INNER JOIN role r ON u.role = r.id
        LEFT JOIN enrollment e ON u.id = e.student_id
        LEFT JOIN course_cycle cc ON e.cycle_id = cc.id
        LEFT JOIN course c ON cc.course_id = c.course_id
        WHERE 1=1`

	var queryParams []interface{}

	if filter.FirstName != nil && *filter.FirstName != "" {
		query += " AND (u.first_name = $" + fmt.Sprint(len(queryParams)+1) + " OR $" + fmt.Sprint(len(queryParams)+1) + " = '')"
		queryParams = append(queryParams, *filter.FirstName)
	}

	if filter.LastName != nil && *filter.LastName != "" {
		query += " AND (u.last_name = $" + fmt.Sprint(len(queryParams)+1) + " OR $" + fmt.Sprint(len(queryParams)+1) + " = '')"
		queryParams = append(queryParams, *filter.LastName)
	}

	if filter.RoleName != nil && *filter.RoleName != "" {
		query += " AND (r.role_name = $" + fmt.Sprint(len(queryParams)+1) + " OR $" + fmt.Sprint(len(queryParams)+1) + " = '')"
		queryParams = append(queryParams, *filter.RoleName)
	}

	if filter.CourseName != nil {
		decodedCourseName, err := url.QueryUnescape(*filter.CourseName)
		if err != nil {
			return err
		}
		filter.CourseName = &decodedCourseName

		query += " AND (TRIM(c.name) = $" + fmt.Sprint(len(queryParams)+1) + " OR $" + fmt.Sprint(len(queryParams)+1) + " = '')"
		queryParams = append(queryParams, *filter.CourseName)
	}

	if filter.CourseCode != nil {
		decodedCourseCode, err := url.QueryUnescape(*filter.CourseCode)
		if err != nil {
			return err
		}

		query += " AND (TRIM(cc.course_code) = $" + fmt.Sprint(len(queryParams)+1) + " OR $" + fmt.Sprint(len(queryParams)+1) + " = '')"
		queryParams = append(queryParams, decodedCourseCode)
	}

	if filter.IsActive != nil {
		if *filter.IsActive {

			query += " AND u.is_active = true"
		} else {
			query += " AND u.is_active = false"
		}
	}

	rows, err := database.Pool.Query(context.Background(), query, queryParams...)
	if err != nil {
		return err
	}

	defer rows.Close()

	var users []models.UsersFilter

	for rows.Next() {
		var user models.UsersFilter
		if err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.ContactNumber,
			&user.DateOfBirth,
			&user.IsActive,
			&user.Role,
			&user.RoleName,
			&user.CycleID,
			&user.StudentID,
			&user.CourseCode,
			&user.CourseCourseID,
			&user.CourseName,
		); err != nil {
			return err
		}
		users = append(users, user)
	}

	return c.JSON(users)
}
