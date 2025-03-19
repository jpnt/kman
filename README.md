# kman (Linux Kernel Manager)

kman aims to automate and unify the Linux Kernel installation from source process, allowing for a
repeatable way of installing a Linux kernel, generating a initramfs image and updating
the bootloader configuration, based on the tools and configurations available of your system,
offering a sane and powerful way to manage kernels on any Linux distribution.

## Features


- [x] Minimal external libraries
- [x] Cross‑distro compatibility
- [x] Automated download and verification of kernel versions
- [ ] Accelerated & cached downloads, incremental updates
- [ ] Embedded tar.gz/tar.xz multi-threaded decompression
- [ ] Embedded key signature verification
- [ ] Ephemeral container build environments
- [ ] Configuration file support
- [ ] Support for multiple bootloaders
    - [ ] GNU grub
    - [ ] systemd-boot
    - [ ] limine
    - [ ] rEFInd
- [ ] Support for multiple initramfs generators
    - [ ] dracut
    - [ ] mkinitcpio
    - [ ] initramfs-tools
    - [ ] booster
- [ ] distcc, ccache, modprobed-db, and unified kernel image support

## Pipeline Steps

Each step in the pipeline can be executed indidually, but requires some parameter/s.

- Download (url)
- Verify (archive)
- Extract (archive)
- Patch (optional: directory)
- Configure (optional: .config file, directory, njobs)
- Compile (optional: directory)
- Install (optional: directory)
- Initramfs (optional: initramfs)
- Bootloader (optional: bootloader)

### Architecture

Some level of software architecture is adopted to make the project maintainability and evolution easier, such
as separation of concerns (e.g. validation of data separated from execution of code) and responsabilities
by layer (gateway, service, core); the use of design patterns for step dependency resolution
and interaction with multiple kinds of outside tools.

As a rule of thumb I like to keep my '.go' files as small as possible, have a component 
based design and have as little external dependencies as possible, even in outside layers.

The final program must work as a unified cohesive experience.

#### Layers

- Gateway: UI (CLI/TUI/GUI), I/O, Interaction with 3rd party programs.
- Service: Execution layer, use cases, orchestration logic.
- Core: Definition of entities, validation, data structures; Does not depend on anything.

Ideally dependencies should flow inward (gateway depends on service, service depends on core),
in practice this is achieved by extensive use of interfaces. Interfaces add overhead and they
are not always needed.

#### Design Patterns

- Strategy Pattern: Used to handle multiple bootloader and initramfs tools. The Service layer
  chooses the right strategy at runtime based on system capabilities and configuration.
- Factory Pattern: To allow for inversion of dependency. For example define the interface of
  how a step is created in core layer (internal/core/step_factory.go), and then implement on service
  layer (internal/service/step_factory.go).
- Pipeline Pattern: Each step (list, download, compile, etc.) is encapsulated as a modular
  component (step). The pipeline runs validation and execution separately.
- Builder Pattern: Flexible way to assemble step-by‑step the pipeline, choosing exacly the steps
  needed and coordenating and resolving dependencies.
