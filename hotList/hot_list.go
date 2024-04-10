package hotList

import "time"

type RankingService interface {
	TopN(ctx context.Context) error // 前一百
}

type BatchRankingService struct {
	svc       ArticleService
	intrSvc   InteractiveService
	batchSize int
	scoreFunc func(likeCnt int64, utime time.Time) float64 // ctime or utime
}

func NewBatchRankingService(svc ArticleService, intrSvc InteractiveService, batchSize int, scoreFunc func(likeCnt int64, utime time.Time) float64) *BatchRankingService {
	return &BatchRankingService{svc: svc, intrSvc: intrSvc, batchSize: batchSize, scoreFunc: scoreFunc}
}

func (b BatchRankingService) TopN(ctx context.Context) error {
	arts, err := b.topN(ctx)
	if err != nil {
		return err
	}
	//TODO implement me
	panic(arts)
	// 最终是放在缓存里面的
}

func (b BatchRankingService) topN(ctx context.Context) ([]domain.Article, error) {
	//TODO implement me
	panic("implement me")
}
