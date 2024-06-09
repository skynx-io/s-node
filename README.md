[![skynx.com](https://github.com/skynx-io/assets/blob/HEAD/images/logo/skynx-logo_black_180x45.png)](https://skynx.com)

[![Discord](https://img.shields.io/discord/654291649572241408?color=%236d82cb&style=flat&logo=discord&logoColor=%23ffffff&label=Chat)](https://skynx.io/discord)
[![GitHub Discussions](https://img.shields.io/badge/GitHub_Discussions-181717?style=flat&logo=github&logoColor=white)](https://github.com/orgs/skynx/discussions)
[![X](https://img.shields.io/badge/Follow_on_X-000000?style=flat&logo=x&logoColor=white)](https://x.com/skynxHQ)
[![Mastodon](https://img.shields.io/badge/Follow_on_Mastodon-2f0c7a?style=flat&logo=mastodon&logoColor=white)](https://mastodon.social/@skynx)

Open source projects from [skynx.com](https://skynx.com).

# skynx-node

[![Go Report Card](https://goreportcard.com/badge/skynx.io/s-node)](https://goreportcard.com/report/skynx.io/s-node)
[![Release](https://img.shields.io/github/v/release/skynx-io/s-node?display_name=tag&style=flat)](https://github.com/skynx-io/s-node/releases/latest)
[![GitHub](https://img.shields.io/github/license/skynx-io/s-node?style=flat)](/LICENSE)

This repository contains the `skynx-node` agent, the component that runs on the machines you want to connect to your [skynx](https://skynx.com) network.

`skynx-node` is available for a variety of Linux platforms, macOS and Windows.

## Minimum Requirements

`skynx-node` has the same [minimum requirements](https://github.com/golang/go/wiki/MinimumRequirements#minimum-requirements) as Go:

- Linux kernel version 2.6.23 or later
- Windows 7 or later
- FreeBSD 11.2 or later
- MacOS 10.11 El Capitan or later

## Getting Started

The instructions in this repo assume you already have a skynx account and are ready to start adding nodes.

See [Quick Start](https://skynx.com/docs/platform/getting-started/quickstart/) to learn how to start building your skynx cloud-agnostic architecture.

The fastest way to add Linux nodes to your skynx network is by [generating a magic link](#linux-installation-with-magic-link) in the skynx web UI or with `skynxctl`:

```shell
skynxctl node add
```

See [Installation](#installation) for more details and other platforms.

## Documentation

For the complete skynx platform documentation visit [skynx.com/docs](https://skynx.com/docs/).

## Installation

### Binary Downloads

Linux, macOS and Windows binary downloads are available from the [Releases](https://github.com/skynx-io/s-node/releases) page.

You can download the pre-compiled binaries and install them with the appropriate tools.

### Linux Installation

#### Linux installation with magic link

The easiest way to add nodes to your skynx network is by generating a magic link in the skynx web UI or with `skynxctl`:

```shell
skynxctl node add
```

You will be able to use the magic link to install the `skynx-node` agent in seconds with no additional configuration required.

Once installed you can review the configuration at `/etc/skynx/skynx-node.yml`.

> See the [skynx-node configuration reference](https://skynx.com/docs/platform/reference/skynx-node.yml/) to find all the configuration options.

#### Linux binary installation with curl

1. Download the latest release.

    ```shell
    curl -LO "https://dl.skynx.com/binaries/stable/latest/linux/amd64/skynx-node"
    ```

2. Validate the binary (optional).

    Download the skynx-node checksum file:

    ```shell
    curl -LO "https://dl.skynx.com/binaries/stable/latest/linux/amd64/skynx-node_checksum.sha256"
    ```

    Validate the skynx-node binary against the checksum file:

    ```bash
    sha256sum --check < skynx-node_checksum.sha256
    ```

    If valid, the output must be:

    ```console
    skynx-node: OK
    ```

    If the check fails, sha256 exits with nonzero status and prints output similar to:

    ```console
    skynx-node: FAILED
    sha256sum: WARNING: 1 computed checksum did NOT match
    ```

3. Install skynx-node and create its configuration file according to your needs.

    ```shell
    sudo install -o root -g root -m 0750 skynx-node /usr/local/bin/skynx-node
    sudo mkdir /etc/skynx
    sudo vim /etc/skynx/skynx-node.yml
    ```

    See the [skynx-node configuration reference](https://skynx.com/docs/platform/reference/skynx-node.yml/) to find all the configuration options.

4. Create the `skynx-node.service` for systemd.

    ```shell
    sudo cat << EOF > /etc/systemd/system/skynx-node.service
    [Unit]
    Description=skynx-node service
    Documentation=https://github.com/skynx-io/s-node
    After=network.target

    [Service]
    Type=simple

    # Another Type: forking

    # User=
    WorkingDirectory=/var/local/skynx
    ExecStart=/usr/local/bin/skynx-node start
    Restart=always

    # Other restart options: always, on-abort, etc

    # The install section is needed to use

    # 'systemctl enable' to start on boot

    # For a user service that you want to enable

    # and start automatically, use 'default.target'

    # For system level services, use 'multi-user.target'

    [Install]
    WantedBy=multi-user.target
    EOF
    ```

5. Ensure the `tun` kernel module is loaded.

    ```shell
    sudo modprobe tun
    ```

6. Start the skynx-node service.

    ```shell
    sudo systemctl daemon-reload
    sudo systemctl enable skynx-node
    sudo systemctl restart skynx-node
    ```

##### Uninstall Linux skynx-node

To remove `skynx-node` from the system, use the following commands:

```shell
sudo systemctl stop skynx-node
sudo systemctl disable skynx-node
sudo rm /etc/systemd/system/skynx-node.service
sudo systemctl daemon-reload
sudo rm /usr/local/bin/skynx-node
sudo rm /etc/skynx/skynx-node.yml
sudo rmdir /etc/skynx
```

#### Package Repository

skynx provides a package repository that contains both DEB and RPM downloads.

For DEB-based platforms (e.g. Ubuntu and Debian) run the following to setup a new APT sources.list entry and install `skynx-node`:

```shell
echo 'deb [trusted=yes] https://repo.skynx.com/apt/ /' | sudo tee /etc/apt/sources.list.d/skynx.list
sudo apt update
sudo apt install skynx-node
```

For RPM-based platforms (e.g. RHEL, CentOS) use the following to create a repo file and install `skynx-node`:

```shell
cat <<EOF | sudo tee /etc/yum.repos.d/skynx.repo
[skynx]
name=skynx Repository - Stable
baseurl=https://repo.skynx.com/yum
enabled=1
gpgcheck=0
EOF
sudo yum install skynx-node
```

### macOS Installation

#### macOS binary installation with curl

1. Download the latest release.

    **Intel**:

    ```shell
    curl -LO "https://dl.skynx.com/binaries/stable/latest/darwin/amd64/skynx-node"
    ```

    **Apple Silicon**:

    ```shell
    curl -LO "https://dl.skynx.com/binaries/stable/latest/darwin/arm64/skynx-node"
    ```

2. Validate the binary (optional).

    Download the skynx-node checksum file:

    **Intel**:

    ```shell
    curl -LO "https://dl.skynx.com/binaries/stable/latest/darwin/amd64/skynx-node_checksum.sha256"
    ```

    **Apple Silicon**:

    ```shell
    curl -LO "https://dl.skynx.com/binaries/stable/latest/darwin/arm64/skynx-node_checksum.sha256"
    ```

    Validate the skynx-node binary against the checksum file:

    ```console
    shasum --algorithm 256 --check skynx-node_checksum.sha256
    ```

    If valid, the output must be:

    ```console
    skynx-node: OK
    ```

    If the check fails, sha256 exits with nonzero status and prints output similar to:

    ```console
    skynx-node: FAILED
    sha256sum: WARNING: 1 computed checksum did NOT match
    ```

3. Install `skynx-node` and create its configuration file according to your needs.

    ```console
    chmod +x skynx-node
    sudo mkdir -p /usr/local/libexec
    sudo mv skynx-node /usr/local/libexec/skynx-node
    sudo chown root: /usr/local/libexec/skynx-node
    sudo mkdir /etc/skynx
    sudo vim /etc/skynx/skynx-node.yml
    sudo chmod 600 /etc/skynx/skynx-node.yml
    ```

    > **IMPORTANT**: In macOS, `iface` must be `utun[0-9]+` in the `skynx-node.yml`, being `utun7` usually a good choice for that setting. Use the command `ifconfig -a` before launching the `skynx-node` service and check that the interface is not in-use.

    See the [skynx-node configuration reference](https://skynx.com/docs/platform/reference/skynx-node.yml/) to find all the configuration options.

4. Install and start the skynx-node agent as a system service.

    ```shell
    sudo /usr/local/libexec/skynx-node service-install
    ```

5. Check the service status.

    ```shell
    launchctl print system/com.skynx.skynx-node
    ```

    You should get an output like this:

    ```console
    system/com.skynx.skynx-node = {
        active count = 1
        path = /Library/LaunchDaemons/com.skynx.skynx-node.plist
        state = running

        program = /usr/local/libexec/skynx-node
        arguments = {
            /usr/local/libexec/skynx-node
            service-start
        }

        working directory = /var/tmp

        stdout path = /usr/local/var/log/com.skynx.skynx-node.out.log
        stderr path = /usr/local/var/log/com.skynx.skynx-node.err.log
        default environment = {
            PATH => /usr/bin:/bin:/usr/sbin:/sbin
        }

        environment = {
            XPC_SERVICE_NAME => com.skynx.skynx-node
        }

        domain = system
        minimum runtime = 10
        exit timeout = 5
        runs = 1
        pid = 3925
        immediate reason = speculative
        forks = 28
        execs = 1
        initialized = 1
        trampolined = 1
        started suspended = 0
        proxy started suspended = 0
        last exit code = (never exited)

        spawn type = daemon (3)
        jetsam priority = 4
        jetsam memory limit (active) = (unlimited)
        jetsam memory limit (inactive) = (unlimited)
        jetsamproperties category = daemon
        submitted job. ignore execute allowed
        jetsam thread limit = 32
        cpumon = default

        properties = keepalive | runatload | inferred program
    }
    ```

##### Uninstall macOS skynx-node

To remove `skynx-node` from the system, use the following commands:

```shell
sudo /usr/local/libexec/skynx-node service-uninstall
sudo rm /usr/local/libexec/skynx-node
sudo rm /etc/skynx/skynx-node.yml
sudo rmdir /etc/skynx
```

### Windows Installation

#### Windows binary installation with curl

1. Open the Command Prompt as Administrator and create a folder for skynx.

    ```shell
    mkdir 'C:\Program Files\skynx'
    ```

2. Download the latest release into the skynx folder.

    ```shell
    curl -LO "https://dl.skynx.com/binaries/stable/latest/windows/amd64/skynx-node.exe"
    ```

3. Validate the binary (optional).

    Download the skynx-node.exe checksum file:

    ```shell
    curl -LO "https://dl.skynx.com/binaries/stable/latest/windows/amd64/skynx-node.exe_checksum.sha256"
    ```

    Validate the skynx-node.exe binary against the checksum file:

    - Using Command Prompt to manually compare CertUtil's output to the checksum file downloaded:

         ```shell
         CertUtil -hashfile skynx-node.exe SHA256
         type skynx-node.exe_checksum.sha256
         ```

    - Using PowerShell to automate the verification using the -eq operator to get a `True` or `False` result:

         ```powershell
         $($(CertUtil -hashfile .\skynx-node.exe SHA256)[1] -replace " ", "") -eq $(type .\skynx-node.exe_checksum.sha256).split(" ")[0]
         ```

4. Download the `wintun` driver from <https://wintun.net>.

5. Unzip the wintun archive and copy the AMD64 binary `wintun.dll` to `C:\Program Files\skynx`.

6. Use an editor to create the skynx-node configuration file `C:\Program Files\skynx\skynx-node.yml`.

    See the [skynx-node configuration reference](https://skynx.com/docs/platform/reference/skynx-node.yml/) to find all the configuration options.

7. Install the skynx-node agent as a Windows service.

    > The instructions below assume that the `wintun.dll`, `skynx-node.exe` and `skynx-node.yml` files are stored in `C:\Program Files\skynx`.

    ```shell
    .\skynx-node.exe service-install --config "C:\Program Files\skynx\skynx-node.yml"
    ```

    Make sure to provide the absolute path of the skynx-node.yml configuration file, otherwise the Windows service may fail to start.

8. Start the service.

    ```shell
    net start "skynx-node"
    ```

##### Uninstall Windows skynx-node

To remove `skynx-node` from the system, open the Command Prompt as Administrator and use the following commands:

```shell
net stop "skynx-node"
cd 'C:\Program Files\skynx'
.\skynx-node.exe service-uninstall
del *.*
cd ..
rmdir 'C:\Program Files\skynx'
```

## Artifacts Verification

### Binaries

All artifacts are checksummed and the checksum file is signed with [cosign](https://github.com/sigstore/cosign).

1. Download the files you want and the `checksums.txt`, `checksum.txt.pem` and `checksums.txt.sig` files from the [Releases](https://github.com/skynx-io/s-node/releases) page:

2. Verify the signature:

    ```shell
    cosign verify-blob \
        --cert checksums.txt.pem \
        --signature checksums.txt.sig \
        checksums.txt
    ```

3. If the signature is valid, you can then verify the SHA256 sums match with the downloaded binary:

    ```shell
    sha256sum --ignore-missing -c checksums.txt
    ```

### Docker Images

Our Docker images are signed with [cosign](https://github.com/sigstore/cosign).

Verify the signatures:

```console
COSIGN_EXPERIMENTAL=1 cosign verify skynx/skynx-node
```

## Configuration

See the [skynx-node configuration reference](https://skynx.com/docs/platform/reference/skynx-node.yml/) to find all the configuration options.

## Running with Docker

You can also run the `skynx-node` agent as a Docker container. See examples below.

Registries:

- `skynx/skynx-node`
- `ghcr.io/skynx-io/skynx-node`

Example usage:

```shell
docker run -d --restart=always \
  --net=host \
  --cap-add=net_admin \
  --device=/dev/net/tun \
  --name skynx-node \
  -v /etc/skynx:/etc/skynx:ro \
  skynx/skynx-node:latest start
```

## Community

Have questions, need support and or just want to talk about skynx?

Get in touch with the skynx community!

[![Discord](https://img.shields.io/badge/Join_us_on_Discord-5865F2?style=flat&logo=discord&logoColor=white)](https://skynx.com/discord)
[![GitHub Discussions](https://img.shields.io/badge/GitHub_Discussions-181717?style=flat&logo=github&logoColor=white)](https://github.com/orgs/skynx/discussions)
[![X](https://img.shields.io/badge/Follow_on_X-000000?style=flat&logo=x&logoColor=white)](https://x.com/skynxHQ)
[![Mastodon](https://img.shields.io/badge/Follow_on_Mastodon-2f0c7a?style=flat&logo=mastodon&logoColor=white)](https://mastodon.social/@skynx)

## Code of Conduct

Participation in the skynx community is governed by the Contributor Covenant [Code of Conduct](https://github.com/skynx-io/.github/blob/HEAD/CODE_OF_CONDUCT.md). Please make sure to read and observe this document.

Please make sure to read and observe this document. By participating, you are expected to uphold this code.

## License

The skynx open source projects are licensed under the [Apache 2.0 License](/LICENSE).

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fskynx-io%2Fs-node.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fskynx-io%2Fs-node?ref=badge_large)
