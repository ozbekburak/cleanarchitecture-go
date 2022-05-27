package interactor

import (
	"github.com/ozbekburak/cleanarch-mongo-inmem/domain"
	"github.com/ozbekburak/cleanarch-mongo-inmem/usecase/repository"
)

// RecordInteractor is the struct that implements the RecordRepository interface
type RecordInteractor struct {
	RecordRepository repository.RecordRepository
}

// Filter includes the logic to filter records
func (ri *RecordInteractor) Filter(f domain.Filter) (domain.Records, error) {
	start, err := f.FormatDate(f.StartDate)
	if err != nil {
		return domain.Records{}, err
	}

	end, err := f.FormatDate(f.EndDate)
	if err != nil {
		return domain.Records{}, err
	}

	// start date must be before end date
	if start.After(end) {
		return domain.Records{}, domain.ErrStartDateLessThanEnd
	}

	// min count must be less than max count
	if f.MinCount > f.MaxCount {
		return domain.Records{}, domain.ErrMaxCountLessThanMin
	}

	records, err := ri.RecordRepository.Filter(f)
	return records, err
}
