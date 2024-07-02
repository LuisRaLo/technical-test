package application

import (
	"net/http"
	"strconv"
	"technical-challenge/internal/core/domain/constants"
	"technical-challenge/internal/core/domain/models"
	"technical-challenge/internal/core/domain/repositories"
	"technical-challenge/internal/utils/helpers"

	"github.com/google/uuid"

	"go.uber.org/zap"
)

type IBondsUseCase struct {
	Logger                 *zap.SugaredLogger
	bondRepository         repositories.BondRepository
	transactionsRepository repositories.TransactionsRepository
	usersRepository        repositories.UsersRepository
}

func NewBondsUseCase(
	logger *zap.SugaredLogger,
	bondRepository repositories.BondRepository,
	transactionsRepository repositories.TransactionsRepository,
	usersRepository repositories.UsersRepository,
) repositories.BondUseCase {
	logger.Info("BondsUseCase created")
	return &IBondsUseCase{
		Logger:                 logger,
		bondRepository:         bondRepository,
		transactionsRepository: transactionsRepository,
		usersRepository:        usersRepository,
	}
}

// CreateBond implements repositories.BondUseCase.
func (i *IBondsUseCase) CreateBond(payload repositories.CreateBondRequest, userID string) models.DevResponse {

	isExists := i.bondRepository.IsExistBond(payload.Name, payload.Price, payload.Quantity)
	if isExists {
		i.Logger.Error("Bond already exists")
		return models.DevResponse{
			StatusCode: http.StatusConflict,
			Response: models.Response409WithResult{
				Message: "Bond already exists",
			},
		}
	}
	timestamp := helpers.GetTimeNow()
	bondID := uuid.New()

	r := i.bondRepository.CreateBond(repositories.Bond{
		ID:        bondID,
		UserID:    userID,
		Name:      payload.Name,
		Quantity:  payload.Quantity,
		Price:     payload.Price,
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
		DeleteAt:  0,
	})

	if r == 0 {
		i.Logger.Error("Error creating bond. ")

		return models.DevResponse{
			StatusCode: http.StatusInternalServerError,
			Response: models.Response500WithResult{
				Message: "ErrorCreatingBond",
			},
		}
	}

	bond, err := i.bondRepository.GetBondByID(bondID)
	if err != nil {
		i.Logger.Error("Error getting bond by ID. ", err)

		return models.DevResponse{
			StatusCode: http.StatusInternalServerError,
			Response: models.Response500WithResult{
				Message: "ErrorGettingBond",
			},
		}
	}

	response200CreateBond := repositories.CreateBondResponse200{
		Response: &models.Response{
			Message: "Bond created",
		},
		Result: repositories.CreateBondResponse{
			ID: bond.ID,
		},
	}

	return models.DevResponse{
		StatusCode: http.StatusOK,
		Response:   response200CreateBond,
	}

}

// GetAllBonds implements repositories.BondUseCase.
func (i *IBondsUseCase) GetAllBonds(userID string, bondType string) models.DevResponse {

	if bondType != "SOLD" && bondType != "AVAILABLE" && bondType != "BOUGHT" {
		return models.DevResponse{
			StatusCode: http.StatusBadRequest,
			Response: models.Response400WithResult{
				Message: "Invalid bond type",
				Details: []string{"Bond type must be SOLD, AVAILABLE or BOUGHT"},
			},
		}
	}

	if bondType == "AVAILABLE" {
		bonds, err := i.bondRepository.GetAllAvailableBonds(userID)
		if err != nil {
			i.Logger.Error("Error getting all available bonds. ", err)
			return models.DevResponse{
				StatusCode: http.StatusInternalServerError,
				Response: models.Response500WithResult{
					Message: "ErrorGettingAllAvailableBonds",
				},
			}
		}

		var bondModels []repositories.BondModel
		for _, bond := range bonds {
			bondModels = append(bondModels, repositories.BondModel{
				BondID:            bond.BondID,
				Seller:            bond.Seller,
				Name:              bond.Name,
				TotalQuantity:     bond.TotalQuantity,
				Price:             bond.Price,
				AvailableQuantity: bond.AvailableQuantity,
			})
		}

		response200GetAllBonds := repositories.GetAllBondsResponse200{
			Response: &models.Response{
				Message: "Bonds retrieved",
			},
			Result: repositories.GetAllBondsResponse{
				Bonds: bondModels,
			},
		}

		return models.DevResponse{
			StatusCode: http.StatusOK,
			Response:   response200GetAllBonds,
		}
	} else if bondType == "BOUGHT" {
		return i.caseBougth(bondType, userID)
	}

	return i.caseSold(userID)

}

