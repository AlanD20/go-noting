package api

import (
	"strings"

	"github.com/aland20/go-noting/app/response"
	"github.com/labstack/echo/v4"
)

type NoteController struct {
	App *BaseApp
}

func BindNoteApi(app *BaseApp, rg *echo.Group) {

	ctrl := &NoteController{App: app}

	rg.GET("/:username/notes", ctrl.IndexNote).Name = "note.index"
	rg.GET("/:username/notes/:id", ctrl.ShowNote).Name = "note.show"
	rg.POST("/:username/notes", ctrl.StoreNote).Name = "note.store"
	rg.PATCH("/:username/notes/:id", ctrl.UpdateNote).Name = "note.update"
	rg.DELETE("/:username/notes/:id", ctrl.DestroyNote).Name = "note.destroy"
}

func (ctrl *NoteController) IndexNote(ctx echo.Context) error {

	username := ctx.Param("username")
	var user User

	// Retrieve user
	if err := ctrl.App.Connection.Where("username = ?", username).Take(&user).Error; err != nil {
		return response.NotFound(ctx)
	}

	notes := &[]Note{}

	// Retrieve note
	err := ctrl.App.Connection.
		Model(&user).
		Select("id", "title", "body", "private", "created_at", "updated_at").
		Association("Notes").
		Find(&notes)

	if err != nil {
		return response.NotFound(ctx)
	}

	return response.Ok(ctx, notes)
}

func (ctrl *NoteController) ShowNote(ctx echo.Context) error {

	username := ctx.Param("username")
	var user User

	// Retrieve user
	if err := ctrl.App.Connection.Where("username = ?", username).Take(&user).Error; err != nil {
		return response.NotFound(ctx)
	}

	noteId := ctx.Param("id")
	note := &Note{}

	// Retrieve note
	err := ctrl.App.Connection.
		Model(&user).
		Select("id", "title", "body", "private", "created_at", "updated_at").
		Where("id", noteId).
		Association("Notes").
		Find(&note)

	if err != nil {
		return response.NotFound(ctx)
	}

	return response.Ok(ctx, note)
}

func (ctrl *NoteController) StoreNote(ctx echo.Context) error {

	username := ctx.Param("username")
	var user User

	// Retrieve user
	if err := ctrl.App.Connection.Where("username = ?", username).Take(&user).Error; err != nil {
		return response.NotFound(ctx)
	}

	n := &NoteSchema{}

	// Bind incoming form to u (NoteSchema) type
	if err := ctx.Bind(&n); err != nil {
		return response.BadRequest(ctx)
	}

	// Create note
	note := Note{
		NoteSchema: NoteSchema{
			Title:   strings.ToTitle(n.Title),
			Body:    n.Body,
			Private: n.Private,
		},
	}

	// Generate UUID
	note.GenerateId()

	// Append new note to user's relation
	if err := ctrl.App.Connection.Model(&user).Association("Notes").Append([]Note{note}); err != nil {
		return response.BadRequest(ctx)
	}

	return response.Created(ctx, &n)
}

func (ctrl *NoteController) UpdateNote(ctx echo.Context) error {

	username := ctx.Param("username")
	var user User

	// Retrieve user
	if err := ctrl.App.Connection.Where("username = ?", username).Take(&user).Error; err != nil {
		return response.NotFound(ctx)
	}

	noteId := ctx.Param("id")
	n := &NoteSchema{}

	// Bind incoming form to u (NoteSchema) type
	if err := ctx.Bind(&n); err != nil {
		return response.BadRequest(ctx)
	}

	note := &Note{}

	// Retrieve note
	err := ctrl.App.Connection.
		Model(&user).
		Select("id", "title", "body", "private", "created_at", "updated_at").
		Where("id", noteId).
		Association("Notes").
		Find(&note)

	if err != nil {
		return response.NotFound(ctx)
	}

	// Update note
	note.Title = strings.ToTitle(n.Title)
	note.Body = n.Body
	note.Private = n.Private

	// Update note for given user
	if err := ctrl.App.Connection.Where("id", noteId).Updates(&note).Error; err != nil {
		return response.UnprocessableEntity(ctx)
	}

	return response.Update(ctx, note)
}

func (ctrl *NoteController) DestroyNote(ctx echo.Context) error {

	username := ctx.Param("username")
	var user User

	// Retrieve user
	if err := ctrl.App.Connection.Where("username = ?", username).Take(&user).Error; err != nil {
		return response.NotFound(ctx)
	}

	noteId := ctx.Param("id")
	note := &Note{}

	// Retrieve note
	err := ctrl.App.Connection.
		Model(&user).
		Select("id", "title", "body", "private", "created_at", "updated_at").
		Where("id", noteId).
		Association("Notes").
		Find(&note)

	if err != nil {
		return response.NotFound(ctx)
	}

	// Delete note for given user
	if err := ctrl.App.Connection.Where("id", noteId).Delete(&note).Error; err != nil {
		return response.UnprocessableEntity(ctx)
	}

	return response.Delete(ctx)
}
