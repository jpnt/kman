package initramfs

type IInitramfs interface {
	Generate() error
}
