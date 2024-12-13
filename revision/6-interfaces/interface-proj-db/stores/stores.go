package stores

import "interface-proj-db/stores/models"

type DataBase interface {
	Create(u models.User) (models.User, bool)
	Update(int, string) (models.User, bool)
	Delete(int) bool
	FetchAll() (map[int]models.User, bool)
	FetchUser(int) (models.User, bool)
}

func Create(d DataBase, u models.User) {
	d.Create(u)
}
