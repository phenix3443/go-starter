# Golang Development with VSCode Devcontainers and GitHub Codespaces

## Table of Contents

1. [Motivation](#motivation)
2. [Setting Up VSCode Devcontainer Locally](#setting-up-vscode-devcontainer-locally)
3. [Using GitHub Codespaces in the Browser](#using-github-codespaces-in-the-browser)
4. [Upgrading golang Version](#upgrading-golang-version)

## Motivation

Setting up a local development environment for go-starter can be time-consuming and error-prone. This devcontainer setup streamlines the process by providing a pre-configured environment with all necessary dependencies. It ensures consistency across development environments, reduces setup time, and allows developers to focus on coding rather than configuration.

Key benefits include:

- Pre-configured environment with golang, and Node.js
- Essential VS Code extensions pre-installed
- Consistent development environment across team members

## Setting Up VSCode Devcontainer Locally

1. Clone the go-starter repository:

   ```bash
   git clone https://github.com/go-starter/go-starter.git
   cd go-starter
   ```

2. Open the project in VS Code:

   ```bash
   code .
   ```

3. Before re-opening in the container, you may find it useful to configure SSH authorization. To do this:

   a. Ensure you have SSH access to GitHub configured on your local machine.

   b. Open `.devcontainer/devcontainer.json`.

   c. Uncomment the `mounts` section:

   ```json
   "mounts": [
     "source=${localEnv:HOME}/.ssh/config,target=/home/vscode/.ssh/config,type=bind,consistency=cached",
     "source=${localEnv:HOME}/.ssh/id_ed25519,target=/home/vscode/.ssh/id_ed25519,type=bind,consistency=cached"
   ],
   ```

   d. Adjust the paths or ssh file name if your SSH keys are stored in a different location.

   e. Use `git update-index --assume-unchanged .devcontainer/devcontainer.json` to prevent the changes to `devcontainer.json` from appearing in `git status` and VS Code's Source Control. To undo the changes, use `git update-index --no-assume-unchanged .devcontainer/devcontainer.json`.

4. When prompted, click "Reopen in Container". If not prompted, press `F1`, type "Remote-Containers: Reopen in Container", and press Enter.

5. VS Code will build the devcontainer. This process includes:
   - Pulling the base Docker image
   - Installing specified VS Code extensions
   - Installing project dependencies

   This may take several minutes the first time.

6. Once the devcontainer is built, you'll be working inside the containerized environment.

7. If you modified the `devcontainer.json` file in step 3, you may want to execute `git update-index --assume-unchanged .devcontainer/devcontainer.json` in a terminal within your devcontainer to prevent the changes to `devcontainer.json` from appearing in `git status` and VS Code's Source Control.

### Signing in to GitHub for Pull Request Extension

1. In the devcontainer, click on the GitHub icon in the Primary sidebar.
2. Click on "Sign in to GitHub" and follow the prompts to authenticate.

## Using GitHub Codespaces in the Browser

To open the project in GitHub Codespaces:

1. Navigate to the go-starter repository on GitHub.
2. Switch to the branch you want to work on.
3. Click the "Code" button.
4. Instead of clicking "Create codespace on [branch]" (which would use the default machine type that may not be sufficient for this golang-based project), click on the three dots (...) next to it.
5. Select "New with options".
6. Choose the "4-core/16GB RAM" machine type for optimal performance.
7. Click "Create codespace".

This will create a new Codespace with the specified resources, ensuring adequate performance for the golang-based project.

Note: After the container opens, you may see an error about the inability to use "GitHub Copilot Chat". This Copilot functionality will not be accessible in the Codespace environment.

## Upgrading golang Version

To upgrade the golang version:

1. Open `.devcontainer/Dockerfile`.
2. Update the `VARIANT` argument with the desired golang version.
3. Rebuild the devcontainer.

Note: Ensure that the version you choose is compatible with the project dependencies.

After testing the new golang version, propagate the corresponding changes in the Dockerfile to the repo <https://github.com/phenix3443/devcontainer-golang>. Once a new release tag is published there and a new docker image `ghcr.io/phenix3443/devcontainer-golang` appears in the GitHub registry, modify the `docker-compose.yml` file in the `.devcontainer` directory to reflect the proper docker image tag.
