package externals

import (
	"context"

	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type JobServiceGRPC struct {
	serviceAddress string
	conn           *grpc.ClientConn
}

func NewJobServiceGRPC(serverAddress string, conn *grpc.ClientConn) *JobServiceGRPC {
	return &JobServiceGRPC{
		serviceAddress: serverAddress,
		conn:           conn,
	}
}

func (js *JobServiceGRPC) GetJob(id string) (*models.Jobs, error) {
	client := pb.NewJobServiceClient(js.conn)

	job, err := client.GetJobByID(context.Background(), &pb.GetJobByIDRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	jobId, _ := uuid.Parse(job.Job.GetId())
	enterpriseId, _ := uuid.Parse(job.Job.GetEnterpriseId())

	return &models.Jobs{
		ID:                 jobId,
		EmployerID:         job.Job.GetEmployerId(),
		Status:             job.Job.GetStatus(),
		Title:              job.Job.GetTitle(),
		Type:               job.Job.GetType(),
		WorkWhenever:       job.Job.GetWorkWhenever(),
		WorkShift:          job.Job.GetWorkShift(),
		Description:        job.Job.GetDescription(),
		Visa:               job.Job.GetVisa(),
		Experience:         job.Job.GetExperience(),
		StartDate:          job.Job.GetStartDate(),
		Currency:           job.Job.GetCurrency(),
		SalaryLevelDisplay: job.Job.GetSalaryLevelDisplay(),
		ExactSalary:        job.Job.GetExactSalary(),
		RangeSalary:        job.Job.GetRangeSalary(),
		PaidPeriod:         job.Job.GetPaidPeriod(),
		ExpiresAt:          job.Job.GetExpiresAt(),

		CreatedAt: job.Job.GetCreatedAt(),
		UpdatedAt: job.Job.GetUpdatedAt(),

		EnterpriseID:      enterpriseId,
		EnterpriseName:    job.Job.GetEnterpriseName(),
		EnterpriseAddress: job.Job.GetEnterpriseAddress(),

		Crawl:         job.Job.GetCrawl(),
		JobURL:        job.Job.GetJobUrl(),
		JobSourceName: job.Job.GetJobSourceName(),
	}, nil
}
