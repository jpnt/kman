package bootloader

type Bootloader interface {
	Configure() error
}