func (i *IBondsUseCase) caseBougth(bondType string, userID string) models.DevResponse {

	bondSoldAndBought := []repositories.BondSoldAndBought{}

	//obtenemos las transacciones por usuario y por status
	transactions, err := i.transactionsRepository.GetTransactionsByUserIDAndStatus(userID, bondType)
	if err != nil {
		i.Logger.Error("Error getting transactions by user ID and status. ", err)

		return models.DevResponse{
			StatusCode: http.StatusInternalServerError,
			Response: models.Response500WithResult{
				Message: constants.INTERNAL_SERVER_ERROR,
			},
		}
	}

	bondIds := make(map[uuid.UUID]bool)
	for _, transaction := range transactions {
		bondIds[transaction.BondID] = true
	}

	//conusltamos los bonos por ID
	bondsData := []repositories.Bond{}
	for bondID := range bondIds {
		bond, err := i.bondRepository.GetBondByID(bondID)
		if err != nil {
			i.Logger.Error("Error getting bond by ID. ", err)

			return models.DevResponse{
				StatusCode: http.StatusInternalServerError,
				Response: models.Response500WithResult{
					Message: constants.INTERNAL_SERVER_ERROR,
				},
			}
		}

		bondsData = append(bondsData, bond)
	}

	usersData := make(map[string]string)
	for _, bondData := range bondsData {
		if _, ok := usersData[bondData.UserID]; !ok {
			user, err := i.usersRepository.GetUserByID(bondData.UserID)
			if err != nil {
				i.Logger.Error("Error getting user by ID. ", err)

				return models.DevResponse{
					StatusCode: http.StatusInternalServerError,
					Response: models.Response500WithResult{
						Message: constants.INTERNAL_SERVER_ERROR,
					},
				}
			}

			usersData[bondData.UserID] = user.Email
		}
	}

	chanel := make(chan repositories.BondSoldAndBought)
	for _, transaction := range transactions {
		go func(transaction repositories.Transactions) {
			bond := repositories.BondSoldAndBought{}
			for _, bondData := range bondsData {
				if transaction.BondID == bondData.ID {
					bond.ID = bondData.ID.String()
					bond.Name = bondData.Name
					bond.Currency = "MXN"
					bond.NumerOfBonds = transaction.Quantity
					totalPrice := (float64(bondData.Price) * float64(transaction.Quantity)) / float64(bondData.Quantity)
					bond.TotalPrice, _ = strconv.ParseFloat(strconv.FormatFloat(totalPrice, 'f', 4, 64), 64)
					bond.SellerOrBuyer = usersData[bondData.UserID]
					chanel <- bond
				}
			}
		}(transaction)
	}

	for range transactions {
		bond := <-chanel
		bondSoldAndBought = append(bondSoldAndBought, bond)
	}

	response200GetAllBondsSoldAndBought := repositories.GetAllBondsSoldAndBoughtResponse200{
		Response: &models.Response{
			Message: "Bonds retrieved",
		},
		Result: repositories.GetAllBondsSoldAndBoughtResponse{
			Bonds: bondSoldAndBought,
		},
	}

	return models.DevResponse{
		StatusCode: http.StatusOK,
		Response:   response200GetAllBondsSoldAndBought,
	}
}

