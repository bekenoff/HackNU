package dbs

import (
	"HackNU/pkg/models"
	"database/sql"
)

type VideoModel struct {
	DB *sql.DB
}

func (m *VideoModel) Insert(video *models.Video) error {
	stmt := `
        INSERT INTO ainur_hacknu.videos 
        ( filename, filepath, likes, client_id) 
        VALUES (?, ?, ?, ?);`

	_, err := m.DB.Exec(stmt, video.Filename, video.Filepath, video.Likes, video.ClientId)
	if err != nil {
		return err
	}

	return nil
}

func (m *VideoModel) GetAllVideos() ([]*models.Video, error) {
	stmt := `SELECT id, filename, filepath, likes, client_id FROM ainur_hacknu.videos;`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	videos := []*models.Video{}

	for rows.Next() {
		var video models.Video
		err := rows.Scan(&video.Id, &video.Filename, &video.Filepath, &video.Likes, &video.ClientId)
		if err != nil {
			return nil, err
		}
		videos = append(videos, &video)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return videos, nil
}

func (m *VideoModel) IncrementLike(videoID, clientID string) error {
	stmt := `INSERT INTO ainur_hacknu.likes (video_id, client_id)
VALUES (?, ?);
`
	_, err := m.DB.Exec(stmt, videoID, clientID)
	return err
}

func (m *VideoModel) GetAllVideosWithLikes() ([]*models.Video, error) {
	stmt := `
    SELECT v.id, v.filename, v.filepath, COALESCE(SUM(l.id IS NOT NULL), 0) AS likes_count
    FROM ainur_hacknu.videos v
    LEFT JOIN ainur_hacknu.likes l ON v.id = l.video_id
    GROUP BY v.id;`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos []*models.Video
	for rows.Next() {
		var v models.Video
		if err := rows.Scan(&v.Id, &v.Filename, &v.Filepath, &v.Likes); err != nil {
			return nil, err
		}
		videos = append(videos, &v)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return videos, nil
}
