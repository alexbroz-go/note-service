package handlers

import (
	"context"
	"log"
	"net"
	"note-service/models"
	"note-service/services"

	pb "note-service/note-service/noteservice" // Импортируйте сгенерированные протобуферы

	"google.golang.org/grpc"
)

type noteServer struct {
	pb.UnimplementedNoteServiceServer // Встраивание интерфейса для gRPC
}

// GetAllNotes обрабатывает gRPC запрос для получения всех заметок
func (s *noteServer) GetAllNotes(ctx context.Context, req *pb.Request) (*pb.NoteList, error) {
	notes, err := services.GetAllNotes()
	if err != nil {
		return nil, err
	}
	noteList := &pb.NoteList{}
	for _, note := range notes {
		noteList.Notes = append(noteList.Notes, &pb.Note{
			Id:   int32(note.ID),
			Date: note.Date,
			Tag:  note.Tag,
			Text: note.Text,
		})
	}
	return noteList, nil
}

// AddNote обрабатывает gRPC запрос для добавления новой заметки
func (s *noteServer) AddNote(ctx context.Context, note *pb.Note) (*pb.Note, error) {
	newNote := models.Note{
		ID:   int(note.Id),
		Date: note.Date,
		Tag:  note.Tag,
		Text: note.Text,
	}
	err := services.AddNote(newNote)
	if err != nil {
		return nil, err
	}
	return note, nil
}

// GetNoteByText обрабатывает gRPC запрос для получения заметки по тексту
func (s *noteServer) GetNoteByText(ctx context.Context, req *pb.NoteTextRequest) (*pb.Note, error) {
	note, err := services.GetNoteByText(req.Text)
	if err != nil {
		return nil, err
	}
	return &pb.Note{
		Id:   int32(note.ID),
		Date: note.Date,
		Tag:  note.Tag,
		Text: note.Text,
	}, nil
}

// GetNoteByDate обрабатывает gRPC запрос для получения заметки по дате
func (s *noteServer) GetNoteByDate(ctx context.Context, req *pb.NoteDateRequest) (*pb.Note, error) {
	note, err := services.GetNoteByDate(req.Date)
	if err != nil {
		return nil, err
	}
	return &pb.Note{
		Id:   int32(note.ID),
		Date: note.Date,
		Tag:  note.Tag,
		Text: note.Text,
	}, nil
}

// GetNoteByTag обрабатывает gRPC запрос для получения заметок по тегу
func (s *noteServer) GetNoteByTag(ctx context.Context, req *pb.NoteTagRequest) (*pb.NoteList, error) {
	notes, err := services.GetNoteByTag(req.Tag)
	if err != nil {
		return nil, err
	}
	noteList := &pb.NoteList{}
	for _, note := range notes {
		noteList.Notes = append(noteList.Notes, &pb.Note{
			Id:   int32(note.ID),
			Date: note.Date,
			Tag:  note.Tag,
			Text: note.Text,
		})
	}
	return noteList, nil
}

// StartGRPCServer запускает gRPC сервер
func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterNoteServiceServer(grpcServer, &noteServer{}) // Регистрация gRPC сервиса
	log.Println("gRPC server listening on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
