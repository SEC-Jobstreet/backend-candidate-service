package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/SEC-Jobstreet/backend-candidate-service/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type applyRequest struct {
	JobID uuid.UUID `json:"job_id" binding:"required"`
}

type applyResponse struct {
	Status string `json:"status"`
}

// @Summary		Apply Job
// @Description	Candidate applies job
// @Param		request	body	applyRequest	true	"Apply Job"
// @Tags		application
// @Accept		json
// @Produce		json
// @Success		200	{object}	applyResponse
// @Router		/apply_job [post]
func (s *Server) apply(ctx *gin.Context) {
	var request applyRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	currentUser, err := utils.GetCurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	id, err := uuid.NewRandom()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	// Status: reviewing, accepted, denied
	arg := &models.Applications{
		ID:          id,
		CandidateID: currentUser.Username,
		JobID:       request.JobID,
		Status:      utils.ACCETPED, // change
	}

	err = s.store.Create(arg).Error
	if err != nil {
		if err == gorm.ErrDuplicatedKey || err == gorm.ErrForeignKeyViolated {
			ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	res := &applyResponse{
		Status: arg.Status,
	}
	ctx.JSON(http.StatusOK, res)
}

type applicationsRequest struct {
	ApplicationID string `uri:"application_id" binding:"required"`
}

// @Summary		Get Application
// @Description	get application by id
// @Param application_id path int true "Application ID"
// @Tags		application
// @Accept		json
// @Produce		json
// @Success		200	{object}	models.Applications
// @Router		/application_by_employer/{application_id} [get]
func (server *Server) getApplicationByEmployer(ctx *gin.Context) {
	var req applicationsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	id, err := uuid.Parse(req.ApplicationID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	app := &models.Applications{}

	err = server.store.Where("id = ?", id).First(app).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	// check this employer post this job of application
	// get candidate info

	ctx.JSON(http.StatusOK, app)
}

type listApplicationsRequest struct {
	CandidateID string `form:"candidate_id"`
	JobID       string `form:"job_id"`
	Status      string `form:"status"`
	PageID      int    `form:"page_id" binding:"required,min=1"`
	PageSize    int    `form:"page_size" binding:"required,min=10,max=20"`
}

type listApplicationsResponse struct {
	Total        int64                 `json:"total"`
	PageID       int                   `json:"page_id"`
	PageSize     int                   `json:"page_size" `
	Applications []models.Applications `json:"applications"`
}

// @Summary		List Applications
// @Description	get Applications
// @Param		q    query    listApplicationsRequest true	"search"
// @Tags		application
// @Accept		json
// @Produce		json
// @Success		200	{array}	listApplicationsResponse
// @Router		/application_list_by_employer [get]
func (server *Server) listApplicationsByEmployer(ctx *gin.Context) {
	var req listApplicationsRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	if req.Status != "" && !utils.IsSupportedStatus(req.Status) {
		err = errors.New("status is not supported")
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	list := []models.Applications{}
	var total int64

	var jobid uuid.UUID
	id, err := uuid.Parse(req.JobID)
	if err == nil {
		jobid = id
	}

	fmt.Println(jobid)
	tx := server.store.Model(&models.Applications{}).
		Where("candidate_id = @candidateid OR @candidateid = ''", sql.Named("candidateid", req.CandidateID)).
		Where("job_id = @jobid OR @jobid = '00000000-0000-0000-0000-000000000000'", sql.Named("jobid", jobid)).
		Where("status = @status OR @status = ''", sql.Named("status", strings.ToUpper(req.Status)))
	if tx.Error != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	tx.Count(&total)
	err = tx.Limit(req.PageSize).Offset((req.PageID - 1) * req.PageSize).Find(&list).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	// get profile from candidate service TO DO
	// for _, item := range list {
	// 	requestURL := fmt.Sprintf("http://localhost:4002/api/v1/profile_by_employer/%s", item.CandidateID)

	// 	request, err := http.NewRequest(http.MethodGet, requestURL, nil)
	// 	request.Header.Set("authorization", ctx.GetHeader("authorization"))
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
	// 		return
	// 	}

	// 	res, err := http.DefaultClient.Do(request)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
	// 		return
	// 	}

	// 	fmt.Printf("client: got response!\n")
	// 	fmt.Printf("client: status code: %d\n", res.StatusCode)

	// 	resBody, err := io.ReadAll(res.Body)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
	// 		return
	// 	}
	// 	fmt.Printf("client: response body: %s\n", resBody)
	// }

	rsp := &listApplicationsResponse{
		Total:        total,
		PageID:       req.PageID,
		PageSize:     req.PageSize,
		Applications: list,
	}

	ctx.JSON(http.StatusOK, rsp)
}

type applicationNumberByJobIdRequest struct {
	JobId string `uri:"job_id" binding:"required"`
}

func (server *Server) getApplicationNumberByJobId(ctx *gin.Context) {
	var req applicationNumberByJobIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	var total int64

	server.store.Model(&models.Applications{}).
		Where("job_id = ?", req.JobId).Count(&total)

	ctx.JSON(http.StatusOK, gin.H{"total": total})
}

type updateStatusRequest struct {
	ID     uuid.UUID `json:"id" binding:"required"`
	Status string    `json:"status" binding:"required,status"`
}

// @Summary		Update tags
// @Description	Update status of application
// @Param		request	body	updateStatusRequest	true	"update status by id"
// @Tags		application
// @Accept		json
// @Produce		json
// @Success		200	{object}	models.Applications
// @Router		/update_status [put]
func (server *Server) updateStatus(ctx *gin.Context) {
	var req updateStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	// check this employer post this job of application

	fmt.Println(req)
	app := models.Applications{
		ID: req.ID,
	}

	err := server.store.Model(&app).Update("status", strings.ToUpper(req.Status)).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, app)
}
