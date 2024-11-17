# kman (Linux Kernel Manager)

kman aims to automate the Linux Kernel installation from source process, allowing for a
repeatable way of downloading, compiling, installing, configuring, patching,
generating a initramfs image and configuring the bootloader, based on the tools 
and configurations available of your system.

## Features

- Cross-distro compatibility
- No external dependencies
- Support for multiple bootloaders
- Support for multiple initramfs generators
- Distributed compilation (distcc)
- Incremental builds

### Bootloaders

- [x] GNU grub
- [ ] systemd-boot
- [ ] limine
- [ ] rEFInd

### Initramfs Generators

- [x] dracut
- [ ] mkinitcpio
- [ ] initramfs-tools
- [ ] booster

### Project Structure

```
kman/
|_ cmd/
|    |_ kman/
|        |_ kman.go
|_ internal/
|    |_ app
|         |_ bootloader
|         |    |_ dp_strategy.go
|         |    |_ grub.go
|         |    |_ systemd-boot.go
|         |    |_ limine.go
|         |   
|         |_ initramfs
|         |    |_ dp_strategy.go
|         |    |_ dracut.go
|         |    |_ mkinitcpio.go
|         |    |_ booster.go
|         |   
|         |_ kernel
|              |_ dp_facade.go
|              |_ dp_builder.go
|              |_ dp_command.go
|              |_ download.go
|              |_ verify.go
|              |_ configure.go
|              |_ compile.go
|              |_ install.go
|              |_ remove.go
|_ pkg
    |_ logger
    |    |_ logger.go
    |_ progress
    |    |_ progress.go
    |_ utils
         |_ utils.go
```

### Design Patterns

Strategy Pattern: The strategy lets the algorithm vary independently from clients
that use it. Used for handling multiple bootloaders (e.g., GRUB, LILO, systemd-boot)
and multiple initramfs tools (e.g., Dracut, mkinitcpio, booster).

Command Pattern: Encapsulates each operation (downloading, configuring, etc)
into a command class, allowing for easy execution and management.

Builder Pattern: Provides a flexible way to construct kernel configuration parameters step-by-step.

Facade Pattern: Manages the kernel context and the flow of commands.
