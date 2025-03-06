# kman (Linux Kernel Manager)

kman aims to automate the Linux Kernel installation from source process, allowing for a
repeatable way of installing a Linux kernel, generating a initramfs image and updating
the bootloader configuration, based on the tools and configurations available of your system.

## Features

- [x] No External Library dependencies
- [x] Automated download of Kernel versions
- [x] Cross-distro compatibility
- [ ] Configuration file
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
- [ ] distcc support for distributed compilation
- [ ] modprobed-db support to reduce compile time
- [ ] Unified kernel image

## Phases

Phases follow this order, all of them are optional and 
defined either with command flags or with a configuration file.

TODO: ensure all of this phases are correct and in right order

- Linux
	- list
	- download
	- verify
	- extract
	- configure
	- patch
	- compile
	- install
- Initramfs
	- generate
- Bootloader
	- configure

### Design Patterns

It is adopted some level of project maintainability by using some
of the following design patterns:

- Strategy Pattern: The strategy lets the algorithm vary independently from clients
  that use it. Used for handling multiple bootloaders (e.g., GRUB, LILO, systemd-boot)
  and multiple initramfs tools (e.g., Dracut, mkinitcpio, booster) depending on the system.
- Pipeline Pattern: Encapsulates each step (downloading, configuring, etc)
  into a set of steps (pipeline), allowing for easy execution and management. Ensures commands
  are ran in the correct order.
- Builder Pattern: Provides a flexible way to construct and add kernel configuration parameters
  step-by-step.
- Facade Pattern: Manages the kernel context and encapsulates the flow of commands into one component.
