ASCII-ART v1.2.1

FUNCTIONALITY
-This is a program that receives:
    One argument in the form of a [STRING] ("example" or 'example')
    OPTIONAL: An additional argument after the text that defines the [BANNER]
        *USAGE: go run . "examplestring" bannername
    OPTIONAL: An additional argument before the text for the [OPTION] to print the output to a text file 
               instead of printing it to the terminal
        *USAGE: go run . --output=examplefile.txt "examplestring"

    and outputs the ASCII art of the text in the chosen banner style

    *GENERAL USAGE: go run . [OPTION] [STRING] [BANNER]


-Attention: Special characters (", ! etc) inside the text that would be considered COMMANDS by the Go terminal
            must be formatted by the user using escapes (\).

CHANGELOG
-v1.0 - Created base ascii printing functionality
-v1.1 - Added the ability to switch between banners
-v1.2 - Added the ability to print the output into a file
-v1.2.1 - Added banners and funcs folders