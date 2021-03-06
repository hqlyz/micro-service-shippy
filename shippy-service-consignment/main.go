package main

import (
	"context"
	"log"
	pb "micro-service-shippy/shippy-service-consignment/proto/consignment"
	"sync"

	"github.com/micro/go-micro"
)

const (
	port = ":5051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

// Create a new consignment
func (repo *Repository) Create(c *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	repo.consignments = append(repo.consignments, c)
	repo.mu.Unlock()
	return c, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type consignmentService struct {
	repo repository
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *consignmentService) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {

	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	resp.Consignment = consignment
	resp.Created = true
	return nil
}

func (s *consignmentService) GetConsignments(ctx context.Context, in *pb.GetRequest, resp *pb.Response) error {
	resp.Consignments = s.repo.GetAll()
	return nil
}

func main() {
	// serv := service.New(service.Name("consignment"))
	serv := micro.NewService(micro.Name("consignment"))
	serv.Init()
	
	repo := &Repository{}
	pb.RegisterShippingServiceHandler(serv.Server(), &consignmentService{repo})

	if err := serv.Run(); err != nil {
		log.Fatal(err)
	}
	// pb.RegisterShippingServiceHandler(service.Server(), &consignmentService{repo})

	// // Register reflection service on gRPC server.
	// reflection.Register(server)

	// log.Println("Running on port:", port)
	// if err := server.Serve(l); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
}
