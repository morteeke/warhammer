package sqlite

import (
	"context"

	"github.com/brittonhayes/warhammer/api"
	"github.com/uptrace/bun"
)

func (r *warhammerRepository) warscrollFilterQuery(query *bun.SelectQuery, params api.GetWarscrollsParams) (*bun.SelectQuery, error) {
	if params.Name != nil {
		query = query.Where("? LIKE ?", bun.Ident("name"), *params.Name+"%")
	}

	if params.Points != nil {
		query = query.Where("? LIKE ?", bun.Ident("points"), *params.Points)
	}

	if params.Size != nil {
		query = query.Where("? LIKE ?", bun.Ident("size"), *params.Size)
	}

	if params.BattlefieldRole != nil {
		query = query.Where("? LIKE ?", bun.Ident("battlefield_role"), "%"+*params.BattlefieldRole+"%")
	}

	if params.Notes != nil {
		query = query.Where("? LIKE ?", bun.Ident("notes"), "%"+*params.Notes+"%")
	}

	return query, nil
}

func (r *warhammerRepository) GetWarscrolls(ctx context.Context, params api.GetWarscrollsParams) ([]api.Warscroll, error) {
	var warscrolls []api.Warscroll

	query, err := r.warscrollFilterQuery(r.db.NewSelect().Model(&warscrolls), params)
	if err != nil {
		return nil, err
	}

	err = query.Scan(ctx, &warscrolls)
	if err != nil {
		return nil, err
	}

	return warscrolls, nil
}

func (r *warhammerRepository) GetWarscrollByID(ctx context.Context, id string) (*api.Warscroll, error) {
	var warscroll api.Warscroll

	err := r.db.NewSelect().Model(&warscroll).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &warscroll, nil
}
