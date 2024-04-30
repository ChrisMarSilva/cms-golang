package services

// import (
// 	"errors"
// 	"net/http"
// 	"strings"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/log"
// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/google/uuid"
// 	"golang.org/x/crypto/bcrypt"
// )

// type AuthController interface {
// 	SignUp(ctx *fiber.Ctx) error
// 	SignIn(ctx *fiber.Ctx) error
// 	GetUser(ctx *fiber.Ctx) error
// 	GetUsers(ctx *fiber.Ctx) error
// 	PutUser(ctx *fiber.Ctx) error
// 	DeleteUser(ctx *fiber.Ctx) error
// }

// type authController struct {
// 	usersRepo repository.UsersRepository
// }

// func NewAuthController(usersRepo repository.UsersRepository) AuthController {
// 	return &authController{usersRepo}
// }

// func (c *authController) SignUp(ctx *fiber.Ctx) error {

// type UserService struct {
// 	userRepo UserRepository
// }

// func NewUserService(userRepo UserRepository) *UserService {
// 	return &UserService{userRepo: userRepo}
// }

// /*

//   func Home(c *gin.Context) {

//    cookie, err := c.Cookie("token")
//       if err != nil {
//           c.JSON(401, gin.H{"error": "unauthorized"})
//           return
//       }

//       claims, err := utils.ParseToken(cookie)
//       if err != nil {
//           c.JSON(401, gin.H{"error": "unauthorized"})
//           return
//       }

//       if claims.Role != "user" && claims.Role != "admin" {
//           c.JSON(401, gin.H{"error": "unauthorized"})
//           return
//       }

//       c.JSON(200, gin.H{"success": "home page", "role": claims.Role})

//   }
// 	    func Premium(c *gin.Context) {
//       cookie, err := c.Cookie("token")
//       if err != nil {
//           c.JSON(401, gin.H{"error": "unauthorized"})
//           return
//       }

//       claims, err := utils.ParseToken(cookie)
//      if err != nil {
//           c.JSON(401, gin.H{"error": "unauthorized"})
//           return
//       }

//       if claims.Role != "admin" {
//           c.JSON(401, gin.H{"error": "unauthorized"})
//           return
//       }

//       c.JSON(200, gin.H{"success": "premium page", "role": claims.Role})
//   }

// */

// func (h UserService) Register(c *fiber.Ctx, payload UserRegisterRequest) (UserRegisterResponse, error) {
// 	// Validate user input (username, email, password)
// 	// Hash the password
// 	// Store user data in the database
// 	// Return a success message or error response

// 	// errors := ValidateStruct(payload)
// 	// if errors != nil {
// 	// 	return UserRegisterResponse{}, err // return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
// 	// }

// 	// if payload.Password != payload.PasswordConfirm {
// 	// 	return UserRegisterResponse{}, err // return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Passwords do not match"})
// 	// }

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return UserRegisterResponse{}, err // return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
// 	}

// 	if err != nil {
// 		switch {
// 		case errors.Is(err, ErrUsernameNotFound):
// 			return c.String(http.StatusNotFound, WarnWhenUsernameNotFound)
// 		case errors.Is(err, ErrUsernameOrPasswordInvalid):
// 			return c.String(http.StatusUnauthorized, WarnNonValidCredentials)
// 		default:
// 			return c.String(http.StatusInternalServerError, WarnInternalServerError)
// 		}
// 	}

// 	user := UserModel{
// 		ID:        uuid.New(),
// 		Nome:      payload.Nome,
// 		Email:     strings.ToLower(payload.Email),
// 		Password:  string(hashedPassword),
// 		IsActive:  true,
// 		CreatedAt: time.Now(),
// 	}

// 	response := UserRegisterResponse{
// 		Nome:  user.Nome,
// 		Email: user.Email,
// 	}

// 	models.DB.Where("email = ?", user.Email).First(&existingUser)
// 	if existingUser.ID != 0 {
// 		c.JSON(400, gin.H{"error": "user already exists"})
// 		return
// 	}

// 	var errHash error
// 	user.Password, errHash = utils.GenerateHashPassword(user.Password)
// 	if errHash != nil {
// 		c.JSON(500, gin.H{"error": "could not generate password hash"})
// 		return
// 	}
// 	// result := initializers.DB.Create(&newUser)

// 	// if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
// 	// 	return err // return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "User with that email already exists"})
// 	// } else if result.Error != nil {
// 	// 	return err // return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
// 	// }

// 	return response, nil
// }

// func (h UserService) Login(c *fiber.Ctx, payload UserLoginRequest) (string, error) {
// 	// errors := ValidateStruct(payload)
// 	// if errors != nil {
// 	//  log.Error("Erro no repository:", err.Error())
// 	// 	return "", err
// 	// }

