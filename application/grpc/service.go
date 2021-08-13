package grpc

import (
	"context"

	"github.com/c-4u/company-service/domain/service"
	"github.com/c-4u/company-service/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcService struct {
	pb.UnimplementedCompanyServiceServer
	Service *service.Service
}

func NewGrpcService(service *service.Service) *GrpcService {
	return &GrpcService{
		Service: service,
	}
}

func (s *GrpcService) CreateCompany(ctx context.Context, in *pb.CreateCompanyRequest) (*pb.CreateCompanyResponse, error) {
	id, err := s.Service.CreateCompany(
		ctx,
		in.Company.CorporateName,
		in.Company.TradeName,
		in.Company.Cnpj,
	)
	if err != nil {
		return &pb.CreateCompanyResponse{}, err
	}

	return &pb.CreateCompanyResponse{Id: *id}, nil
}

func (s *GrpcService) FindCompany(ctx context.Context, in *pb.FindCompanyRequest) (*pb.FindCompanyResponse, error) {
	company, err := s.Service.FindCompany(ctx, in.Id)
	if err != nil {
		return &pb.FindCompanyResponse{}, err
	}

	return &pb.FindCompanyResponse{
		Company: &pb.Company{
			Id:            company.ID,
			CorporateName: company.CorporateName,
			TradeName:     company.TradeName,
			Cnpj:          company.Cnpj,
			CreatedAt:     timestamppb.New(company.CreatedAt),
			UpdatedAt:     timestamppb.New(company.UpdatedAt),
		},
	}, nil
}

func (s *GrpcService) SearchCompanies(ctx context.Context, in *pb.SearchCompaniesRequest) (*pb.SearchCompaniesResponse, error) {
	nextPageToken, companies, err := s.Service.SearchCompanies(ctx, in.Filter.CorporateName, in.Filter.TradeName, in.Filter.Cnpj, int(in.Filter.PageSize), in.Filter.PageToken)
	if err != nil {
		return &pb.SearchCompaniesResponse{}, err
	}

	var result []*pb.Company
	for _, company := range companies {
		result = append(
			result,
			&pb.Company{
				Id:            company.ID,
				CorporateName: company.CorporateName,
				TradeName:     company.TradeName,
				Cnpj:          company.Cnpj,
				CreatedAt:     timestamppb.New(company.CreatedAt),
				UpdatedAt:     timestamppb.New(company.UpdatedAt),
			},
		)
	}

	return &pb.SearchCompaniesResponse{NextPageToken: *nextPageToken, Companies: result}, nil
}

func (s *GrpcService) UpdateCompany(ctx context.Context, in *pb.UpdateCompanyRequest) (*pb.StatusResponse, error) {
	err := s.Service.UpdateCompany(ctx, in.Id, in.CorporateName, in.TradeName)
	if err != nil {
		return &pb.StatusResponse{
			Code:    uint32(status.Code(err)),
			Message: "not updated",
			Error:   err.Error(),
		}, err
	}

	return &pb.StatusResponse{
		Code:    uint32(codes.OK),
		Message: "successfully updated",
	}, nil
}
