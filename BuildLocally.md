# Build locally

**Important notice:**

> The build infrastructure is based on tools available on *nix-sh operating
> systems like MacOS or Linux. Building under Windows is not supported. The
> provider can be build, installed and used in the Windows Subsytems for Linux
> (WSL) though.

## Prerequisites

1. Install `Go`

   Linux:

   ```sh
   $ sudo -E add-apt-repository ppa:longsleep/golang-backports
   $ sudo -E apt-get update
   $ sudo -E apt-get install golang-go
   ```
   
2. [Install `Pulumi`](https://www.pulumi.com/docs/get-started/azure/install-pulumi/)

    Linux:

    ```sh
    $ curl -fsSL https://get.pulumi.com | sh
    ```

3. Install `tf2pulumi` 

   https://github.com/pulumi/tf2pulumi

   >Ensure that you install the binary in a directory that is already in your `PATH`
   >variable and make sure that the binary is executable.

   Linux:

   ```sh
   $ curl -s "https://api.github.com/repos/pulumi/tf2pulumi/releases" \
     | jq -r ".[0].assets[] | select(.browser_download_url | test(\"$(uname | tr '[:upper:]' '[:lower:]')\")) | .browser_download_url" \
     | xargs curl -s -L \
     | tar xvz --directory="$(go env GOPATH)/bin"
   ```

4. Install `nodejs`

   Linux:

   ```sh
   $ curl -sL https://deb.nodesource.com/setup_12.x | sudo bash
   ```

5. Install `yarn`

   https://classic.yarnpkg.com/en/docs/install/#debian-stable

   Linux:

   ```
   $ curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
   $ echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list
   $ sudo -E apt update && sudo -E apt install --no-install-recommends yarn
   ```

6. Install `Python3` and `Pip`

   Linux:

   ```sh
   $ sudo -E apt install python3-pip
   ```

7. Install Dotnet

   https://docs.microsoft.com/en-us/dotnet/core/install/linux-package-manager-ubuntu-1804

   Linux:

   ```sh
   $ curl -L https://packages.microsoft.com/config/ubuntu/18.04/packages-microsoft-prod.deb --output packages-microsoft-prod.deb
   $ sudo dpkg -i packages-microsoft-prod.deb
   $ rm packages-microsoft-prod.deb
   $ sudo -E add-apt-repository universe
   $ sudo -E apt update
   $ sudo -E apt install apt-transport-https
   $ sudo -E apt update
   $ sudo -E apt install dotnet-sdk-3.1
   ```

8. Install `golangci-lint`

   https://github.com/golangci/golangci-lint

   Linux:

   ```sh
   $ curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.25.0
   ```

## Build and Install

To build and install th eprovider locally

1. Clone repository https://github.com/tmeckel/pulumi-azuredevops

2. Run `make build`

3. To perform a local install do `make install`

Samples for testing your installation can be found in in the [examples repository](https://github.com/tmeckel/pulumi-azuredevops-samples).
