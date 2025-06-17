package service

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/sirupsen/logrus"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model/constants"
	"go-data-repository/src/main/go/com/mccusa/datarepository/repository"
	"sync"
	"time"
)

type DocusignService interface {
	// ScheduleDocusignJob schedules a job to run every 30 minutes.
	ProcessDocusignEnvelopes() error
}

type docusignServiceImpl struct {
	docusignEnvelopesRepository repository.DocusignEnvelopesRepository
}

func NewDocusignService(docusignEnvelopesRepository repository.DocusignEnvelopesRepository) DocusignService {
	return &docusignServiceImpl{docusignEnvelopesRepository: docusignEnvelopesRepository}
}

func (d docusignServiceImpl) ProcessDocusignEnvelopes() error {
	start := time.Now()
	logrus.Println("Processing Docusign envelopes...")
	//get the docusign envelopes from the database
	envelopes, err := d.docusignEnvelopesRepository.FindAllByEnvelopeStatusAndClientEmailIsNotNull(constants.EnvelopeStatusSent)
	if err != nil {
		return fmt.Errorf("fetching envelopes: %w", err)
	}

	var (
		wg          sync.WaitGroup
		mu          sync.Mutex // protects 'combinedErr'
		combinedErr error      // collects all errs
	)
	wg.Add(len(envelopes))
	//process the envelopes
	/*
		for index, envelope := range envelopes {
			handleSingleEnvelope(context.Background(), envelope, index)
		}
	*/

	// 3) Launch one goroutine per envelope
	for index, envelope := range envelopes {
		envelope := envelope // capture loop variable
		go func() {
			defer wg.Done()
			if err := handleSingleEnvelope(context.Background(), envelope, index); err != nil {
				mu.Lock()
				combinedErr = multierror.Append(combinedErr, fmt.Errorf("envelope %s: %w", envelope.EnvelopeID, err))
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	//*/

	elapsed := time.Since(start)
	logrus.Infof("Finished processing Docusign envelopes in %f seconds.", elapsed.Seconds())
	return combinedErr
}

func handleSingleEnvelope(ctx context.Context, envelope model.DocusignEnvelope, index int) error {
	logrus.Printf("Processing envelope %d with ID: %s", index, envelope.EnvelopeID)
	return nil
}
