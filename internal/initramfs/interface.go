package initramfs

type Initramfs interface {
	Generate() error
}
