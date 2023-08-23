package class

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"iwexlmsapi/models"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)


func getEnrollment(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	rows, err := db.Query(`
		SELECT e.*, u.*
		FROM enrollment AS e
		INNER JOIN users AS u ON e.student_id = u.id
		WHERE e.cycle_id = $1;
	`, id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}
	defer rows.Close()

	enrollments := []models.Enrollment{}
	for rows.Next() {
		var enrollment models.Enrollment
		err := rows.Scan(&enrollment.CycleID, &enrollment.StudentID) // ... остальные поля
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
		}
		enrollments = append(enrollments, enrollment)
	}

	return ctx.JSON(enrollments)
}

func addEnrollment(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var requestBody struct {
		Users []int `json:"users"`
	}
	if err := ctx.BodyParser(&requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "Invalid request format"})
	}

	tx, err := db.Begin()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO enrollment (cycle_id, student_id) VALUES ($1, $2)")
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}
	defer stmt.Close()

	for _, userID := range requestBody.Users {
		_, err := stmt.Exec(id, userID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
		}
	}

	if err := tx.Commit(); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]string{"message": "Успешно добавлено в enrollment"})
}

func findMany(ctx *fiber.Ctx) error {
	rows, err := db.Query(`
		SELECT course_cycle.id, course_cycle.description, course_cycle.start_date, course_cycle.end_date,
		       course_cycle.open_for_enrollment, course_cycle.course_code, branch_office.name as branch_name,
		       course.name as course_name
		FROM course_cycle
		INNER JOIN branch_office ON course_cycle.branch_id = branch_office.id
		INNER JOIN course ON course_cycle.course_id = course.course_id
	`)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}
	defer rows.Close()

	classes := []models.Class{}
	for rows.Next() {
		var class models.Class
		err := rows.Scan(&class.ID, &class.Description, &class.StartDate, &class.EndDate, &class.OpenForEnrollment,
			&class.CourseCode, &class.BranchName, &class.CourseName)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
		}
		classes = append(classes, class)
	}

	if len(classes) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(map[string]string{"error": "Классы не найдены"})
	}

	return ctx.JSON(classes)
}

func findOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var class models.Class
	err := db.QueryRow(`
		SELECT course_cycle.id, course_cycle.description, course_cycle.start_date, course_cycle.end_date,
		       course_cycle.open_for_enrollment, course_cycle.course_code, branch_office.name as branch_name,
		       course.name as course_name
		FROM course_cycle
		INNER JOIN branch_office ON course_cycle.branch_id = branch_office.id
		INNER JOIN course ON course_cycle.course_id = course.course_id
		WHERE course_cycle.id = $1
	`, id).Scan(&class.ID, &class.Description, &class.StartDate, &class.EndDate, &class.OpenForEnrollment,
		&class.CourseCode, &class.BranchName, &class.CourseName)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.Status(fiber.StatusNotFound).JSON(map[string]string{"error": fmt.Sprintf("Класс с id %s не найден", id)})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	return ctx.JSON(class)
}

func updateOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var requestBody models.ClassReqBodySchema
	if err := ctx.BodyParser(&requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "Invalid request format"})
	}

	tx, err := db.Begin()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}
	defer tx.Rollback()

	var oldClass models.Class
	err = tx.QueryRow(`
		SELECT * from course_cycle where id = $1
	`, id).Scan(&oldClass.ID, &oldClass.Description, &oldClass.StartDate, &oldClass.EndDate,
		&oldClass.OpenForEnrollment, &oldClass.CourseCode, &oldClass.BranchName, &oldClass.CourseName)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.Status(fiber.StatusNotFound).JSON(map[string]string{"error": "Класс не найден"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	newClass := models.Class{
		ID:                oldClass.ID,
		Description:       requestBody.Description,
		StartDate:         requestBody.StartDate,
		EndDate:           requestBody.EndDate,
		OpenForEnrollment: requestBody.OpenForEnrollment,
		CourseCode:        requestBody.CourseCode,
		BranchName:        oldClass.BranchName,
		CourseName:        oldClass.CourseName,
	}

	_, err = tx.Exec(`
		UPDATE course_cycle 
		SET course_code = $1, 
		description = $2,
		start_date = $3,
		end_date = $4,
		open_for_enrollment = $5,
		branch_id = $6,
		course_id = $7
		WHERE id = $8
	`, newClass.CourseCode, newClass.Description, newClass.StartDate, newClass.EndDate,
		newClass.OpenForEnrollment, oldClass.BranchID, oldClass.CourseID, id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": "Ошибка изменения данных"})
	}

	if err := tx.Commit(); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]string{"message": "Успешно обновлено"})
}

func deleteOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	tx, err := db.Begin()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM course_cycle WHERE id = $1", id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	if err := tx.Commit(); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]string{"message": "Успешно удалено"})
}


var db *sql.DB

func InitDB(connStr string) error {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	return nil
}
func ToggleOpenForEnrollment(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "Invalid id parameter"})
	}

	var body struct {
		OpenForEnrollment bool `json:"open_for_enrollment"`
	}
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "Invalid request body"})
	}

	query := `UPDATE course_cycle SET open_for_enrollment = $1 WHERE id = $2`
	result, err := db.Exec(query, body.OpenForEnrollment, id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": "Failed to update"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected < 1 {
		return ctx.Status(fiber.StatusNotFound).JSON(map[string]string{"error": fmt.Sprintf("Class with id %d not found", id)})
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]string{"message": "Successfully updated"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}