package main

import (
	"math/rand"
)

type Api struct {
	Full   string `bson:"full" json:"full"`
	Secret string `bson:"secret" json:"secret"`
	Slug   string `bson:"slug" json:"slug"`
}

const secretCharset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"0123456789" +
	"~!@#$%^&*?<>"

const shortCharset = "abcdefghijklmnopqrstuvwxyz" +
	"0123456789"

func (api *Api) createSecret() string {
	secret := RandString(rand.Intn(5)+10, secretCharset)
	api.Secret = secret
	return secret
}

func (api *Api) createSlug() string {
	slug := RandString(rand.Intn(3)+5, shortCharset)
	api.Slug = slug
	return slug
}

func RandString(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
