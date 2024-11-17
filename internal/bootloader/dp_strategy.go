package bootloader

type IBootloader interface {
	Configure() error
}