// 	user, err := h.userRepo.GetByEmail(c.Context(), strings.ToLower(payload.Email))
// 	if err != nil {
// 		log.Error("Erro no repository:", err.Error())
// 		return "", err
// 	}

// 	// if user.Password != payload.Password {
// 	// 	log.Error("Erro no Password: Senha inválida.")
// 	// 	return "", errors.New("Senha inválida.")
// 	// }

// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
// 	if err != nil {
// 		log.Error("Erro no repository:", err.Error())
// 		return "", err // return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid email or Password"})
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"sub":       user.ID,
// 		"nome":      user.Nome,
// 		"email":     user.Email,
// 		"is_active": user.IsActive,
// 		"exp":       time.Now().UTC().Add(time.Hour * 24 * 7).Unix(),
// 		"iat":       time.Now().UTC().Unix(),
// 		"nbf":       time.Now().UTC().Unix(),
// 	})

// 	// now := time.Now().UTC()
// 	// tokenByte := jwt.New(jwt.SigningMethodHS256)
// 	// claims := tokenByte.Claims.(jwt.MapClaims)
// 	// claims["sub"] = user.ID
// 	// claims["exp"] = now.Add(time.Hour * 24 * 7).Unix()
// 	// claims["iat"] = now.Unix()
// 	// claims["nbf"] = now.Unix()

// 	tokenStr, err := token.SignedString([]byte(SecretKey))
// 	if err != nil {
// 		log.Error("Erro no SignedString: ", err.Error())
// 		return "", err // c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
// 	}

// 	sess, err := store.Get(c)
// 	if err != nil {
// 		log.Error("Erro no store: ", err.Error())
// 		return "", err // c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
// 	}

// 	sess.Set("jwt", tokenStr)
// 	if err := sess.Save(); err != nil {
// 		log.Error("Erro no token: ", err.Error())
// 		return "", err // c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
// 	}

// 	sess.SetExpiry(time.Hour * 24)

// 	c.Cookie(&fiber.Cookie{
// 		Name:     "token",
// 		Value:    tokenStr,
// 		Path:     "/",
// 		MaxAge:   60 * 60 * 24 * 7, // config.JwtMaxAge * 60,
// 		Secure:   false,
// 		HTTPOnly: true,
// 		Domain:   "localhost",
// 	})

// 	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)
// 	if !errHash {
// 		c.JSON(400, gin.H{"error": "invalid password"})
// 		return
// 	}
// 	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)

// 	return tokenStr, err
// }

// func (h UserService) Logout(c *fiber.Ctx) error {
// 	sess, err := store.Get(c)
// 	if err != nil {
// 		return err // c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
// 	}

// 	sess.Destroy()

// 	if err := sess.Save(); err != nil {
// 		return err // c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
// 	}

// 	c.SetCookie("token", "", -1, "/", "localhost", false, true)

// 	c.Cookie(&fiber.Cookie{
// 		Name:    "token",
// 		Value:   "",
// 		Expires: time.Now().Add(-time.Hour * 24),
// 	})

// 	return nil // c.SendStatus(fiber.StatusOK)
// }

// func (h *handler) Logout(c echo.Context) error {
// 	cookie := new(http.Cookie)
// 	cookie.Name = "token"
// 	cookie.Value = ""
// 	cookie.MaxAge = 0
// 	c.SetCookie(cookie)
// 	return c.String(http.StatusOK, "You have successfully logout")
// }

// func (h UserService) Refresh(c *fiber.Ctx) (jwt.MapClaims, error) {
// 	sess, err := store.Get(c)
// 	if err != nil {
// 		return nil, err // c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
// 	}

// 	tokenStr := sess.Get("jwt")
// 	if tokenStr == nil {
// 		return nil, err // c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "No token found"})
// 	}

// 	token, err := jwt.Parse(tokenStr.(string), func(token *jwt.Token) (interface{}, error) {
// 		return []byte(SecretKey), nil
// 	})

// 	if err != nil {
// 		return nil, err // c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || !token.Valid {
// 		return nil, errors.New("Invalid token.") // c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
// 	}

// 	//sess.SetExpiry(time.Second * 2)
// 	// if err := sess.Save(); err != nil {
// 	//     return nil, err // c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
// 	// }

// 	return claims, nil
// }

// func (h UserService) Verify(c *fiber.Ctx) error {
// 	sess, err := store.Get(c)
// 	if err != nil {
// 		return err // c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
// 	}

