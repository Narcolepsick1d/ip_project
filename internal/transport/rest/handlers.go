package rest

import (
	"awesomeProject2/internal/models"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type User interface {
	SignUp(inp models.SignUpInput) error
	SignIn(inp models.SignInInput) (string, string, error)
	ChooseRole(userID int, role string) error
	UpdateUserInfo(userId int, user models.UserUpdate) error
	ParseToken(accessToken string) (int64, error)
	RefreshTokens(refreshToken string) (string, string, error)
}
type Agent interface {
	Create(agent models.Agent) error
}
type Client interface {
	Create(client_user_id int64) error
}
type Owner interface {
	Create(client_user_id int64) error
}
type Handler struct {
	usersService  User
	clientService Client
	agentService  Agent
	ownerService  Owner
}

func NewHandler(users User, client Client, agent Agent, owner Owner) *Handler {
	return &Handler{
		clientService: client,
		usersService:  users,
		agentService:  agent,
		ownerService:  owner,
	}
}
func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()

	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/sign-up", h.signUp).Methods(http.MethodPost)
		auth.HandleFunc("/sign-in", h.signIn).Methods(http.MethodGet)
		auth.HandleFunc("/refresh", h.refresh).Methods(http.MethodGet)
	}
	user := r.PathPrefix("/user").Subrouter()
	{
		user.Use(h.authMiddleware)
		user.HandleFunc("/choose_role", h.chooseRole).Methods(http.MethodPost)
		user.HandleFunc("/update_user_info", h.updateUser).Methods(http.MethodPut)
	}
	return r
}

func getIdFromRequest(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, errors.New("id can't be 0")
	}

	return id, nil
}
