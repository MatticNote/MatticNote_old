package apiV1

import (
	"context"
	"github.com/MatticNote/MatticNote/db"
	"github.com/google/uuid"
	"github.com/savsgio/atreugo/v11"
)

func GetEntryUser(ctx *atreugo.RequestCtx) error {
	targetUuid := ctx.UserValue("uuid").(string)

	parse, err := uuid.Parse(targetUuid)
	if err != nil {
		ctx.SetStatusCode(400)
		return nil
	}

	var data MNAPIV1User
	query := db.DB.QueryRow(
		context.Background(),
		"SELECT uuid, username, display_name, summary, created_at, updated_at, is_bot FROM \"user\" WHERE uuid = $1",
		parse.String(),
	)
	err = query.Scan(
		&data.Uuid,
		&data.Username,
		&data.DisplayName,
		&data.Summary,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.IsBot,
	)
	if err != nil {
		if err.Error() == "no rows in result set" {
			ctx.SetStatusCode(404)
			return nil
		} else {
			return err
		}
	}

	return ctx.JSONResponse(data, 200)
}
