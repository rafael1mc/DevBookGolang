package repositories

import (
	"api/src/models"
	"database/sql"
)

type Publications struct {
	db *sql.DB
}

func NewPublicationRepository(db *sql.DB) *Publications {
	return &Publications{db}
}

func (repository Publications) Create(publication models.Publication) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO publications (title, content, author_id) VALUES (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(publication.Title, publication.Content, publication.AuthorID)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}

func (repository Publications) GetByID(publicationID uint64) (models.Publication, error) {
	line, err := repository.db.Query(`
		SELECT p.*, u.nick FROM
		publications p INNER JOIN users u
		ON u.id = p.author_id WHERE p.id = ?`,
		publicationID,
	)
	if err != nil {
		return models.Publication{}, err
	}
	defer line.Close()

	var publication models.Publication

	if line.Next() {
		if err = line.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return models.Publication{}, err
		}
	}

	return publication, nil
}

func (repository Publications) Get(userID uint64) ([]models.Publication, error) {
	lines, err := repository.db.Query(`
	SELECT p.*, u.nick FROM publications p 
	INNER JOIN users u ON u.id = p.author_id 
	WHERE u.id = ? ORDER BY 1 DESC`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var publications []models.Publication

	for lines.Next() {
		var publication models.Publication

		if err = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (repository Publications) Update(publicationID uint64, publication models.Publication) error {
	statement, err := repository.db.Prepare("UPDATE publications SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publication.Title, publication.Content, publicationID); err != nil {
		return err
	}

	return nil
}

func (repository Publications) Delete(publicationId uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM publications WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicationId); err != nil {
		return err
	}

	return nil
}

func (repository Publications) GetByUser(userID uint64) ([]models.Publication, error) {
	lines, err := repository.db.Query(`
		SELECT p.*, u.nick FROM publications p
		JOIN users u ON u.id = p.author_id
		WHERE p.author_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var publications []models.Publication

	for lines.Next() {
		var publication models.Publication

		if err = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (repository Publications) Like(publicationID uint64) error {
	statement, err := repository.db.Prepare("UPDATE publications SET likes = likes + 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}

func (repository Publications) Unlike(publicationID uint64) error {
	statement, err := repository.db.Prepare(`
		UPDATE publications SET likes = 
		CASE 
			WHEN likes > 0 THEN likes - 1
			ELSE 0 
		END
		WHERE id = ?
	`)
	if err != nil {
		return err
	}

	if _, err = statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}
