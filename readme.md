# gh_explorer

A GUI stands for Graphical User Interface. In contrast, a TUI is a Terminal User Interface. This TUI application was made to help manage your github repositories on your local machine

From any directory in the terminal, type `jgh`, ( or the command of your choice) and a menu will pop up for you in the terminal. With this menu you can select the repository you would like to either clone or open! If you don't have the repository on your local machine the repository will be cloned to `~/Documents/your-github-account-name`. If you already have the repository cloned the selected repository will open the select repository with VSCode. Note: make sure you have VScode and its terminal feature installed.

## Use case

Let's say you want to open `my-repo` in your github account, and you don't have the repository cloned.

- Type `jgh` ( or the command of your choice)
- Type `/` to search for the repository (or scroll with arrow keys or i/j/k/l)
- Press `Enter` once to clone the repository
- Press `Enter` once more to open the folder

And BAM! You know have your github repo cloned and ready to use. You never have to login to get your repositories ever again!

## Installation

To install the application, follow the below steps:

1. Download the latest binary from the repository release page into your `~/Documents/your-github-user-account-name` folder.
2. Create an alias inside your `~/.zshrc` file equal to whatever name you would like. This name will be the command to run the application. See the below code snippet for an example.

```bash
# Example aliases
# alias zshconfig="mate ~/.zshrc"
# alias ohmyzsh="mate ~/.oh-my-zsh"

# my gh_explorer alias
# alias jgh="~/Documents/jonathon/gh_explorer --PAT ghp_1234567890 --USER jonathon_gh"
alias my-alias-name="~/Documents/your-github-username/gh_explorer --PAT <your-personal-access-token> --USER <your-github-username>"

```

3. Once the command has been added to your `~/.zshrc` file, quit all open terminals and re open a terminal. Or you can source your `~/.zshrc` file with this command `source ~/.zshrc`

4. After reloading your terminal in step 3, run the `jgh` command ( or your alias name you have chosen ) in thet erminal and the application should be ready to use!
