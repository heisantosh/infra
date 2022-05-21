# Orchestration Services

Backend services for handling VM sessions, environments' pipeline, and the API for them.

## Development
Nomad dashboard - http://34.149.1.201/ui
Consule dashboard - https://35.244.214.226/ui

Architecture - https://www.figma.com/file/pr02o1okRpScOmNpAmgvCL/Architecture

### GCE Terraform troubleshooting
If you get stuck with a non-working Nomad on all VMs and therefore you are not able to change the infrastructure because the Nomad jobs that are supposedly running there cannot be reached use the:
```
sudo nomad agent -dev -bind 0.0.0.0 -log-level INFO
```

to start temporary Nomad on any server and tear it all down with `terraform apply -destroy`.

If the Nomad is running but it has no leader, delete Load balancer and Instance group then destroy the infrastructure.

### Subtrees
#### firecracker-task-driver
FC task driver is a submodule made from https://github.com/devbookhq/firecracker-task-driver repository.

The subtree commands you need for controling this repo are:
```bash
git subtree add --prefix modules/cluster/disk-image/firecracker-task-driver https://github.com/devbookhq/firecracker-task-driver.git master
```

```bash
git subtree pull --prefix modules/cluster/disk-image/firecracker-task-driver https://github.com/devbookhq/firecracker-task-driver.git master
```

```bash
git subtree push --prefix modules/cluster/disk-image/firecracker-task-driver https://github.com/devbookhq/firecracker-task-driver.git master
```

#### shared
FC task driver is a submodule made from https://github.com/devbookhq/shared repository.

The subtree commands you need for controling this repo are:
```bash
git subtree add --prefix shared https://github.com/devbookhq/shared.git master
```

```bash
git subtree pull --prefix shared https://github.com/devbookhq/shared.git master
```

```bash
git subtree push --prefix shared https://github.com/devbookhq/shared.git master
```

## Resources
### Clock drift
- https://github.com/firecracker-microvm/firecracker/blob/eb8de3ba1f7cb636d2aaa632fe96b234f3a302e6/FAQ.md#my-guest-wall-clock-is-drifting-how-can-i-fix-it

### Fs overlay
- https://www.kernel.org/doc/Documentation/device-mapper/snapshot.txt
- https://github.com/firecracker-microvm/firecracker-containerd/issues/75
- https://github.com/firecracker-microvm/firecracker-containerd/pull/153
- https://jvns.ca/blog/2019/11/18/how-containers-work--overlayfs/
- https://gruchalski.com/posts/2021-02-16-live-resize-firecracker-vmm-drive/
- https://github.com/firecracker-microvm/firecracker/pull/2743
- https://github.com/firecracker-microvm/firecracker/blob/34955a935c59c19361e0652b43e5bb77bca92da7/docs/overlay-filesystem.md
- https://github.com/firecracker-microvm/firecracker-containerd/blob/main/tools/image-builder/files_debootstrap/sbin/overlay-init#L39
