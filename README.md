# kman -- Linux Kernel Manager

kman aims to streamline the Linux Kernel installation from source process, allowing for a
"reproducible" way of downloading, compiling, installing, configuring, patching the kernel,
generating a initramfs image and configuring the bootloader, based on the tools available
of your system.

## Project Structure

```
kman/
|_ cmd/
|    |_ kman/
|        |_ kman.go
|_ internal/
|    |_ app
|    |    |_ bootloader
|    |    |    |_ dp_strategy.go
|    |    |    |_ grub.go
|    |    |    |_ systemd-boot.go
|    |    |   
|    |    |_ initramfs
|    |    |    |_ dp_strategy.go
|    |    |    |_ dracut.go
|    |    |    |_ mkinitcpio.go
|    |    |    |_ booster.go
|    |    |   
|    |    |_ kernel
|    |         |_ dp_facade.go
|    |         |_ dp_builder.go
|    |         |_ download.go
|    |         |_ verify.go
|    |         |_ configure.go
|    |         |_ compile.go
|    |         |_ install.go
|    |         |_ remove.go
|    |        
|    |_ pkg
|       |_ dp_command.go
|       |_ dp_state.go
|_ pkg
    |_ logger
    |    |_ logger.go
    |_ progress
    |    |_ progress.go
    |_ utils
         |_ utils.go
```

### Design Patterns

Command Pattern: Encapsulates each operation (downloading, configuring, etc)
into a command class, allowing for easy execution and management.

Strategy Pattern: The strategy lets the algorithm vary independently from clients
that use it. Used for handling multiple bootloaders (e.g., GRUB, LILO, systemd-boot)
and multiple initramfs tools (e.g., Dracut, mkinitcpio, booster).

State Machine Pattern: Used to represent different phases of the kernel management
process (downloading, installing, configuring, etc.) and manage transitions between these states.

Facade Pattern: Simplifies user interactions with the kernel management process,
providing a clear interface for common operations.

Builder Pattern: Provides a flexible way to construct kernel configuration parameters step-by-step.