// 	tokenStr := sess.Get("jwt")
// 	if tokenStr == nil {
// 		return err // c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "No token found"})
// 	}

// 	token, err := jwt.Parse(tokenStr.(string), func(token *jwt.Token) (interface{}, error) {
// 		return []byte(SecretKey), nil
// 	})

// 	if err != nil || !token.Valid {
// 		return err // c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
// 	}

// 	return nil
// }

// type AuthService struct {
// 	repo   auth.AuthRepository
// 	config *config.Config
// }

// func NewAuthService(repo auth.AuthRepository, config *config.Config) *AuthService {
// 	return &AuthService{repo, config}
// }

// func (s *AuthService) Authenticate(username, password string) (string, error) {
// 	user, err := s.repo.GetUserByUsername(username)
// 	if err != nil {
// 		return "", err
// 	}

// 	if user == nil || user.Password != password {
// 		return "", errors.New("invalid username or password")
// 	}

// 	claims := JWTClaims{
// 		jwt.StandardClaims{},
// 		user.Username,
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString([]byte(s.config.JWTSecret))
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokenString, nil
// }

// func (s *AuthService) Register(username, password string) error {
// 	user, err := s.repo.GetUserByUsername(username)
// 	if err != nil {
// 		return err
// 	}

// 	if user != nil {
// 		return errors.New("username already taken")
// 	}

// 	user = &auth.User{
// 		Username: username,
// 		Password: password,
// 	}

// 	err = s.repo.SaveUser(user)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func Login(c *fiber.Ctx) error {
//     var data map[string]string

//     if err := c.BodyParser(&data); err != nil {
//         return err
//     }

//     var user models.User

//     database.DB.Where("email = ?", data["email"]).First(&user) //Check the email is present in the DB

//     if user.ID == 0 { //If the ID return is '0' then there is no such email present in the DB
//         c.Status(fiber.StatusNotFound)
//         return c.JSON(fiber.Map{
//             "message": "user not found",
//         })
//     }

//     if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
//         c.Status(fiber.StatusBadRequest)
//         return c.JSON(fiber.Map{
//             "message": "incorrect password",
//         })
//     } // If the email is present in the DB then compare the Passwords and if incorrect password then return error.

//     claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
//         Issuer:    strconv.Itoa(int(user.ID)), //issuer contains the ID of the user.
//         ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //Adds time to the token i.e. 24 hours.
//     })

//     token, err := claims.SignedString([]byte(SecretKey))

//     if err != nil {
//         c.Status(fiber.StatusInternalServerError)
//         return c.JSON(fiber.Map{
//             "message": "could not login",
//         })
//     }

//     cookie := fiber.Cookie{
//         Name:     "jwt",
//         Value:    token,
//         Expires:  time.Now().Add(time.Hour * 24),
//         HTTPOnly: true,
//     } //Creates the cookie to be passed.

//     c.Cookie(&cookie)

//     return c.JSON(fiber.Map{
//         "message": "success",
//     })
// }

// func User(c *fiber.Ctx) error {
//     cookie := c.Cookies("jwt")

//     token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
//         return []byte(SecretKey), nil //using the SecretKey which was generated in th Login function
//     })

//     if err != nil {
//         c.Status(fiber.StatusUnauthorized)
//         return c.JSON(fiber.Map{
//             "message": "unauthenticated",
//         })
//     }

//     claims := token.Claims.(*jwt.StandardClaims)

//     var user models.User

//     database.DB.Where("id = ?", claims.Issuer).First(&user)

//     return c.JSON(user)

// }

// func Logout(c *fiber.Ctx) error {
//     cookie := fiber.Cookie{
//         Name:     "jwt",
//         Value:    "",
//         Expires:  time.Now().Add(-time.Hour), //Sets the expiry time an hour ago in the past.
//         HTTPOnly: true,
//     }

//     c.Cookie(&cookie)

//     return c.JSON(fiber.Map{
//         "message": "success",
//     })

// }

// err = dbRepo.Transaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
// 	id, err = userRepo.CreateAUser(ctx, tx, user)
// 	if err != nil {
// 		return err
// 	}

// 	userRepo.UpdateAUsersName(ctx, tx, id, body.FirstName, "test")
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// })

// var id int

// err = dbRepo.Transaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
// 	id, err = userRepo.CreateAUser(ctx, tx, user)
// 	if err != nil {
// 		return err
// 	}

// 	userRepo.UpdateAUsersName(ctx, tx, id, body.FirstName, "test")
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// })

// if err != nil {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusInternalServerError)
// 	return
// }