func (i *IBondsUseCase) caseSold(userID string) models.DevResponse {
	bondSoldAndBought := []repositories.BondSoldAndBought{}

	// Obtener las transacciones de ventas por el usuario
	transactions, err := i.transactionsRepository.GetTransactionsByStatusAndNotUserID(repositories.StatusBought, userID)
	if err != nil {
		i.Logger.Error("Error getting transactions by user ID and status. ", err)
		return models.DevResponse{
			StatusCode: http.StatusInternalServerError,
			Response: models.Response500WithResult{
				Message: constants.INTERNAL_SERVER_ERROR,
			},
		}
	}

	// Obtener los bonos del usuario
	myBonds, err := i.bondRepository.GetBondsByUserID(userID)
	if err != nil {
		i.Logger.Error("Error getting bonds by user ID. ", err)
		return models.DevResponse{
			StatusCode: http.StatusInternalServerError,
			Response: models.Response500WithResult{
				Message: constants.INTERNAL_SERVER_ERROR,
			},
		}
	}

	//filtramos de transactions myBonds
	bondIds := make(map[uuid.UUID]bool)
	for _, transaction := range transactions {
		for _, myBond := range myBonds {
			if transaction.BondID == myBond.ID {
				bondIds[transaction.BondID] = true
			}
		}
	}

	//conusltamos los bonos por ID
	bondsData := []repositories.Bond{}
	for bondID := range bondIds {
		bond, err := i.bondRepository.GetBondByID(bondID)
		if err != nil {
			i.Logger.Error("Error getting bond by ID. ", err)
			return models.DevResponse{
				StatusCode: http.StatusInternalServerError,
				Response: models.Response500WithResult{
					Message: constants.INTERNAL_SERVER_ERROR,
				},
			}
		}

		bondsData = append(bondsData, bond)
	}

	usersData := make(map[string]string)
	for _, transaction := range transactions {
		if _, ok := usersData[transaction.UserID]; !ok {
			user, err := i.usersRepository.GetUserByID(transaction.UserID)
			if err != nil {
				i.Logger.Error("Error getting user by ID. ", err)
				return models.DevResponse{
					StatusCode: http.StatusInternalServerError,
					Response: models.Response500WithResult{
						Message: constants.INTERNAL_SERVER_ERROR,
					},
				}
			}

			usersData[transaction.UserID] = user.Email
		}
	}

	chanel := make(chan repositories.BondSoldAndBought)
	for _, transaction := range transactions {
		go func(transaction repositories.Transactions) {
			bond := repositories.BondSoldAndBought{}
			for _, bondData := range bondsData {
				if transaction.BondID == bondData.ID {
					bond.ID = bondData.ID.String()
					bond.Name = bondData.Name
					bond.Currency = "MXN"
					bond.NumerOfBonds = transaction.Quantity
					totalPrice := (float64(bondData.Price) * float64(transaction.Quantity)) / float64(bondData.Quantity)
					bond.TotalPrice, _ = strconv.ParseFloat(strconv.FormatFloat(totalPrice, 'f', 4, 64), 64)
					bond.SellerOrBuyer = usersData[transaction.UserID]
					chanel <- bond
				}
			}
		}(transaction)
	}

	for range transactions {
		bond := <-chanel
		bondSoldAndBought = append(bondSoldAndBought, bond)
	}

	// Preparar la respuesta final
	response200GetAllBondsSoldAndBought := repositories.GetAllBondsSoldAndBoughtResponse200{
		Response: &models.Response{
			Message: "Bonds retrieved",
		},
		Result: repositories.GetAllBondsSoldAndBoughtResponse{
			Bonds: bondSoldAndBought,
		},
	}

	return models.DevResponse{
		StatusCode: http.StatusOK,
		Response:   response200GetAllBondsSoldAndBought,
	}
}

// SellBond implements repositories.BondUseCase.
func (i *IBondsUseCase) SellBond(payload repositories.SellBondRequest, userID string) models.DevResponse {
	var sellBondResponse repositories.SellBondResponse200 = repositories.SellBondResponse200{}

	bond, err := i.bondRepository.GetBondByIDAndQuantity(payload.BondID, payload.Quantity)
	if err != nil {
		i.Logger.Error("Error getting bond by ID and quantity. ", err)

		if err.Error() == "sql: no rows in result set" {
			i.Logger.Error("Bond not found or quantity not enough")

			return models.DevResponse{
				StatusCode: http.StatusNotFound,
				Response: models.Response404WithResult{
					Message: constants.REQUEST_INVALID,
					Details: []string{"Bond not found or quantity not enough"},
				},
			}
		}

		return models.DevResponse{
			StatusCode: http.StatusInternalServerError,
			Response: models.Response500WithResult{
				Message: constants.INTERNAL_SERVER_ERROR,
			},
		}
	}

	timestamp := helpers.GetTimeNow()

	createTransaction, err := i.transactionsRepository.CreateTransaction(repositories.Transactions{
		ID:        uuid.New(),
		BondID:    bond.BondID,
		UserID:    userID,
		Quantity:  payload.Quantity,
		Price:     bond.Price,
		Status:    repositories.StatusBought,
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
		DeletedAt: 0,
	})
	if err != nil {
		i.Logger.Error("Error creating transaction. ", err)

		return models.DevResponse{
			StatusCode: http.StatusInternalServerError,
			Response: models.Response500WithResult{
				Message: constants.INTERNAL_SERVER_ERROR,
			},
		}
	}

	i.Logger.Info("Transaction created: ", createTransaction)

	// si la cantidad de bonos es igual a la cantidad vendida
	if bond.TotalQuantity == payload.Quantity {

		r, err := i.bondRepository.DeleteBond(payload.BondID)
		if err != nil {
			i.Logger.Error("Error deleting bond. ", err)

			return models.DevResponse{
				StatusCode: http.StatusInternalServerError,
				Response: models.Response500WithResult{
					Message: constants.INTERNAL_SERVER_ERROR,
				},
			}
		}

		i.Logger.Info("Bond deleted: ", r)

	}
	//TODO: realizar la venta
	//TODO: actualizar la cantidad del bono

	//TODO: actualizar el estado del bono
	//TODO: guardar la transaccion
	//TODO: guardar la venta

	//creamos respuesta exitosa
	sellBondResponse = repositories.SellBondResponse200{
		Response: &models.Response{
			Message: "Bond sold",
		}, Result: repositories.SellBondResponse{
			ID: payload.BondID,
		},
	}

	return models.DevResponse{
		StatusCode: http.StatusOK,
		Response:   sellBondResponse,
	}
}
