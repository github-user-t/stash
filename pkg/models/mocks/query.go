package mocks

import (
	context "context"

	"github.com/stashapp/stash/pkg/models"
)

type galleryResolver struct {
	galleries []*models.Gallery
}

func (g galleryResolver) FindMany(ctx context.Context, ids []int) ([]*models.Gallery, error) {
	return g.galleries, nil
}

func (g galleryResolver) Find(ctx context.Context, id int) (*models.Gallery, error) {
	panic("not implemented")
}

type sceneResolver struct {
	scenes []*models.Scene
}

func (s *sceneResolver) Find(ctx context.Context, id int) (*models.Scene, error) {
	panic("not implemented")
}

func (s *sceneResolver) FindMany(ctx context.Context, ids []int) ([]*models.Scene, error) {
	return s.scenes, nil
}

func GalleryQueryResult(galleries []*models.Gallery, count int) *models.GalleryQueryResult {
	ret := models.NewGalleryQueryResult(&galleryResolver{
		galleries: galleries,
	})

	ret.Count = count
	return ret
}

func SceneQueryResult(scenes []*models.Scene, count int) *models.SceneQueryResult {
	ret := models.NewSceneQueryResult(&sceneResolver{
		scenes: scenes,
	})

	ret.Count = count
	return ret
}

type imageResolver struct {
	images []*models.Image
}

func (s *imageResolver) Find(ctx context.Context, id int) (*models.Image, error) {
	panic("not implemented")
}

func (s *imageResolver) FindMany(ctx context.Context, ids []int) ([]*models.Image, error) {
	return s.images, nil
}

func ImageQueryResult(images []*models.Image, count int) *models.ImageQueryResult {
	ret := models.NewImageQueryResult(&imageResolver{
		images: images,
	})

	ret.Count = count
	return ret
}
