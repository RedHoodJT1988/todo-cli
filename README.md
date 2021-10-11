# todo-cli

This is a CLI application that allows you to run a complete ToDo app from your terminal application.
As a user you can:
- Create a list of todo items
- Get your list of items
- Mark your items as complete. 

Features that are missing but I am hoping to add at some point are:
- Using -del to delete an item from the list. This will utilize the Delete() method in the API portion 
of the application.

- Would like to add another flag for a verbose mode which would show information like date/time
- A flag to show the completed items
- Update the usage function to display the additional instructions on how to provide new tasks to the tool
- Update the getTask() function allowing it to handle multi-line input from STDIN. Each line would be a new task

And then of course add new tests such as:
- Include test cases for the remaining options like -complete
- Test the TODO_FILENAME environemnt variable instead of hard coding the test file name so it does not cause conflicts with an existing file
- Test for the getTask() update for multi-line input