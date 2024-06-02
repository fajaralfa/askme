package repo

import (
	"database/sql"

	"github.com/fajaralfa/askme/internal/model"
)

type QuestionInterface interface {
	FindAssociatedWithUser(email string) ([]model.Question, error)
	Find(id string) (*model.Question, error)
	Create(text string, userId string) (*model.Question, error)
	Delete(id string) error
}

type Question struct {
	Db *sql.DB
}

func (q *Question) FindAssociatedWithUser(email string) ([]model.Question, error) {
	questions := make([]model.Question, 0)
	rows, err := q.Db.Query("SELECT questions.id, questions.question, questions.created_at, questions.user_id FROM questions LEFT JOIN users ON users.id = questions.user_id  WHERE users.email = ?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var question model.Question
		if err := rows.Scan(&question.Id, &question.Text, &question.Date, &question.UserId); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return questions, nil
}

func (q *Question) Find(id string) (*model.Question, error) {
	question := new(model.Question)
	return question, nil
}

func (q *Question) Create(text string, userId string) (*model.Question, error) {
	question := new(model.Question)
	return question, nil
}

func (q *Question) Delete(id string) error {
	return nil
}