// transacaoResponse, err := services.CreateTransacao(clienteId, &transacao)
// if err != nil {
// 	if errors.Is(err, services.ErroClienteNaoExiste) {
// 		return c.Status(http.StatusNotFound).JSON(&fiber.Map{"erro": "Cliente não existe."})
// 	}
// 	if errors.Is(err, services.ErroValorDaTransacao) {
// 		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"erro": "O valor da transação deve ser maior que zero."})
// 	}
// 	if errors.Is(err, services.ErroTipoDaTransacao) {
// 		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"erro": "A transação deve ser do tipo 'c' (crédito) ou 'd' (débito)."})
// 	}
// 	if errors.Is(err, services.ErroDescricao) {
// 		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"erro": "A descrição deve ter de 1 a 10 caractéres."})
// 	}
// 	if errors.Is(err, services.ErroTransacaoDebito) {
// 		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"erro":       "A transação do tipo 'd' (débito) nunca pode deixar o saldo do cliente menor que seu limite disponível.", "observacao": "Transações do tipo 'c' (crédito) ainda serão processadas."})
// 	}
// 	return c.Status(http.StatusBadRequest).JSON(&fiber.Map{"erro": err.Error()})
// }

// TransactionQuery(db *sql.DB, baseQuery string) (rows *sql.Rows, code int, err error) {
// 	tx, txErr := db.Begin()
// 	if txErr != nil {
// 		return nil, -1, txErr
// 	}

// 	selectStmt, prepErr := tx.Prepare(baseQuery)
// 	if prepErr != nil {
// 		return nil, -1, fmt.Errorf("Failed to prepare statment: %s Error: %v", baseQuery, prepErr)
// 	}

// 	defer func() {
// 		if stmtErr := selectStmt.Close(); stmtErr != nil {
// 			rows = nil
// 			code = -2
// 			err = fmt.Errorf("Failed to close statement: %v.", stmtErr)
// 		}
// 	}()

// 	rows, err = selectStmt.Query()
// 	if err != nil {
// 		fmt.Errorf("Failed to retrieve data: %v", err)
// 		return nil, -1, err
// 	}
// 	if txCloseErr := tx.Commit(); txErr != nil {
// 		rows = rows
// 		code = -3
// 		err = txCloseErr
// 	}
// 	return rows, 0, nil
// }

// type Service struct {
// 	db   *sqlx.DB
// 	repo *srepo
//   }

//   func NewService(db *sqlx.DB, repo *srepo) *Service {
// 	return &Service{repo: repo, db: db}
//   }

//   func (s *Service) CancelSubscription(ctx context.Context, id int64) (*Subscription, error) {
// 	tx, err := s.db.BeginTxx(ctx, nil)
// 	if err != nil {
// 	  return nil, err
// 	}

// 	defer func() {
// 	  // !!! This would not work if the subscriptions is already canceled
// 	  // and the error is not returned
// 	  if err != nil {
// 		_ = tx.Rollback()
// 		return
// 	  }
// 	}()

// 	sub, err := s.repo.GetSubscription(tx, id)
// 	if err != nil {
// 	  return nil, err
// 	}

// 	if sub.Status != "active" {
// 	  return &sub, nil
// 	}

// 	if sub.CanceledAt.Valid {
// 	  return &sub, nil
// 	}

// 	sub, err = s.repo.CancelSubscription(tx, id)
// 	if err != nil {
// 	  return nil, err
// 	}

// 	err = tx.Commit()

// 	return &sub, err
//   }

//   user, err := h.iDao.GetByCondition(ctx, condition)
//     if err != nil {
//         if errors.Is(err, model.ErrRecordNotFound) {
//             logger.Warn("Login not found", logger.Err(err), logger.Any("form", req), middleware.GCtxRequestIDField(c))
//             response.Error(c, ecode.ErrLoginUsers)
//         } else {
//             logger.Error("Login error", logger.Err(err), logger.Any("form", req), middleware.GCtxRequestIDField(c))
//             response.Output(c, ecode.InternalServerError.ToHTTPCode())
//         }
//         return
//     }

//     if !gocrypto.VerifyPassword(req.Password, user.Password) {
//         logger.Warn("password error", middleware.CtxRequestIDField(ctx))
//         response.Error(c, ecode.ErrLoginUsers)
//     }

//     token, err := jwt.GenerateToken(utils.Uint64ToStr(user.ID), user.Username)
//     if err != nil {
//         logger.Error("jwt.GenerateToken error", logger.Err(err), middleware.CtxRequestIDField(ctx))
//         response.Output(c, ecode.InternalServerError.ToHTTPCode())
//     }

//     // TODO: save token to cache

//     response.Success(c, gin.H{
//         "id":    user.ID,
//         "token": token,
//     })
