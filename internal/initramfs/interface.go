package initramfs

type Initramfs interface {
	Configure() error
	Generate() error
}
