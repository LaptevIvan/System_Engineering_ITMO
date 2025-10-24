package library

import (
	"context"

	"github.com/project/library/pkg/logger"

	"github.com/project/library/generated/api/library"
	"github.com/project/library/internal/entity"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertBook(book *entity.Book) *library.Book {
	return &library.Book{
		Id:        book.ID,
		Name:      book.Name,
		AuthorId:  book.AuthorIDs,
		CreatedAt: timestamppb.New(book.CreatedAt),
		UpdatedAt: timestamppb.New(book.UpdatedAt),
	}
}

func (l *libraryImpl) AddBook(ctx context.Context, name string, authorIDs []string) (*library.AddBookResponse, error) {
	book, err := l.booksRepository.AddBook(ctx, entity.Book{
		Name:      name,
		AuthorIDs: authorIDs,
	})

	if logger.CheckError(err, l.logger, "Failed adding book", zap.Error(err)) {
		return nil, err
	}
	if l.logger != nil {
		l.logger.Info("Added book", zap.String("id", book.ID))
	}

	return &library.AddBookResponse{
		Book: convertBook(&book),
	}, nil
}

func (l *libraryImpl) GetBookInfo(ctx context.Context, bookID string) (*library.GetBookInfoResponse, error) {
	book, err := l.booksRepository.GetBook(ctx, bookID)

	if logger.CheckError(err, l.logger, "Failed get book info", zap.String("id of book", bookID), zap.Error(err)) {
		return nil, err
	}
	if l.logger != nil {
		l.logger.Info("Get the book", zap.String("id of book", bookID))
	}

	return &library.GetBookInfoResponse{
		Book: convertBook(&book),
	}, nil
}

func (l *libraryImpl) UpdateBook(ctx context.Context, id, newName string, newAuthorIDs []string) error {
	err := l.booksRepository.UpdateBook(ctx, entity.Book{
		ID:        id,
		Name:      newName,
		AuthorIDs: newAuthorIDs,
	})

	if !logger.CheckError(err, l.logger, "Failed update book", zap.Error(err)) {
		if l.logger != nil {
			l.logger.Info("Updated the book with id", zap.String("id of book", id))
		}
	}

	return err
}

func (l *libraryImpl) GetAuthorBooks(ctx context.Context, idAuthor string) (<-chan *library.Book, error) {
	books, err := l.booksRepository.GetAuthorBooks(ctx, idAuthor)

	if logger.CheckError(err, l.logger, "Failed get author books", zap.Error(err)) {
		return nil, err
	}
	if l.logger != nil {
		l.logger.Info("Got the author's book", zap.String("author's id", idAuthor))
	}

	ans := make(chan *library.Book)
	go func() {
		defer close(ans)
		for b := range books {
			ans <- &library.Book{
				Id:        b.ID,
				Name:      b.Name,
				AuthorId:  b.AuthorIDs,
				CreatedAt: timestamppb.New(b.CreatedAt),
				UpdatedAt: timestamppb.New(b.UpdatedAt),
			}
		}
	}()

	return ans, err
}
