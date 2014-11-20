package auctionrunner

import (
	"os"
	"time"

	"github.com/cloudfoundry/gunk/timeprovider"

	"github.com/cloudfoundry-incubator/auction/auctiontypes"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
	"github.com/cloudfoundry/gunk/workpool"
)

type AuctionRunner interface {
	AddLRPStartAuction(models.LRPStartAuction)
	AddLRPStopAuction(models.LRPStopAuction)
}

type AuctionRunnerDelegate interface {
	FetchAuctionRepClients() (map[string]auctiontypes.AuctionRep, error)
	DistributedBatch(WorkResults)
}

type WorkResults struct {
	SuccessfulStarts []auctiontypes.StartAuction
	SuccessfulStops  []auctiontypes.StopAuction
	FailedStarts     []auctiontypes.StartAuction
	FailedStops      []auctiontypes.StopAuction
}

type auctionRunner struct {
	delegate     AuctionRunnerDelegate
	batch        *Batch
	timeProvider timeprovider.TimeProvider
	workPool     *workpool.WorkPool
	maxRetries   int
}

func New(delegate AuctionRunnerDelegate, timeProvider timeprovider.TimeProvider, maxRetries int, workPool *workpool.WorkPool) *auctionRunner {
	return &auctionRunner{
		delegate:     delegate,
		batch:        NewBatch(timeProvider),
		timeProvider: timeProvider,
		workPool:     workPool,
		maxRetries:   maxRetries,
	}
}

func (a *auctionRunner) Run(signals <-chan os.Signal, ready chan<- struct{}) error {
	close(ready)

	var hasWork chan struct{}
	hasWork = a.batch.HasWork

	for {
		select {
		case <-hasWork:
			clients, err := a.delegate.FetchAuctionRepClients()
			if err != nil {
				time.Sleep(time.Second)
				hasWork = make(chan struct{}, 1)
				hasWork <- struct{}{}
				break
			}
			hasWork = a.batch.HasWork
			cells := FetchStateAndBuildCells(a.workPool, clients)
			startAuctions, stopAuctions := a.batch.DedupeAndDrain()
			if len(startAuctions) == 0 && len(stopAuctions) == 0 {
				break
			}
			workResults := DistributeWork(a.workPool, cells, a.timeProvider, startAuctions, stopAuctions)
			workResults = ResubmitFailedWork(a.batch, workResults, a.maxRetries)

			go a.delegate.DistributedBatch(workResults)
		case <-signals:
			return nil
		}
	}
}

func (a *auctionRunner) AddLRPStartAuction(start models.LRPStartAuction) {
	a.batch.AddLRPStartAuction(start)
}

func (a *auctionRunner) AddLRPStopAuction(stop models.LRPStopAuction) {
	a.batch.AddLRPStopAuction(stop)
}
