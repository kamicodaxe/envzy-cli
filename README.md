# Envzy Command-Line Interface (CLI)

Envzy CLI is a powerful command-line tool for managing your environment variables and secrets easily. This command-line interface provides intuitive commands and options to help you create, organize, and maintain your projects secrets effectively.

## Installation (LINUX)

To use the Envzy CLI, you'll need to install it on your system. You can download the latest release for your platform from the [Envzy GitHub releases page](https://github.com/kamicodaxe/envzy-cli/releases).

After downloading, make the binary executable and place it in a directory included in your system's PATH.

```sh
chmod +x envzy
sudo mv envzy /usr/local/bin
```

Now, you can use the `envzy` command from your terminal.

## Usage

Envzy CLI supports several ways to interact with your projects, branches, and secrets.

## Available Commands

Here are the commands available in the Envzy CLI along with their short forms (aliases):

| Command                                                  | Short Form | Description                                                                             |
| -------------------------------------------------------- | ---------- | --------------------------------------------------------------------------------------- |
| `envzy project create [projectName] [-d\|--description]` | `pc`       | Creates a new project with an optional description.                                     |
| `envzy project list`                                     | `pl`       | Lists all available projects.                                                           |
| `envzy project select [projectName]`                     | `ps`       | Selects a project as the active project.                                                |
| `envzy project update [projectName] [-d\|--description]` | `pu`       | Updates the details of an existing project.                                             |
| `envzy project delete [projectName]`                     | `pd`       | Deletes an existing project.                                                            |
| `envzy branch list`                                      | `bl`       | Lists all branches within the active project.                                           |
| `envzy branch create [branchName] [-d\|--description]`   | `bc`       | Creates a new branch within the active project with an optional description.            |
| `envzy branch select [branchName]`                       | `bs`       | Selects a branch within the active project.                                             |
| `envzy branch update [branchName] [-d\|--description]`   | `bu`       | Updates the details of an existing branch.                                              |
| `envzy branch delete [branchName]`                       | `bd`       | Deletes an existing branch.                                                             |
| `envzy secret list [-b\|--branch]`                       | `sl`       | Lists all secrets within the active branch or a specified branch.                       |
| `envzy secret create [secretName] [value]`               | `sc`       | Adds a new secret to the active branch.                                                 |
| `envzy secret update [secretName] [newValue]`            | `su`       | Updates the value of an existing secret.                                                |
| `envzy secret delete [secretName]`                       | `sd`       | Deletes an existing secret.                                                             |
| `envzy help [command]`                                   | `h`        | Provides help and usage information for a specific command.                             |
| `envzy version`                                          | `v`        | Displays the version information of the Envzy CLI tool.                                 |
| `envzy config push path/to/dotenv`                       |            | Writes all env variables in path/to/dotenv to databse selected project and branch       |
| `envzy config pull path/to/dotenv`                       |            | Writes all env variables from selected project and branch in database to path/to/dotenv |

## Examples

**Project Management:**

1. Create a new project with a description:

   ```shell
   envzy project create myproject -d "My project description"
   ```

   - Creates a project named "myproject" with an optional description.

2. List all available projects:

   ```shell
   envzy project list
   ```

   - Lists all existing projects.

3. Select a project as the active project:

   ```shell
   envzy project select myproject
   ```

   - Sets "myproject" as the active project.

4. Update the details of an existing project:

   ```shell
   envzy project update myproject -d "Updated project description"
   ```

   - Updates the description of the "myproject" project.

5. Delete an existing project:

   ```shell
   envzy project delete myproject
   ```

   - Deletes the "myproject" project.

**Branch Management:**

6. List all branches within the active project:

   ```shell
   envzy branch
   ```

   - Lists all branches in the active project.

7. Create a new branch within the active project with a description:

   ```shell
   envzy branch create mybranch -d "My branch description"
   ```

   - Creates a branch named "mybranch" within the active project with an optional description.

8. Select a branch within the active project:

   ```shell
   envzy branch select mybranch
   ```

   - Sets "mybranch" as the active branch within the active project.

9. Update the details of an existing branch:

   ```shell
   envzy branch update mybranch -d "Updated branch description"
   ```

   - Updates the description of the "mybranch" branch.

10. Delete an existing branch:

    ```shell
    envzy branch delete mybranch
    ```

    - Deletes the "mybranch" branch.

**Secret Management:**

11. List all secrets within the active branch:

    ```shell
    envzy secret list
    ```

    - Lists all secrets within the active branch.

12. Add a new secret to the active branch:

    ```shell
    envzy secret add mysecret "mysecretvalue"
    ```

    - Adds a secret named "mysecret" with the value "mysecretvalue" to the active branch.

13. Update the value of an existing secret:

    ```shell
    envzy secret update mysecret "newsecretvalue"
    ```

    - Updates the value of the "mysecret" secret to "newsecretvalue."

14. Delete an existing secret:

    ```shell
    envzy secret delete mysecret
    ```

    - Deletes the "mysecret" secret.

**Configuration File:**

15. Pushing local configuration changes to cli:

    ```shell
    envzy config push path/to/dotenv/file
    ```

16. Pulling configuration files from cli and updating your local environment:

    ```shell
    envzy config pull path/to/save/dotenv/file
    ```

**Global Commands:**

17. Display version information of the Envzy CLI tool:

    ```shell
    envzy version
    ```

- Displays the current version of the Envzy CLI tool.

## Contributing

We welcome contributions from the community! If you'd like to contribute to Envzy, please follow our [contribution guidelines](CONTRIBUTING.md).

## License

Envzy is open-source software licensed under the [GNU General Public License (GPL)](https://www.gnu.org/licenses/gpl-3.0.md) version 3.0.

This software is free to use, modify, and distribute, as long as you comply with the terms and conditions of the GPL 3.0 license.

The GPL v3.0 ensures that any modifications or improvements made to Envzy will also be open source and available to the community. If you modify Envzy and make it available to others, you must also provide the corresponding source code.

Please read the full text of the GPL 3.0 license to understand your rights and obligations.

For the full license text, visit: [GNU GPL 3.0](https://www.gnu.org/licenses/gpl-3.0.md)
