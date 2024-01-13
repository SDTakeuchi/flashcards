package spreadsheet

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var (
	ErrSpreadsheetInvalidFormat = errors.New("spreadsheet format is invalid")
)

type SpreadsheetService interface {
	Read(spreadsheetID string) ([]*Value, error)
	Update(spreadsheetID string, row uint, values ...string) error
}

type spreadsheetService struct {
	tabName         string
	cellTopLeft     string
	cellBottomRight string
	service         *sheets.Service
}

var _ SpreadsheetService = (*spreadsheetService)(nil)

func InitService(ctx context.Context, pathCredential, tabName string) (*spreadsheetService, error) {
	credential := option.WithCredentialsFile(pathCredential)

	srv, err := sheets.NewService(ctx, credential)
	if err != nil {
		log.Fatal(err)
	}

	return &spreadsheetService{
		tabName:         tabName,
		cellTopLeft:     "A1",
		cellBottomRight: "H",
		service:         srv,
	}, nil
}

func (s spreadsheetService) getReadRange() string {
	return fmt.Sprintf("%s!%s:%s", s.tabName, s.cellTopLeft, s.cellBottomRight)
}

func (s *spreadsheetService) getSheet(spreadsheetID string) (*sheets.ValueRange, error) {
	resp, err := s.service.Spreadsheets.Values.Get(
		spreadsheetID,
		s.getReadRange(),
	).Do()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrSpreadsheetInvalidFormat, err)
	}
	if len(resp.Values) == 0 {
		return nil, fmt.Errorf("%w: %s", ErrSpreadsheetInvalidFormat, "data not found")
	}
	return resp, nil
}

type Value struct {
	Word         string
	Description  string
	LastSeen     time.Time
	Status       int
	PartOfSpeech string
	Example      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (s *spreadsheetService) Read(spreadsheetID string) ([]*Value, error) {
	sheet, err := s.getSheet(spreadsheetID)
	if err != nil {
		return nil, err
	}

	values := make([]*Value, len(sheet.Values)-1)
	for i, row := range sheet.Values {
		if i == 0 || len(row) < 2 {
			// skip header
			// or skip row that does not have both word and description at least
			continue
		}

		word := strings.Trim(row[0].(string), " ")

		description := strings.Trim(row[1].(string), " ")

		var partOfSpeech string
		if len(row) >= 3 {
			partOfSpeech = row[2].(string)
		}

		var example string
		if len(row) >= 4 {
			example = row[3].(string)
		}

		var lastSeen time.Time
		if len(row) >= 5 {
			lastSeen, _ = time.Parse(time.RFC3339, row[4].(string))
		}

		var status int
		if len(row) >= 6 {
			statusRaw, _ := row[5].(string)
			status, _ = strconv.Atoi(statusRaw)
		}

		var createdAt time.Time
		if len(row) >= 7 {
			createdAt, _ = time.Parse(time.RFC3339, row[6].(string))
		}

		var updatedAt time.Time
		if len(row) >= 8 {
			updatedAt, _ = time.Parse(time.RFC3339, row[7].(string))
		}

		values[i-1] = &Value{ // i-1 because skip header
			Word:         word,
			Description:  description,
			LastSeen:     lastSeen,
			Status:       status,
			PartOfSpeech: partOfSpeech,
			Example:      example,
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
		}
	}
	return values, nil
}

func (s *spreadsheetService) Update(spreadsheetID string, row uint, values ...string) error {
	innerData := make([]interface{}, len(values))
	for i, v := range values {
		innerData[i] = v
	}

	_, err := s.service.Spreadsheets.Values.Update(
		spreadsheetID,
		fmt.Sprintf("%s!A%d", s.tabName, row),
		&sheets.ValueRange{
			Values: [][]interface{}{
				innerData,
			},
		}).
		ValueInputOption("USER_ENTERED").
		Do()
	if err != nil {
		return err
	}
	return nil
}
