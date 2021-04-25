package apiV1

import (
	"context"
	"github.com/MatticNote/MatticNote/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
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
	var isActive bool
	var isSuspend bool
	query := db.DB.QueryRow(
		context.Background(),
		"SELECT uuid, username, display_name, summary, created_at, updated_at, is_bot, is_active, is_suspend FROM \"user\" WHERE uuid = $1",
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
		&isActive,
		&isSuspend,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.SetStatusCode(404)
			return nil
		} else {
			return err
		}
	}
	if !isActive {
		ctx.SetStatusCode(410)
		return nil
	}
	if isSuspend {
		ctx.SetStatusCode(403)
		return nil
	}

	return ctx.JSONResponse(data, 200)
}
