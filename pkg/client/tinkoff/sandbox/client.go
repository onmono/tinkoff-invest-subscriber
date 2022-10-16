package tinkoff

import (
	"context"
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/onmono/clean-architecture/internal/config/tinkoff"
	"github.com/onmono/clean-architecture/internal/domain/entity/tinkoff/sandbox"
	"log"
	"time"
)

type SandboxClient struct {
	client    *sdk.SandboxRestClient
	account   sdk.Account
	portfolio sdk.Portfolio
	accounts  []sdk.Account
	// TODO: inject logger
}

func NewClient(cfg config.TinkoffInvestConfig) *SandboxClient {
	client := sdk.NewSandboxRestClient(cfg.Token())
	return &SandboxClient{client: client}
}

func (s *SandboxClient) RegisterAccount(c context.Context) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	log.Println("Регистрация обычного счета в песочнице")
	var err error
	s.account, err = s.client.Register(ctx, sdk.AccountTinkoff)
	if err != nil {
		log.Fatalln(errorHandle(err))
	}
	log.Printf("%+v\n", s.account)
}

func (s *SandboxClient) Dial(ctx context.Context, balance sandbox.Balance) {
	s.RegisterAccount(ctx)
	s.DefaultCurrencyBalance(ctx, balance)
	s.GetPortfolioList(ctx)
	s.GetAllAccounts(ctx)
	s.SetPositionBalance(ctx, "TSLA")
}

func (s *SandboxClient) GetInstrument(c context.Context, ticker string) ([]sdk.Instrument, error) {
	return s.client.InstrumentByTicker(c, ticker)
}

func (s *SandboxClient) DefaultCurrencyBalance(c context.Context, balance sandbox.Balance) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	log.Println("Рисуем себе 100500 рублей в портфеле песочницы")
	var err error

	err = s.client.SetCurrencyBalance(ctx, s.account.ID, sdk.Currency(balance.Ticker), balance.Balance)
	if err != nil {
		log.Fatalln(err)
	}
}

func (s *SandboxClient) GetPortfolioList(c context.Context) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	log.Println("Получение списка валютных и НЕ валютных активов портфеля для счета по-умолчанию")
	// Метод является совмещением PositionsPortfolio и CurrenciesPortfolio
	var err error
	s.portfolio, err = s.client.Portfolio(ctx, sdk.DefaultAccount)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", s.portfolio)
}

func (s *SandboxClient) GetAllAccounts(c context.Context) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	log.Println("Получение всех брокерских счетов")
	var err error
	s.accounts, err = s.client.Accounts(ctx)
	if err != nil {
		log.Fatalln(errorHandle(err))
	}
	log.Printf("%+v\n", s.accounts)
}

func (s *SandboxClient) SetPositionBalance(c context.Context, ticker string) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	log.Println("Рисуем себе 100 акций TSLA в портфеле песочницы")
	instruments, err := s.GetInstrument(ctx, ticker)

	err = s.client.SetPositionsBalance(ctx, s.account.ID, instruments[0].FIGI, 100)
	if err != nil {
		log.Fatalln(err)
	}
}

func errorHandle(err error) error {
	if err == nil {
		return nil
	}

	if tradingErr, ok := err.(sdk.TradingError); ok {
		if tradingErr.InvalidTokenSpace() {
			tradingErr.Hint = "Do you use sandbox token in production environment or vise verse?"
			return tradingErr
		}
	}

	return err
}
