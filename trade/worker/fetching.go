package worker

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stryukovsky/go-backend-learn/trade"
	"github.com/stryukovsky/go-backend-learn/trade/protocols"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type FetchEnvironment struct {
	chainId           string
	db                *gorm.DB
	trackedWallets    []trade.TrackedWallet
	participants      []string
	erc20Handlers     []protocols.DeFiProtocolHandler[trade.ERC20Transfer, trade.Deal]
	aaveHandlers      []protocols.DeFiProtocolHandler[trade.AaveEvent, trade.AaveInteraction]
	compoundHandlers  []protocols.DeFiProtocolHandler[trade.Compound3Event, trade.Compound3Interaction]
	uniswapv3Handlers []protocols.DeFiProtocolHandler[trade.UniswapV3Event, trade.UniswapV3Deal]
}

func NewFetchEnvironment(
	chainId string,
	db *gorm.DB,
	wallets []trade.TrackedWallet,
	participants []string,
	erc20Handlers []protocols.DeFiProtocolHandler[trade.ERC20Transfer, trade.Deal],
	aaveHandlers []protocols.DeFiProtocolHandler[trade.AaveEvent, trade.AaveInteraction],
	compoundHandlers []protocols.DeFiProtocolHandler[trade.Compound3Event, trade.Compound3Interaction],
	uniswapv3Handlers []protocols.DeFiProtocolHandler[trade.UniswapV3Event, trade.UniswapV3Deal],
) *FetchEnvironment {
	return &FetchEnvironment{
		chainId,
		db,
		wallets,
		participants,
		erc20Handlers,
		aaveHandlers,
		compoundHandlers,
		uniswapv3Handlers,
	}
}

func fetchInteractionsFromEthJSONRPC[BlockchainInteractions any, FinancialInteractions any](
	chainId string,
	startBlock uint64,
	endBlock uint64,
	handlers []protocols.DeFiProtocolHandler[BlockchainInteractions, FinancialInteractions],
	participants []string,
) ([]FinancialInteractions, error) {
	resultsBlockchain := make([]BlockchainInteractions, 0)
	resultsFinancial := make([]FinancialInteractions, 0)
	chBlockchain := make(chan BlockchainInteractions)
	chFinancial := make(chan FinancialInteractions)
	eg, ctx := errgroup.WithContext(context.Background())
	for _, handler := range handlers {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				slog.Warn("Cancelled fetchInteractionsFromEthJSONRPC")
				return ctx.Err()
			default:

				blockchainInteractions, err := handler.FetchBlockchainInteractions(
					chainId,
					participants,
					startBlock,
					endBlock,
				)
				if err != nil {
					slog.Warn(fmt.Sprintf("[%s] Cannot fetch blockchain interactions: %s", handler.Name(), err.Error()))
					return err
				}
				if len(blockchainInteractions) == 0 {
					slog.Info(fmt.Sprintf("[%s] No blockchain interactions found", handler.Name()))
					return nil
				}
				for _, interaction := range blockchainInteractions {
					chBlockchain <- interaction
				}

				slog.Info(fmt.Sprintf(
					"[%s] Found %d blockchain interactions where tracked wallets participated",
					handler.Name(),
					len(blockchainInteractions)))
				financialInteractions, err := handler.PopulateWithFinanceInfo(blockchainInteractions)
				if err != nil {
					slog.Warn(fmt.Sprintf("[%s] Cannot fetch financial interactions: %s", handler.Name(), err.Error()))
					return err
				}
				slog.Info(fmt.Sprintf("[%s] Fetched %d financial interactions for corresponding %d blockchain interactions", handler.Name(), len(financialInteractions), len(blockchainInteractions)))

				for _, interaction := range financialInteractions {
					chFinancial <- interaction
				}
				return nil
			}
		})
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for value := range chBlockchain {
			resultsBlockchain = append(resultsBlockchain, value)
		}
	}()
	go func() {
		defer wg.Done()
		for value := range chFinancial {
			resultsFinancial = append(resultsFinancial, value)
		}
	}()
	err := eg.Wait()
	close(chBlockchain)
	close(chFinancial)
	wg.Wait()
	if err != nil {
		return nil, err
	}

	return resultsFinancial, nil
}

func saveInteractions[T any](db *gorm.DB, interactions []T, handlerName string) {
	for _, item := range interactions {
		err := db.Create(&item).Error
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				slog.Warn(fmt.Sprintf("[%s] Error on postgres, possibly duplicate: %v", handlerName, err))
				// Duplicate key – ignore
				continue
			}
			slog.Warn(fmt.Sprintf("[%s] Cannot save interaction: %v", handlerName, err))
		}
	}
}

func (f *FetchEnvironment) Fetch(startBlock, endBlock uint64) {
	type task struct {
		name     string
		handlers any
		run      func() error
	}

	tasks := []task{
		{
			name:     "ERC20",
			handlers: f.erc20Handlers,
			run: func() error {
				financial, err := fetchInteractionsFromEthJSONRPC(
					f.chainId, startBlock, endBlock, f.erc20Handlers, f.participants)
				if err != nil {
					return err
				}
				saveInteractions(f.db, financial, "ERC20")
				return nil
			},
		},
		{
			name:     "Aave",
			handlers: f.aaveHandlers,
			run: func() error {
				financial, err := fetchInteractionsFromEthJSONRPC(
					f.chainId, startBlock, endBlock, f.aaveHandlers, f.participants)
				if err != nil {
					return err
				}
				saveInteractions(f.db, financial, "Aave")
				return nil
			},
		},
		{
			name:     "Compound3",
			handlers: f.compoundHandlers,
			run: func() error {
				financial, err := fetchInteractionsFromEthJSONRPC(
					f.chainId, startBlock, endBlock, f.compoundHandlers, f.participants)
				if err != nil {
					return err
				}
				saveInteractions(f.db, financial, "Compound3")
				return nil
			},
		},
		{
			name:     "UniswapV3",
			handlers: f.uniswapv3Handlers,
			run: func() error {
				financial, err := fetchInteractionsFromEthJSONRPC(
					f.chainId, startBlock, endBlock, f.uniswapv3Handlers, f.participants)
				if err != nil {
					return err
				}
				saveInteractions(f.db, financial, "UniswapV3")
				return nil
			},
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)
	for _, t := range tasks {
		g.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			if err := t.run(); err != nil {
				slog.Warn(fmt.Sprintf("Cannot fetch interactions at task %s: %v", t.name, err))
				return err
			}
			return nil
		})
	}

	if err := g.Wait(); err == nil {
		for i := range f.trackedWallets {
			f.trackedWallets[i].LastBlock = endBlock
		}
		slog.Info(fmt.Sprintf("Successfully fetched blockchain events so mark wallets as indexed on block %d", endBlock))
		f.db.Save(f.trackedWallets)
	}
}
