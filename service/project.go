package service

import (
	config "back-end-website/config"
	"back-end-website/graph/model"
	"back-end-website/graph/mypkg"
	db "back-end-website/internal"
	"back-end-website/utils"
	"context"
	"errors"

	"github.com/google/uuid"
)

type IProjectService interface {
	CreateProject(ctx context.Context, input *model.CreateProjectInput) (*model.Project, error)
	UpdateProject(ctx context.Context, input *model.UpdateProjectInput) (*model.Project, error)
	DeleteProject(ctx context.Context, id mypkg.UUID) (*bool, error)
	GetProjects(ctx context.Context, limit int, offset int) ([]*model.Project, error)
	GetProject(ctx context.Context, id mypkg.UUID) (*model.Project, error)
}

type ProjectService struct {
	server *config.Server
}

func SqlProjectToGraphProject(sqlProject *db.Project) *model.Project {
	if sqlProject == nil {
		return nil
	}
	return &model.Project{
		ID:             mypkg.UUID(sqlProject.ID.String()),
		UserID:         mypkg.UUID(sqlProject.UserID.String()),
		Title:          sqlProject.Title,
		Content:        sqlProject.Content,
		CreatedAt:      sqlProject.CreatedAt,
		UpdatedAt:      sqlProject.UpdatedAt,
		DeletedAt:      &sqlProject.DeletedAt.Time,
		ImgCover:       sqlProject.ImgCover,
		ImgDescription: sqlProject.ImgDescription,
		URL:            sqlProject.Url,
		Language:       sqlProject.Language.String,
		Tag:            model.ProjectTag(sqlProject.Tag),
	}
}

func NewProjectService(server *config.Server) *ProjectService {
	return &ProjectService{
		server: server,
	}
}

func (s *ProjectService) CreateProject(ctx context.Context, input *model.CreateProjectInput) (*model.Project, error) {
	var res *model.Project

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		// check if user_id exists
		isUser, err := q.CheckUserByID(ctx, uuid.MustParse(string(input.UserID)))
		if err != nil {
			return err
		}
		if !isUser {
			return utils.ErrorResponse("USER_NOT_FOUND", errors.New("User not found"))
		}
		// create Project
		newProject, err := q.InsertProject(ctx, db.InsertProjectParams{
			Title:          input.Title,
			Content:        input.Content,
			Language:       utils.NullS(input.Language),
			Url:            input.URL,
			Tag:            db.ProjectTag(input.Tag),
			ImgCover:       input.ImgCover,
			ImgDescription: input.ImgDescription,
			UserID:         uuid.MustParse(string(input.UserID)),
		})
		if err != nil {
			return err
		}

		// get Project
		p, err := q.GetProjectByID(ctx, uuid.MustParse(string(newProject.ID.String())))
		if err != nil {
			return err
		}

		// convert to graphql model
		res = SqlProjectToGraphProject(&p)
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_CREATE_PROJECT", err)
	}
	return res, nil
}

func (s *ProjectService) UpdateProject(ctx context.Context, input *model.UpdateProjectInput) (*model.Project, error) {
	var res *model.Project

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		// update Project
		if err := q.UpdateProject(ctx, db.UpdateProjectParams{
			ID:             uuid.MustParse(string(input.ID)),
			Title:          input.Title,
			Content:        input.Content,
			Language:       utils.NullS(input.Language),
			Tag:            db.ProjectTag(input.Tag),
			Url:            input.URL,
			ImgCover:       input.ImgCover,
			ImgDescription: input.ImgDescription,
			UserID:         uuid.MustParse(string(input.UserID)),
		}); err != nil {
			return err
		}

		// get $
		p, err := q.GetProjectByID(ctx, uuid.MustParse(string(input.ID)))
		if err != nil {
			return err
		}

		// convert to graphql model
		res = SqlProjectToGraphProject(&p)
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_UPDATE_PROJECT", err)
	}
	return res, nil
}

func (s *ProjectService) DeleteProject(ctx context.Context, id mypkg.UUID) (*bool, error) {
	var res bool = false

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		// check if Project exists
		isProject, err := q.CheckProjectByID(ctx, uuid.MustParse(string(id)))
		if err != nil {
			return err
		}
		if !isProject {
			return utils.ErrorResponse("PROJECT_NOT_FOUND", errors.New("Project not found"))
		}
		if err := q.DeleteProjectByID(ctx, uuid.MustParse(string(id))); err != nil {
			res = false
			return err
		}
		res = true
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_DELETE_PROJECT", err)
	}
	return &res, nil
}

func (s *ProjectService) GetProject(ctx context.Context, id mypkg.UUID) (*model.Project, error) {
	var res *model.Project

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		project, err := q.GetProjectByID(ctx, uuid.MustParse(string(id)))
		if err != nil {
			return err
		}
		res = SqlProjectToGraphProject(&project)
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_GET_PROJECT", err)
	}
	return res, err
}

func (s *ProjectService) GetProjects(ctx context.Context, limit int, offset int) ([]*model.Project, error) {
	var res []*model.Project

	err := s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		// check if limit is valid
		if limit < 0 {
			return utils.ErrorResponse("INVALID_LIMIT", errors.New("limit must be greater than 0"))
		}
		// check if offset is valid
		if offset < 0 {
			return utils.ErrorResponse("INVALID_OFFSET", errors.New("offset must be greater than 0"))
		}
		ps, err := q.GetAllProject(ctx, db.GetAllProjectParams{
			Limit:    int32(limit),
			Offset:   int32(offset),
			TitleAsc: utils.FilterOrderBy("title", "asc", "TitleAsc"),
		})
		if err != nil {
			return err
		}

		// convert to graphql model
		for _, p := range ps {
			res = append(res, SqlProjectToGraphProject(&p))
		}
		return nil
	})

	if err != nil {
		return nil, utils.ErrorResponse("TX_GET_PROJECTS", err)
	}
	return res, nil
}
