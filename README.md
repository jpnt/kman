# kman (Linux Kernel Manager)

kman aims to automate the Linux Kernel installation from source process, allowing for a
repeatable way of downloading, configuring, patching, compiling, installing,
generating a initramfs image and configuring the bootloader, based on the tools 
and configurations available of your system.

## Features

- [x] No External Library dependencies
- [x] Automated download of Kernel versions
- [x] Cross-distro compatibility
- [ ] Incremental updates via patching
- [ ] Support multiple bootloaders configuration
    - [ ] GNU grub
    - [ ] systemd-boot
    - [ ] limine
    - [ ] rEFInd
- [ ] Support multiple initramfs generators
    - [ ] dracut
    - [ ] mkinitcpio
    - [ ] initramfs-tools
    - [ ] booster
- [ ] Distributed compilation (distcc)
- [ ] Support modprobed-db to reduce compilation time
- [ ] Unified Kernel Image

### Design Patterns

Strategy Pattern: The strategy lets the algorithm vary independently from clients
that use it. Used for handling multiple bootloaders (e.g., GRUB, LILO, systemd-boot)
and multiple initramfs tools (e.g., Dracut, mkinitcpio, booster).

Command Pattern: Encapsulates each operation (downloading, configuring, etc)
into a command class, allowing for easy execution and management.

Builder Pattern: Provides a flexible way to construct kernel configuration parameters step-by-step.

Facade Pattern: Manages the kernel context and the flow of commands.
