package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	"fmt"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestJobRepository_INSERT(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	id := uuid.NewV4()
	video.ID = id.String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDB{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("outputh_path", "PENDING", video)

	require.Nil(t, err)

	jobRepo := repositories.JobRepositoryDB{Db: db}
	jobRepo.Insert(job)

	j, err := jobRepo.Find(job.ID)

	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.Video.ID, video.ID)

}

func TestJobRepositoryDBUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	id := uuid.NewV4().String()
	video.ID = id
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDB{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("outputh_path", "PENDING", video)

	require.Nil(t, err)

	jobRepo := repositories.JobRepositoryDB{Db: db}
	jobRepo.Insert(job)

	job.Status = "Complete"

	fmt.Println(*job)
	jobRepo.Update(job)
	j, err := jobRepo.Find(job.ID)
	fmt.Println(*job)
	fmt.Println(*j)

	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.Status, job.Status)
}
