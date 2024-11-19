package v1

import (
	"context"
	"encoding/json"

	"connectrpc.com/connect"
	"go.uber.org/zap"

	"github.com/crlssn/getstronger/apps/backend/pkg/jwt"
	v1 "github.com/crlssn/getstronger/apps/backend/pkg/pb/api/v1"
	"github.com/crlssn/getstronger/apps/backend/pkg/pb/api/v1/apiv1connect"
	"github.com/crlssn/getstronger/apps/backend/pkg/repo"
	"github.com/crlssn/getstronger/apps/backend/pkg/xzap"
)

var _ apiv1connect.RoutineServiceHandler = (*routineHandler)(nil)

type routineHandler struct {
	log  *zap.Logger
	repo *repo.Repo
}

func NewRoutineHandler(log *zap.Logger, r *repo.Repo) apiv1connect.RoutineServiceHandler {
	return &routineHandler{log, r}
}

func (h *routineHandler) Create(ctx context.Context, req *connect.Request[v1.CreateRoutineRequest]) (*connect.Response[v1.CreateRoutineResponse], error) {
	log := h.log.With(xzap.FieldRPC(apiv1connect.RoutineServiceCreateProcedure))
	log.Info("request received")

	userID := jwt.MustExtractUserID(ctx)
	log = log.With(xzap.FieldUserID(userID))

	routine, err := h.repo.CreateRoutine(ctx, repo.CreateRoutineParams{
		UserID:      userID,
		Name:        req.Msg.GetName(),
		ExerciseIDs: req.Msg.GetExerciseIds(),
	})
	if err != nil {
		log.Error("create routine failed", zap.Error(err))
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	log.Info("routine created")
	return connect.NewResponse(&v1.CreateRoutineResponse{
		Id: routine.ID,
	}), nil
}

func (h *routineHandler) Get(ctx context.Context, req *connect.Request[v1.GetRoutineRequest]) (*connect.Response[v1.GetRoutineResponse], error) {
	log := h.log.With(xzap.FieldRPC(apiv1connect.RoutineServiceGetProcedure))
	log.Info("request received")

	userID := jwt.MustExtractUserID(ctx)
	log = log.With(xzap.FieldUserID(userID))

	routine, err := h.repo.GetRoutine(ctx,
		repo.GetRoutineWithID(req.Msg.GetId()),
		repo.GetRoutineWithUserID(userID),
		repo.GetRoutineWithExercises(),
	)
	if err != nil {
		log.Error("get routine failed", zap.Error(err))
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	log.Info("routine returned")
	return connect.NewResponse(&v1.GetRoutineResponse{
		Routine: parseRoutineToPB(routine),
	}), nil
}

func (h *routineHandler) Update(ctx context.Context, req *connect.Request[v1.UpdateRoutineRequest]) (*connect.Response[v1.UpdateRoutineResponse], error) {
	log := h.log.With(xzap.FieldRPC(apiv1connect.RoutineServiceUpdateProcedure))
	log.Info("request received")

	userID := jwt.MustExtractUserID(ctx)
	log = log.With(xzap.FieldUserID(userID))

	routine, err := h.repo.GetRoutine(ctx, repo.GetRoutineWithID(req.Msg.GetRoutine().GetId()))
	if err != nil {
		log.Error("find exercise failed", zap.Error(err))
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	if routine.UserID != userID {
		log.Error("exercise does not belong to user")
		return nil, connect.NewError(connect.CodePermissionDenied, nil)
	}

	var updateRoutineOpts []repo.UpdateRoutineOpt
	for _, path := range req.Msg.GetUpdateMask().GetPaths() {
		switch path {
		case "name":
			updateRoutineOpts = append(updateRoutineOpts, repo.UpdateRoutineName(req.Msg.GetRoutine().GetName()))
		default:
			log.Error("invalid update mask path", zap.String("path", path))
			return nil, connect.NewError(connect.CodeInvalidArgument, errInvalidUpdateMaskPath)
		}
	}

	if err = h.repo.UpdateRoutine(ctx, routine.ID, updateRoutineOpts...); err != nil {
		log.Error("update exercise failed", zap.Error(err))
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	log.Info("routine updated")
	return connect.NewResponse(&v1.UpdateRoutineResponse{
		Routine: parseRoutineToPB(routine),
	}), nil
}

func (h *routineHandler) Delete(ctx context.Context, req *connect.Request[v1.DeleteRoutineRequest]) (*connect.Response[v1.DeleteRoutineResponse], error) {
	log := h.log.With(xzap.FieldRPC(apiv1connect.RoutineServiceDeleteProcedure))
	log.Info("request received")

	userID := jwt.MustExtractUserID(ctx)
	log = log.With(xzap.FieldUserID(userID))

	routine, err := h.repo.GetRoutine(ctx, repo.GetRoutineWithID(req.Msg.GetId()))
	if err != nil {
		log.Error("find routine failed", zap.Error(err))
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	if routine.UserID != userID {
		log.Error("routine does not belong to user")
		return nil, connect.NewError(connect.CodePermissionDenied, nil)
	}

	if err = h.repo.DeleteRoutine(ctx, req.Msg.GetId()); err != nil {
		log.Error("delete routine failed", zap.Error(err))
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	log.Info("routine deleted")
	return connect.NewResponse(&v1.DeleteRoutineResponse{}), nil
}

func (h *routineHandler) List(ctx context.Context, req *connect.Request[v1.ListRoutinesRequest]) (*connect.Response[v1.ListRoutinesResponse], error) { //nolint:dupl
	log := h.log.With(xzap.FieldRPC(apiv1connect.RoutineServiceListProcedure))
	log.Info("request received")

	userID := jwt.MustExtractUserID(ctx)
	log = log.With(xzap.FieldUserID(userID))

	limit := int(req.Msg.GetPageLimit())
	routines, err := h.repo.ListRoutines(ctx,
		repo.ListRoutinesWithName(req.Msg.GetName()),
		repo.ListRoutinesWithLimit(limit+1),
		repo.ListRoutinesWithUserID(userID),
		repo.ListRoutinesWithPageToken(req.Msg.GetPageToken()),
	)
	if err != nil {
		log.Error("list routines failed", zap.Error(err))
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	var nextPageToken []byte
	if len(routines) > limit {
		routines = routines[:limit]
		if nextPageToken, err = json.Marshal(repo.PageToken{
			CreatedAt: routines[len(routines)-1].CreatedAt,
		}); err != nil {
			log.Error("marshal page token failed", zap.Error(err))
			return nil, connect.NewError(connect.CodeInternal, nil)
		}
	}

	log.Info("routines listed")
	return connect.NewResponse(&v1.ListRoutinesResponse{
		Routines:      parseRoutinesToPB(routines),
		NextPageToken: nextPageToken,
	}), nil
}

func (h *routineHandler) AddExercise(ctx context.Context, req *connect.Request[v1.AddExerciseRequest]) (*connect.Response[v1.AddExerciseResponse], error) { //nolint:dupl
	log := h.log.With(xzap.FieldRPC(apiv1connect.RoutineServiceAddExerciseProcedure))
	log.Info("request received")

	userID := jwt.MustExtractUserID(ctx)
	log = log.With(xzap.FieldUserID(userID))

	routine, err := h.repo.GetRoutine(ctx,
		repo.GetRoutineWithID(req.Msg.GetRoutineId()),
		repo.GetRoutineWithUserID(userID),
	)
	if err != nil {
		log.Error("find routine failed", zap.Error(err))
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	exercise, err := h.repo.GetExercise(ctx,
		repo.GetExerciseWithID(req.Msg.GetExerciseId()),
		repo.GetExerciseWithUserID(userID),
	)
	if err != nil {
		log.Error("find exercise failed", zap.Error(err))
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	if err = h.repo.AddExerciseToRoutine(ctx, exercise, routine); err != nil {
		log.Error("add exercise to routine failed", zap.Error(err))
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	log.Info("exercise added to routine")
	return connect.NewResponse(&v1.AddExerciseResponse{}), nil
}

func (h *routineHandler) RemoveExercise(ctx context.Context, req *connect.Request[v1.RemoveExerciseRequest]) (*connect.Response[v1.RemoveExerciseResponse], error) { //nolint:dupl
	log := h.log.With(xzap.FieldRPC(apiv1connect.RoutineServiceRemoveExerciseProcedure))
	log.Info("request received")

	userID := jwt.MustExtractUserID(ctx)
	log = log.With(xzap.FieldUserID(userID))

	routine, err := h.repo.GetRoutine(ctx,
		repo.GetRoutineWithID(req.Msg.GetRoutineId()),
		repo.GetRoutineWithUserID(userID),
	)
	if err != nil {
		log.Error("find routine failed", zap.Error(err))
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	exercise, err := h.repo.GetExercise(ctx,
		repo.GetExerciseWithID(req.Msg.GetExerciseId()),
		repo.GetExerciseWithUserID(userID),
	)
	if err != nil {
		log.Error("find exercise failed", zap.Error(err))
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	if err = h.repo.RemoveExerciseFromRoutine(ctx, exercise, routine); err != nil {
		log.Error("remove exercise from routine failed", zap.Error(err))
		return nil, connect.NewError(connect.CodeInternal, nil)
	}

	log.Info("exercise removed from routine")
	return connect.NewResponse(&v1.RemoveExerciseResponse{}), nil
}
