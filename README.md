# Expense Tracker

This is a personal project designed to help me maintain control over household expenses while allowing me to do it while doing something I love: developing tech solutions through programming.

# How Is it Used?
After any expense(s), I open a terminal and quickly type my command `expense-tracker add` which will run me through questions for defining the expense(s). At the end of the month, I can run `expense-tracker summarize` and see some details on how our spending was for the month.

# Cool Features
### Excel Spreadsheet Export
* Adding --export to the `summarize` command will open a temporary excel spreadhseet populated with the data being used in the summary. This allows for any further exploration of the data to be done easily
### Short Args for CLI Commands
* I have made using this tool more convenient in my life by adding shorthand arguments to define values rather than going through the questionnaire style for everything

## Tech Used
### Golang
* Programming language used to create this tool
### Cobra
* Golang CLI library used to define structure of CLI
### MongoDB
* Database for storing expenses and categories of expenses