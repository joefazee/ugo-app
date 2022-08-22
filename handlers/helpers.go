package handlers

import (
	"context"
	"github.com/joefazee/ugo"
	"net/http"
)

func (h *Handler) render(w http.ResponseWriter, r *http.Request, tmpl string, variables, data interface{}) error {
	return h.App.Render.Page(w, r, tmpl, variables, data)
}

func (h *Handler) sessionPut(ctx context.Context, key string, data interface{}) {
	h.App.Session.Put(ctx, key, data)
}

func (h *Handler) sessionHas(ctx context.Context, key string) bool {
	return h.App.Session.Exists(ctx, key)
}

func (h *Handler) sessionGet(ctx context.Context, key string) interface{} {
	return h.App.Session.Get(ctx, key)
}

func (h *Handler) sessionRemove(ctx context.Context, key string) {
	h.App.Session.Remove(ctx, key)
}

func (h *Handler) sessionDestroy(ctx context.Context) error {
	return h.App.Session.Destroy(ctx)
}

func (h *Handler) sessionRenew(ctx context.Context) error {
	return h.App.Session.RenewToken(ctx)
}

func (h *Handler) randomString(n int) string {
	return h.App.GenerateRandomString(n)
}

func (h *Handler) encrypt(text string) (string, error) {
	enc := ugo.Encryption{
		Key: []byte(h.App.EncryptionKey),
	}
	encrypted, err := enc.Encrypt(text)
	if err != nil {
		return "", err
	}
	return encrypted, nil
}

func (h *Handler) decrypt(encryptedText string) (string, error) {
	enc := ugo.Encryption{
		Key: []byte(h.App.EncryptionKey),
	}

	decrypted, err := enc.Decrypt(encryptedText)
	if err != nil {
		return "", err
	}
	return decrypted, nil
}
