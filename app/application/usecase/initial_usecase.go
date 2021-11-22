package usecase

// ALLの場合は最初に作成される？
// Groupの場合は、Group作成時に作成される
// DMの場合はメッセージ送信時に無ければ作成される

type InitialUsecase interface {
	DataBaseInitialize() error
}
