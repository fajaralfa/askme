package repo

import (
	"database/sql"
	"time"

	"github.com/fajaralfa/askme/internal/model"
)

type QuestionInterface interface {
	FindAllByUserEmail(userid string) ([]model.Question, error)
	FindAllByUserId(userid string) ([]model.Question, error)
	FindByUserId(qid, userid string) (*model.Question, error)
	FindByUserEmail(qid, email string) (*model.Question, error)
	Create(questionText string, userid int64) (*model.Question, error)
	DeleteByQIdUserId(id, userid string) (int64, error)
}

type Question struct {
	Db *sql.DB
}

func (q *Question) FindAllByUserEmail(email string) ([]model.Question, error) {
	questions := make([]model.Question, 0)
	rows, err := q.Db.Query("SELECT questions.* FROM questions LEFT JOIN users ON user_id = users.id WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var question model.Question
		if err := rows.Scan(&question.Id, &question.Question, &question.CreatedAt, &question.UserId); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return questions, nil
}

func (q *Question) FindAllByUserId(userid string) ([]model.Question, error) {
	questions := make([]model.Question, 0)
	rows, err := q.Db.Query("SELECT * FROM questions WHERE user_id = ?", userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var question model.Question
		if err := rows.Scan(&question.Id, &question.Question, &question.CreatedAt, &question.UserId); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return questions, nil
}

func (q *Question) FindByUserEmail(qid, email string) (*model.Question, error) {
	var question *model.Question
	rows, err := q.Db.Query("SELECT questions.* FROM questions LEFT JOIN users ON user_id = users.id WHERE questions.id = ? AND email = ?", qid, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		question = new(model.Question)
		if err := rows.Scan(&question.Id, &question.Question, &question.CreatedAt, &question.UserId); err != nil {
			return nil, err
		}
	}

	return question, nil
}

func (q *Question) FindByUserId(qid, userid string) (*model.Question, error) {
	var question *model.Question
	rows, err := q.Db.Query("SELECT * FROM questions WHERE id = ? AND user_id = ?", qid, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		question = new(model.Question)
		if err := rows.Scan(&question.Id, &question.Question, &question.CreatedAt, &question.UserId); err != nil {
			return nil, err
		}
	}

	return question, nil
}

func (q *Question) Create(questionText string, userid int64) (*model.Question, error) {
	question := new(model.Question)
	created_at := time.Now()
	res, err := q.Db.Exec("INSERT INTO questions (question, created_at, user_id) VALUES (?,?,?)", questionText, time.Now(), userid)
	if err != nil {
		return nil, err
	}
	lastid, _ := res.LastInsertId()

	question.Id = lastid
	question.Question = questionText
	question.CreatedAt = created_at.String()
	question.UserId = userid
	return question, nil
}

func (q *Question) DeleteByQIdUserId(id, userid string) (int64, error) {
	res, err := q.Db.Exec("DELETE FROM questions WHERE id = ? and user_id = ?", id, userid)
	affected, _ := res.RowsAffected()
	return affected, err
}
