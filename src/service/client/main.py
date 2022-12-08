import time
import sys


from shutil import copyfile
from signal import signal, SIGINT
from lib.prependShell import prependInterface


userHistory = []


def helpHandler(args: [] or None = None):
    commands = {
        "help": "Display this help message.",
        "swarm": "Perform swarm related operations.",
        "payload": "Perform payload generation and other payload operations."
    }

    commandsMenu = ""

    for command, description in commands.items():
        commandsMenu += f"{command} - {description}\n"

    if args is None or args == []:
        print("\n", commandsMenu)
    elif args is not None or args != []:
        commandsMenu = f"{args[0]} - {commands[args[0]]}\n"
        print(commandsMenu)


def payloadHandler(args: [] or None = None):
    if args is None or args == [] or args[0] == "-h" or args[0] == "help":
        print("""
Usage: payload [command] [arguments]

Commands:
    generate         Dynamically generates a payload 

Options:
    -O, --output     Output filename
    -o, --obfuscate  Whether to obfuscate the payload or not.
    -h, --help       Display this help menu.
            
        """)

    if args != None and len(args) != 0:
        if args[0] == "generate":
            print("Generating payload...")
    elif args == None and len(args) == 0:
        print("")

commandsRegistry = {
    "help": helpHandler,
    "payload": payloadHandler
}


def getChar():
    first_char = getch()
    if first_char == '\x1b':
        return {'[A': 'up', '[B': 'down', '[C': 'right', '[D': 'left'}[getch() + getch()]
    else:
        return first_char


def runCommand(command, args):
    commandsRegistry[command](args)
    argsStr = ""
    for arg in args:
        argsStr += arg

    userHistory.append(f"{command} {argsStr}".strip())


def main():
    while True:
        prependInterface()
        command = input("")

        if command != ("" or None):
            try:
                argv = command.split(" ")
                command = argv[0]
                argv.remove(command)
                runCommand(command, argv)
            except:
                print("Unknown command.")
                helpHandler()

def handleClose(sig, frame):
    print("\nQuitting client.")
    sys.exit(0)


def handleHistory():
    orderedUserHistory = userHistory[::-1]
    print(orderedUserHistory[len(orderedUserHistory) - 1], end="")


signal(SIGINT, handleClose)
main()
