package controllers

import (
	"net/http"

	"github.com/rinrey24/ptmmlj-be/database"
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
	"github.com/rinrey24/ptmmlj-be/helper"

	"github.com/rinrey24/ptmmlj-be/models"
)

func Login(c *fiber.Ctx) error {
	var userInput models.User

	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(userInput)
	// // mengambil inputan json
	// var userInput models.User
	// decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(&userInput); err != nil {
	// 	response := map[string]string{"message": err.Error()}
	// 	helper.ResponseJSON(w, http.StatusBadRequest, response)
	// 	return
	// }
	// defer r.Body.Close()

	// // ambil data user berdasarkan username
	// var user models.User
	// if err := database.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
	// 	switch err {
	// 	case gorm.ErrRecordNotFound:
	// 		response := map[string]string{"message": "Username atau password salah"}
	// 		helper.ResponseJSON(w, http.StatusUnauthorized, response)
	// 		return
	// 	default:
	// 		response := map[string]string{"message": err.Error()}
	// 		helper.ResponseJSON(w, http.StatusInternalServerError, response)
	// 		return
	// 	}
	// }

	// // cek apakah password valid
	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
	// 	response := map[string]string{"message": "Username atau password salah"}
	// 	helper.ResponseJSON(w, http.StatusUnauthorized, response)
	// 	return
	// }

	// // proses pembuatan token jwt
	// expTime := time.Now().Add(time.Minute * 1)
	// claims := &config.JWTClaim{
	// 	Username: user.Username,
	// 	RegisteredClaims: jwt.RegisteredClaims{
	// 		Issuer:    "go-jwt-mux",
	// 		ExpiresAt: jwt.NewNumericDate(expTime),
	// 	},
	// }

	// // medeklarasikan algoritma yang akan digunakan untuk signing
	// tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// // signed token
	// token, err := tokenAlgo.SignedString(config.JWT_KEY)
	// if err != nil {
	// 	response := map[string]string{"message": err.Error()}
	// 	helper.ResponseJSON(w, http.StatusInternalServerError, response)
	// 	return
	// }

	// // set token yang ke cookie
	// http.SetCookie(w, &http.Cookie{
	// 	Name:     "token",
	// 	Path:     "/",
	// 	Value:    token,
	// 	HttpOnly: true,
	// })

	// response := map[string]string{"message": "login berhasil"}
	// helper.ResponseJSON(w, http.StatusOK, response)
}

func Register(c *fiber.Ctx) error {

	var userInput models.User
	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// hash pass menggunakan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	if err := database.DB.Create(&userInput).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Akun berhasil dibuat!",
	})
	//return c.JSON(userInput)

	// // mengambil inputan json
	// var userInput models.User
	// decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(&userInput); err != nil {
	// 	response := map[string]string{"message": err.Error()}
	// 	helper.ResponseJSON(w, http.StatusBadRequest, response)
	// 	return
	// }
	// defer r.Body.Close()

	// // hash pass menggunakan bcrypt
	// hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	// userInput.Password = string(hashPassword)

	// // insert ke database
	// if err := database.DB.Create(&userInput).Error; err != nil {
	// 	response := map[string]string{"message": err.Error()}
	// 	helper.ResponseJSON(w, http.StatusInternalServerError, response)
	// 	return
	// }

	// response := map[string]string{"message": "success"}
	// helper.ResponseJSON(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// hapus token yang ada di cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "logout berhasil"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
